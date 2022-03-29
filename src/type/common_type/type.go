package common_type

type Type struct {
	Label     string    `json:"label"`
	Scope     string    `json:"Scope,omitempty"`
	Encoding  Encoding  `json:"encoding,omitempty"`
	ValueType ValueType `json:"value_type,omitempty"`
	Root      bool      `json:"root,omitempty"`
}
