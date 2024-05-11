package main

import (
	"log"
	"net/http"

	"github.com/arty-notes/api-core/internal/core"
	"github.com/arty-notes/api-core/internal/userhandler"
)

func main() {
	publicRouter := http.NewServeMux()

	publicRouter.HandleFunc(core.UserPost.Url(), userhandler.Post)
	publicRouter.HandleFunc(core.UserPut.Url(), userhandler.Put)
	publicRouter.HandleFunc(core.UserGet.Url(), userhandler.Get)
	publicRouter.HandleFunc(core.UserDelete.Url(), userhandler.Delete)

	// :: SERVER ::

	s := http.Server{
		Addr:    ":4000",
		Handler: publicRouter,
	}

	// run server
	log.Println("running on port 4000")
	err := s.ListenAndServe()
	if err != nil {
		log.Println("failed to run server", err)
	}
}
