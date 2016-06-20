// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type WebSocketRequest struct {
	Seq    int64                  `json:"seq"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

func (o *WebSocketRequest) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func WebSocketRequestFromJson(data io.Reader) *WebSocketRequest {
	decoder := json.NewDecoder(data)
	var o WebSocketRequest
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
