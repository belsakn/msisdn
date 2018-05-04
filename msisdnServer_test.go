package main

import "testing"

func TestParse1(t *testing.T) {

	want := successResponse{
		MnoIdentifier:     "A2",
		CountryCode:       386,
		SubscriberNumber:  "040 579 602",
		CountryIdentifier: "SI",
	}

	status, got := getResponse("+38640579602")
	if want != got && status != 200 {
		t.Errorf("want: 200 %v, got: %d %v", want, status, got)
	}

}

func TestParse2(t *testing.T) {

	want := errorResponse{"error parsing phone number", "The string supplied is too short to be a phone number."}

	status, got := getResponse("+386")
	if want != got && status != 500 {
		t.Errorf("want: 200 %v, got: %d %v", want, status, got)
	}

}
