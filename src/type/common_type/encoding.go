package common_type

type Encoding uint8 // 255 possible values

const (
	THING_TYPE     Encoding = 0
	ENTITY_TYPE    Encoding = 1
	RELATION_TYPE  Encoding = 2
	ATTRIBUTE_TYPE Encoding = 3
	ROLE_TYPE      Encoding = 4
)
