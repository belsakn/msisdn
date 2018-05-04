package main

import (
	"log"
	"net/http"
)

const indexBody = `
<html>
  <head>
	<title>PhoneNumberServer</title>
	<script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
  </head>
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
				$("#results").prepend("<div>" + $("#phoneNumber").val() + "</div>");
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

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexBody))
}

func parse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phoneNumber := r.Form.Get("phoneNumber")

	log.Printf("Phone number to be parsed %s", phoneNumber)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/parse", parse)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
