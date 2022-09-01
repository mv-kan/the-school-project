package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func mustConnectToDB() *gorm.DB {
	godotenv.Load("./.env")
	var (
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_HOST     = os.Getenv("DB_HOST")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_NAME     = os.Getenv("DB_NAME")
	)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	// TODO add connection to db
	// TODO New repoisotries
	// TODO New services
	// TODO New routers

	// TODO add logger +
	// TODO add documentation for all internal tools
	// TODO add documentation for all routes in swagger
	godotenv.Load("./.env")
	var (
		LISTENING_PORT = os.Getenv("LISTENING_PORT")
	)
	// open db connection
	db := mustConnectToDB()
	// init all services and the logger
	var (
		log               = logger.New()
		dormServ          = service.New(repo.New[entity.Dormitory](db))
		invoiceServ       = service.NewInvoice(db)
		pupilServ         = service.New(repo.New[entity.Pupil](db))
		roomTypeServ      = service.New(repo.New[entity.RoomType](db))
		roomServ          = service.New(repo.New[entity.Room](db))
		schoolClassServ   = service.New(repo.New[entity.SchoolClass](db))
		supervisorServ    = service.New(repo.New[entity.Supervisor](db))
		typeOfServiceServ = service.New(repo.New[entity.TypeOfService](db))
		enrollServ        = service.NewEnroll(db)
		financialServ     = service.NewFinancial(db)
		roomStatServ      = service.NewRoomStat(db)
	)
	log.Debug(fmt.Sprint("connected to db and initilized services", LISTENING_PORT))

	// get router
	r := router.New(log, dormServ, invoiceServ, pupilServ, roomTypeServ, roomServ, schoolClassServ, supervisorServ, typeOfServiceServ, enrollServ, financialServ, roomStatServ)

	http.Handle("/", r)
	log.Debug(fmt.Sprint("starting server on port ", LISTENING_PORT))
	// start server
	log.Error(http.ListenAndServe(fmt.Sprint(":", LISTENING_PORT), nil).Error())
}
