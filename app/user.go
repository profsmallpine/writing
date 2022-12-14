package app

import (
	"github.com/profsmallpine/mid/http/routes"
	"github.com/xy-planning-network/trails/http/middleware"
)

type userStorer struct{}

type user struct{}

func (u user) HasAccess() bool {
	return true
}

func (u user) HomePath() string {
	return routes.RootURL
}

func (store userStorer) GetByID(id uint) (middleware.User, error) {
	return user{}, nil
}
