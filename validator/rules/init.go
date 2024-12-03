package rules

import "github.com/weiloon1234/gokit/validator"

func init() {
	// Register all core rules
	validator.RegisterRule("unique", Unique)
	validator.RegisterRule("exists", Exists)
	validator.RegisterRule("required", Required)
	validator.RegisterRule("email", Email)
	validator.RegisterRule("min", Min)
	validator.RegisterRule("max", Max)
	validator.RegisterRule("alphanumeric", Alphanumeric)
	validator.RegisterRule("isNumber", IsNumber)
	validator.RegisterRule("isPercentage", IsPercentage)
	validator.RegisterRule("isMoney", IsMoney)
	validator.RegisterRule("isContactNumber", IsContactNumber)
	validator.RegisterRule("isUsername", IsUsername)
	validator.RegisterRule("isPassword", IsPassword)
	validator.RegisterRule("isPassword2", IsPassword2)
}
