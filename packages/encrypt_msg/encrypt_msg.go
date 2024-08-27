package encrypt_msg

import (
	"strings"
)

type EncryptMsg struct {
	Key     string
	Message string
}

func (em EncryptMsg) EncryptMessage() string {
	if em.Key == "" {
		em.Key = "DCJ"
	}

	if em.Message == "" {
		return ""
	}

	var result strings.Builder
	vowels := "aeiouAEIOU"

	for _, char := range em.Message {
		if strings.ContainsRune(vowels, char) {
			result.WriteString(em.Key)
		}

		result.WriteRune(char)
	}

	return result.String()
}
