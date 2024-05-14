package user

import (
	"fmt"
	"net/http"

	"github.com/AnirudhBathala/ecom-api/models"
	"github.com/AnirudhBathala/ecom-api/services/auth"
	"github.com/AnirudhBathala/ecom-api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Store models.UserStore
}

func (h *Handler) RigesterRoutes(router chi.Router) {
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello to Login"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// check for the payload in request
	var user models.RegisterUserPayload

	if err:=utils.ParseJSON(r,&user); err!=nil {
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("error while parsing json: %v",err.Error()))
		return
	}


	// validate the payload
	if err:=utils.Validate.Struct(user); err!=nil{
		errors:=err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid payload: %v",errors))
		return
	}


	//check if the user exist
	_,err:=h.Store.GetUserByEmail(user.Email)
	if err!=nil {
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user already exists with this email: %s",user.Email))
		return 
	}
	
	hashedPassword,err:=auth.HashPassword(user.Password)
	if err!=nil {
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	err=h.Store.CreateUser(models.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: hashedPassword,
	})

	if err!=nil {
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	utils.WriteJSON(w,http.StatusCreated,"user created sucessfully")
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{
		Store: store,
	}
}
