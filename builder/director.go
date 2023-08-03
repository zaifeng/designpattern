package builder

// Director（导演者）定义了构建过程的顺序
type CarDirector struct {
	builder CarBuilder
}

func NewCarDirector(builder CarBuilder) *CarDirector {
	return &CarDirector{
		builder: builder,
	}
}

// Director负责封装复杂构建，并返回
func (m *CarDirector) Construct() *Car {
	m.builder.SetBrand("Ferrari")
	m.builder.SetColor("Red")
	m.builder.SetEngineVolume("4.0L")
	return m.builder.GetCar()
}
