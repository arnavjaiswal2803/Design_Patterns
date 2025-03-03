package main

import (
	"observer/items"
	"observer/subscribers"
)

func main() {
	iPhone := items.NewItem("iPhone")
	headphone := items.NewItem("headphone")

	emailSubscriber := &subscribers.EmailSubscriber{
		ID:        1,
		Username:  "abcd@gmail.com",
		Publisher: iPhone,
	}
	phoneSubscriber := &subscribers.SmsSubscriber{
		ID:        2,
		Phone:     "999xxxxx99",
		Publisher: headphone,
	}

	iPhone.Register(emailSubscriber)
	headphone.Register(phoneSubscriber)

	iPhone.SetData(3)
	headphone.SetData(10)
}
