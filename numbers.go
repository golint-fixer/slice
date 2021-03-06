// +build generate
//go:generate genny -in=$GOFILE -out=numbers_genny.go gen "NumberType=NUMBERS"

package slice

import "github.com/cheekybits/genny/generic"

type NumberType generic.Number

// A PredicateNumberTypeFunc represents a function that defines a criteria and
// determines whether specified NumberType meets that criteria.
type PredicateNumberTypeFunc func(NumberType) bool

// A NNumberType represents a slice of NumberType.
type NNumberType []NumberType

// Average returns the average value from all elements of current slice.
func (n NNumberType) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NNumberType) Equal(other []NumberType) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NNumberType) Except(a ...NumberType) NNumberType {
	var result NNumberType
	aBox := NNumberType(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified NumberType exists into current slice.
func (n NNumberType) Exists(a NumberType) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified NumberTypes exists into
// current slice.
func (n NNumberType) ExistsAll(a ...NumberType) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified NumberTypes exists into current
// slice.
func (n NNumberType) ExistsAny(a ...NumberType) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified NumberType into current slice.
func (n NNumberType) IndexOf(a NumberType) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NNumberType) Sum() NumberType {
	var sum NumberType
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NNumberType) TrueForAll(pred PredicateNumberTypeFunc) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NNumberType) TrueForAny(pred PredicateNumberTypeFunc) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NNumberType) Where(pred PredicateNumberTypeFunc) NNumberType {
	var result NNumberType
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}
