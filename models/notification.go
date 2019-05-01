package models

import (
	"fmt"
	"gopkg.in/maddevsio/fcm.v1"
)

const key = "AAAA4jJugEs:APA91bFqAhgDTpNeg-bua9PHKjAxvbzihjPy_h-LZxO9rRd-EHETqgm0lURvLhSJ8TSU62wXW1Ogh-V-LZ4zhC8Jb_fQ821lvX843ES4ssUBSy9L9YJmm4a12XKNLFe-20AYix0qp2wx"

var notif = fcm.NewFCM(key)

type Notification struct {
	Title string                 `json:"title"`
	Body  string                 `json:"body"`
	Data  map[string]interface{} `json:"data"`
}

func (n *Notification) SendNotification(token string) error {
	response, err := notif.Send(fcm.Message{
		Data:             n.Data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
	})
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)
	return err
}

func (n *Notification) SendNotificationTopic(topic string) error {
	response, err := notif.Send(fcm.Message{
		Data:             n.Data,
		To:               topic,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
	})
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)
	return err
}
