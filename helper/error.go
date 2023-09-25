package helper

func WriteMsgForTag(tag string) string {
	switch tag {
	case "required":
		return "this field is required"
	case "min":
		return "the data is too short"
	case "max":
		return "the data is too long"

	}
	return ""
}
