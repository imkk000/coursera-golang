package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow int

func (Cow) Eat() {
	fmt.Println("grass")
}
func (Cow) Move() {
	fmt.Println("walk")
}
func (Cow) Speak() {
	fmt.Println("moo")
}

type Bird int

func (Bird) Eat() {
	fmt.Println("worms")
}
func (Bird) Move() {
	fmt.Println("fly")
}
func (Bird) Speak() {
	fmt.Println("peep")
}

type Snake int

func (Snake) Eat() {
	fmt.Println("mice")
}
func (Snake) Move() {
	fmt.Println("slither")
}
func (Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	names := make(map[string]Animal)
	fmt.Println("Enter 'newanimal <name> <type>' for creating new animal (type=[cow,bird,snake])")
	fmt.Println("Enter 'query <name> <command>' for giving order to animal (command=[eat,move,speak])")
	var cmd, name, kind string
	for {
		fmt.Print("> ")
		fmt.Scan(&cmd)
		if cmd == "x" {
			break
		}
		fmt.Scan(&name, &kind)

		switch cmd {
		case "newanimal":
			if _, exists := names[name]; exists {
				fmt.Println(name, "already exists")
				continue
			}

			switch kind {
			case "cow":
				names[name] = new(Cow)
			case "bird":
				names[name] = new(Bird)
			case "snake":
				names[name] = new(Snake)
			default:
				fmt.Println("animal types can be only [cow,bird,snake]")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			if _, exists := names[name]; !exists {
				fmt.Println(name, "not exists")
				continue
			}

			switch kind {
			case "eat":
				names[name].Eat()
			case "move":
				names[name].Move()
			case "speak":
				names[name].Speak()
			}
		}
	}
}
