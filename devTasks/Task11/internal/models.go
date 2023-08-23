package internal

import (
	"L2/devTasks/Task11/config"
	"net/http"
	"time"
)

type Model struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Info string    `json:"info"`
	Date time.Time `json:"date"`
}

type UseCase interface {
	Create(model []byte) error
	Update(model []byte) error
	Delete(model []byte) error
	Get(flag string) ([]byte, error)
}

type Repository interface {
	Create(model Model) error
	Update(model Model) error
	Delete(model Model) error
	Get() map[int]Model
}

type Rest interface {
	Hearing(conf *config.Config) error
	Post(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetForDay(w http.ResponseWriter, r *http.Request)
	GetForWeek(w http.ResponseWriter, r *http.Request)
	GetForMonth(w http.ResponseWriter, r *http.Request)
}
