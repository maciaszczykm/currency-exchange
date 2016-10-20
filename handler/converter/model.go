package converter

import (
	"encoding/xml"
	"fmt"
)

// converterResponse structure containing converted rates, requested amount and currency.
type converterResponse struct {
	XMLName   xml.Name     `json:"-" xml:"converterResponse"`
	Amount    float64      `json:"amount" xml:"amount"`
	Currency  string       `json:"currency" xml:"currency"`
	Converted convertedMap `json:"converted" xml:"converted"`
}

// convertedMap contains string and float64 pairs.
type convertedMap map[string]float64

// MarshalXML marshals convertedMap into XML.
func (s convertedMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(fmt.Sprintf("%v", value)), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}
