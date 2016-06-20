// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api

import (
	l4g "github.com/alecthomas/log4go"
	goi18n "github.com/nicksnyder/go-i18n/i18n"

	"github.com/mattermost/platform/model"
	"github.com/mattermost/platform/utils"
)

type WebSocketContext struct {
	Session model.Session
	Seq     int64
	Conn    *WebConn
	Action  string
	Err     *model.AppError
	T       goi18n.TranslateFunc
	Locale  string
}

func ApiWebSocketHandler(wh func(*WebConn, *model.WebSocketRequest)) *webSocketHandler {
	return &webSocketHandler{wh}
}

type webSocketHandler struct {
	handlerFunc func(*WebConn, *model.WebSocketRequest)
}

func (wh *webSocketHandler) ServeWebSocket(conn *WebConn, r *model.WebSocketRequest) {
	l4g.Debug(utils.T("websocket request: %v"), r.Action)
}
