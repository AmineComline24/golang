package migration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	req := httptest.NewRequest("GET", "/example", nil)
	req.Header.Set("Authorization", "Bearer validToken")

	rr := httptest.NewRecorder()

	authenticationMiddleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	reqWithoutToken := httptest.NewRequest("GET", "/example", nil)
	rrWithoutToken := httptest.NewRecorder()

	authenticationMiddleware(handler).ServeHTTP(rrWithoutToken, reqWithoutToken)

	assert.Equal(t, http.StatusUnauthorized, rrWithoutToken.Code)
}

func TestRoleVerificationMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	req := httptest.NewRequest("GET", "/example", nil)
	req = req.WithContext(SetUserRoleContext(req.Context(), "admin"))

	rr := httptest.NewRecorder()

	roleVerificationMiddleware("admin")(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	reqWithWrongRole := httptest.NewRequest("GET", "/example", nil)
	reqWithWrongRole = reqWithWrongRole.WithContext(SetUserRoleContext(reqWithWrongRole.Context(), "user"))

	rrWithWrongRole := httptest.NewRecorder()

	roleVerificationMiddleware("admin")(handler).ServeHTTP(rrWithWrongRole, reqWithWrongRole)

	assert.Equal(t, http.StatusForbidden, rrWithWrongRole.Code)
}
