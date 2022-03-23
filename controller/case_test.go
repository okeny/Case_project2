 package controller

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"project-cases/repository"
// 	"project-cases/service"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	caseRepo       repository.CaseRepository = repository.NewMemoryRepository()
// 	caseServ       service.CaseService       = service.NewCaseService(caseRepo)
// 	caseController CaseController            = NewCaseCrontroller(caseServ)
// )

// func TestGetCases(t *testing.T) {
// 	router := routing.SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/cases/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "item deleted", w.Body.String())
// }

// func TestDeleteCase(t *testing.T) {

//     router := routing.SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/cases/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "item deleted", w.Body.String())
// }

// func TestGetCase(t *testing.T) {

//     router := router.SetupRouter()
// 	//create new http POST request
// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("GET", "/cases",nil)
// 	router.ServeHTTP(w, req)

// 	//Assertion
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestPostCase(t *testing.T) {

// 	router := router.SetupRouter()
// 	//create new http POST request
// 	var json = []byte(`{"title":"title1","description":"text 1"}`)

// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("POST", "/cases", bytes.NewBuffer(json))
// 	router.ServeHTTP(w, req)

// 	//Assertion
// 	assert.Equal(t, http.StatusOK, w.Code)

// }

// func TestPutCase(t *testing.T) {

// 	router := router.SetupRouter()
// 	//create new http POST request
// 	var json = []byte(`{"title":"title1","description":"text 1"}`)

// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("PUT", "/cases", bytes.NewBuffer(json))
// 	router.ServeHTTP(w, req)

// 	//Assertion
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "item deleted", w.Body.String())

// }
