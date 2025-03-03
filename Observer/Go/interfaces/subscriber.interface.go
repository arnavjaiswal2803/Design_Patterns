package interfaces

type Subscriber interface {
	GetID() int
	Update()
}
