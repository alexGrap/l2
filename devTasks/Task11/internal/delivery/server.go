package delivery

import (
	"L2/devTasks/Task11/config"
	models "L2/devTasks/Task11/internal"
	"L2/devTasks/Task11/internal/middleware"
	"L2/devTasks/Task11/internal/usecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Server  *mux.Router
	UseCase models.UseCase
}

func Fabric() models.Rest {
	server := Server{}
	server.Server = mux.NewRouter().StrictSlash(true)
	server.Server.HandleFunc("/create_event", middleware.MiddlewareLogger(server.Post)).Methods("POST")
	server.Server.HandleFunc("/update_event", middleware.MiddlewareLogger(server.Update)).Methods("POST")
	server.Server.HandleFunc("/delete_event", middleware.MiddlewareLogger(server.Delete)).Methods("POST")
	server.Server.HandleFunc("/events_for_day", middleware.MiddlewareLogger(server.GetForDay)).Methods("GET")
	server.Server.HandleFunc("/events_for_week", middleware.MiddlewareLogger(server.GetForWeek)).Methods("GET")
	server.Server.HandleFunc("/events_for_month", middleware.MiddlewareLogger(server.GetForMonth)).Methods("GET")
	server.UseCase = usecase.InitUseCase()
	return &server
}

func (server *Server) Hearing(config *config.Config) error {
	log.Printf("Server started on %s", config.Listen.Port)
	log.Fatal(http.ListenAndServe(config.Listen.Port, server.Server))
	return nil
}

func (server *Server) Post(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	err := server.UseCase.Create(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		err = json.NewEncoder(w).Encode("Event wasn't created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Event is was created")
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Delete(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	err := server.UseCase.Delete(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		err = json.NewEncoder(w).Encode("Event wasn't deleted")
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Event is was deleted")
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Update(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	err := server.UseCase.Update(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Event is was updated")
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) GetForDay(w http.ResponseWriter, r *http.Request) {
	result, err := server.UseCase.Get("day")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (server *Server) GetForWeek(w http.ResponseWriter, r *http.Request) {
	result, err := server.UseCase.Get("week")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (server *Server) GetForMonth(w http.ResponseWriter, r *http.Request) {
	result, err := server.UseCase.Get("month")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
