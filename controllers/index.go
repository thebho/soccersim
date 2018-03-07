package controllers

import (
	"SoccerSim/data"
	"net/http"

	"github.com/go-pg/pg"
)

var controller Controller

type Controller interface {
	getDB() *pg.DB
}

type ControllerImp struct {
	db *pg.DB
}

func init() {
	controller = ControllerImp{}
}

func (c ControllerImp) getDB() *pg.DB {
	if c.db == nil {
		c.db = data.ConnectToDB()
	}
	return c.db
}

func setReturnDefaults(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
