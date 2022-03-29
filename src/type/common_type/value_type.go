package common_type

type ValueType uint8 // 255 possible values

const (
	OBJECT   ValueType = 0
	BOOLEAN  ValueType = 1
	LONG     ValueType = 2
	DOUBLE   ValueType = 3
	STRING   ValueType = 4
	DATETIME ValueType = 5
)
