package common_type

type Thing[T Value] struct {
	IID      []byte       `json:"Iid,omitempty"`
	Type     Type         `json:"type,omitempty"`
	Value    ValueItem[T] `json:"value,omitempty"`
	Inferred bool         `json:"inferred,omitempty"`
}
