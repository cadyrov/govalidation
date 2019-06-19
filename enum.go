package validation

func MsgByCode(code int) string {
	codes := map[int]string{
		1000: "internal",
		1001: "only a pointer to a struct can be validated",
		1002: "field #%v must be specified as a pointer",
		1003: "field #%v cannot be found in the struct",
		1010: "cannot get the length of %v",
		1011: "must be either a string or byte slice",
		1012: "must be a valid value",
		1100: "must be a valid date",
		1150: "the data is out of range",
		1200: "the value must be empty",
		1300: "must be in a valid format",
		1400: "must be multiple of %v",
		1500: "must not be in list",

		1600: "is required",
		1700: "cannot be blank",

		1210: "the length must be no more than %v",
		1220: "the length must be no less than %v",
		1230: "the length must be exactly %v",
		1240: "the length must be between %v and %v",

		1800: "must be a valid email address",
		1801: "must be a valid URL",
		1802: "must be a valid request URL",
		1803: "must be a valid request URI",

		1804: "must contain English letters only",
		1805: "must contain digits only",
		1806: "must contain English letters and digits only",
		1807: "must contain unicode letter characters only",
		1808: "must contain unicode decimal digits only",
		1809: "must contain unicode letters and numbers only",
		1810: "must contain unicode number characters only",

		1811: "must be in lower case",
		1812: "must be in upper case",

		1813: "must be a valid hexadecimal number",
		1814: "must be a valid hexadecimal color code",
		1815: "must be a valid RGB color code",

		1816: "must be an integer number",
		1817: "must be a floating point number",

		1818: "must be a valid UUID v3",
		1819: "must be a valid UUID v4",
		1820: "must be a valid UUID v5",
		1821: "must be a valid UUID",

		1822: "must be a valid credit card number",
		1823: "must be a valid ISBN-10",
		1824: "must be a valid ISBN-13",
		1825: "must be a valid ISBN",

		1826: "must be in valid JSON format",

		1827: "must contain ASCII characters only",
		1828: "must contain printable ASCII characters only",
		1829: "must contain multibyte characters",
		1830: "must contain full-width characters",
		1831: "must contain half-width characters",
		1832: "must contain both full-width and half-width characters",
		1833: "must be encoded in Base64",
		1834: "must be a Base64-encoded data URI",

		1835: "must be a valid E164 number",
		1836: "must be a valid two-letter country code",
		1837: "must be a valid three-letter country code",

		1838: "must be a valid dial string",
		1839: "must be a valid MAC address",
		1840: "must be a valid IP address",
		1841: "must be a valid IPv4 address",
		1842: "must be a valid IPv6 address",
		1843: "must be a valid subdomain",
		1844: "must be a valid domain",
		1845: "must be a valid DNS name",
		1846: "must be a valid IP address or DNS name",
		1847: "must be a valid port number",

		1848: "must be a valid hex-encoded MongoDB ObjectId",

		1849: "must be a valid latitude",
		1850: "must be a valid longitude",

		1851: "must be a valid social security number",
		1852: "must be a valid semantic version",

		1900: "is not correct",
	}
	return codes[code]

}
