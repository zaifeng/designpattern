package factory

// Factory是抽象工厂接口
type Factory interface {
	CreateProduct() Product
}
