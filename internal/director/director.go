package director

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"serviceweaver/internal/gen"
)

type DirectorServerInterface interface {
	GetDirectorById(ctx context.Context, request *gen.GetDirectorByIdRequest) (*gen.GetDirectorResponse, error)
}

var _ DirectorServerInterface = (*DirectorServerComponent)(nil)

type DirectorServerComponent struct {
	weaver.Implements[DirectorServerInterface]

	//// Embed the unimplemented Director Server
	//UnimplementedDirectorServiceServer
}

func (d DirectorServerComponent) GetDirectorById(ctx context.Context, request *gen.GetDirectorByIdRequest) (*gen.GetDirectorResponse, error) {
	//TODO implement me
	panic("implement me")
}

//func (d DirectorServerComponent) mustEmbedUnimplementedDirectorServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}
