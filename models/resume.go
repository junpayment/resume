package models

type Resume struct {
	Base               `yaml:",inline"`
	User               *User               `json:"user"yaml:"user"`
	EducationHistories []*EducationHistory `json:"education_history"yaml:"education_history"`
	WorkHistories      []*WorkHistory      `json:"work_history"yaml:"work_history"`
	ReasonsForApplying string              `json:"reasons_for_applying"yaml:"reasons_for_applying"`
	Licences           []*Licence          `json:"licences"yaml:"licences"`
	Extends            []*Extend           `json:"extends"yaml:"extends"`
}

type Extend struct {
	Base `yaml:",inline"`
}
