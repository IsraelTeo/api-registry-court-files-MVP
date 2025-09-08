package route

import (
	"net/http"

	"github.com/IsraelTeo/api-registry-court-files-MVP/database"
	"github.com/IsraelTeo/api-registry-court-files-MVP/handler"
	"github.com/IsraelTeo/api-registry-court-files-MVP/repository"
	"github.com/IsraelTeo/api-registry-court-files-MVP/service"
	"github.com/gorilla/mux"
)

const (
	idPath   = "/{id}"
	voidPath = ""
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	setUpJudicialFile(api)

	return r
}

func setUpJudicialFile(api *mux.Router) {
	// Repositorio
	judicialFileRepo := repository.NewJudicialFileRepository(database.GDB)
	personRepo := repository.NewPersonRepository(database.GDB)
	lawyerRepo := repository.NewLawyerRepository(database.GDB)
	courtRepo := repository.NewCourtRepository(database.GDB)

	// Servicio
	judicialFileService := service.NewJudicialFileService(judicialFileRepo, personRepo, lawyerRepo, courtRepo)

	// Handler
	judicialFileHandler := handler.NewJudicialFileHandler(judicialFileService)

	files := api.PathPrefix("/judicial-files").Subrouter()

	files.HandleFunc(voidPath, judicialFileHandler.GetAll).Methods(http.MethodGet)
	files.HandleFunc(idPath, judicialFileHandler.GetByID).Methods(http.MethodGet)

	files.HandleFunc(voidPath, judicialFileHandler.Create).Methods(http.MethodPost)

	files.HandleFunc(idPath, judicialFileHandler.Update).Methods(http.MethodPut)

	files.HandleFunc(idPath, judicialFileHandler.Delete).Methods(http.MethodDelete)

	//files.HandleFunc("/{id}/persons/{personId}", judicialFileHandler.AddPerson).Methods(http.MethodPost)
	//files.HandleFunc("/{id}/lawyers/{lawyerId}", judicialFileHandler.AddLawyer).Methods(http.MethodPost)
}
