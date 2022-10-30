package main

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"time"
)

type JsonResp struct {
	OK bool `json:"ok"`
}

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     VERSION,
	}

	err := app.writeJSON(w, http.StatusOK, currentStatus, "status")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type GetParser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (app *application) getOneMenu(w http.ResponseWriter, r *http.Request) {

	var parser GetParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.logger.Println("id is", parser.ID)

	menu, err := app.models.DB.Get(parser.ID)
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, menu, "menu")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) testOK(w http.ResponseWriter, r *http.Request) {

	res := JsonResp{
		OK: true,
	}

	err := app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

	res := JsonResp{
		OK: true,
	}

	err := app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type CreateParser struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
	FileString string `json:"fileString"`
}

func (app *application) createMenu(w http.ResponseWriter, r *http.Request) {

	var parser CreateParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	var request models.Menu
	request.Name = parser.Name
	request.Type = parser.Type
	request.Memo = parser.Memo
	request.FileString = parser.FileString
	request.CreatedAt = time.Now().Format("2006-01-02")
	request.UpdatedAt = time.Now().Format("2006-01-02")
	request.Opened = false

	err = app.models.DB.Create(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getAllMenu(w http.ResponseWriter, r *http.Request) {
	menu, err := app.models.DB.AllMenu()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, menu, "menu")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type UpdateParser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	CloseAt string `json:"closeAt"`
}

func (app *application) updateOpen(w http.ResponseWriter, r *http.Request) {
	var parser UpdateParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	err = app.models.DB.UpdateOpen(parser.ID, parser.Name, parser.CloseAt)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getOpenedMenu(w http.ResponseWriter, r *http.Request) {
	menu, err := app.models.DB.OpenedMenu()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, menu, "menu")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) addOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	order.UpdatedAt = time.Now().Format("2006-01-02")

	err = app.models.DB.AddOrder(order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getAllOrder(w http.ResponseWriter, r *http.Request) {
	orders, err := app.models.DB.AllOrder()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, orders, "orders")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type ID struct {
	ID int `json:"id"`
}

func (app *application) getOrderById(w http.ResponseWriter, r *http.Request) {
	var id ID
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	orders, err := app.models.DB.GetOrderById(id.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, orders, "orders")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) updateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	order.UpdatedAt = time.Now().Format("2006-01-02")

	err = app.models.DB.UpdateOrder(order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) updateMenu(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	menu.UpdatedAt = time.Now().Format("2006-01-02")

	err = app.models.DB.UpdateMenu(menu)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) deleteOpenMenu(w http.ResponseWriter, r *http.Request) {

	var parser GetParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.models.DB.DeleteOpenMenu(parser.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) deleteOrder(w http.ResponseWriter, r *http.Request) {

	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.models.DB.DeleteOrder(order.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type UpdateRating struct {
	ID    int     `json:"id"`
	Score float64 `json:"score"`
}

func (app *application) updateMenuRating(w http.ResponseWriter, r *http.Request) {
	var score UpdateRating
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.models.DB.UpdateMenuRating(score.ID, score.Score)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
