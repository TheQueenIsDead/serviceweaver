package director

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"serviceweaver/internal/gen"
)

type DirectorServerInterface gen.DirectorServiceServer

type DirectorServerComponent struct {
	weaver.Implements[DirectorServerInterface]
}

func (d DirectorServerComponent) GetDirectorById(ctx context.Context, request *gen.GetDirectorByIdRequest) (*gen.GetDirectorResponse, error) {
	//TODO implement me
	id := request.GetId()

	if name, ok := DirectorDatabase[id]; ok {
		return &gen.GetDirectorResponse{
			Data: &gen.Director{
				Name: name,
				Id:   id,
			},
		}, nil
	}
	return nil, errors.New("could not locate director")
}
