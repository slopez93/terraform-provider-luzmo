package utils

import "encoding/json"

func NormalizeString(input string) (string, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}

	normalizedBytes, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}

	return string(normalizedBytes), nil
}

func NormalizeMap(input map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	normalizedContents, err := NormalizeString(string(jsonData))
	if err != nil {
		return "", err
	}

	return normalizedContents, nil
}
