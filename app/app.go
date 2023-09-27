package app

import (
	"GoRealEstateManagement/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func New() (*App, error) {
	dsn := "host=localhost user=dev password=abcde dbname=GoRealEstateManagement sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.Property{})
	_ = db.AutoMigrate(&model.Unit{})
	if err != nil {
		return nil, err
	}
	return &App{
		DB:     db,
		Router: mux.NewRouter(),
	}, nil
}

func (a *App) Run() {
	log.Println("Serving on port 8000")
	log.Fatalln(http.ListenAndServe(":8000", a.Router))
}
