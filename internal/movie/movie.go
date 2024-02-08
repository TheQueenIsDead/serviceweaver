package movie

import (
	"context"
	"github.com/ServiceWeaver/weaver"
)

var _ MovieServiceServer = (*Component)(nil)

type Component struct {
	weaver.Implements[MovieServiceServer]
}

func (m *Component) mustEmbedUnimplementedMovieServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (m *Component) GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Component) GetMovieByDirector(context.Context, *GetMovieByDirectorRequest) (*GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}
