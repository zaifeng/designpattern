package simplefactory

type YamlParser struct{}

func (YamlParser) Parse(data []byte) {
	println("YAML:", data)
}
