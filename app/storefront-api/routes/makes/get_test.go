package makes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockserver "github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver/mock"
	mockstore "github.com/dudleycodes/golang-microservice-structure/internal/storefront/mock"

	"github.com/gorilla/mux"
)

func TestGet(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		expectedCode     int
		storefrontResult mockstore.Result
	}{
		"found": {
			expectedCode: 200,
		},
		"not found": {
			expectedCode:     404,
			storefrontResult: mockstore.GetMakeResult(errors.New("it wasn't found")),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			srv := mockserver.New().WithStorefront(test.storefrontResult)

			r := mux.NewRouter()
			r.HandleFunc("/makes/{makeID}", Get(srv)).Methods(http.MethodGet)

			req, err := http.NewRequest(http.MethodGet, "/makes/some-id", nil)
			if err != nil {
				t.Fatalf("couldn't create test HTTP request: %s", err.Error())
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if rr.Code != test.expectedCode {
				t.Fatalf("expected status code %03d but got %03d (body: %s)", test.expectedCode, rr.Code, rr.Body)
			}
		})
	}
}
