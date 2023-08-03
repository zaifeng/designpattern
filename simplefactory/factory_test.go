package simplefactory

import (
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	type args struct {
		cate string
	}
	tests := []struct {
		name string
		args args
		want iParser
	}{
		{"json", args{cate: "json"}, JsonParser{}},
		{"xml", args{cate: "xml"}, XMLParser{}},
		{"yaml", args{cate: "yaml"}, YamlParser{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(tt.args.cate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
