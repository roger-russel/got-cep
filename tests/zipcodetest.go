package tests

import (
	"github.com/revel/revel/testing"
)

type ZipCodeTest struct {
	testing.TestSuite
}

func (t *ZipCodeTest) Before() {
	println("Set up")
}

func (t *ZipCodeTest) TestThatZipCodeRouteWorks() {

	t.Get("/zipcode/03816030")
	t.AssertStatus(401) // Access have been revoked :(
}

func (t *ZipCodeTest) After() {
	println("Tear down")
}

