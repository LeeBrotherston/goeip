package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

var dbCity, dbASN, dbCountry *geoip2.Reader

func main() {
	//var dbCity, dbASN, dbCountry *geoip2.Reader
	var err error
	dbCity, err = geoip2.Open("db/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	dbASN, err = geoip2.Open("db/GeoLite2-ASN.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	dbCountry, err = geoip2.Open("db/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	defer dbCity.Close()
	defer dbASN.Close()
	defer dbCountry.Close()

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		log.Fatal("This doesn't look like a valid IP address")
	}

	recordASN, err := dbASN.ASN(ip)
	//recordCountry, err := dbCountry.Country(ip)
	recordCity, err := dbCity.City(ip)
	// ASN Lookups

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s:", recordASN.AutonomousSystemOrganization)

	//Country Lookups

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v:%v\n", recordCity.City.Names["en"], recordCity.Continent.Names["en"])
}

func geoLookup(ip string) {

}
