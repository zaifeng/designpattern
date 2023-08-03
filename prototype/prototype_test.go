package prototype

import (
	"log"
	"testing"
)

func TestCarClone(*testing.T) {
	// 创建一个原型对象
	prototypeCar := NewCar("Toyota", "Red")
	log.Println("Original Car:")
	log.Printf("Brand: %s\nColor: %s\n", prototypeCar.GetBrand(), prototypeCar.GetColor())

	// 复制原型对象
	cloneCar := prototypeCar.Clone().(*Car)
	log.Println("\nCloned Car:")
	log.Printf("\nBrand: %s\nColor: %s\n", cloneCar.GetBrand(), cloneCar.GetColor())
}
