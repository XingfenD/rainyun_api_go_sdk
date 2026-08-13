package config

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// ReadSecret prompts the user with msg and reads a line from stdin.
// When stdin is a terminal, echo is disabled so the typed value stays
// invisible. When stdin is piped (non-TTY) it falls back to a normal
// read so scripting still works.
func ReadSecret(msg string) (string, error) {
	fmt.Fprint(os.Stderr, msg)
	var value string
	if term.IsTerminal(int(os.Stdin.Fd())) {
		b, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		value = string(b)
	} else {
		sc := bufio.NewScanner(os.Stdin)
		if !sc.Scan() {
			return "", sc.Err()
		}
		value = sc.Text()
	}
	fmt.Fprintln(os.Stderr)
	return value, nil
}

// ReadLine prompts the user with msg and reads a normal (echo-enabled) line.
func ReadLine(msg string) (string, error) {
	fmt.Fprint(os.Stderr, msg)
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		return "", sc.Err()
	}
	return sc.Text(), nil
}
