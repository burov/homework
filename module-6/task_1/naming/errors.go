package naming

import "github.com/pkg/errors"

var ErrNameIsEmpty = errors.New("given name is empty")
var ErrEmailInvalidFormat = errors.New("email is in wrong format")
var ErrAlreadyEmailWrongDomain = errors.New("given string is already an email, and it has wrong domain")
var ErrInvalidNameForEmail = errors.New("given string seems to broken")
var ErrMaxNamesCountHit = errors.New("max users with given name reached")
