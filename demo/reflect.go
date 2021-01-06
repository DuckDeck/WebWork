package demo

import (
	"fmt"
	"reflect"
)

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

func TestReflect() {

	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Printf("Hero's type is %s, kind is %s \n", typeOfHero, typeOfHero.Kind())
	typeOfHeroPtr := reflect.TypeOf(&Hero{})
	fmt.Printf("Hero's type is %s, kind is %s \n", typeOfHeroPtr, typeOfHeroPtr.Kind())
	newTypeOfHero := typeOfHeroPtr.Elem()
	fmt.Printf("typeOfHeroPtr elem to TypeOfHero ,hero's type is %s, kind is %s \n", newTypeOfHero, newTypeOfHero.Kind())

	for i := 0; i < typeOfHero.NumField(); i++ {
		fmt.Printf("field's name is %s,type is %s,kind is %s\n", typeOfHero.Field(i).Name, typeOfHero.Field(i).Type, typeOfHero.Field(i).Type.Kind())
	}

	nameField, _ := typeOfHero.FieldByName("Name")
	fmt.Printf("field's name is %s,type is %s,kind is %s\n", nameField.Name, nameField.Type, nameField.Type.Kind())

	var person Person = &Hero{}
	typeOfPerson := reflect.TypeOf(person)

	for i := 0; i < typeOfPerson.NumMethod(); i++ {
		fmt.Printf("method's name is %s,type is %s,kind is %s\n", typeOfPerson.Method(i).Name, typeOfPerson.Method(i).Type, typeOfPerson.Method(i).Type.Kind())
	}

	method, res := typeOfPerson.MethodByName("Run")
	if res {
		fmt.Printf("method is %s,type is %s,kind is %s.\n", method.Name, method.Type, method.Type.Kind())
	}

	//使用反射对象来生成原始对象
	fmt.Printf("使用反射对象来生成原始对象\n")
	newReflectHeroObj := reflect.New(typeOfHero)
	fmt.Printf("使用反射生成的对象 new hero type is %s ,kind is %s \n", newReflectHeroObj.Type(), newReflectHeroObj.Kind())

	//修改name为菲夜
	fmt.Printf("使用反射来修改对象和字段\n")
	name := "菲夜"
	hero := &Hero{
		Name: "老王",
	}
	fmt.Printf("the hero name is : %s\n", hero.Name)
	valueOfHero := reflect.ValueOf(hero).Elem()
	valueOfName := valueOfHero.FieldByName("Name")
	if valueOfName.CanSet() {
		valueOfName.Set(reflect.ValueOf(name))
	}
	fmt.Printf("after alter the hero name is : %s\n", hero.Name)
	fmt.Printf("使用反射来调用对象方法\n")

	person = &Hero{
		Name:  "老王",
		Speed: 100,
	}

	valueOfPerson := reflect.ValueOf(person)
	sayHelloMethod := valueOfPerson.MethodByName("SayHello")
	sayHelloMethod.Call([]reflect.Value{reflect.ValueOf("小王")})

	runMethod := valueOfPerson.MethodByName("Run")
	callResult := runMethod.Call([]reflect.Value{})
	fmt.Printf("result of the run method: %s", callResult[0])
}
