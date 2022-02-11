package controller

import (
	"Go-Dispatch-Bootcamp/mocks"
	"Go-Dispatch-Bootcamp/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	user = types.User{
		Id:         1,
		Username:   "test",
		Identifier: "test",
		FirstName:  "John",
		LastName:   "Doe",
	}
	feedUsers = [][]string{
		{"test", "test", "test", "test"},
		{"test", "test", "test", "test"},
		{"test", "test", "test", "test"},
		{"test", "test", "test", "test"},
	}
)

func Test_demoController_Fetch(t *testing.T) {
	tests := []struct {
		name    string
		ucData  []types.User
		ucError error
		want    int
	}{
		{
			name:    "Fetch users. Success story.",
			ucData:  []types.User{user},
			ucError: nil,
			want:    http.StatusOK,
		},
	}

	uc := &mocks.DemoUsecase{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/fetch", nil)

			uc.On("Fetch").Return(&tt.ucData, tt.ucError)

			ct := demoController{
				usecase: uc,
			}

			ct.Fetch(rw, req)

			assert.Equal(t, tt.want, rw.Code)
		})
	}
}

func Test_demoController_FetchById(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		pathParams map[string]string
		ucData     types.User
		ucError    error
		want       int
	}{
		{
			name:       "Fetch user by id. Success story.",
			id:         1,
			pathParams: map[string]string{"id": "1"},
			ucData:     user,
			ucError:    nil,
			want:       http.StatusOK,
		},
	}

	uc := &mocks.DemoUsecase{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/fetch/:id", nil)
			req = mux.SetURLVars(req, tt.pathParams)

			uc.On("FetchById", tt.id).Return(&tt.ucData, tt.ucError)

			ct := demoController{
				usecase: uc,
			}

			ct.FetchById(rw, req)

			assert.Equal(t, tt.want, rw.Code)
		})
	}
}

func Test_demoController_Feed(t *testing.T) {
	tests := []struct {
		name    string
		ucData  [][]string
		ucError error
		want    int
	}{
		{
			name:    "Feed. Success story.",
			ucData:  feedUsers,
			ucError: nil,
			want:    http.StatusOK,
		},
	}

	uc := &mocks.DemoUsecase{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/feed", nil)

			uc.On("Feed").Return(tt.ucData, tt.ucError)

			ct := demoController{
				usecase: uc,
			}

			ct.Feed(rw, req)

			assert.Equal(t, tt.want, rw.Code)
		})
	}
}

func Test_demoController_UpdateUsersFromFeed(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		ucData  bool
		ucError error
		want    int
	}{
		{
			name:    "Update users from feed. Success story.",
			ucData:  true,
			ucError: nil,
			want:    http.StatusOK,
		},
	}

	uc := &mocks.DemoUsecase{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/update-users-from-feed", nil)

			uc.On("UpdateUsersFromFeed").Return(tt.ucData, tt.ucError)

			ct := demoController{
				usecase: uc,
			}

			ct.UpdateUsersFromFeed(rw, req)

			assert.Equal(t, tt.want, rw.Code)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		uc usecase
	}
	tests := []struct {
		name string
		args args
		want *demoController
	}{
		{
			name: "New controller test",
			args: args{
				uc: &mocks.DemoUsecase{},
			},
			want: &demoController{
				usecase: &mocks.DemoUsecase{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
