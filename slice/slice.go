package slice

import "errors"

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 |
	~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~uintptr | ~float32 | ~float64 | ~string
}

var IndexIllegalError = errors.New("Index is illegal")

// Delete the indexed element.
// Return the new slice and error.
func Delete[T ordered](src []T, idx int) ([]T, error) {
	lens := len(src)
	if idx < 0 || idx >= lens {
		return nil, IndexIllegalError
	}
	copy(src[idx:], src[idx+1:])
	src = src[:lens-1]
	src = shrinkSlice(src)
	return src, nil
}

// If the length of slice is smaller than capacity/3, shrink its capacity to half.
func shrinkSlice[T ordered](src []T) []T {
	lens := len(src)
	caps := cap(src)
	if lens >= caps/3 {
		return src
	}
	newSrc := make([]T, caps/2)
	copy(newSrc, src)
	return newSrc
}
