package country

import (
	"fmt"
	"testing"
)

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func TestCountry(t *testing.T) {
	region, continent := location("Boulder")
	fmt.Printf("John works in %s, %s\n", region, continent)
}
