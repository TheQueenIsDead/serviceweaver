package movie

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"serviceweaver/internal/gen"
)

type MovieServerInterface interface {
	GetMovieById(context.Context, *gen.GetMovieByIdRequest) (*gen.GetMovieResponse, error)
	GetMovieByDirector(context.Context, *gen.GetMovieByDirectorRequest) (*gen.GetMovieResponse, error)
}

var _ MovieServerInterface = (*MovieServerComponent)(nil)

type MovieServerComponent struct {
	weaver.Implements[MovieServerInterface]
}

func (m MovieServerComponent) GetMovieById(context.Context, *gen.GetMovieByIdRequest) (*gen.GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MovieServerComponent) GetMovieByDirector(context.Context, *gen.GetMovieByDirectorRequest) (*gen.GetMovieResponse, error) {
	//TODO implement me
	panic("implement me")
}

//func (m *MovieServerComponent) mustEmbedUnimplementedMovieServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}
