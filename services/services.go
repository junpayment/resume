package services

import (
	"resume/repositories"
)

type Template interface {
	GetText(yaml []byte) ([]byte, error)
	GetPdf(yaml []byte) ([]byte, error)
}

type TemplateImpl struct {
	ResumeRepository repositories.ResumeRepository
}

func (t *TemplateImpl) GetText(b []byte) (string, error) {
	t.ResumeRepository.Get()
}

func (t *TemplateImpl) GetPdf(b []byte) (string, error) {
	t.ResumeRepository.Get()
}
