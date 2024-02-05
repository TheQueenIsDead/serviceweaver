package internal

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	gen "test-proto/internal/gen"
)

var _ gen.MovieServiceServer = (*movieComponent)(nil)

// Implementation of the Reverser component.
type movieComponent struct {
	weaver.Implements[gen.MovieServiceServer]
}

func (m *movieComponent) GetMovieById(context.Context, *gen.GetMovieByIdRequest) (*gen.GetMovieResponse, error) {
	return nil, nil
}

func (m *movieComponent) GetMovieByDirector(context.Context, *gen.GetMovieByDirectorRequest) (*gen.GetMovieResponse, error) {
	return nil, nil
}

func (m *movieComponent) mustEmbedUnimplementedMovieServiceServer() {}
