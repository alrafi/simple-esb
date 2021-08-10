package getcapitalcity

import "encoding/xml"

type GetCapitalCityResponseXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                string `xml:",chardata"`
		CapitalCityResponse struct {
			Text              string `xml:",chardata"`
			M                 string `xml:"m,attr"`
			CapitalCityResult string `xml:"CapitalCityResult"`
		} `xml:"CapitalCityResponse"`
	} `xml:"Body"`
}

type GetCapitalCityResponseJSON struct {
	CapitalCity string `json:"capital_city"`
}

type GetCapitalCityResponse struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    GetCapitalCityResponseJSON `json:"data"`
}
