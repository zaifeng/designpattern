package factory

// ConcreteProductB是具体产品类B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B used"
}

// ConcreteFactoryB是具体工厂类B，用于创建产品B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}
