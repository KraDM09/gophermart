package util

import (
	"math/rand"
	"strconv"
	"strings"
	"unicode"

	UUID "github.com/KraDM09/gophermart/internal/app/util/uuid"
)

func CreateHash() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz1234567890"
	hash := ""

	for i := 0; i < 6; i++ {
		randomNumber := rand.Intn(36)
		char := string(alphabet[randomNumber])

		if rand.Intn(2) == 1 {
			char = strings.ToUpper(char)
		}

		hash = hash + char
	}

	return hash
}

func CreateUUID() string {
	uuid := &UUID.GoogleUUID{}

	return uuid.New()
}

func CheckLuna(number string) bool {
	var sum int
	alt := false

	for i := len(number) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(number[i])) {
			return false
		}

		n, _ := strconv.Atoi(string(number[i]))

		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}

	return sum%10 == 0
}
