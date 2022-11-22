package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := map[string]Animal{
		"cow": Animal{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
		"bird": Animal{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
		"snake": Animal{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}
	fmt.Println("enter the name of animal and information requested (i.e. cow eat) and exit (x)")
	fmt.Print("current animals: ")
	for k := range animals {
		fmt.Print(k, " ")
	}
	fmt.Println()

	var name, req string
	for {
		fmt.Print("> ")
		fmt.Scan(&name)
		if _, exists := animals[name]; !exists {
			break
		}
		fmt.Scan(&req)

		switch req {
		case "eat":
			animals[name].Eat()
		case "move":
			animals[name].Move()
		case "speak":
			animals[name].Speak()
		}
	}
}
