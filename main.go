package main

import (
	"battleship/app/api/coordinates"
	"battleship/app/api/fight"
	"battleship/app/api/game"
	"battleship/app/api/ship"
	"battleship/app/api/user"
	"battleship/app/config"
	"battleship/app/processors"
	"battleship/app/repository"
)
import "battleship/app/generated"
import "battleship/app/system/db"

const (
	serviceName = "github.com/rimma2828/battleShip"
	logIndex    = "rimma"
)

var VERSION = "undefined"

func main() {
	srv := server.New(VERSION, logIndex)
	appConfig, err := config.Init()
	if err != nil {
		panic(err)
	}
	//db, err := db.NewPQDatabase("postgres://testapp@localhost:6543/testapp?sslmode=disable&binary_parameters=yes", 5, 10, 10)
	db, err := db.NewPQDatabase(appConfig.GetDSN(), 5, 10, 10)
	if err != nil {
		panic(err)
	}

	repoUser := repository.NewUserRepository()
	userProcessor := processors.NewUserProcessor(db, repoUser, srv.Log)
	userApi := user.NewUser(userProcessor, srv.Log)
	srv.MustRegisterServiceWithHTTPPost(userApi)

	repoShip := repository.NewShipRepository()
	shipProcessor := processors.NewShipProcessor(db, repoShip, srv.Log)
	shipApi := ship.NewShip(shipProcessor, srv.Log)
	srv.MustRegisterServiceWithHTTPPost(shipApi)

	repoCoordinates := repository.NewCoordinatesRepository()
	coordinatesProcessor := processors.NewCoordinatesProcessor(db, repoCoordinates, repoUser, repoShip, srv.Log)
	coordinatesApi := coordinates.NewCoordinates(coordinatesProcessor, srv.Log)
	srv.MustRegisterServiceWithHTTPPost(coordinatesApi)

	repoFight := repository.NewFightRepository()
	fightProcessor := processors.NewFightProcessor(db, repoFight, repoUser, repoCoordinates, srv.Log)
	fightApi := fight.NewFight(fightProcessor, srv.Log)
	srv.MustRegisterServiceWithHTTPPost(fightApi)

	gameProcessor := processors.NewGameProcessor(db, repoFight, repoCoordinates, srv.Log)
	gameApi := game.NewGame(gameProcessor, srv.Log)
	srv.MustRegisterServiceWithHTTPPost(gameApi)

	srv.Run()
}
