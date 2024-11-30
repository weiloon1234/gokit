package rules

import "github.com/weiloon1234/gokit/validator"

func init() {
	// Register all core rules
	validator.RegisterRule("required", Required)
	validator.RegisterRule("email", Email)
	validator.RegisterRule("min", Min)
	validator.RegisterRule("max", Max)
	validator.RegisterRule("alphanumeric", Alphanumeric)
}
