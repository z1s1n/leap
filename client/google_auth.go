package client

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/ssh"
)

func keyboardInteractivePassword(secret string) ssh.KeyboardInteractiveChallenge {
	return func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
		token, err := totp.GenerateCode(secret, time.Now().UTC())
		fmt.Println(token)
		return []string{token}, err
	}
}
