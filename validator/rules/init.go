package rules

import "github.com/weiloon1234/gokit/validator"

func init() {
	// Register all core rules
	validator.RegisterRule("alphanumeric", Alphanumeric)
	validator.RegisterRule("confirmed", Confirmed)
	validator.RegisterRule("date", Date)
	validator.RegisterRule("date_format", DateFormat)
	validator.RegisterRule("datetime", DatetimeISO8601)
	validator.RegisterRule("email", Email)
	validator.RegisterRule("exists", Exists)
	validator.RegisterRule("file", File)
	validator.RegisterRule("image", Image)
	validator.RegisterRule("is_contact_number", IsContactNumber)
	validator.RegisterRule("is_money", IsMoney)
	validator.RegisterRule("is_number", IsNumber)
	validator.RegisterRule("is_password", IsPassword)
	validator.RegisterRule("is_password2", IsPassword2)
	validator.RegisterRule("is_percentage", IsPercentage)
	validator.RegisterRule("is_username", IsUsername)
	validator.RegisterRule("max", Max)
	validator.RegisterRule("max_amount", MaxAmount)
	validator.RegisterRule("min", Min)
	validator.RegisterRule("min_amount", MinAmount)
	validator.RegisterRule("required", Required)
	validator.RegisterRule("unique", Unique)
}
