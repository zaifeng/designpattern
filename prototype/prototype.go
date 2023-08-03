package prototype

// Cloneable是原型接口，定义了克隆方法
type Cloneable interface {
	Clone() Cloneable
}

// Car是具体原型类，实现了Cloneable接口
type Car struct {
	brand string
	color string
}

func NewCar(brand, color string) *Car {
	return &Car{
		brand: brand,
		color: color,
	}
}

// Clone实现了克隆方法，用于复制Car对象
func (c *Car) Clone() Cloneable {
	return &Car{
		brand: c.brand,
		color: c.color,
	}
}

func (c *Car) GetBrand() string {
	return c.brand
}

func (c *Car) GetColor() string {
	return c.color
}
