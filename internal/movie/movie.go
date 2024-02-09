package movie

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"serviceweaver/internal/gen"
)

type MovieServerInterface gen.MovieServiceServer

type MovieServerComponent struct {
	weaver.Implements[MovieServerInterface]
}

func (m MovieServerComponent) GetMovieById(context.Context, *gen.GetMovieByIdRequest) (*gen.GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MovieServerComponent) GetMovieByDirector(_ context.Context, request *gen.GetMovieByDirectorRequest) (*gen.GetMovieResponse, error) {
	//TODO implement me

	//director := request.GetId()
	panic("implement me")
}
