package subscribers

import (
	"fmt"
	"observer/interfaces"
)

type SmsSubscriber struct {
	ID        int
	Phone     string
	Publisher interfaces.Publisher
}

func (e *SmsSubscriber) GetID() int {
	return e.ID
}

func (e *SmsSubscriber) Update() {
	e.sendSms(e.Publisher.GetData())
}

func (e *SmsSubscriber) sendSms(itemName string) {
	fmt.Println("sending sms to : ", e.Phone, ". ", itemName, " is back in stock")
}
