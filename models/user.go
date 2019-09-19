package models

import (
	"net/url"
	"time"
)

const (
	GenderMale       = "male"
	GenderFemale     = "female"
	GenderUnAnswered = "un_answered"
)

const (
	PhoneTypeHome   = "home"
	PhoneTypeOffice = "office"
	PhoneTypeMobile = "mobile"
	PhoneTypeFax    = "fax"
)

const (
	SnsTwitter  = "twitter"
	SnsGithub   = "github"
	SnsFacebook = "facebook"
)

type User struct {
	Base         `yaml:",inline"`
	FirstName    string    `json:"first_name"yaml:"first_name"`
	LastName     string    `json:"last_name"yaml:"last_name"`
	Gender       string    `json:"gender"yaml:"gender"`
	Birth        time.Time `json:"birth"yaml:"birth"`
	Address      *Address  `json:"address"yaml:"address"`
	PhoneNumbers []*Phone  `json:"phone_numbers"yaml:"phone_numbers"`
	SNSs         []*Sns    `json:"snss"yaml:"snss"`
}

type Address struct {
	Base     `yaml:",inline"`
	Country  string `json:"country"yaml:"country"`
	City     string `json:"city"yaml:"city"`
	State    string `json:"state"yaml:"state"`
	ZipCode  string `json:"zip_code"yaml:"zip_code"`
	Address1 string `json:"address_1"yaml:"address_1"`
	Address2 string `json:"address_2"yaml:"address_2"`
}

type Phone struct {
	Base        `yaml:",inline"`
	PhoneNumber string `json:"phone_number"yaml:"phone_number"`
	Type        string `json:"type"yaml:"type"`
}

type Sns struct {
	Base    `yaml:",inline"`
	Type    string   `json:"type"yaml:"type"`
	Account string   `json:"account"yaml:"account"`
	URL     *url.URL `json:"url"yaml:"url"`
}
