// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api

import (
	"testing"
	"time"
)

func TestWebSocket(t *testing.T) {
	th := Setup().InitBasic()
	//Client := th.BasicClient
	WebSocketClient, err := th.CreateWebSocketClient()
	if err != nil {
		t.Fatal(err)
	}
	defer WebSocketClient.Close()

	time.Sleep(300 * time.Millisecond)

	// Test closing and reconnecting
	WebSocketClient.Close()
	if err := WebSocketClient.Connect(); err != nil {
		t.Fatal(err)
	}

	WebSocketClient.Listen()
	WebSocketClient.SendMessage("ping", nil)

	time.Sleep(300 * time.Millisecond)

	if resp := <-WebSocketClient.ResponseChannel; resp.Data["text"].(string) != "pong" {
		t.Fatal("wrong response")
	}
}

/*func TestSocket(t *testing.T) {
	th := Setup().InitBasic()
	Client := th.BasicClient
	team := th.BasicTeam
	channel1 := th.BasicChannel
	channel2 := th.CreateChannel(Client, team)
	Client.Must(Client.AddChannelMember(channel1.Id, th.BasicUser2.Id))

	url := "ws://localhost" + utils.Cfg.ServiceSettings.ListenAddress + model.API_URL_SUFFIX + "/users/websocket"

	header1 := http.Header{}
	header1.Set(model.HEADER_AUTH, "BEARER "+Client.AuthToken)

	c1, _, err := websocket.DefaultDialer.Dial(url, header1)
	if err != nil {
		t.Fatal(err)
	}

	th.LoginBasic2()

	header2 := http.Header{}
	header2.Set(model.HEADER_AUTH, "BEARER "+Client.AuthToken)

	c2, _, err := websocket.DefaultDialer.Dial(url, header2)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(300 * time.Millisecond)

	var rmsg model.WebSocketMessage

	// Test sending message without a channelId
	m := model.NewWebSocketEvent(team.Id, "", "", model.EVENT_TYPING)
	m.Add("RootId", model.NewId())
	m.Add("ParentId", model.NewId())

	c1.WriteJSON(m)

	if err := c2.ReadJSON(&rmsg); err != nil {
		t.Fatal(err)
	}

	t.Log(rmsg.ToJson())

	if team.Id != rmsg.TeamId {
		t.Fatal("Ids do not match")
	}

	if m.Props["RootId"] != rmsg.Props["RootId"] {
		t.Fatal("Ids do not match")
	}

	// Test sending messsage to Channel you have access to
	m = model.NewWebSocketEvent(team.Id, channel1.Id, "", model.EVENT_TYPING)
	m.Add("RootId", model.NewId())
	m.Add("ParentId", model.NewId())

	c1.WriteJSON(m)

	if err := c2.ReadJSON(&rmsg); err != nil {
		t.Fatal(err)
	}

	if team.Id != rmsg.TeamId {
		t.Fatal("Ids do not match")
	}

	if m.Props["RootId"] != rmsg.Props["RootId"] {
		t.Fatal("Ids do not match")
	}

	// Test sending message to Channel you *do not* have access too
	m = model.NewWebSocketEvent("", channel2.Id, "", model.EVENT_TYPING)
	m.Add("RootId", model.NewId())
	m.Add("ParentId", model.NewId())

	c1.WriteJSON(m)

	go func() {
		if err := c2.ReadJSON(&rmsg); err != nil {
			t.Fatal(err)
		}

		t.Fatal(err)
	}()

	time.Sleep(2 * time.Second)
}*/
