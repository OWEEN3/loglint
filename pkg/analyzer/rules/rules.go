package rules

import (
	"fmt"
	"strings"
	"unicode"
)

// Проверяем что сообщение не пустое
func IsEmpty(msg string) bool {
	return len(strings.TrimSpace(msg)) == 0
}

// Проверяем что первый символ - маленькая латинская буква.
// Если первая символ не буква, то возвращаем false
// Проверяется только то, что первый символ маленькая буква, остальные могут быть любыми.
func IsFirstLower(msg string) bool {
	first := []rune(msg)[0]
	value := unicode.IsLower(first) && unicode.Is(unicode.Latin, first)
	return value


	// Если нужно проверять все символы
	// if value := unicode.IsLower([]rune(msg)[0]) && unicode.Is(unicode.Latin, []rune(msg)[0]); !value {
	// 	return false
	// }
	// for _, r := range msg {
	// 	if unicode.Is(unicode.Latin, r) && !unicode.IsLower(r) {
	// 		return false
	// 	}
	// }
	// return true
}

// Проверяем что сообщение состоит только из латинских букв, цифр и пробелов
// В данной реализации допускаются '\t', '\n', '\v', '\f', '\r', ' '
func IsValidChars(msg string) bool {
	for _, r := range msg {
		if unicode.Is(unicode.Latin, r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}
		return false
	}
	return true
}

// Ничего лучше этого не смог придумать. Просто проверяем содержит ли строка упоминание чувствительных данных
func ContainsSensitive(msg string) string {
    lowerMsg := strings.ToLower(msg)
    sensitiveWords := []string{
        "password", "passwd", "pass", "passphrase",
        "token", "secret", "pwd", "key", "keyid", "apikey",
        "secretkey", "privatekey", "sshkey", "ssh-rsa",
        "ssh-dss", "ssh-ed25519",
    }
    for _, word := range sensitiveWords {
        if strings.Contains(lowerMsg, word) {
            return fmt.Sprintf("contains keyword %q", word)
        }
    }
    return ""
}