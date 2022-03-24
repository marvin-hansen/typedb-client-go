package v2

import "github.com/marvin-hansen/typedb-client-go/common"

// Proto utils
// 728

func protoBoolAttributeValue(value bool) *common.Attribute_Value_Boolean {
	return &common.Attribute_Value_Boolean{Boolean: value}
}

func protoLongAttributeValue(value int) *common.Attribute_Value_Long {
	return &common.Attribute_Value_Long{Long: int64(value)}
}

func protoDoubleAttributeValue(value float64) *common.Attribute_Value_Double {
	return &common.Attribute_Value_Double{Double: value}
}

func protoStringAttributeValue(value string) *common.Attribute_Value_String_ {
	return &common.Attribute_Value_String_{String_: value}
}

func protoDatetimeAttributeValue(value int64) *common.Attribute_Value_DateTime {
	return &common.Attribute_Value_DateTime{DateTime: value}
}
