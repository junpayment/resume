package repositories

import (
	"fmt"
	"io"
	"io/ioutil"
	"resume/models"

	"gopkg.in/yaml.v2"
)

type ResumeRepository interface {
	Get(io.Reader) (*models.Resume, error)
}

type ResumeRepositoryImpl struct {
}

func (rep *ResumeRepositoryImpl) Get(r io.Reader) (*models.Resume, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("b, err := ioutil.ReadAll(r): %w", err)
	}
	resume := &models.Resume{}
	err = yaml.Unmarshal(b, resume)
	if err != nil {
		return nil, fmt.Errorf("err = yaml.Unmarshal(b, resume): %w", err)
	}
	return resume, nil
}
