package director

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

var _ DirectorServiceServer = (*Component)(nil)

// Implementation of the Reverser component.
type Component struct {
	weaver.Implements[DirectorServiceServer]

	// Embed the unimplemented Director Server
	UnimplementedDirectorServiceServer
}

func (d Component) GetDirectorById(ctx context.Context, request *GetDirectorByIdRequest) (*GetDirectorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d Component) mustEmbedUnimplementedDirectorServiceServer() {
	//TODO implement me
	panic("implement me")
}
