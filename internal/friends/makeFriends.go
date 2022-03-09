package friends

import (
	"encoding/json"
	"lastModule/internal/user"
	"log"
	"net/http"
	"strconv"
)

type SelectUser struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

func (su *SelectUser) DeleteUser(userStorage *user.Storage, w http.ResponseWriter, r http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&su)

	for i, u := range userStorage.Users {
		for j, f := range u.Friends {
			if f == su.TargetId {
				userStorage.Users[i].Friends = append(u.Friends[:j], u.Friends[j+1:]...)
			}
		}
	}
	for index, u := range userStorage.Users {
		if u.Id == su.TargetId {
			userStorage.Users = append(userStorage.Users[:index], userStorage.Users[index+1:]...)
			_, err := json.MarshalIndent(&userStorage, "", "  ")
			if err != nil {
				log.Fatal("JSON marshaling failed:", err)
			}

			_, err = w.Write([]byte(u.Name + " was delete. Status: " + strconv.Itoa(http.StatusOK)))
			if err != nil {
				return
			}
			return
		}
	}
	_, err := w.Write([]byte("User not found"))
	if err != nil {
		return
	}
}

func (su SelectUser) AddFriends(userStorage user.Storage, w http.ResponseWriter, r http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&su)
	var name1 string
	var name2 string
	_, err := strconv.Atoi(su.TargetId)
	if err != nil {
		_, err := w.Write([]byte("ID must be int type"))
		if err != nil {
			return
		}
		return
	}
	_, err = strconv.Atoi(su.SourceId)
	if err != nil {
		_, err := w.Write([]byte("ID must be int type"))
		if err != nil {
			return
		}
		return
	}
	for _, u := range userStorage.Users {
		if u.Id == su.TargetId {
			name1 = u.Name
		}
		if u.Id == su.SourceId {
			name2 = u.Name
		}
	}
	if name1 == "" || name2 == "" {
		_, err := w.Write([]byte("Users not found"))
		if err != nil {
			return
		}
		return
	}
	for index, u := range userStorage.Users {
		if u.Id == su.TargetId {
			userStorage.Users[index].Friends = append(userStorage.Users[index].Friends, su.SourceId)
		}
		if u.Id == su.SourceId {
			userStorage.Users[index].Friends = append(userStorage.Users[index].Friends, su.TargetId)
		}
	}
	_, err = w.Write([]byte("User " + name1 + " and User " + name2 + " now friends! Status: " + strconv.Itoa(http.StatusOK)))
	if err != nil {
		return
	}
}
