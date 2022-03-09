package user

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
)

type User struct {
	Id      string   `json:"id" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"Name"`
	Age     string   `json:"age" bson:"Age"`
	Friends []string `json:"friends" bson:"Friends"`
}

func (us *User) AddUserId(userStorage Storage) string {
	us.Id = strconv.Itoa(len(userStorage.Users) + 1)
	for i, u := range userStorage.Users {
		if u.Id != strconv.Itoa(i+1) {
			us.Id = strconv.Itoa(i + 1)
			break
		}
	}
	return us.Id
}

type Storage struct {
	Users []User
}

var Db Storage

func (s Storage) UpdateStorage(u User) Storage {
	s.Users = append(s.Users, u)
	sort.SliceStable(s.Users, func(i, j int) bool {
		return s.Users[i].Id < s.Users[j].Id
	})
	return s
}

func (s Storage) GetFriends(params string, w http.ResponseWriter) {
	for _, u := range s.Users {
		if u.Id == params {
			id, _ := strconv.Atoi(u.Id)
			_ = json.NewEncoder(w).Encode(&s.Users[id-1].Friends)
			return
		}
	}
	_, err := w.Write([]byte("User not find"))
	if err != nil {
		return
	}
}
