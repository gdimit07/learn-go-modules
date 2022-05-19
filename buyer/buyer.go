package main

import (
	"fmt"
	"log"

	"example.com/cars"
)

func main() {

	log.SetPrefix("buyer: ")
	log.SetFlags(log.Ldate | log.Ltime)

	brand := "Tesla"
	model := "Cybertruck"
	//color := "silver"
	//milage := 17.7 //value set above 15 in purpose to check the error handling

	fmt.Println("Car details:")
	mes, err := cars.Brand(brand)
	if err != nil {
		log.Fatal(err) //prints the error and stops the program
	}
	fmt.Println(mes)

	mes, err = cars.Model(model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mes)

	mes, err = cars.Color()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mes)

	mes, err = cars.Milage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mes)
}
