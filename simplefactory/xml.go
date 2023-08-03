package simplefactory

type XMLParser struct{}

func (XMLParser) Parse(data []byte) {
	println("XML:", data)
}
