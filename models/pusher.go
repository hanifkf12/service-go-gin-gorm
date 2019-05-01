package models

import "github.com/pusher/pusher-http-go"

type Pusher struct {
	Channel string                 `json:"channel"`
	Event   string                 `json:"event"`
	Data    map[string]interface{} `json:"data"`
}

//app_id = "764811"
//key = "9c3a2b819bac150a68ce"
//secret = "7b1658cb27bf2eb19420"
//cluster = "ap1"
func (p *Pusher) TriggerClient() error {
	client := pusher.Client{
		AppId:   "764811",
		Key:     "9c3a2b819bac150a68ce",
		Secret:  "7b1658cb27bf2eb19420",
		Cluster: "ap1",
	}
	_, err := client.Trigger(p.Channel, p.Event, p.Data)
	return err
}
