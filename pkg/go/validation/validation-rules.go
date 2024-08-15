package validation

import (
	"fmt"
	"regexp"
)

type Rule string

const (
	RuleType      Rule = "[^:#@\\*\\s]{1,254}"
	RuleRelation  Rule = "[^:#@\\*\\s]{1,50}"
	RuleCondition Rule = "[^\\*\\s]{1,50}"
	RuleID        Rule = "[^#:\\*\\s]+"
	RuleObject    Rule = "[^\\s]{2,256}"
)

func ValidateObject(object string) bool {
	typeMatch, _ := regexp.MatchString(fmt.Sprintf("^%s:%s$", RuleType, RuleID), object)
	objectMatch, _ := regexp.MatchString(fmt.Sprintf("^%s$", RuleObject), object)

	return typeMatch && objectMatch
}

func ValidateRelation(relation string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("^%s$", RuleRelation), relation)

	return match
}

func ValidateUserSet(userSet string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("^%s:%s#%s$", RuleType, RuleID, RuleRelation), userSet)

	return match
}

func ValidateUserObject(userObject string) bool {
	typeMatch, _ := regexp.MatchString(fmt.Sprintf("^%s:%s$", RuleType, RuleID), userObject)
	objectMatch, _ := regexp.MatchString(fmt.Sprintf("^%s$", RuleObject), userObject)

	return typeMatch && objectMatch
}

func ValidateUserWildcard(userWildcard string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("^%s:\\*$", RuleType), userWildcard)

	return match
}

func ValidateUser(user string) bool {
	return ValidateUserSet(user) || ValidateObject(user) || ValidateUserWildcard(user)
}

func ValidateRelationshipCondition(condition string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("^%s$", RuleCondition), condition)

	return match
}

func ValidateType(typeString string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("^%s$", RuleType), typeString)

	return match
}
