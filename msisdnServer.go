package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/liuzl/phonenumbers"
)

const indexBody = `
<html>
  <head>
	<title>PhoneNumberServer</title>
	<script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
  </head>
  <style>
  #results div { display: inline-block; border: solid 3px green; padding: 5px; }
  #results div.error { border: solid 3px red; }
  </style>
  <body>
	<form>
	    <label for="phoneNumber">Enter Phone Number: </label>
		<input id="phoneNumber" type="text" name="phoneNumber" value="+38640579602" />
		<input type="submit" value="Parse" class="button"/>
	</form>
	<div id="results">
	</div>
  </body>
  <script>
    $("form").submit(function(e){
		event.preventDefault();
		$.ajax({
			"url": "/parse?" + $("form").serialize(), 
			"success": function(data, status, xhr){
				$("#results").prepend("<pre>" + JSON.stringify(data, null, 4) + "</pre>");
				$("#results").prepend("<div> Phone number OK:  " + $("#phoneNumber").val() + "</div>");
			},
			"error": function(request, status, error){
				$("#results").prepend("<pre class='error'>" + JSON.stringify(JSON.parse(request.responseText), null, 4) + "</pre>");
				$("#results").prepend("<div class='error'> ERROR: " + $("#phoneNumber").val() + "</div>");
			}
		});
	})
  </script>
</html>
`

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type successResponse struct {
	NationalNumber    uint64 `json:"national_number"`
	CountryCode       int32  `json:"country_code"`
	NationalFormatted string `json:"national_formatted"`
	CarrierForNumber  string `json:"carrier_for_number"`
	CountryCodeName   string `json:"country_code_name"`
}

func writeResponse(w http.ResponseWriter, status int, body interface{}) {
	js, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func parse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phoneNumber := r.Form.Get("phoneNumber")

	log.Printf("Phone number to be parsed %s", phoneNumber)

	status, interfaceResponse := getResponse(phoneNumber)
	writeResponse(w, status, interfaceResponse)
}

func getResponse(phoneNumber string) (int, interface{}) {

	if phoneNumber == "" {
		return http.StatusBadRequest, errorResponse{"missing phoneNumber", "missing 'phoneNumber' parameter"}
	}

	phoneNumberMetadata, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return http.StatusBadRequest, errorResponse{"error parsing phone number", err.Error()}
	}

	carrier, err := phonenumbers.GetCarrierForNumber(phoneNumberMetadata, phonenumbers.GetRegionCodeForNumber(phoneNumberMetadata))
	if err != nil {
		return http.StatusBadRequest, errorResponse{"error parsing carrier from phone number", err.Error()}
	}

	return http.StatusOK, successResponse{
		NationalNumber:    *phoneNumberMetadata.NationalNumber,
		CountryCode:       *phoneNumberMetadata.CountryCode,
		NationalFormatted: phonenumbers.Format(phoneNumberMetadata, phonenumbers.NATIONAL),
		CarrierForNumber:  carrier,
		CountryCodeName:   phonenumbers.GetRegionCodeForNumber(phoneNumberMetadata),
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexBody))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/parse", parse)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
