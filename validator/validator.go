package validator

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/localization"
)

type Validator struct {
	errors map[string][]string
}

type Rule struct {
	Name   string
	Params []string
}

type RuleFunc func(value interface{}, params []string, c *gin.Context) bool

var rules = make(map[string]RuleFunc)

// RegisterRule allows projects or the core module to register custom validation rules
func RegisterRule(name string, ruleFunc RuleFunc) {
	rules[name] = ruleFunc
}

// ValidateAll validates the input data against the provided rules
func (v *Validator) ValidateAll(input map[string]interface{}, validationRules map[string][]string, c *gin.Context) {
	v.errors = make(map[string][]string)

	for field, fieldRules := range validationRules {
		value, exists := input[field]

		for _, ruleString := range fieldRules {
			rule := parseRule(ruleString)
			if !exists || !v.applyRule(field, value, rule, c) {
				break
			}
		}
	}
}

// HasErrors checks if there are validation errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors retrieves the validation errors
func (v *Validator) Errors() map[string][]string {
	return v.errors
}

// parseRule parses a rule string into a Rule struct
func parseRule(ruleString string) Rule {
	parts := strings.Split(ruleString, ":")
	ruleName := parts[0]
	var ruleParams []string
	if len(parts) > 1 {
		ruleParams = strings.Split(parts[1], ",")
	}

	return Rule{
		Name:   ruleName,
		Params: ruleParams,
	}
}

// applyRule applies a single rule to a field
func (v *Validator) applyRule(field string, value interface{}, rule Rule, c *gin.Context) bool {
	// Check if the rule exists
	ruleFunc, exists := rules[rule.Name]
	if !exists {
		// Add an error if the rule is unknown
		v.addError(field, c, "validation.unknown", map[string]string{"rule": rule.Name})
		return false
	}

	// Apply the rule function
	if !ruleFunc(value, rule.Params, c) {
		// Add an error if the rule fails
		v.addError(field, c, fmt.Sprintf("validation.%s", rule.Name), map[string]string{
			"attribute": field,
			"params":    strings.Join(rule.Params, ", "),
		})
		return false
	}

	return true
}

// addError adds a validation error message for a field
func (v *Validator) addError(field string, c *gin.Context, messageKey string, attributes map[string]string) {
	message := localization.Translate(c, messageKey, attributes)
	v.errors[field] = append(v.errors[field], message)
}
