package main

import (
	"github.com/go-chi/chi"
	"lastModule/internal/httpReqests"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = os.Remove("userStorage.json")

	nr := chi.NewRouter()
	nr.MethodFunc("GET", "/users", httpRequests.HttpGetUsers)
	nr.MethodFunc("POST", "/create", httpRequests.HttpCreateUser)
	nr.MethodFunc("GET", "/friends/{id}", httpRequests.HttpGetUserFriends)
	nr.MethodFunc("PUT", "/{id}", httpRequests.HttpUpdateUserAge)
	nr.MethodFunc("DELETE", "/user", httpRequests.HttpDeleteUser)
	nr.MethodFunc("POST", "/make_friends", httpRequests.HttpMakeFriends)
	log.Fatal(http.ListenAndServe(":8080", nr))

}
