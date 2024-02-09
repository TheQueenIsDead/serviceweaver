package director

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"serviceweaver/internal/gen"
)

type DirectorServerInterface gen.DirectorServiceServer

type DirectorServerComponent struct {
	weaver.Implements[DirectorServerInterface]
}

func (d DirectorServerComponent) GetDirectorById(ctx context.Context, request *gen.GetDirectorByIdRequest) (*gen.GetDirectorResponse, error) {
	//TODO implement me
	panic("implement me")
}
