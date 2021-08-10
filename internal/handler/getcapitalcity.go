package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"simple-esb/pkg/model/getcapitalcity"
	"strings"
)

func GetCapitalCity(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /capitalcity")

	countryID := r.URL.Query().Get("country_id")

	url := fmt.Sprintf("%s",
		"http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso")

	// request payload
	rawPayload := fmt.Sprintf(`
	<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<CapitalCity xmlns="http://www.oorsprong.org/websamples.countryinfo">
				<sCountryISOCode>%s</sCountryISOCode>
			</CapitalCity>
		</soap:Body>
	</soap:Envelope>
	`, countryID)
	payload := []byte(strings.TrimSpace(rawPayload))

	httpMethod := "POST"

	log.Println("-> Preparing the request")

	// prepare the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}

	// set the content type header, as well as the oter required headers
	req.Header.Set("Content-type", "text/xml")

	// prepare the client request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(getcapitalcity.GetCapitalCityResponseXML)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	log.Println("-> Everything is good, printing data")

	// print the users data
	capitalCity := result.Body.CapitalCityResponse.CapitalCityResult

	// var continentList []getcontinentlist.GetContinentListResponseJSON
	var capitalCityResult getcapitalcity.GetCapitalCityResponseJSON
	var response getcapitalcity.GetCapitalCityResponse

	capitalCityResult.CapitalCity = capitalCity

	// // set response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = capitalCityResult

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
