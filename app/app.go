package app

import (
	"GoRealEstateManagement/model"
	"encoding/json"
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

type Response struct {
	Msg string
}

type ResponseToken struct {
	Token string
}

func New() (*App, error) {
	dsn := "host=localhost user=dev password=abcde dbname=GoRealEstateManagement sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
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

func (a *App) Respond(w http.ResponseWriter, r *http.Request, data any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}
