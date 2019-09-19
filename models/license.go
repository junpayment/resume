package models

import "time"

type Licence struct {
	Base  `yaml:",inline"`
	GotAt time.Time `json:"got_at"yaml:"got_at"`
}
