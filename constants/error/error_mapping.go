package error

func ErrorMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(append(GeneralErrors[:], UserErrors[:]...))

	for _, item := range allErrors {
		if item.Error() == err.Error() {
			return true
		}
	}

	return false
}
