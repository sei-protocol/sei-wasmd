package x

import "github.com/pkg/errors"

func ErrorOut() error {
	return errors.WithStack(errors.New("dumb"))
}
