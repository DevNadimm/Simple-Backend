package hospital

import (
	"net/http"
	"test/rest/middleware"
)

func (handler *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Public APIs - we apply standard middleware like Cors, Logger etc. via manager
	mux.Handle("POST /api/doctors/search", manager.With(http.HandlerFunc(handler.SearchDoctors)))
	mux.Handle("POST /api/appointments/availability", manager.With(http.HandlerFunc(handler.CheckAvailability)))
	mux.Handle("POST /api/appointments/book", manager.With(http.HandlerFunc(handler.BookAppointment)))
	mux.Handle("POST /api/cost/estimate", manager.With(http.HandlerFunc(handler.CheckTreatmentCost)))
}
