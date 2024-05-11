package userhandler

import (
	"github.com/arty-notes/api-core/internal/core"
)

const UsersCap = 100

var users = map[string]core.User{}
