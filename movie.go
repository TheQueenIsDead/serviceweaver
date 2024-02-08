package main

import (
	"context"
	"github.com/ServiceWeaver/weaver"
)

var _ MovieServiceServer = (*MovieServerComponent)(nil)

type MovieServerComponent struct {
	weaver.Implements[MovieServiceServer]
}

func (m *MovieServerComponent) mustEmbedUnimplementedMovieServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (m *MovieServerComponent) GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieServerComponent) GetMovieByDirector(context.Context, *GetMovieByDirectorRequest) (*GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}
