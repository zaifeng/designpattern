package builder

// Product（汽车）结构体，表示最终构建的复杂对象
type Car struct {
	brand        string
	color        string
	engineVolume string
}

// Builder（抽象建造者）接口
type CarBuilder interface {
	SetBrand(brand string)
	SetColor(color string)
	SetEngineVolume(engineVolume string)
	GetCar() *Car
}
