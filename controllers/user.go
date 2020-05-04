package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/models"
)

// UserController is a user controller struct.
type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.GetAll(w, r)
		case http.MethodPost:
			uc.Post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		// The following line breaks with an unhandled panic (http: panic
		// serving [::1]:64100: runtime error: index out of range [1] with length 0)
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.Get(id, w)
		case http.MethodPut:
			uc.Put(id, w, r)
		case http.MethodDelete:
			uc.Delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// GetAll returns a collection of user objects
func (uc *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	EncodeResponseAsJSON(models.GetUsers(), w)
}

// Get returns a user object by id
func (uc *UserController) Get(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	EncodeResponseAsJSON(user, w)
}

// Post creates a user object
func (uc *UserController) Post(w http.ResponseWriter, r *http.Request) {
	user, error := uc.ParseRequest(r)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	user, err := models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	EncodeResponseAsJSON(user, w)
}

// Put updates a user object
func (uc *UserController) Put(id int, w http.ResponseWriter, r *http.Request) {
	user, error := uc.ParseRequest(r)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	user, err := models.UpdateUserByID(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	EncodeResponseAsJSON(user, w)
}

// Delete removes a user object pointer from the users slice
func (uc *UserController) Delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

// ParseRequest parses an http request for a user object
func (uc *UserController) ParseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var user models.User
	err := dec.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// NewUserController returns a new user controller object
func NewUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^users/(\d+)/?`),
	}
}
