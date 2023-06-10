package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "0xff0000",
		"green": "0x00ff00",
		"blue":  "0x0000ff",
	}
	fmt.Println(colors)

	// var moreColors map[string]string
	// fmt.Println(moreColors)
	// moreColors["red"] = "0xff0000"
	// fmt.Println(moreColors)

	color := make(map[string]string)
	color["red"] = "0xff0000"
	fmt.Println(color["red"])
	for key, values := range colors {
		fmt.Println(key, " : ", values)
	}
	delete(colors, "red")

	for key, values := range colors {
		fmt.Println(key, " : ", values)
	}

}

// func (m map[string]string) print() {
// 	fmt.Printf("%+v", m)
// }
