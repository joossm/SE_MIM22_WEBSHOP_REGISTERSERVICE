package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"SE_MIM22_WEBSHOP_MONO/model"
	_ "github.com/go-sql-driver/mysql"
)

const post = "POST"

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	case post:
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
				return
			}
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Id, Username, Password FROM users WHERE Username = ? AND Password = ?", user.Username, user.Password)
			errorHandler(err)
			var users []model.User
			for result.Next() {
				var user model.User
				err = result.Scan(&user.Id, &user.Username, &user.Password)
				errorHandler(err)
				users = append(users, user)
			}
			for _, iUser := range users {
				fmt.Println(user.Username + " " + user.Password)
				fmt.Println(iUser.Username + " " + iUser.Password)
				if iUser.Username == user.Username && iUser.Password == user.Password {
					responseWriter.Write([]byte("{true}"))
					return
				}
			}
			responseWriter.Write([]byte("{false}"))
			return
		}
		responseWriter.Write([]byte("{false}"))
	}
}
