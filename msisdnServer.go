package main

import (
	"log"
	"net/http"
)

const indexBody = `
<html>
  <head>
	<title>PhoneNumberServer</title>
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
</html>
`

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexBody))
}

func main() {
	http.HandleFunc("/", index)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}