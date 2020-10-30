package verror

import (
	"fmt"
	"net/http"

	"github.com/cadyrov/goerr"
)

var mpErr map[int]string = map[int]string{
	1000: "internal",
	1001: "only_a_pointer_to_a_struct_can_be_validated",
	1002: "field_must_be_specified_as_a_pointer",
	1003: "field_cannot_be_found_in_the_struct",
	1004: "cannot_get_the_length %s",
	1005: "must_be_either_a_string_or_byte_slice",
	1006: "type_not_supported",

	1101: "must_be_a_valid_value",
	1102: "must_be_a_valid_date",
	1103: "the_data_is_out_of_range",
	1104: "the_value_must_be_empty",
	1105: "must_be_in_a_valid_format",
	1106: "must_be_multiple_of_%v",
	1107: "must_not_be_in_list",

	1201: "is_required",
	1202: "cannot_be_blank",

	1300: "is_not_correct",
	1301: "the_length_must_be_no_more_than_%v",
	1302: "the_length_must_be_no_less_than_%v",
	1303: "the_length_must_be_exactly_%v",
	1304: "the_length_must_be_between_%v_and_%v",

	1401: "must_be_a_valid_email_address",

	1501: "must_contain_English_letters_only",
	1502: "must_contain_digits_only",
	1503: "must_contain_English_letters_and_digits_only",
	1504: "must_contain_unicode_letter_characters_only",
	1505: "must_contain_unicode_decimal_digits_only",
	1506: "must_contain_unicode_letters_and_numbers_only",
	1507: "must_contain_unicode_number_characters_only",

	1601: "must_be_in_lower_case",
	1602: "must_be_in_upper_case",

	1701: "must_be_a_valid_hexadecimal_number",
	1702: "must_be_a_valid_hexadecimal_color_code",
	1703: "must_be_a_valid_RGB_color_code",

	1801: "must_be_an_integer_number",
	1802: "must_be_a_floating_point_number",

	1901: "must_be_a_valid_UUID_v3",
	1902: "must_be_a_valid_UUID_v4",
	1903: "must_be_a_valid_UUID_v5",
	1904: "must_be_a_valid_UUID",

	2001: "must_be_a_valid_credit_card_number",
	2002: "must_be_a_valid_ISBN_10",
	2003: "must_be_a_valid_ISBN_13",
	2004: "must_be_a_valid_ISBN",

	2101: "must_be_in_valid_JSON_format",

	2201: "must_contain_ASCII_characters_only",
	2202: "must_contain_printable_ASCII_characters_only",
	2203: "must_contain_multibyte_characters",
	2204: "must_contain_full_width_characters",
	2205: "must_contain_half_width_characters",
	2206: "must_contain_both_full_width_and_half_width_characters",
	2207: "must_be_encoded_in_Base64",
	2208: "must_be_a_Base64_encoded_data_URI",

	2301: "must_be_a_valid_E164_number",
	2302: "must_be_a_valid_two_letter_country_code",
	2303: "must_be_a_valid_three_letter_country_code",

	2401: "must_be_a_valid_URL",
	2402: "must_be_a_valid_request_URL",
	2403: "must_be_a_valid_request_URI",
	2404: "must_be_a_valid_dial_string",
	2405: "must_be_a_valid_MAC_address",
	2406: "must_be_a_valid_IP_address",
	2407: "must_be_a_valid_IPv4_address",
	2408: "must_be_a_valid_IPv6_address",
	2409: "must_be_a_valid_subdomain",
	2410: "must_be_a_valid_domain",
	2411: "must_be_a_valid_DNS_name",
	2412: "must_be_a_valid_IP_address_or_DNS_name",
	2413: "must_be_a_valid_port_number",

	2501: "must_be_a_valid_hex_encoded_MongoDB_ObjectId",

	2601: "must_be_a_valid_latitude",
	2602: "must_be_a_valid_longitude",

	2701: "must_be_a_valid_social_security_number",
	2702: "must_be_a_valid_semantic_version",

	2810: "inn_10_simbols_not_correct",
	2811: "only 10 digits",
	2812: "control_sum_is_invalid",
	2820: "inn_12_simbols_not_correct",
	2821: "can't_parse_value",
	2822: "only_12_digits",
	2823: "control_sum_is_invalid",
	2830: "inn_not_correct",
	2840: "ogrn_Law_not_correct",
	2841: "only_13_digits",
	2850: "ogrn_IP_not_correct",
	2851: "only_15_digits",
	2852: "control_sum_is_invalid",
	2860: "ogrn_not_correct",
	2870: "okato_not_correct",
	2880: "snils_not_correct",
}

type ErrStack struct {
	Stack map[string]goerr.IError `json:"stack"`
	Error goerr.IError            `json:"error"`
}

func NewErrStack(message string) ErrStack {
	mp := make(map[string]goerr.IError)
	e := goerr.New(message)
	return ErrStack{
		Stack: mp,
		Error: e,
	}
}

func NewGoErr(code int, args ...interface{}) goerr.IError {
	errtxt, ok := mpErr[code]
	if !ok {
		errtxt = "UnknownError"
	}
	return goerr.New(fmt.Sprintf(errtxt, args...)).HTTP(http.StatusBadRequest)
}
