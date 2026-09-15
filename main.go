package main

import "fmt"

type TYpechecker interface {
	Hello(h string)
}

type Person struct {
	Name   string
	age    int
	weight float32
}

func (p Person) Greet() {
	fmt.Print(p.Name)

}
func (p Person) Hello(s string) {
	fmt.Print("hello")

}

func main() {

	var arr []int = []int{1, 2, 3}
	fmt.Println(arr[0])

	dict := map[string]int{"Aswin": 1}
	fmt.Print(dict["aswin"])

	var Aswin Person
	Aswin.Name = "aswin"
	Aswin.age = 12
	Aswin.weight = 55
	fmt.Print(Aswin)

	Kiran := Person{Name: "Kiran",
		age:    12,
		weight: Aswin.weight}
	fmt.Print(Kiran)

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)

	// }
	// i := 1
	// for i < 2 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// nums := []int{1, 2, 3}
	// for index, value := range nums {
	// 	fmt.Printf("%d %d", index, value)

	// }
	Kiran.Greet()
}
