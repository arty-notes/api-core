package userhandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arty-notes/api-core/internal/core"
	"github.com/arty-notes/api-core/pkg/api"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("deleting user")
	var body core.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("done with error", core.DeleteBadStructure)
		api.JsonReply(w, http.StatusBadRequest, core.DeleteBadStructure)
		return
	}

	if body.Id == "" {
		log.Println("done with error", core.DeleteMissingId)
		api.JsonReply(w, http.StatusBadRequest, core.DeleteMissingId)
		return
	}

	_, ok := users[body.Id]
	if !ok {
		log.Println("done with error", core.DeleteUserNotFound)
		api.JsonReply(w, http.StatusNotFound, core.DeleteUserNotFound)
		return
	}
	delete(users, body.Id)

	log.Println("done")
	api.JsonReply(w, http.StatusOK, struct{}{})
}
