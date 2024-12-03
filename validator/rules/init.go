package rules

import "github.com/weiloon1234/gokit/validator"

func init() {
	// Register all core rules
	validator.RegisterRule("alphanumeric", Alphanumeric)
	validator.RegisterRule("confirmed", Confirmed)
	validator.RegisterRule("email", Email)
	validator.RegisterRule("exists", Exists)
	validator.RegisterRule("isContactNumber", IsContactNumber)
	validator.RegisterRule("isMoney", IsMoney)
	validator.RegisterRule("isNumber", IsNumber)
	validator.RegisterRule("isPassword", IsPassword)
	validator.RegisterRule("isPassword2", IsPassword2)
	validator.RegisterRule("isPercentage", IsPercentage)
	validator.RegisterRule("isUsername", IsUsername)
	validator.RegisterRule("max", Max)
	validator.RegisterRule("min", Min)
	validator.RegisterRule("required", Required)
	validator.RegisterRule("unique", Unique)
}
