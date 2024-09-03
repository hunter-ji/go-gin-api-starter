// @Title user_info_validator.go
// @Description
// @Author Hunter 2024/9/3 17:29

package formatValidator

import (
	"errors"
	"net"
	"regexp"
)

// ValidateMobileNumber
// @Description: Validates mobile phone number format using regex
// @param mobileNumber The mobile phone number to validate
// @return error
func ValidateMobileNumber(mobileNumber string) error {
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(mobileNumber) {
		return errors.New("invalid mobile phone number format")
	}
	return nil
}

// ValidateIDNumber
// @Description: Validates ID card number format using regex
// @param idNumber The ID card number to validate
// @return error
func ValidateIDNumber(idNumber string) error {
	pattern := `(^\d{8}(0\d|10|11|12)([0-2]\d|30|31)\d{3}$)|(^\d{6}(18|19|20)\d{2}(0[1-9]|10|11|12)([0-2]\d|30|31)\d{3}(\d|X|x)$)`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(idNumber) {
		return errors.New("invalid ID card number format")
	}
	return nil
}

// ValidateEmail
// @Description: Validates email address format using regex
// @param email The email address to validate
// @return error
func ValidateEmail(email string) error {
	pattern := `^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(email) {
		return errors.New("invalid email address format")
	}
	return nil
}

// ValidateChineseName
// @Description: Validates Chinese name format using regex
// @param name The Chinese name to validate
// @return error
func ValidateChineseName(name string) error {
	pattern := `^(?:[\u4e00-\u9fa5·]{2,16})`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(name) {
		return errors.New("invalid Chinese name format")
	}
	return nil
}

// ValidateAccountName
// @Description: Validates account name format, allowing numbers, letters, and symbols
// @param accountName The account name to validate
// @return error
func ValidateAccountName(accountName string) error {
	pattern := `^[[:graph:]]{1,50}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(accountName) {
		return errors.New("invalid account name format")
	}
	return nil
}

// ValidateRoleName
// @Description: Validates role name format, allowing Chinese characters, letters, and numbers
// @param roleName The role name to validate
// @return error
func ValidateRoleName(roleName string) error {
	pattern := `^[\u4e00-\u9fa5a-zA-Z0-9]{1,50}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(roleName) {
		return errors.New("invalid role name format")
	}
	return nil
}

// ValidateIPWithWildcard
// @Description: Validates IP address format, allowing wildcards
// @param ip The IP address to validate
// @return error
func ValidateIPWithWildcard(ip string) error {
	pattern := `^(([0,1]?\d{1,2}|2([0-4][0-9]|5[0-5]))|\*)(\.(([0,1]?\d{1,2}|2([0-4][0-9]|5[0-5]))|\*)){3}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(ip) {
		return errors.New("invalid IP address format (with wildcard)")
	}
	return nil
}

// ValidateIP
// @Description: Validates standard IP address format
// @param ip The IP address to validate
// @return error
func ValidateIP(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}
	return errors.New("invalid IP address format")
}
