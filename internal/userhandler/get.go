package userhandler

import (
	"log"
	"net/http"

	"github.com/arty-notes/api-core/internal/core"
	"github.com/arty-notes/api-core/pkg/api"
)

func Get(w http.ResponseWriter, r *http.Request) {
	log.Println("getting user")
	id := r.URL.Query().Get("id")

	if id == "" {

		s := make([]core.User, 0)
		for _, u := range users {
			s = append(s, u)
		}
		log.Println("done")
		api.JsonReply(w, http.StatusBadRequest, s)
		return
	}

	user, ok := users[id]
	if !ok {
		log.Println("done with error", core.GetUserNotFound)
		api.JsonReply(w, http.StatusNotFound, core.GetUserNotFound)
		return
	}

	log.Println("done")
	api.JsonReply(w, http.StatusNotFound, []core.User{user})
}
