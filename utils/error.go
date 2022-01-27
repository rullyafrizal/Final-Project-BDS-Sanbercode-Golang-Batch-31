package utils

func StringifyErrors(err ...error) []string {
	var errors []string

	for _, e := range err {
		errors = append(errors, e.Error())
	}

	return errors
}
