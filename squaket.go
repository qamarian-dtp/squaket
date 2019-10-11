package squaket

import (
	"errors"
	"reflect"
	"regexp"
)

// New () creates a new squaket. All elements must be a struct pointer, otherwise, a nil
// squaket and an error would be returned.
func New (element ... interface{}) (output *Squaket, err error) {
	if buggyPackage == true {
		return nil, errors.New ("Package is buggy.")
	}

	for _, someElement := range element {
		pointer := reflect.ValueOf (someElement)
		if pointer.Kind () != reflect.Ptr {
			return nil, errors.New ("An element is not a pointer.")
		}
		if pointer.Elem ().Kind () != reflect.Struct {
			return nil, errors.New ("A pointer is not a struct pointer.")
		}
	}
	output = &Squaket {element}

	return output, nil
}

type Squaket struct {
	element []interface {}
}

// Group () groups a squaket. If not all elements have the grouping property, a nil
// grouping and an error would be returned.
func (s *Squaket) Group (property string) (g map[interface {}][]interface {}, e error) {
	if buggyPackage == true {
		return nil, errors.New ("Package is buggy.")
	}

	if propertyPattern.MatchString (property) == false {
		return nil, errors.New ("Invalid property name.")
	}

	g = map[interface {}][]interface {} {}

	for _, element := range s.element {
		value := reflect.ValueOf (element).Elem ().FieldByName (property)
		if value.IsValid () == false {
			return nil, errors.New ("An element does not have that " +
				"property.")
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

func init () {
	var errX error

	propertyPattern, errX = regexp.Compile ("^[A-Z][A-Za-z0-9_]*$")
	if errX != nil {
		buggyPackage = true
	}
}
