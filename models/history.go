package models

import "time"

const (
	EducationEnter    = "enter"
	EducationGraduate = "graduate"
	EducationDropout  = "dropout"
	EducationReturn   = "return"
)

const (
	WorkJoin  = "join"
	WorkLeave = "leave"
)

type History struct {
	Base  `yaml:",inline"`
	GotAt time.Time `json:"got_at"yaml:"got_at"`
}

type EducationHistory struct {
	History `yaml:",inline"`
	Type    string `json:"type"yaml:"type"`
}

type WorkHistory struct {
	History `yaml:",inline"`
	Type    string `json:"type"yaml:"type"`
}

type Company struct {
	Base `yaml:",inline"`
}
