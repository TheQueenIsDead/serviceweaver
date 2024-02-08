package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

var _ DirectorServiceServer = (*DirectorServerComponent)(nil)

// Implementation of the Reverser component.
type DirectorServerComponent struct {
	weaver.Implements[DirectorServiceServer]

	// Embed the unimplemented Director Server
	UnimplementedDirectorServiceServer
}

func (d DirectorServerComponent) GetDirectorById(ctx context.Context, request *GetDirectorByIdRequest) (*GetDirectorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d DirectorServerComponent) mustEmbedUnimplementedDirectorServiceServer() {
	//TODO implement me
	panic("implement me")
}
