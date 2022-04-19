package stringoperations

// IsOneOf checks whether a goal string equals at least one of the options.
func IsOneOf(goal string, options []string) bool {
	for _, option := range options {
		if goal == option {
			return true
		}
	}

	return false
}
