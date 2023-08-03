package simplefactory

type JsonParser struct{}

func (JsonParser) Parse(data []byte) {
	println("json:", data)
}
