package subscribers

import (
	"fmt"
	"observer/interfaces"
)

type EmailSubscriber struct {
	ID        int
	Username  string
	Publisher interfaces.Publisher
}

func (e *EmailSubscriber) GetID() int {
	return e.ID
}

func (e *EmailSubscriber) Update() {
	e.sendEmail(e.Publisher.GetData())
}

func (e *EmailSubscriber) sendEmail(itemName string) {
	fmt.Println("sending email to : ", e.Username, ". ", itemName, " is back in stock")
}
