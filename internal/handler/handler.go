package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"simple-esb/pkg/model/getcontinentlist"
	"strings"
)

func GetContinentList(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /continents")

	url := fmt.Sprintf("%s",
		"http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso")

	// request payload
	payload := []byte(strings.TrimSpace(`
		<?xml version="1.0" encoding="utf-8"?>
		<soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
			<soap12:Body>
				<ListOfContinentsByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
				</ListOfContinentsByName>
			</soap12:Body>
		</soap12:Envelope>`,
	))

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
	result := new(getcontinentlist.GetContinentListResponseXML)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	log.Println("-> Everything is good, printing data")

	// print the users data
	continents := result.Body.ListOfContinentsByNameResponse.ListOfContinentsByNameResult.TContinent

	var continentList []getcontinentlist.GetContinentListResponseJSON
	var continentItem getcontinentlist.GetContinentListResponseJSON
	var response getcontinentlist.GetContinentListResponse

	for _, value := range continents {
		continentItem.Name = value.SName
		continentItem.Code = value.SCode
		continentList = append(continentList, continentItem)
	}

	// set response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = continentList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
