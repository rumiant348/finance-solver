package controllers

import (
	"encoding/json"
	"finance-solver.com/context"
	"finance-solver.com/models"
	"finance-solver.com/response"
	"net/http"
)

func NewLists(ls models.ListService) *Lists {
	return &Lists{
		ls: ls,
	}
}

type Lists struct {
	ls models.ListService
}

func (l *Lists) ByID(w http.ResponseWriter, r *http.Request) {
	//var j response.JsonResponse

	//lists, err := l.ls.ByID(id)
	//if err != nil {
	//	j.Error(w, err, http.StatusInternalServerError)
	//}
}

type getResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (l *Lists) GetAll(w http.ResponseWriter, r *http.Request) {
	var j response.JsonResponse

	user := context.User(r.Context())
	if user == nil {
		j.ErrorWithMessage(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	lists, err := l.ls.ByUserID(user.ID)
	if err != nil {
		j.Error(w, err, http.StatusInternalServerError)
		return
	}

	j.OkData(w, convert(lists), http.StatusOK)
}

type createRequest struct {
	Title string `json:"title"`
}

type createResponse struct {
	ID uint `json:"id"`
}

func (l *Lists) Create(w http.ResponseWriter, r *http.Request) {
	var j response.JsonResponse

	user := context.User(r.Context())
	if user == nil {
		j.ErrorWithMessage(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var lr createRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lr)
	if err != nil {
		j.Error(w, err, http.StatusBadRequest)
		return
	}

	list := &models.List{
		UserID: user.ID,
		Title:  lr.Title,
	}
	list.Title = lr.Title
	err = l.ls.Create(list)

	if err != nil {
		j.Error(w, err, http.StatusInternalServerError)
		return
	}
	cr := &createResponse{ID: list.ID}

	j.OkData(w, convertOne(cr), http.StatusCreated)
}

func convert(lists []models.List) *[]interface{} {
	res := make([]interface{}, len(lists))
	for i, list := range lists {
		res[i] = getResponse{
			ID:    list.ID,
			Title: list.Title,
		}
	}
	return &res
}

func convertOne(item interface{}) *[]interface{} {
	res := make([]interface{}, 1)
	res[0] = item
	return &res
}

func (l *Lists) Update(list *models.List) error {
	//TODO implement me
	panic("implement me")
}

func (l *Lists) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
