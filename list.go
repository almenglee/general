package general

import (
	"sort"
)

type List[T any] []T

func Join[T any](a List[T], b List[T]) *List[T] {
	a = append([]T(a), b...)
	return &a
}

func (l List[T]) Len() int {
	return len([]T(l))
}

func (l List[T]) Slice() []T {
	return []T(l)
}

func AsList[T any](l []T) *List[T] {
	return NewList(l...)
}

func EmptyList[T any]() *List[T] {
	return &List[T]{}
}

func NewList[T any](e ...T) *List[T] {
	r := List[T](e)
	return &r
}

func (l *List[T]) Each(f func(T)) {
	for _, t := range *l {
		f(t)
	}
}

func (l *List[T]) Sort(cmp func(T, T) int) {
	_list := []T(*l)
	sort.Slice(_list, func(i, j int) bool {
		return cmp(_list[i], _list[j]) < 0
	})
	*l = _list

}

func (l *List[T]) SortReverse(cmp func(T, T) int) {
	_cmp := func(a, b T) int {
		return cmp(b, a)
	}
	l.Sort(_cmp)
}

func (l *List[T]) Reverse() {
	for i, j := 0, l.Len()-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}

func (l *List[T]) Iter(f func(int, T)) {
	for i, t := range *l {
		f(i, t)
	}
}

func (l *List[T]) Append(e T) {
	*l = List[T](append([]T(*l), e))
}

func (l *List[T]) First() *T {
	if l.Len() == 0 {
		return nil
	}
	return &(*l)[0]
}

func (l *List[T]) Filter(cond func(int, T) bool) *List[T] {
	//var rtn []T
	rtn := new(List[T])
	for i, v := range *l {
		if cond(i, v) {
			rtn.Append(v)
		}
	}
	return rtn
}

func (l *List[T]) Take(n int) *List[T] {
	rtn := new(List[T])
	if len([]T(*l)) < n {
		panic("index out of bound")
	}
	for i := 0; i < n; i++ {
		rtn.Append((*l)[i])
	}
	return rtn
}
