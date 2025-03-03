package items

import "observer/interfaces"

type Item struct {
	stockCount     int
	name           string
	subscriberList []interfaces.Subscriber
}

func NewItem(name string) *Item {
	return &Item{
		stockCount:     0,
		name:           name,
		subscriberList: make([]interfaces.Subscriber, 0),
	}
}

func (i *Item) Register(subscriber interfaces.Subscriber) {
	i.subscriberList = append(i.subscriberList, subscriber)
}

func (i *Item) Unregister(subscriber interfaces.Subscriber) {
	i.removeFromSubscriberList(i.subscriberList, subscriber)
}

func (i *Item) Notify() {
	for _, subscriber := range i.subscriberList {
		subscriber.Update()
	}
}

func (i *Item) SetData(stockCount int) {
	if i.stockCount == 0 {
		i.Notify()
	}
	i.stockCount = stockCount
}

func (i *Item) GetData() string {
	return i.name
}

func (i *Item) removeFromSubscriberList(subscriberList []interfaces.Subscriber, subscriberToRemove interfaces.Subscriber) {
	indexToRemove := -1
	for idx, subscriber := range subscriberList {
		if subscriber.GetID() == subscriberToRemove.GetID() {
			indexToRemove = idx
		}
	}

	i.subscriberList = append(i.subscriberList[:indexToRemove], i.subscriberList[indexToRemove+1:]...)
}
