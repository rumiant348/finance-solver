package controllers

import (
	"encoding/json"
	"finance-solver.com/models"
	"finance-solver.com/rand"
	"finance-solver.com/response"
	"fmt"
	"log"
	"net/http"
)

func NewUsers(us models.UserService) *Users {
	return &Users{
		us: us,
	}
}

type Users struct {
	us models.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// New is used to render a new form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	//u.NewView.Render(w, r, nil)
}

type AuthRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create
//
//	@Summary		Signup
//	@Description	Create is used to process the signup form when a user
//	@Description	tries to create a new user account.
//	@Tags			users
//	@Accept			json
//	@Param			request	body	AuthRequestBody	true	"required object"
//	@Produce		json
//	@Success		201	{object}	response.JsonResponse	"created"
//	@Router			/signup [POST]
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var j response.JsonResponse

	var a AuthRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		j.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := &models.User{
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
	}
	err := u.us.Create(user)
	if err != nil {
		j.Error(w, err, http.StatusBadRequest)
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		j.Error(w, err, http.StatusInternalServerError)
		return
	}

	j.Success(w, "User created", http.StatusCreated)
}

// Login
//
//	@Summary		Login
//	@Description	Login is used to process the login form when a user
//	@Description	tries to log in as an existing user (via email & pw)
//	@Tags			users
//	@Accept			json
//	@Param			request	body	AuthRequestBody	true	"required object"
//	@Produce		json
//	@Success		200	{object}	response.JsonResponse	"success"
//	@Router			/login [POST]
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	var j response.JsonResponse
	var a AuthRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a)
	if err != nil {
		j.Error(w, err, http.StatusBadRequest)
		return
	}

	user, err := u.us.Authenticate(a.Email, a.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			msg := "No user exists with that email address"
			j.ErrorWithMessage(w, msg, http.StatusNotFound)
		default:
			j.Error(w, err, http.StatusInternalServerError)
		}
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		j.ErrorWithMessage(w, "Authorization error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	j.Success(w, "Authorization successful", http.StatusOK)
}

func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
	}

	cookie := http.Cookie{
		Name:  "remember_token",
		Value: user.Remember,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// CookieTest is used to display cookies on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user.Email, cookie.Value)
}
