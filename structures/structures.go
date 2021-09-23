package structures

type Color struct {
	Name string
	Code []byte
}

type Banner struct {
	Id          int
	AsciiSymbol [8][]rune
	Color       Color
}

type Flag struct {
	Class string
	Value string
}
