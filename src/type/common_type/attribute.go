package common_type

type Attribute[T Value] struct {
	Value ValueItem[T]
}

type ValueItem[T Value] struct {
	Item T
}

func (i ValueItem[T]) GetValue() (item T) {
	return i.Item
}

type Value interface {
	StringValue | BooleanValue | LongValue | DoubleValue | DateTimeValue
}

type StringValue struct {
	Value string
}

type BooleanValue struct {
	Value bool
}

type LongValue struct {
	Value int64
}

type DoubleValue struct {
	Value float64
}

type DateTimeValue struct {
	Value int64
}
