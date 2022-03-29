// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package _type

import "fmt"

func NewLabel(name, scope string) *Label {
	return &Label{
		name:  &name,
		scope: &scope,
	}
}

type Label struct {
	name  *string
	scope *string
}

func (s Label) HasName() bool {
	if s.name != nil {
		return true
	} else {
		return false
	}
}

func (s Label) HasScope() bool {
	if s.scope != nil {
		return true
	} else {
		return false
	}
}

func (s Label) GetName() string {
	return *s.name
}

func (s Label) GetScope() string {
	return *s.scope
}

func (s Label) String() string {
	return fmt.Sprintf("Name: %v, Scope: %v",
		*s.name,
		*s.scope,
	)
}
