package movie

import (
	"context"
	"github.com/ServiceWeaver/weaver"
)

type MovieServerInterface interface {
	GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieResponse, error)
	GetMovieByDirector(context.Context, *GetMovieByDirectorRequest) (*GetMovieResponse, error)
}

var _ MovieServerInterface = (*MovieServerComponent)(nil)

type MovieServerComponent struct {
	weaver.Implements[MovieServerInterface]
}

func (m MovieServerComponent) GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MovieServerComponent) GetMovieByDirector(context.Context, *GetMovieByDirectorRequest) (*GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

//func (m *MovieServerComponent) mustEmbedUnimplementedMovieServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}
