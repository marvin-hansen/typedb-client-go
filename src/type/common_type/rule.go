package common_type

type Rule struct {
	Label string `json:"label,omitempty"`
	When  string `json:"when,omitempty"`
	Then  string `json:"then,omitempty"`
}

type Explanation struct {
	Rule       Rule       `json:"rule,omitempty"`
	VarMapping VarList    `json:"var_mapping,omitempty"`
	Condition  ConceptMap `json:"condition,omitempty"`
	Conclusion ConceptMap `json:"conclusion,omitempty"`
}

type VarList struct {
	Vars []string `json:"vars,omitempty"`
}
