package es

import "encoding/json"

//var mapping = `{
//  "mappings": {
//    "properties": {
//        "title": {
//          "type": "text"
//        },
//        "brand_name": {
//          "type": "text"
//        }
//      }
//  }
//}`

type Mapping struct {
	Mappings Mappings `json:"mappings"`
}

func (m *Mapping) MappingToJsonString() ([]byte, error) {
	return json.Marshal(m)
}

type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties interface {
	SetProperties(title string, pt *PropertiesType)
}

type PropertiesType struct {
	Type string `json:"type"`
}

type baseIndexProperties map[string]*PropertiesType

func (p baseIndexProperties) SetProperties(title string, pt *PropertiesType) {
	p[title] = pt
}

func NewBaseIndexProperties() *baseIndexProperties {
	ip := baseIndexProperties{}
	ip = make(map[string]*PropertiesType)
	return &ip
}

func NewMapping(pt Properties) *Mapping {
	mapping := new(Mapping)
	mapping.Mappings.Properties = pt
	return mapping
}
