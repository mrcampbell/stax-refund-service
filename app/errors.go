package app

import "github.com/pkg/errors"

var ErrorNotFound = errors.New("not found")
var ErrorAlreadyExists = errors.New("already exists")
