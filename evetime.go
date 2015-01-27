package eveapi

import (
	"encoding/xml"
	"time"
)

//Source: https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields
type eveTime struct {
	time.Time
}

func (c eveTime) String() string {
	return c.Format(dateFormat)
}
func (c *eveTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if len(attr.Value) == 0 {
		*c = eveTime{time.Unix(0, 0)}
		return nil
	}

	parse, err := time.Parse(dateFormat, attr.Value)
	if err != nil {
		return err
	}
	*c = eveTime{parse}
	return nil
}
func (c *eveTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

	if len(v) == 0 {
		*c = eveTime{time.Unix(0, 0)}
		return nil
	}

	parse, err := time.Parse(dateFormat, v)
	if err != nil {
		return nil
	}
	*c = eveTime{parse}
	return nil
}
