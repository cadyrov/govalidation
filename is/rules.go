// Package is provides a list of commonly used string validation rules.
package is

import (
	"github.com/asaskevich/govalidator"
	"github.com/cadyrov/govalidation"
	"regexp"
	"unicode"
)

var (
	// Email validates if a string is an email or not.
	Email = validation.NewStringRule(govalidator.IsEmail, 1401)
	// URL validates if a string is a valid URL
	URL = validation.NewStringRule(govalidator.IsURL, 2401)
	// RequestURL validates if a string is a valid request URL
	RequestURL = validation.NewStringRule(govalidator.IsRequestURL, 2402)
	// RequestURI validates if a string is a valid request URI
	RequestURI = validation.NewStringRule(govalidator.IsRequestURI, 2403)
	// Alpha validates if a string contains English letters only (a-zA-Z)
	Alpha = validation.NewStringRule(govalidator.IsAlpha, 1501)
	// Digit validates if a string contains digits only (0-9)
	Digit = validation.NewStringRule(isDigit, 1502)
	// Alphanumeric validates if a string contains English letters and digits only (a-zA-Z0-9)
	Alphanumeric = validation.NewStringRule(govalidator.IsAlphanumeric, 1503)
	// UTFLetter validates if a string contains unicode letters only
	UTFLetter = validation.NewStringRule(govalidator.IsUTFLetter, 1504)
	// UTFDigit validates if a string contains unicode decimal digits only
	UTFDigit = validation.NewStringRule(govalidator.IsUTFDigit, 1505)
	// UTFLetterNumeric validates if a string contains unicode letters and numbers only
	UTFLetterNumeric = validation.NewStringRule(govalidator.IsUTFLetterNumeric, 1506)
	// UTFNumeric validates if a string contains unicode number characters (category N) only
	UTFNumeric = validation.NewStringRule(isUTFNumeric, 1507)
	// LowerCase validates if a string contains lower case unicode letters only
	LowerCase = validation.NewStringRule(govalidator.IsLowerCase, 1601)
	// UpperCase validates if a string contains upper case unicode letters only
	UpperCase = validation.NewStringRule(govalidator.IsUpperCase, 1602)
	// Hexadecimal validates if a string is a valid hexadecimal number
	Hexadecimal = validation.NewStringRule(govalidator.IsHexadecimal, 1701)
	// HexColor validates if a string is a valid hexadecimal color code
	HexColor = validation.NewStringRule(govalidator.IsHexcolor, 1702)
	// RGBColor validates if a string is a valid RGB color in the form of rgb(R, G, B)
	RGBColor = validation.NewStringRule(govalidator.IsRGBcolor, 1703)
	// Int validates if a string is a valid integer number
	Int = validation.NewStringRule(govalidator.IsInt, 1801)
	// Float validates if a string is a floating point number
	Float = validation.NewStringRule(govalidator.IsFloat, 1802)
	// UUIDv3 validates if a string is a valid version 3 UUID
	UUIDv3 = validation.NewStringRule(govalidator.IsUUIDv3, 1901)
	// UUIDv4 validates if a string is a valid version 4 UUID
	UUIDv4 = validation.NewStringRule(govalidator.IsUUIDv4, 1902)
	// UUIDv5 validates if a string is a valid version 5 UUID
	UUIDv5 = validation.NewStringRule(govalidator.IsUUIDv5, 1903)
	// UUID validates if a string is a valid UUID
	UUID = validation.NewStringRule(govalidator.IsUUID, 1904)
	// CreditCard validates if a string is a valid credit card number
	CreditCard = validation.NewStringRule(govalidator.IsCreditCard, 2001)
	// ISBN10 validates if a string is an ISBN version 10
	ISBN10 = validation.NewStringRule(govalidator.IsISBN10, 2002)
	// ISBN13 validates if a string is an ISBN version 13
	ISBN13 = validation.NewStringRule(govalidator.IsISBN13, 2003)
	// ISBN validates if a string is an ISBN (either version 10 or 13)
	ISBN = validation.NewStringRule(isISBN, 2004)
	// JSON validates if a string is in valid JSON format
	JSON = validation.NewStringRule(govalidator.IsJSON, 2101)
	// ASCII validates if a string contains ASCII characters only
	ASCII = validation.NewStringRule(govalidator.IsASCII, 2201)
	// PrintableASCII validates if a string contains printable ASCII characters only
	PrintableASCII = validation.NewStringRule(govalidator.IsPrintableASCII, 2202)
	// Multibyte validates if a string contains multibyte characters
	Multibyte = validation.NewStringRule(govalidator.IsMultibyte, 2203)
	// FullWidth validates if a string contains full-width characters
	FullWidth = validation.NewStringRule(govalidator.IsFullWidth, 2204)
	// HalfWidth validates if a string contains half-width characters
	HalfWidth = validation.NewStringRule(govalidator.IsHalfWidth, 2205)
	// VariableWidth validates if a string contains both full-width and half-width characters
	VariableWidth = validation.NewStringRule(govalidator.IsVariableWidth, 2206)
	// Base64 validates if a string is encoded in Base64
	Base64 = validation.NewStringRule(govalidator.IsBase64, 2207)
	// DataURI validates if a string is a valid base64-encoded data URI
	DataURI = validation.NewStringRule(govalidator.IsDataURI, 2208)
	// E164 validates if a string is a valid ISO3166 Alpha 2 country code
	E164 = validation.NewStringRule(isE164Number, 2301)
	// CountryCode2 validates if a string is a valid ISO3166 Alpha 2 country code
	CountryCode2 = validation.NewStringRule(govalidator.IsISO3166Alpha2, 2302)
	// CountryCode3 validates if a string is a valid ISO3166 Alpha 3 country code
	CountryCode3 = validation.NewStringRule(govalidator.IsISO3166Alpha3, 2303)
	// DialString validates if a string is a valid dial string that can be passed to Dial()
	DialString = validation.NewStringRule(govalidator.IsDialString, 2404)
	// MAC validates if a string is a MAC address
	MAC = validation.NewStringRule(govalidator.IsMAC, 2405)
	// IP validates if a string is a valid IP address (either version 4 or 6)
	IP = validation.NewStringRule(govalidator.IsIP, 2406)
	// IPv4 validates if a string is a valid version 4 IP address
	IPv4 = validation.NewStringRule(govalidator.IsIPv4, 2407)
	// IPv6 validates if a string is a valid version 6 IP address
	IPv6 = validation.NewStringRule(govalidator.IsIPv6, 2408)
	// Subdomain validates if a string is valid subdomain
	Subdomain = validation.NewStringRule(isSubdomain, 2409)
	// Domain validates if a string is valid domain
	Domain = validation.NewStringRule(isDomain, 2410)
	// DNSName validates if a string is valid DNS name
	DNSName = validation.NewStringRule(govalidator.IsDNSName, 2411)
	// Host validates if a string is a valid IP (both v4 and v6) or a valid DNS name
	Host = validation.NewStringRule(govalidator.IsHost, 2412)
	// Port validates if a string is a valid port number
	Port = validation.NewStringRule(govalidator.IsPort, 2413)
	// MongoID validates if a string is a valid Mongo ID
	MongoID = validation.NewStringRule(govalidator.IsMongoID, 2501)
	// Latitude validates if a string is a valid latitude
	Latitude = validation.NewStringRule(govalidator.IsLatitude, 2601)
	// Longitude validates if a string is a valid longitude
	Longitude = validation.NewStringRule(govalidator.IsLongitude, 2602)
	// SSN validates if a string is a social security number (SSN)
	SSN = validation.NewStringRule(govalidator.IsSSN, 2701)
	// Semver validates if a string is a valid semantic version
	Semver = validation.NewStringRule(govalidator.IsSemver, 2702)
)

var (
	reDigit = regexp.MustCompile("^[0-9]+$")
	// Subdomain regex source: https://stackoverflow.com/a/7933253
	reSubdomain = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9\-]{0,61}[A-Za-z0-9])?$`)
	// E164 regex source: https://stackoverflow.com/a/23299989
	reE164 = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	// Domain regex source: https://stackoverflow.com/a/7933253
	// Slightly modified: Removed 255 max length validation since Go regex does not
	// support lookarounds. More info: https://stackoverflow.com/a/38935027
	reDomain = regexp.MustCompile(`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-z]{1,63}| xn--[a-z0-9]{1,59})$`)
)

func isISBN(value string) bool {
	return govalidator.IsISBN(value, 10) || govalidator.IsISBN(value, 13)
}

func isDigit(value string) bool {
	return reDigit.MatchString(value)
}

func isE164Number(value string) bool {
	return reE164.MatchString(value)
}

func isSubdomain(value string) bool {
	return reSubdomain.MatchString(value)
}

func isDomain(value string) bool {
	if len(value) > 255 {
		return false
	}

	return reDomain.MatchString(value)
}

func isUTFNumeric(value string) bool {
	for _, c := range value {
		if unicode.IsNumber(c) == false {
			return false
		}
	}
	return true
}
