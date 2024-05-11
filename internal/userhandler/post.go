package userhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/arty-notes/api-core/internal/core"
	"github.com/arty-notes/api-core/pkg/api"
)

func Post(w http.ResponseWriter, r *http.Request) {
	log.Println("creating user")
	var body core.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("done with error", core.PostDecodeErr)
		api.JsonReply(w, http.StatusBadRequest, core.PostDecodeErr)
		return
	}

	if body.Email == "" || body.FullName == "" {
		log.Println("done with error", core.PostBadData)
		api.JsonReply(w, http.StatusBadRequest, core.PostBadData)
		return
	}

	if UsersCap <= len(users) {
		log.Println("done with error", core.PostReachedCap)
		api.JsonReply(w, http.StatusBadRequest, core.PostReachedCap)
		return
	}

	body.Id = fmt.Sprintf("u%v", rand.IntN(1000))
	users[body.Id] = body

	log.Println("done")
	api.JsonReply(w, http.StatusBadRequest, body)
}
