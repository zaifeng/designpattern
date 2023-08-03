package simplefactory

type iParser interface {
	Parse(b []byte)
}

func NewParser(cate string) iParser {
	switch cate {
	case "xml":
		return XMLParser{}
	case "json":
		return JsonParser{}
	case "yaml":
		return YamlParser{}
	default:
		return XMLParser{}
	}
}
