package api

import "math/rand"

func Sort(comparator Comparator, elems Container) Container {

	if len(elems) < 2 {
		return elems
	}

	left, right := 0, len(elems)-1
	pivot := rand.Int() % len(elems)

	elems[pivot], elems[right] = elems[right], elems[pivot]

	for i := range elems {
		if comparator(elems[i], elems[right]) {
			elems[left], elems[i] = elems[i], elems[left]
			left++
		}

	}

	elems[left], elems[right] = elems[right], elems[left]

	Sort(comparator, elems[:left])
	Sort(comparator, elems[left+1:])

	return elems
}
