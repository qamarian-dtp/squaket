package squaket

import (
	"errors"
	"reflect"
	"regexp"
)

func init () {
	var errX error

	propertyPattern, errX = regexp.Compile ("^[A-Z][A-Za-z0-9_]*$")
	if errX != nil {
		buggyPackage = true
	}
}

func New (element []interface{}) (*Squaket, error) {
	if buggyPackage == true {
		return nil, errors.New ("Package is buggy.")
	}

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
	if buggyPackage == true {
		return nil, errors.New ("Package is buggy.")
	}

	if propertyPattern.MatchString (property) == false {
		return nil, errors.New ("Invalid property name.")
	}

	g = map[interface {}][]interface {} {}

	for _, element := range s.element {
		value := reflect.ValueOf (element).FieldByName (property)
		if value.IsValid () == false {
			return nil, errors.New ("An element does not have that property.")
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

var (
	propertyPattern *regexp.Regexp
	buggyPackage bool = false
)
