package factory

import (
	"log"
	"testing"
)

func TestFactory(t *testing.T) {
	// 使用具体工厂类A来创建产品A
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	log.Println(productA.Use()) // Output: Product A used

	// 使用具体工厂类B来创建产品B
	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	log.Println(productB.Use()) // Output: Product B used
}
