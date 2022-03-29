package common_type

type Numeric[T NumericType] struct {
	Value T `json:"value,omitempty"`
}

type NumericType interface {
	int64 | float64 | bool
}

type NumericGroup[T NumericType, V Value] struct {
	Owner   Concept[V] `json:"owner,omitempty"`
	Numeric Numeric[T] `json:"numeric,omitempty"`
}
