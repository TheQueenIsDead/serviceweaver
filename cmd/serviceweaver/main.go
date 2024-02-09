package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"serviceweaver/internal/director"
	"serviceweaver/internal/gen"
	"serviceweaver/internal/movie"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	director weaver.Ref[director.DirectorServerInterface]
	movie    weaver.Ref[movie.MovieServerInterface]
	http     weaver.Listener `weaver:"http"`
	grpc     weaver.Listener `weaver:"grpc"`
}

func serve(ctx context.Context, app *app) error {
	//fmt.Printf("http listener available on %v\n", app.http)

	logger := app.Logger(ctx)

	/*--------------------------------------------------------------------------------
	gRPC

	Despite service weaver handling GRPC between instances for us, we start a gRPC
	server with all of our services registered to it. This does not need to be exposed
	as part of the application, but is used to proxy HTTP requests through after they
	are transposed to gRPC messages using gRPC-Gateway.
	--------------------------------------------------------------------------------*/

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach concrete server implementations to the GRPC server
	gen.RegisterDirectorServiceServer(s, &director.DirectorServerComponent{})
	gen.RegisterMovieServiceServer(s, &movie.MovieServerComponent{})
	// Serve gRPC server
	go func() {
		logger.Info(fmt.Sprintf("grpc listener available on %v", app.grpc))

		err := s.Serve(app.grpc)
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	// TODO: How do we do this the best way? Can we service.Get
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	logger.Debug("Dialling, bring bring!")
	conn, err := grpc.DialContext(
		context.Background(),
		app.grpc.String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	logger.Debug("Dialled in!")

	/*--------------------------------------------------------------------------------
	HTTP
	--------------------------------------------------------------------------------*/
	gwmux := runtime.NewServeMux()

	// TODO: Need to figure out how to let grpc gateway talk to other services via builtin serviceweaver comms.
	registerServiceHandlerFuncs := []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		gen.RegisterDirectorServiceHandler,
		gen.RegisterMovieServiceHandler,
	}
	for _, serviceHandlerFunc := range registerServiceHandlerFuncs {
		err = serviceHandlerFunc(context.Background(), gwmux, conn)
		if err != nil {
			log.Fatalln("Failed to register service handler:", err)
		}
	}

	// TODO: Verify if this addr is correct and happy.
	gwServer := &http.Server{
		Addr:    fmt.Sprintf("%v", app.http.Addr()),
		Handler: gwmux,
	}

	logger.Info(fmt.Sprintf("http listener available on %v", app.http))
	err = gwServer.ListenAndServe()

	return http.Serve(app.http, gwmux)

	//// Serve the /hello endpoint.
	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	name := r.URL.Query().Get("name")
	//	if name == "" {
	//		name = "World"
	//	}
	//	fmt.Fprintf(w, "Hello, %s!\n", name)
	//})
	//return http.Serve(app.http, nil)
}
