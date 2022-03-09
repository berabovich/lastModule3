package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	httpRequests "lastModule/internal/httpReqests"
	"lastModule/internal/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//_ = os.Remove(httpRequests.UserStorageFile)
	rawDataIn, _ := ioutil.ReadFile(httpRequests.UserStorageFile)
	_ = json.Unmarshal(rawDataIn, &user.Db)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	nr := chi.NewRouter()
	nr.MethodFunc("GET", "/users", httpRequests.HttpGetUsers)
	nr.MethodFunc("POST", "/create", httpRequests.HttpCreateUser)
	nr.MethodFunc("GET", "/friends/{id}", httpRequests.HttpGetUserFriends)
	nr.MethodFunc("PUT", "/{id}", httpRequests.HttpUpdateUserAge)
	nr.MethodFunc("DELETE", "/user", httpRequests.HttpDeleteUser)
	nr.MethodFunc("POST", "/make_friends", httpRequests.HttpMakeFriends)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nr))
	}()
	<-done
	//_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		rawDataOut, err := json.MarshalIndent(user.Db, "", "  ")
		if err != nil {
			log.Fatal("JSON marshaling failed:", err)
		}
		err = ioutil.WriteFile("userStorage.json", rawDataOut, 0644)
		if err != nil {
			log.Fatal("Cannot write updated file:", err)
		}
		fmt.Printf("%s file successfuly saved", httpRequests.UserStorageFile)
		//cancel()
	}()

}
