package cars

import (
	"fmt"
	"regexp"
	"testing"
)

func TestBrand(t *testing.T) {
	test_value := "TestBrand"
	expected_value := regexp.MustCompile(`\b` + test_value + `\b`)
	brand, err := Brand("TestBrand")

	if err != nil || !expected_value.MatchString(brand) {
		t.Errorf("Brand was not returned succesfully or is not matching the expected format/value")
	}

	fmt.Println(brand)
}

func TestBrandEmpty(t *testing.T) {

	brand, err := Brand("")

	if err == nil || brand != "" {
		t.Errorf("Error handling code is not performing as expected")
	}
}
