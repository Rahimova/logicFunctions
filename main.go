package main

func Index(slice []int, predicate func(int) bool) int {
	for i, elem := range slice {
		if predicate(elem) {
			return i
		}
	}
	return -1
}

func All(slice []int, predicate func(int) bool) bool {
	return Index(slice, func(i int) bool {
		return !predicate(i)
	}) == -1
}

func Any(slice []int, predicate func(int) bool) bool {
	return Index(slice, predicate) != -1
}

func None(slice []int, predicate func(int) bool) bool {
	return Index(slice, predicate) == -1
}

func Find(slice []int, predicate func(int) bool) int {
	return slice[Index(slice, predicate)]
}

// trivial realization (faster)-----------------------------------------
func All0(slice []int, predicate func(int) bool) bool {
	for _, elem := range slice {
		if !predicate(elem) {
			return false
		}
	}
	return true
}

func Any0(slice []int, predicate func(int) bool) bool {
	for _, elem := range slice {
		if predicate(elem) {
			return true
		}
	}
	return false
}

func None0(slice []int, predicate func(int) bool) bool {
	for _, elem := range slice {
		if predicate(elem) {
			return false
		}
	}
	return true
}

func Find0(slice []int, predicate func(int) bool) int {
	for _, elem := range slice {
		if predicate(elem) {
			return elem
		}
	}
	return -1
}

func main() {

}