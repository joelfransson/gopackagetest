package p1

type greeter interface {
	SayHello()
}

func SayHello(g greeter) {
	g.SayHello()
}
