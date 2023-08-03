package builder

// Director（导演者）定义了构建过程的顺序
type CarManufacturer struct {
	builder CarBuilder
}

func NewCarManufacturer(builder CarBuilder) *CarManufacturer {
	return &CarManufacturer{
		builder: builder,
	}
}

func (m *CarManufacturer) Construct() *Car {
	m.builder.SetBrand("Ferrari")
	m.builder.SetColor("Red")
	m.builder.SetEngineVolume("4.0L")
	return m.builder.GetCar()
}
