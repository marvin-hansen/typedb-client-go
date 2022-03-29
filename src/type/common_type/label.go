// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package common_type

import "fmt"

func NewLabel(name, scope string) *Label {
	return &Label{
		Name:  &name,
		Scope: &scope,
	}
}

type Label struct {
	Name  *string `json:"Name,omitempty"`
	Scope *string `json:"Scope,omitempty"`
}

func (s Label) HasName() bool {
	if s.Name != nil {
		return true
	} else {
		return false
	}
}

func (s Label) HasScope() bool {
	if s.Scope != nil {
		return true
	} else {
		return false
	}
}

func (s Label) GetName() string {
	return *s.Name
}

func (s Label) GetScope() string {
	return *s.Scope
}

func (s Label) String() string {
	return fmt.Sprintf("Name: %v, Scope: %v",
		*s.Name,
		*s.Scope,
	)
}
