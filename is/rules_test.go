package is

import (
	"testing"

	"github.com/cadyrov/govalidation"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	tests := []struct {
		tag            string
		rule           validation.Rule
		valid, invalid string
		err            string
	}{
		{"Email", Email, "test@example.com", "example.com", "must be a valid email address, ErrCode: 1401"},
		{"URL", URL, "http://example.com", "examplecom", "must be a valid URL, ErrCode: 2401"},
		{"RequestURL", RequestURL, "http://example.com", "examplecom", "must be a valid request URL, ErrCode: 2402"},
		{"RequestURI", RequestURI, "http://example.com", "examplecom", "must be a valid request URI, ErrCode: 2403"},
		{"Alpha", Alpha, "abcd", "ab12", "must contain English letters only, ErrCode: 1501"},
		{"Digit", Digit, "123", "12ab", "must contain digits only, ErrCode: 1502"},
		{"Alphanumeric", Alphanumeric, "abc123", "abc.123", "must contain English letters and digits only, ErrCode: 1503"},
		{"UTFLetter", UTFLetter, "ａｂｃ", "１２３", "must contain unicode letter characters only, ErrCode: 1504"},
		{"UTFDigit", UTFDigit, "１２３", "ａｂｃ", "must contain unicode decimal digits only, ErrCode: 1505"},
		{"UTFNumeric", UTFNumeric, "１２３", "ａｂｃ.１２３", "must contain unicode number characters only, ErrCode: 1507"},
		{"UTFLetterNumeric", UTFLetterNumeric, "ａｂｃ１２３", "ａｂｃ.１２３", "must contain unicode letters and numbers only, ErrCode: 1506"},
		{"LowerCase", LowerCase, "ａｂc", "Aｂｃ", "must be in lower case, ErrCode: 1601"},
		{"UpperCase", UpperCase, "ABC", "ABｃ", "must be in upper case, ErrCode: 1602"},
		{"IP", IP, "74.125.19.99", "74.125.19.999", "must be a valid IP address, ErrCode: 2406"},
		{"IPv4", IPv4, "74.125.19.99", "2001:4860:0:2001::68", "must be a valid IPv4 address, ErrCode: 2407"},
		{"IPv6", IPv6, "2001:4860:0:2001::68", "74.125.19.99", "must be a valid IPv6 address, ErrCode: 2408"},
		{"MAC", MAC, "0123.4567.89ab", "74.125.19.99", "must be a valid MAC address, ErrCode: 2405"},
		{"Subdomain", Subdomain, "example-subdomain", "example.com", "must be a valid subdomain, ErrCode: 2409"},
		{"Domain", Domain, "example-domain.com", "localhost", "must be a valid domain, ErrCode: 2410"},
		{"DNSName", DNSName, "example.com", "abc%", "must be a valid DNS name, ErrCode: 2411"},
		{"Host", Host, "example.com", "abc%", "must be a valid IP address or DNS name, ErrCode: 2412"},
		{"Port", Port, "123", "99999", "must be a valid port number, ErrCode: 2413"},
		{"Latitude", Latitude, "23.123", "100", "must be a valid latitude, ErrCode: 2601"},
		{"Longitude", Longitude, "123.123", "abc", "must be a valid longitude, ErrCode: 2602"},
		{"SSN", SSN, "100-00-1000", "100-0001000", "must be a valid social security number, ErrCode: 2701"},
		{"Semver", Semver, "1.0.0", "1.0.0.0", "must be a valid semantic version, ErrCode: 2702"},
		{"ISBN", ISBN, "1-61729-085-8", "1-61729-085-81", "must be a valid ISBN, ErrCode: 2004"},
		{"ISBN10", ISBN10, "1-61729-085-8", "1-61729-085-81", "must be a valid ISBN-10, ErrCode: 2002"},
		{"ISBN13", ISBN13, "978-4-87311-368-5", "978-4-87311-368-a", "must be a valid ISBN-13, ErrCode: 2003"},
		{"UUID", UUID, "a987fbc9-4bed-3078-cf07-9141ba07c9f1", "a987fbc9-4bed-3078-cf07-9141ba07c9f3a", "must be a valid UUID, ErrCode: 1904"},
		{"UUIDv3", UUIDv3, "b987fbc9-4bed-3078-cf07-9141ba07c9f3", "b987fbc9-4bed-4078-cf07-9141ba07c9f3", "must be a valid UUID v3, ErrCode: 1901"},
		{"UUIDv4", UUIDv4, "57b73598-8764-4ad0-a76a-679bb6640eb1", "b987fbc9-4bed-3078-cf07-9141ba07c9f3", "must be a valid UUID v4, ErrCode: 1902"},
		{"UUIDv5", UUIDv5, "987fbc97-4bed-5078-af07-9141ba07c9f3", "b987fbc9-4bed-3078-cf07-9141ba07c9f3", "must be a valid UUID v5, ErrCode: 1903"},
		{"MongoID", MongoID, "507f1f77bcf86cd799439011", "507f1f77bcf86cd79943901", "must be a valid hex-encoded MongoDB ObjectId, ErrCode: 2501"},
		{"CreditCard", CreditCard, "375556917985515", "375556917985516", "must be a valid credit card number, ErrCode: 2001"},
		{"JSON", JSON, "[1, 2]", "[1, 2,]", "must be in valid JSON format, ErrCode: 2101"},
		{"ASCII", ASCII, "abc", "ａabc", "must contain ASCII characters only, ErrCode: 2201"},
		{"PrintableASCII", PrintableASCII, "abc", "ａabc", "must contain printable ASCII characters only, ErrCode: 2202"},
		{"E164", E164, "+19251232233", "+00124222333", "must be a valid E164 number, ErrCode: 2301"},
		{"CountryCode2", CountryCode2, "US", "XY", "must be a valid two-letter country code, ErrCode: 2302"},
		{"CountryCode3", CountryCode3, "USA", "XYZ", "must be a valid three-letter country code, ErrCode: 2303"},
		{"DialString", DialString, "localhost.local:1", "localhost.loc:100000", "must be a valid dial string, ErrCode: 2404"},
		{"DataURI", DataURI, "data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", "image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", "must be a Base64-encoded data URI, ErrCode: 2208"},
		{"Base64", Base64, "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", "image", "must be encoded in Base64, ErrCode: 2207"},
		{"Multibyte", Multibyte, "ａｂｃ", "abc", "must contain multibyte characters, ErrCode: 2203"},
		{"FullWidth", FullWidth, "３ー０", "abc", "must contain full-width characters, ErrCode: 2204"},
		{"HalfWidth", HalfWidth, "abc123い", "００１１", "must contain half-width characters, ErrCode: 2205"},
		{"VariableWidth", VariableWidth, "３ー０123", "abc", "must contain both full-width and half-width characters, ErrCode: 2206"},
		{"Hexadecimal", Hexadecimal, "FEF", "FTF", "must be a valid hexadecimal number, ErrCode: 1701"},
		{"HexColor", HexColor, "F00", "FTF", "must be a valid hexadecimal color code, ErrCode: 1702"},
		{"RGBColor", RGBColor, "rgb(100, 200, 1)", "abc", "must be a valid RGB color code, ErrCode: 1703"},
		{"Int", Int, "100", "1.1", "must be an integer number, ErrCode: 1801"},
		{"Float", Float, "1.1", "a.1", "must be a floating point number, ErrCode: 1802"},
		{"VariableWidth", VariableWidth, "", "", ""},
	}

	for _, test := range tests {
		err := test.rule.Validate("")
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(test.valid)
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(&test.valid)
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(test.invalid)
		assertError(t, test.err, err, test.tag)
		err = test.rule.Validate(&test.invalid)
		assertError(t, test.err, err, test.tag)
	}
}

func assertError(t *testing.T, expected string, err error, tag string) {
	if expected == "" {
		assert.Nil(t, err, tag)
	} else if assert.NotNil(t, err, tag) {
		assert.Equal(t, expected, err.Error(), tag)
	}
}
