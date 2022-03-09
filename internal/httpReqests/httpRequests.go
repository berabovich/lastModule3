package httpRequests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"lastModule/internal/friends"
	"lastModule/internal/updateUser"
	"lastModule/internal/user"
	"net/http"
	"strconv"
)

const UserStorageFile = "userStorage.json"

func HttpMakeFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var makeFriend friends.SelectUser
	makeFriend.AddFriends(user.Db, w, *r)

}

func HttpGetUsers(w http.ResponseWriter, _ *http.Request) {

	_ = json.NewEncoder(w).Encode(user.Db)

}

func HttpCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u user.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	_, err := strconv.Atoi(u.Age)
	if err != nil {
		_, err := w.Write([]byte("Age must be int type"))
		if err != nil {
			return
		}
		return
	}

	u.Id = u.AddUserId(user.Db)

	user.Db = user.Storage.UpdateStorage(user.Db, u)

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(u.Id)
}

func HttpGetUserFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	user.Storage.GetFriends(user.Db, params, w)

}

func HttpUpdateUserAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateAge updateUser.Update
	(*updateUser.Update).UpdateAge(&updateAge, user.Db, w, *r)

}

func HttpDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var del friends.SelectUser
	del.DeleteUser(&user.Db, w, *r)

}
