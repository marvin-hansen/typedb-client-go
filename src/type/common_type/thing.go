package common_type

type Thing[T Value] struct {
	IID      []byte
	Type     Type
	Value    ValueItem[T]
	Inferred bool
}
