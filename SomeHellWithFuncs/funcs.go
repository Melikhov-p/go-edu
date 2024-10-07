package SomeHellWithFuncs

import "fmt"

func main() {
	var do func(string) func(string) string = whatDo

	Print("dog", do("voice"))
	Print("dog", do("color"))
}

func Print(who string, how func(string) string) {
	fmt.Println(how(who))
}

func whatDo(do string) func(string) string {
	switch do {
	case "voice":
		return Say
	case "color":
		return Color
	}
	return nil
}

func Color(animal string) (color string) {
	switch animal {
	case "dog":
		color = "black"
	case "cat":
		color = "grey"
	case "gopher":
		color = "blue"
	default:
		color = "white"
	}
	return
}

func Say(animal string) (sound string) {
	switch animal {
	case "dog":
		sound = "gav"
	case "cat":
		sound = "moo"
	case "gopher":
		sound = "moo"
	default:
		sound = "heh"
	}
	return
}
