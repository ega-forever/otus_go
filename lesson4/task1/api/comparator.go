package api

type Comparator func(a interface{}, b interface{}) bool

type Container []interface{}

func Max(comparator Comparator, elems Container) interface{} {

	elemsCopy := elems[:]

	Sort(comparator, elemsCopy)

	return elemsCopy[0]
}
