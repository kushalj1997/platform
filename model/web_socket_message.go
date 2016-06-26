// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

const (
	EVENT_TYPING             = "typing"
	EVENT_POSTED             = "posted"
	EVENT_POST_EDITED        = "post_edited"
	EVENT_POST_DELETED       = "post_deleted"
	EVENT_CHANNEL_DELETED    = "channel_deleted"
	EVENT_CHANNEL_VIEWED     = "channel_viewed"
	EVENT_DIRECT_ADDED       = "direct_added"
	EVENT_NEW_USER           = "new_user"
	EVENT_USER_ADDED         = "user_added"
	EVENT_USER_REMOVED       = "user_removed"
	EVENT_PREFERENCE_CHANGED = "preference_changed"
	EVENT_EPHEMERAL_MESSAGE  = "ephemeral_message"
	EVENT_STATUS_CHANGE      = "status_change"
)

type WebSocketMessage interface {
	ToJson() string
}

type WebSocketEvent struct {
	TeamId    string                 `json:"team_id"`
	ChannelId string                 `json:"channel_id"`
	UserId    string                 `json:"user_id"`
	Event     string                 `json:"event"`
	Data      map[string]interface{} `json:"data"`
}

func (m *WebSocketEvent) Add(key string, value interface{}) {
	m.Data[key] = value
}

func NewWebSocketEvent(teamId string, channelId string, userId string, event string) *WebSocketEvent {
	return &WebSocketEvent{TeamId: teamId, ChannelId: channelId, UserId: userId, Event: event, Data: make(map[string]interface{})}
}

func (o *WebSocketEvent) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func WebSocketEventFromJson(data io.Reader) *WebSocketEvent {
	decoder := json.NewDecoder(data)
	var o WebSocketEvent
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

type WebSocketResponse struct {
	Status   string                 `json:"status"`
	SeqReply int64                  `json:"seq_reply"`
	Data     map[string]interface{} `json:"data"`
}

func (m *WebSocketResponse) Add(key string, value interface{}) {
	m.Data[key] = value
}

func NewWebSocketResponse(status string, seqReply int64, data map[string]interface{}) *WebSocketResponse {
	return &WebSocketResponse{Status: status, SeqReply: seqReply, Data: data}
}

func (o *WebSocketResponse) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func WebSocketResponseFromJson(data io.Reader) *WebSocketResponse {
	decoder := json.NewDecoder(data)
	var o WebSocketResponse
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

type WebSocketError struct {
	Status   string    `json:"status"`
	SeqReply int64     `json:"seq_reply"`
	Error    *AppError `json:"error"`
}

func NewWebSocketError(seqReply int64, err *AppError) *WebSocketError {
	return &WebSocketError{Status: STATUS_FAIL, SeqReply: seqReply, Error: err}
}

func (o *WebSocketError) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func WebSocketErrorFromJson(data io.Reader) *WebSocketError {
	decoder := json.NewDecoder(data)
	var o WebSocketError
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
