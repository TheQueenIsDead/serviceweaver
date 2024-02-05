package internal

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type DirectorComponent interface {
	Reverse(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type directorComponent struct {
	weaver.Implements[DirectorComponent]
}
