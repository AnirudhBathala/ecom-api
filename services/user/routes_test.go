package user

import (
	// "bytes"
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"
	"testing"

	// "github.com/AnirudhBathala/ecom-api/models"
	// "github.com/go-chi/chi/v5"
)

func TestUserServiceHandlers(t *testing.T) {
	// userStore:=mockUserStore{}
	// handler:=NewHandler(userStore)

	// t.Run("should fail if the user payload is invalid",func(t *testing.T) {
	// 	payload:=models.RegisterUserPayload{
	// 		FirstName: "John",
	// 		LastName: "Doe",
	// 		Email: "",
	// 		Password: "john#doe",
	// 	}

	// 	marshalled,_:=json.Marshal(payload)

	// 	req,err:= http.NewRequest(http.MethodPost,"/api/v1/register",bytes.NewBuffer(marshalled))
	// 	if err!=nil {
	// 		t.Fatal(err)
	// 	}

	// 	rr:=httptest.NewRecorder()
	// 	router:=chi.NewRouter()

	// 	router.HandleFunc("/")

	// })
}

// type mockUserStore struct{}

// func (m *mockUserStore) GetUserByEmail(email string) (*models.User,error){
// 	return nil,nil
// }

// func (m *mockUserStore) GetUserByID(id int) (*models.User,error){
// 	return nil,nil
// }

// func (m *mockUserStore) CreateUser(user models.User) error{
// 	return nil
// }