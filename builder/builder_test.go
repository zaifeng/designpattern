package builder

import (
	"log"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := NewCarManufacturer(NewSportsCarBuilder())
	car := builder.Construct()
	log.Printf("%v %v %v\b", car.brand, car.color, car.engineVolume)
}
