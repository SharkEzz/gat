package utils

import "regexp"

func GetExtension(filePath string) string {
	regex := regexp.MustCompile(`\.([a-zA-Z]+)`)

	result := regex.FindAllStringSubmatch(filePath, -1)

	if len(result) == 0 {
		return ""
	}

	return result[0][1]
}
