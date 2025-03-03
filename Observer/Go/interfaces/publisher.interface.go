package interfaces

type Publisher interface {
	Register(Subscriber)
	Unregister(Subscriber)
	Notify()
	SetData(int)
	GetData() string
}
