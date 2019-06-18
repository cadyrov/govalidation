package validation

func MsgByCode(code int) string {
	codes := map[int]string{
		1100: "must be a valid date",
		1150: "the data is out of range",
		1200: "the value must be empty",
		1210: "the length must be no more than %v",
		1220: "the length must be no less than %v",
		1230: "the length must be exactly %v",
		1240: "the length must be between %v and %v",
		1300: "must be in a valid format",
		1400: "must be multiple of %v",
		1500: "must not be in list",
		1600: "is required",
		1700: "cannot be blank",
		1800: "must be a valid email address",
		1810: "must be a valid URL",
		1820: "must be a valid request URL",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
		1800: "must be a valid email address",
	}
	return codes[code]

}
