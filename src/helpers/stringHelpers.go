package helpers

import (
	"log"
	"os"
	"regexp"
	"strings"

	"strconv"

	"github.com/joho/godotenv"

	valid "github.com/asaskevich/govalidator"
)

func ValidateEmail(email string) (bool, string) {

	if len(email) < 5 || !strings.Contains(email, "@") || !strings.Contains(email, ".") || strings.Contains(email, " ") {
		return false, "Email inválido."
	}

	parts := strings.Split(email, "@")

	beforeAtSign := ""
	atSign := parts[len(parts)-1]

	for idx, ele := range parts {
		if idx >= len(parts)-1 {
			continue
		}
		beforeAtSign += ele
	}

	if len(beforeAtSign) == 0 {
		return false, "A primeira parte do email está inválida."
	}

	if !strings.Contains(atSign, ".") || len(atSign) < 3 {
		return false, "O domínio do email está inválido."
	}

	return true, ""
}

func ValidatePhone(phone string) (bool, string) {

	if isInt := valid.IsInt(phone); isInt == false {
		return false, "O telefone só pode conter números."
	}

	return true, ""
}

func ValidateName(name string) (bool, string) {

	if len(name) < 2 {
		return false, "O nome deve conter pelo menos duas letras."
	}

	if isValid := regexp.MustCompile(`^[A-Za-z]+$`).MatchString(name); !isValid {
		return false, "O nome só pode conter letras."
	}

	return true, ""
}

func ValidateMessage(message string) (bool, string) {

	limit := 500

	if len(message) > limit {
		return false, "A mensagem deve conter menos do que " + strconv.FormatInt(int64(limit), 10) + " caracteres."
	}

	return true, ""

}

func DecodePassword(password map[int]int, r string) (bool, string) {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Configure a variável de ambiente .env antes de iniciar.")
	}

	if len(os.Getenv("SECRET_REDUCE")) == 0 {
		log.Fatal("Configure a variável de ambiente com o número SECRET_REDUCE antes de iniciar.")
	}

	reverseR := ""

	secretReduce, _ := strconv.Atoi(os.Getenv("SECRET_REDUCE"))

	for i, _ := range r {
		reverseR += r[len(r)-(i+1) : len(r)-(i)]
	}

	decodeR, err := strconv.Atoi(reverseR)

	if err != nil {
		return false, ""
	}

	decodeR = decodeR - secretReduce

	decodedPas := ""

	for i := 0; i < len(password); i++ {
		password[i] += decodeR
		decodedPas += string(password[i])
	}

	if len(decodedPas) != len(password) {
		return false, ""
	}

	return true, decodedPas

}
