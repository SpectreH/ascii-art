package slices

import "ascii-art/structures"

func RemoveSliceElement(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

func Insert(a []structures.Banner, index int, value structures.Banner) []structures.Banner {
	if len(a) == index {
		return append(a, value)
	}

	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
