package movie

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"serviceweaver/internal/gen"
)

type MovieServerInterface gen.MovieServiceServer

type MovieServerComponent struct {
	weaver.Implements[MovieServerInterface]
}

func (m MovieServerComponent) GetMovieByName(_ context.Context, request *gen.GetMovieByNameRequest) (*gen.GetMovieResponse, error) {
	name := request.GetName()

	if director, ok := MovieDatabase[name]; ok {
		return &gen.GetMovieResponse{
			Data: []*gen.Movie{
				{
					Name: name,
					Director: &gen.Director{
						Name: "", // TODO: Inflate this
						Id:   director,
					},
				},
			},
		}, nil
	}
	return nil, errors.New("could not locate director")
}

func (m MovieServerComponent) GetMovieByDirector(_ context.Context, request *gen.GetMovieByDirectorRequest) (*gen.GetMovieResponse, error) {
	id := request.GetId()

	results := []*gen.Movie{}
	for name, director := range MovieDatabase {
		if id == director {
			results = append(results, &gen.Movie{
				Name:     name,
				Director: nil, // TODO: Inflate this
			})
		}
	}

	return &gen.GetMovieResponse{
		Data: results,
	}, nil
}
