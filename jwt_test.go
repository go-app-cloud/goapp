package goapp

import (
	"log"
	"testing"
)

func TestJWT(t *testing.T) {
	token := Token{
		Secret: "1234567890",
	}
	d := Claims{}
	tt, err := token.Build(d)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tt)
}

func TestJWTParse(t *testing.T) {
	str := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBLZXkiOiJhZG1pbiIsIlNlY3JldEtleSI6ImFkbWluIiwiZXhwIjoxNTcxNjQ4MDMxfQ.w6NJp_NH-sseKQklVc7ybcHobPapoGttnCyv2ZouvJg`
	token := Token{
		Secret: "1234567890",
	}
	tt, err := token.Parse(str)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tt.Id)

}
