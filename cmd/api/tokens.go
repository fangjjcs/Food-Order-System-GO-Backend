package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialsEasy struct {
	Username string `json:"username"`
}

func (app *application) Signin(w http.ResponseWriter, r *http.Request) {

	var creds CredentialsEasy

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusForbidden)
		return
	}

	user, err := app.models.DB.CheckUserWithNumber(creds.Username)
	if err != nil {
		app.errorJSON(w, errors.New("employee number is wrong"), http.StatusForbidden)
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(user.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJSON(w, errors.New("error signing"), http.StatusForbidden)
		return
	}

	type Response struct {
		EmployeeID string `json:"employeeId"`
		Username   string `json:"username"`
		JWT        string `json:"jwt"`
	}
	res := Response{
		EmployeeID: creds.Username,
		Username:   user.Username,
		JWT:        string(jwtBytes),
	}

	app.writeJSON(w, http.StatusOK, res, "response")

}
