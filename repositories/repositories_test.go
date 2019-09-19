package repositories

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResumeRepositoryImpl_Get(t *testing.T) {
	t.Run("load yaml", func(t *testing.T) {
		b, _ := ioutil.ReadFile("samples/sample.yml")
		r := bytes.NewReader(b)
		rep := &ResumeRepositoryImpl{}
		resume, err := rep.Get(r)
		assert.Nil(t, err)
		assert.NotNil(t, resume)
	})
}
