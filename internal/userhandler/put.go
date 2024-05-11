package userhandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arty-notes/api-core/internal/core"
	"github.com/arty-notes/api-core/pkg/api"
)

func Put(w http.ResponseWriter, r *http.Request) {
	log.Println("updating user")
	var body core.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("done with error", core.PutBadStructure)
		api.JsonReply(w, http.StatusBadRequest, core.PutBadStructure)
		return
	}

	if body.Id == "" || body.Email == "" || body.FullName == "" {
		log.Println("done with error", core.PutBadData)
		api.JsonReply(w, http.StatusBadRequest, core.PutBadData)
		return
	}

	_, ok := users[body.Id]
	if !ok {
		log.Println("done with error", core.PutUserNotFound)
		api.JsonReply(w, http.StatusNotFound, core.PutUserNotFound)
		return
	}
	users[body.Id] = body

	log.Println("done")
	api.JsonReply(w, http.StatusBadRequest, body)
}
