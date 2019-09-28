package squaket

import (
	"errors"
	"reflect"
)

func New (element []interface{}) (*Squaket, error) {
	for _, someElement := range element {
		if reflect.ValueOf (someElement).Kind () != reflect.Struct {
			return nil, errors.New ("An element is not a struct.")
		}
	}
	return &Squaket {element}, nil
}

type Squaket struct {
	element []interface {}
}

func (s *Squaket) Group (property string) (g map[interface {}][]interface {}, e error) {
	g = map[interface {}][]interface {} {}

	for _, element := range s.element {
		value := reflect.ValueOf (element).FieldByName (property)
		if value.IsValid () == false {
			return nil, errors.New ("An element does not have that field.")
		}
		elements, okX := g [value.Interface ()]
		if okX == false {
			elements = []interface {} {}
		}
		elements = append (elements, element)
		g [value.Interface ()] = elements
	}
	return
}
