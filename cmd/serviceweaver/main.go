package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"log"
	"net/http"
	"serviceweaver/internal/director"
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
	// The hello listener will listen on a random port chosen by the operating
	// system. This behavior can be changed in the config file.
	fmt.Printf("http listener available on %v\n", app.http)

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Hello, %s!\n", name)
	})
	return http.Serve(app.http, nil)
}
