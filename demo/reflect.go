package demo

import "fmt"

type Person interface {
	SayHello(name string)
	Run() string
}

type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Printf("Hello "+name, ", I am"+hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed", hero.Speed)
	return "Running"
}
