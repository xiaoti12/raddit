package jwt

import "testing"

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNzExOTI0NzczMzg4Njk3NiwidXNlcm5hbWUiOiJ4aWFvdGkiLCJpc3MiOiJyYWRkaXQiLCJleHAiOjE3MDg1MDE1NjF9.ADfO1SCRnMml6sPVU2tx0UeVI2gNvhw80nZrMUrPNcY"
	claims, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(claims)
}
