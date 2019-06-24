package validation

var codes = map[int]string{
	1000: "internal",
	1001: "only a pointer to a struct can be validated",
	1002: "field #%v must be specified as a pointer",
	1003: "field #%v cannot be found in the struct",
	1004: "cannot get the length of %v",
	1005: "must be either a string or byte slice",

	1101: "must be a valid value",
	1102: "must be a valid date",
	1103: "the data is out of range",
	1104: "the value must be empty",
	1105: "must be in a valid format",
	1106: "must be multiple of %v",
	1107: "must not be in list",

	1201: "is required",
	1202: "cannot be blank",

	1300: "is not correct",
	1301: "the length must be no more than %v",
	1302: "the length must be no less than %v",
	1303: "the length must be exactly %v",
	1304: "the length must be between %v and %v",

	1401: "must be a valid email address",

	1501: "must contain English letters only",
	1502: "must contain digits only",
	1503: "must contain English letters and digits only",
	1504: "must contain unicode letter characters only",
	1505: "must contain unicode decimal digits only",
	1506: "must contain unicode letters and numbers only",
	1507: "must contain unicode number characters only",

	1601: "must be in lower case",
	1602: "must be in upper case",

	1701: "must be a valid hexadecimal number",
	1702: "must be a valid hexadecimal color code",
	1703: "must be a valid RGB color code",

	1801: "must be an integer number",
	1802: "must be a floating point number",

	1901: "must be a valid UUID v3",
	1902: "must be a valid UUID v4",
	1903: "must be a valid UUID v5",
	1904: "must be a valid UUID",

	2001: "must be a valid credit card number",
	2002: "must be a valid ISBN-10",
	2003: "must be a valid ISBN-13",
	2004: "must be a valid ISBN",

	2101: "must be in valid JSON format",

	2201: "must contain ASCII characters only",
	2202: "must contain printable ASCII characters only",
	2203: "must contain multibyte characters",
	2204: "must contain full-width characters",
	2205: "must contain half-width characters",
	2206: "must contain both full-width and half-width characters",
	2207: "must be encoded in Base64",
	2208: "must be a Base64-encoded data URI",

	2301: "must be a valid E164 number",
	2302: "must be a valid two-letter country code",
	2303: "must be a valid three-letter country code",

	2401: "must be a valid URL",
	2402: "must be a valid request URL",
	2403: "must be a valid request URI",
	2404: "must be a valid dial string",
	2405: "must be a valid MAC address",
	2406: "must be a valid IP address",
	2407: "must be a valid IPv4 address",
	2408: "must be a valid IPv6 address",
	2409: "must be a valid subdomain",
	2410: "must be a valid domain",
	2411: "must be a valid DNS name",
	2412: "must be a valid IP address or DNS name",
	2413: "must be a valid port number",

	2501: "must be a valid hex-encoded MongoDB ObjectId",

	2601: "must be a valid latitude",
	2602: "must be a valid longitude",

	2701: "must be a valid social security number",
	2702: "must be a valid semantic version",

	2801: "inn 10 simbols not correct",
	2802: "inn 12 simbols not correct",
	2803: "inn not correct",
	2804: "ogrn Law not correct",
	2805: "ogrn IP not correct",
	2806: "ogrn not correct",
	2807: "okato not correct",
	2808: "snils not correct",
}

func MsgByCode(code int) string {
	return codes[code]
}

func Dictionary() map[int]string {
	return codes
}
