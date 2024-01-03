package models

import "github.com/a-h/templ"

type InputProps struct {
	Type         string
	Name         string
	Attributes   templ.Attributes
	MainAttributes   templ.Attributes
	IsRequired   bool
	PlaceHolder  string
	StatsMessage string
    StatsMessageColor string
}
