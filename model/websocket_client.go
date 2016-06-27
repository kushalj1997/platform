// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocketClient struct {
	Url             string          // The location of the server like "ws://localhost:8065"
	ApiUrl          string          // The api location of the server like "ws://localhost:8065/api/v3"
	Conn            *websocket.Conn // The WebSocket connection
	AuthToken       string          // The token used to open the WebSocket
	Sequence        int64           // The ever-incrementing sequence attached to each WebSocket action
	EventChannel    chan *WebSocketEvent
	ResponseChannel chan *WebSocketResponse
	ErrorChannel    chan *WebSocketError
}

// NewWebSocketClient constructs a new WebSocket client with convienence
// methods for talking to the server.
func NewWebSocketClient(url, authToken string) (*WebSocketClient, *AppError) {
	header := http.Header{}
	header.Set(HEADER_AUTH, "BEARER "+authToken)
	conn, _, err := websocket.DefaultDialer.Dial(url+API_URL_SUFFIX+"/users/websocket", header)
	if err != nil {
		return nil, NewLocAppError("NewWebSocketClient", "model.websocket_client.connect_fail.app_error", nil, err.Error())
	}

	return &WebSocketClient{
		url,
		url + API_URL_SUFFIX,
		conn,
		authToken,
		1,
		make(chan *WebSocketEvent),
		make(chan *WebSocketResponse),
		make(chan *WebSocketError),
	}, nil
}

func (wsc *WebSocketClient) Connect() *AppError {
	header := http.Header{}
	header.Set(HEADER_AUTH, "BEARER "+wsc.AuthToken)

	var err error
	wsc.Conn, _, err = websocket.DefaultDialer.Dial(wsc.ApiUrl+"/users/websocket", header)
	if err != nil {
		return NewLocAppError("NewWebSocketClient", "model.websocket_client.connect_fail.app_error", nil, err.Error())
	}

	return nil
}

func (wsc *WebSocketClient) Close() {
	wsc.Conn.Close()
}

func (wsc *WebSocketClient) Listen() {
	go func() {
		for {
			var msg WebSocketMessage
			if err := wsc.Conn.ReadJSON(&msg); err != nil {
				return
			}

			switch m := msg.(type) {
			case *WebSocketEvent:
				wsc.EventChannel <- m
			case *WebSocketResponse:
				wsc.ResponseChannel <- m
			case *WebSocketError:
				wsc.ErrorChannel <- m
			}
		}
	}()
}

func (wsc *WebSocketClient) SendMessage(action string, data map[string]interface{}) {
	req := &WebSocketRequest{}
	req.Seq = wsc.Sequence
	req.Action = action
	req.Data = data

	wsc.Sequence++

	wsc.Conn.WriteJSON(req)
}

func (wsc *WebSocketClient) UserTyping(teamId, channelId, parentId string) {
	data := map[string]interface{}{
		"team_id":    teamId,
		"channel_id": channelId,
		"parent_id":  parentId,
	}

	wsc.SendMessage("user_typing", data)
}
