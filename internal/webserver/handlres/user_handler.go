package handlres

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/guiyones/brimos/internal/database"
	"github.com/guiyones/brimos/internal/dto"
	"github.com/guiyones/brimos/internal/entity"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// Get user JWT godoc
// @Summary 		Get user JWT
// @Description 	Get user JWT
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			request 	body 	dto.GetJWTInput  true  "user credentials"
// @Success 		200		{object}	dto.GetJWTOutput
// @Failure 		400		{object}	Error
// @Failure 		500		{object} 	Error
// @Router 			/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput

	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpires := r.Context().Value("jwtExpiresIn").(int)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindUserByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	if u.ValidatePassword(user.Password) {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println(u)
	//Esse any vazia é usada pois aceita o tipo de dado que vou usar
	//O encode é o payload , as informações que serão passadas para o Token
	_, tokenString, _ := jwt.Encode(map[string]any{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpires)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary 		Create user
// @Description 	Create user
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			request 	body 	dto.CreateUserInput  true  "user request"
// @Success 		201
// @Failure 		500			{object}  Error
// @Router 			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.CreateUser(u)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
