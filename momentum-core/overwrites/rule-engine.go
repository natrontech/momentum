package overwrites

func ApplyRule(rule *OverwriteRule) bool {

	// TODO implement applying rules
	return false
}

// rules are applied in reversed order, which leads to the behavior
// that rules occuring earlier in the list have higher priority than
// items later in the list.
func ApplyRules(rules []*OverwriteRule) bool {

	numberOfRules := len(rules)
	for i := range rules {
		if !ApplyRule(rules[numberOfRules-1-i]) {
			return false
		}
	}
	return true
}
