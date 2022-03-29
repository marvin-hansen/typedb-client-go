package common_type

type Explainable struct {
	Conjunction string `json:"conjunction,omitempty"`
	ID          int64  `json:"id,omitempty"`
}

type Owned struct {
	Owned map[string]Explainable `json:"owned,omitempty"`
}

type Explainables struct {
	Relations  map[string]Explainable `json:"relations,omitempty"`
	Attributes map[string]Explainable `json:"attributes,omitempty"`
	Ownerships map[string]Owned       `json:"ownerships,omitempty"`
}
