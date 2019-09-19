package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"resume/models"
	"time"

	"gopkg.in/yaml.v2"
)

func Marshal() {
	now := time.Now()
	base := models.Base{
		ID:          "__id",
		Name:        "__name",
		Description: "__description",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	u, _ := url.Parse("https://test.test")
	user := &models.User{
		Base:      base,
		FirstName: "__first_name",
		LastName:  "__last_name",
		Gender:    "__gender",
		Birth:     now,
		Address: &models.Address{
			Country:  "__country",
			City:     "__city",
			State:    "__state",
			ZipCode:  "__zip_code",
			Address1: "__address1",
			Address2: "__address2",
		},
		PhoneNumbers: []*models.Phone{
			{
				Base:        base,
				PhoneNumber: "__phone_number",
				Type:        "__type",
			},
		},
		SNSs: []*models.Sns{
			{
				Base:    base,
				Type:    "__type",
				Account: "__account",
				URL:     u,
			},
		},
	}
	history := models.History{
		Base:  base,
		GotAt: now,
	}
	resume := &models.Resume{
		Base: base,
		User: user,
		EducationHistories: []*models.EducationHistory{
			{
				History: history,
				Type:    "__type",
			},
		},
		WorkHistories: []*models.WorkHistory{
			{
				History: history,
				Type:    "__type",
			},
		},
		ReasonsForApplying: "__reasons_for_applying",
		Licences: []*models.Licence{
			{
				Base:  base,
				GotAt: now,
			},
		},
		Extends: []*models.Extend{
			{
				Base: base,
			},
		},
	}

	b, err := yaml.Marshal(resume)
	if err != nil {
		log.Fatalln(fmt.Errorf("b, err := yaml.Marshal(resume): %w", err))
	}
	fmt.Println(string(b))
}

func UnMarshal() {
	b, err := ioutil.ReadFile("samples/sample.yml")
	if err != nil {
		log.Fatalln(fmt.Errorf("b, err := ioutil.ReadFile(\"./samples/samples.yml\") : %w", err))
	}
	resume := &models.Resume{}
	err = yaml.Unmarshal(b, resume)
	if err != nil {
		log.Fatalln(fmt.Errorf("err = yaml.Unmarshal(b, resume) : %w", err))
	}
	log.Println(resume)
}

func main() {
	Marshal()
	UnMarshal()
}
