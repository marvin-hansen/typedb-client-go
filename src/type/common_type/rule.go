package common_type

type Rule struct {
	Label string `json:"label,omitempty"`
	When  string `json:"when,omitempty"`
	Then  string `json:"then,omitempty"`
}

type Explanation[T Value] struct {
	Rule       Rule          `json:"rule,omitempty"`
	VarMapping VarList       `json:"var_mapping,omitempty"`
	Condition  ConceptMap[T] `json:"condition,omitempty"`
	Conclusion ConceptMap[T] `json:"conclusion,omitempty"`
}

type VarList struct {
	Vars []string `json:"vars,omitempty"`
}
