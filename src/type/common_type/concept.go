package common_type

// TODO: Add Thing type
type Concept[T Value] struct {
	Concept T
}

type ConceptMap[V Value] struct {
	Map          map[string]Concept[V] `json:"map,omitempty"`
	Explainables Explainables          `json:"explainables,omitempty"`
}

type ConceptMapGroup[V Value] struct {
	Owner       Concept[V]      `json:"owner,omitempty"`
	ConceptMaps []ConceptMap[V] `json:"concept_maps,omitempty"`
}
