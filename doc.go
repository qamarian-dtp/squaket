// Package squaket implements the squaket data type. A squaket is simply a collection of
// composite data you could group based on a specific property.
//
//
// Take for example:
//
//	data1: name="a", age=2, location="q"
//	data2: name="b", age=2, location="m"
//	data3: name="c", age=5, location="m"
//	data4: name="d", age=3, location="d"
//	data5: name="e", age=1, location="m"
//	data6: name="f", age=2, location="q"
//
// The data you see are composite data having three properties, namely: name, age, and
// location. If we are to group them based on the similarity of their locations, we could
// have the following groups:
//
//	group1: data1, data6
//	group2: data2, data3, data5
//	group3: data4
//
// This kind of collection of data is what we mean by squaket. In practice, elements of a
// squaket are expected to structs.
//
//
// However, the elements of a squaket could have different properties, as long as they all
// have the grouping property.
//
// Real world example:
//
//	type struct human {
//		Name string
//		Location string
//	}
//
//	type struct company
//		FounderName string
//		ceo string
//		Location string
//	}
//
//	element0 := human {"Qamardeen", "Lagos"}
//	element1 := company {"Apple", "CA"}
//	element2 := human {"Ibrahim", "CA"}
//	element3 := company {"Twitter", "Lagos"}
//
//	squaketX, err := squaket.New ([]interface {element0, element1, element2, element3})
//	groups, err2 := squaketX.Group ("Location")
//	fmt.Println (groups ["Lagos"], groups ["CA"]) // [element0, element3], [element1, element2]
//
//
// Things to note:
//
//	1. The name of a property must be a valid Golang identifier.
//	2. The name of a property must start with an uppercase letter, like "Location" in
// 		the example above.
//
package squaket
