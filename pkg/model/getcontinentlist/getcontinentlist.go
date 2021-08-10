package getcontinentlist

import "encoding/xml"

type GetContinentListResponseXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                           string `xml:",chardata"`
		ListOfContinentsByNameResponse struct {
			Text                         string `xml:",chardata"`
			M                            string `xml:"m,attr"`
			ListOfContinentsByNameResult struct {
				Text       string `xml:",chardata"`
				TContinent []struct {
					Text  string `xml:",chardata"`
					SCode string `xml:"sCode"`
					SName string `xml:"sName"`
				} `xml:"tContinent"`
			} `xml:"ListOfContinentsByNameResult"`
		} `xml:"ListOfContinentsByNameResponse"`
	} `xml:"Body"`
}

type GetContinentListResponseJSON struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetContinentListResponse struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    []GetContinentListResponseJSON `json:"data"`
}
