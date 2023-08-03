package builder

// ConcreteBuilder（具体建造者）实现Builder接口
type SportsCarBuilder struct {
	car *Car
}

func NewSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{
		car: &Car{},
	}
}

func (b *SportsCarBuilder) SetBrand(brand string) {
	b.car.brand = brand
}

func (b *SportsCarBuilder) SetColor(color string) {
	b.car.color = color
}

func (b *SportsCarBuilder) SetEngineVolume(engineVolume string) {
	b.car.engineVolume = engineVolume
}

func (b *SportsCarBuilder) GetCar() *Car {
	return b.car
}
