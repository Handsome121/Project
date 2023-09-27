/*
简单工厂模式
*/
package main

import "fmt"

type Car interface {
	run()
}

type BMW struct{}

func (c *BMW) run() {
	fmt.Println("BMW is running")
}

type AUDI struct{}

func (e *AUDI) run() {
	fmt.Println("AUDI is running")

}

func GetCar(carType string) Car {
	switch carType {
	case "bmw":
		return &BMW{}
	case "audi":
		return &AUDI{}
	default:
		return nil
	}
}

func main() {
	bmw := GetCar("bmw")
	bmw.run()

	audi := GetCar("audi")
	audi.run()
}
