package common_type

type Answer[T Value] struct {
	Answers ConceptMap[T] `json:"answers,omitempty"`
}
