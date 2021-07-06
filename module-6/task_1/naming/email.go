package naming

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func validateName(name string) error {
	if len(name) == 0 {
		return ErrNameIsEmpty
	}
	return nil
}

func emalifyName(s string) string {
	s = strings.ToLower(s)
	pattern := regexp.MustCompile("([^a-z])")
	return pattern.ReplaceAllString(s, ".")
}

func addDomain(emlName, domain string) (string, error) {
	parts := strings.Split(emlName, "@")
	if len(parts) > 2 {
		return "", ErrInvalidNameForEmail
	}
	if len(parts) == 2 {
		if parts[1] == domain {
			return emlName, nil
		}
		return "", ErrAlreadyEmailWrongDomain
	}
	return fmt.Sprintf("%s@%s", emlName, domain), nil
}

func NameToEmail(name string) (eml string, err error) {
	if err := validateName(name); err != nil {
		return "", errors.Wrap(err, "can't convert name to email")
	}
	emlName := emalifyName(name)
	if eml, err = addDomain(emlName, "acme.com"); err != nil {
		return "", errors.Wrap(err, "can't convert name to email")
	}
	return eml, nil
}


func NameFromEmail(e string) (string, error) {
	parts := strings.Split(e, "@")
	if len(parts) != 2 {
		return "", errors.Wrap(ErrEmailInvalidFormat, e)
	}
	return parts[0], nil
}
