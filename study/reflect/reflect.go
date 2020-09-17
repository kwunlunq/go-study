package main

import (
	"fmt"
	"reflect"
)

var (
	b = Person{
		Age:       23,
		Weight:    65,
		Height:    169,
		MM:        55,
		Interests: []string{"swim", "dance"},
	}
)

type Person struct {
	Age       int
	Weight    int
	Height    int
	MM        int
	Interests []string
}

type Being interface {
	Love(b Being)
}

func (*Person) Love(b Being) {

}

func main() {
	var i interface{}
	fmt.Printf("%v\n", reflect.ValueOf(i))
	i = "3"
	fmt.Println(reflect.TypeOf(i))
	//reflectType()
	//reflectValue()

	//var list []Person
	//setSlice(&list)
	//fmt.Println("Result slice: ", list)

	var sourceList []int
	err := Set(&sourceList, []int{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Slice: ", sourceList)

	var b2 Person
	Set(&b2, b)
	fmt.Println("b2: ", b2)

	//sampleCheckType()
}

func Set(item, setValue interface{}) (err error) {
	fmt.Printf("item: %v, setValue: %v\n", reflect.TypeOf(item), reflect.TypeOf(setValue))
	t, v := reflect.TypeOf(item), reflect.ValueOf(item)
	if !v.IsValid() || t.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		err = fmt.Errorf("err not a settable type: %v", t)
		return
	}
	v.Elem().Set(reflect.ValueOf(setValue))
	return
}

func setSlice(iList interface{}) {
	sourceList := []Person{{Age: 29}, {Age: 30}}

	t, v := reflect.TypeOf(iList), reflect.ValueOf(iList)
	fmt.Println(v.Elem().CanSet())
	v.Elem().Set(reflect.ValueOf(sourceList))
	_ = t
}

func reflectValue() {
	i := 10
	v := reflect.ValueOf(i)
	fmt.Printf("\n ===  value  ===\nint:\nInt:%v\n", v.Int())

	v = reflect.ValueOf(b)
	fmt.Printf("\nstruct:\nage: %d, weight: %d\n", v.Field(0).Int(), v.Field(1).Int())

	v = reflect.ValueOf(&b)
	fmt.Printf("\nptr:\nage: %d\n", v.Elem().Field(0).Int())

	v = reflect.ValueOf(&b).Elem()
	v.Field(0).SetInt(99)
	fmt.Printf("\nset value:\nnew age = %v\n", v.Field(0))
}

func reflectType() {

	fmt.Println("\n=== type ===")

	// struct
	t := reflect.TypeOf(b)
	fmt.Printf("struct:\nType: %s, Kind: %s, Name: %s, String: %s, Size: %d, Field(0): %v\n",
		t, t.Kind(), t.Name(), t.String(), t.Size(), t.Field(0))

	// ptr
	t = reflect.TypeOf(&b)
	fmt.Printf("\nptr:\nType: %s, Kind: %s, Name: %s, String: %s, Size: %d, Elem: %s, Method(0): %v\n",
		t, t.Kind(), t.Name(), t.String(), t.Size(), t.Elem(), t.Method(0))
	t = reflect.TypeOf(t.Elem().Field(0))

	// interface
	var being Being = &b
	t = reflect.TypeOf(being)
	fmt.Printf("\ninterface:\nType: %s, Kind: %s, Name: %s, String: %s, Size: %d, Elem: %s, Method(0): %v\n",
		t, t.Kind(), t.Name(), t.String(), t.Size(), t.Elem(), t.Method(0))

	// slice
	bs := []Person{b}
	t = reflect.TypeOf(bs)
	fmt.Printf("\nslice: \nKind: %v, String: %s, Size: %v, Elem: %v\n", t.Kind(), t.String(), t.Size(), t.Elem())

	// array
	bAry := [...]Person{b}
	t = reflect.TypeOf(bAry)
	fmt.Printf("\narray: \nKind: %v, String: %s, Size: %v, Len: %v, Elem: %v\n", t.Kind(), t.String(), t.Size(), t.Len(), t.Elem())

}

func sampleCheckType() {
	values := [...]interface{}{
		Person{Age: 10},
		Person{Age: 20},
		[...]int{1, 2, 3, 4, 5},
		map[string]int{"caterpillar": 123456, "monica": 54321},
		10,
	}

	for _, value := range values {
		switch t := reflect.TypeOf(value); t.Kind() {
		case reflect.Struct:
			fmt.Println("it's a struct.")
		case reflect.Array:
			fmt.Println("it's a array.")
		case reflect.Map:
			fmt.Println("it's a map.")
		case reflect.Int:
			fmt.Println("it's a integer.")
		default:
			fmt.Println("非預期之型態")
		}
	}
}
