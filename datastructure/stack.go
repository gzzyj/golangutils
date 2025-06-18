package datastructure 


type Stack[T any] []T

func (r *Stack[T]) IsEmpty() bool {
	return len(*r) == 0
}

func (r *Stack[T]) Push(item T) {
	*r = append(*r, item)
}

func (r *Stack[T]) Pop() T {
	n := len(*r)
	var zero T
	if n < 0 {
		return zero
	}
	x := (*r)[n-1]
	*r = (*r)[0 : n-1]
	return x
}
