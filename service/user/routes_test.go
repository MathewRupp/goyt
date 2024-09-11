package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com.mathewrupp/goyt/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "name",
			Email:     "invalid", // Empty email should trigger the failure
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		t.Logf("Response Body: %s", rr.Body.String())
		t.Logf("Status Code: %d", rr.Code)

		// Expecting BadRequest (400) because of the invalid payload
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("should correctly register user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "name",
			Email:     "test@email.com", // Empty email should trigger the failure
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		t.Logf("Response Body: %s", rr.Body.String())
		t.Logf("Status Code: %d", rr.Code)

		// Expecting BadRequest (400) because of the invalid payload
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{ shouldReturnUser bool }

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	if email == "" {
		return nil, fmt.Errorf("user not found") // Simulate user not found when the email is empty
	}

	if m.shouldReturnUser {
		// Simulate user already exists
		return &types.User{
			Email: email,
		}, nil
	}

	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(types.User) error { return nil }
