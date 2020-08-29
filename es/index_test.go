package es

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapping(t *testing.T) {
	pt := NewBaseIndexProperties()
	pt.SetProperties("name", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("des", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("content", &PropertiesType{
		Type: "text",
	})
	mapping := NewMapping(pt)
	b, err := mapping.MappingToJsonString()
	fmt.Println(string(b))
	assert.Equal(t, err, nil)
}
