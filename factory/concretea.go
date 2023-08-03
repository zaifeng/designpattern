package factory

// ConcreteProductA是具体产品类A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A used"
}

// ConcreteFactoryA是具体工厂类A，用于创建产品A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}
