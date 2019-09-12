package test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"../controllers"

// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// )

// func Router() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/attendances", controllers.GetAttendances).Methods("GET")
// 	return router
// }

// func TestAttendances(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/attendances", nil)
// 	response := httptest.NewRecorder()
// 	Router().ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "OK response is expected")
// }
