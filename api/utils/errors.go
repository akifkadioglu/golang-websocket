package utils

func HandleError(msg string) map[string]string {
	result := "something went wrong"
	if msg != "" {
		result = msg
	}
	return map[string]string{
		"error": result,
	}
}
