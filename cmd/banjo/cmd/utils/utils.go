package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/bgentry/speakeasy"
	isatty "github.com/mattn/go-isatty"
)

var buf *bufio.Reader

func GetPassword(prompt string) (password string, err error) {
	if inputIsTty() {
		password, err = speakeasy.Ask(prompt)
	} else {
		password, err = stdinPassword()
	}
	return
}

func inputIsTty() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
}

func stdinPassword() (string, error) {
	if buf == nil {
		buf = bufio.NewReader(os.Stdin)
	}
	password, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(password), nil
}
