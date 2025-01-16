package validators

import (
	"errors"
	"fmt"
	"log"

	"github.com/Aadil-Nabi/cmgarage/internal/config"
	"github.com/go-passwd/validator"
)

func GetPasswordValidated() {
	// MustLoad function from config package loads password from the provided config.yaml file in command line
	var configs = config.MustLoad()
	var password = configs.Cm_password
	ValidatePassword(password)
}

func ValidatePassword(password string) {

	passwordValidator := validator.New(
		validator.MinLength(14, errors.New("password validation failed, must have password of length 15")),
		validator.CommonPassword(errors.New("password validation failed, must not be common")),
		validator.Similarity([]string{"username", "PaSSw0rd@1234", "Asdf@1234", "Temp@1234", "Welcome@1234"}, nil, errors.New("password validation failed: must not a basic pattern, make it more complex to guess")),
		validator.ContainsAtLeast("0123456789", 4, errors.New("password validation failed, must contain at least four numbers 0-9")),
		validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 2, errors.New("password validation failed, must contain atleast 2 lowercase letters(a-z)")),
		validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2, errors.New("password validation failed, must contain at least 2 uppercase letters (A-Z)")),
		validator.ContainsAtLeast("! \" # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _ ` { | } ~", 4, errors.New("password validation failed, must contain atleast 4 special characters")))
	if err := passwordValidator.Validate(password); err != nil {
		log.Fatal(err)
	}

	fmt.Println("=> ✔ Password meets all complexity requirements, Password Validation Passed. ")

}
