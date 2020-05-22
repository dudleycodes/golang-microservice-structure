package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/routes"

	mockserver "github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver/mock"
	mockauth "github.com/dudleycodes/golang-microservice-structure/pkg/authentication/mock"

	"github.com/gorilla/mux"
)

func Test_Authentication(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		authResult     mockauth.Result
		expectedStatus int
	}{
		"Auth Passes": {
			expectedStatus: http.StatusOK,
		},
		"Auth Fails": {

			authResult:     mockauth.ValidateJWTFail(),
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			srv := mockserver.New().WithAuthentication(test.authResult)

			req, err := http.NewRequest(http.MethodGet, "/ping", nil)
			if err != nil {
				t.Errorf("test failed while creating new HTTP request, %w", err)
			}

			r := mux.NewRouter()
			r.Use(Authentication(srv))

			r.HandleFunc("/ping", routes.Ping(srv)).Methods(http.MethodGet)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if rr.Code != test.expectedStatus {
				t.Errorf("Expected status code `%d` but got `%03d`.", test.expectedStatus, rr.Code)
			}
		})
	}
}
