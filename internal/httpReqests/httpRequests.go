package httpRequests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"lastModule/internal/makeFriends"
	"lastModule/internal/updateUser"
	"lastModule/internal/user"
	"log"
	"net/http"
	"sort"
	"strconv"
)

const userStorageFile = "userStorage.json"

func HttpMakeFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawDataIn, err := ioutil.ReadFile(userStorageFile)
	if err != nil {
		_, _ = w.Write([]byte("No users in Storage"))
	}
	var userStorage user.Storage
	_ = json.Unmarshal(rawDataIn, &userStorage)

	var makeFriend makeFriends.MakeFriends
	_ = json.NewDecoder(r.Body).Decode(&makeFriend)
	var name1 string
	var name2 string
	_, err = strconv.Atoi(makeFriend.TargetId)
	if err != nil {
		_, err := w.Write([]byte("ID must be int type"))
		if err != nil {
			return
		}
		return
	}
	_, err = strconv.Atoi(makeFriend.SourceId)
	if err != nil {
		_, err := w.Write([]byte("ID must be int type"))
		if err != nil {
			return
		}
		return
	}
	for _, u := range userStorage.Users {
		if u.Id == makeFriend.TargetId {
			name1 = u.Name
		}
		if u.Id == makeFriend.SourceId {
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
		if u.Id == makeFriend.TargetId {
			userStorage.Users[index].Friends = append(userStorage.Users[index].Friends, makeFriend.SourceId)
		}
		if u.Id == makeFriend.SourceId {
			userStorage.Users[index].Friends = append(userStorage.Users[index].Friends, makeFriend.TargetId)
		}
	}
	rawDataOut, err := json.MarshalIndent(&userStorage, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(userStorageFile, rawDataOut, 0644)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}

	_, err = w.Write([]byte("User " + name1 + " and User " + name2 + " now friends! Status: " + strconv.Itoa(http.StatusOK)))
	if err != nil {
		return
	}
}

func HttpGetUsers(w http.ResponseWriter, _ *http.Request) {
	rawDataIn, err := ioutil.ReadFile(userStorageFile)
	if err != nil {
		_, _ = w.Write([]byte("No users in Storage"))
	}
	var userStorage user.Storage
	_ = json.Unmarshal(rawDataIn, &userStorage)
	_ = json.NewEncoder(w).Encode(userStorage)

	//var response string
	//for _, u := range user.Users {
	//	response += u.ToString() + "\n"
	//}
	//_ = json.NewEncoder(w).Encode(user.Users)
	//_, err = w.Write([]byte(response))
	//if err != nil {
	//	return
	//}

}

func HttpCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var us user.User
	_ = json.NewDecoder(r.Body).Decode(&us)

	rawDataIn, _ := ioutil.ReadFile(userStorageFile)

	var userStorage user.Storage
	_ = json.Unmarshal(rawDataIn, &userStorage)

	_, err := strconv.Atoi(us.Age)
	if err != nil {
		_, err := w.Write([]byte("Age must be int type"))
		if err != nil {
			return
		}
		return
	}
	us.Id = strconv.Itoa(len(userStorage.Users) + 1)
	for i, u := range userStorage.Users {
		if u.Id != strconv.Itoa(i+1) {
			us.Id = strconv.Itoa(i + 1)
			break
		}
		if u.Id == u.Id {
			id, _ := strconv.Atoi(u.Id)
			u.Id = strconv.Itoa(id + 1)
		}
	}

	userStorage.Users = append(userStorage.Users, us)
	sort.SliceStable(userStorage.Users, func(i, j int) bool {
		return userStorage.Users[i].Id < userStorage.Users[j].Id
	})
	rawDataOut, err := json.MarshalIndent(&userStorage, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(userStorageFile, rawDataOut, 0644)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}

	w.WriteHeader(http.StatusCreated)

	//_, err = w.Write([]byte("User ID: " + u.Id + " Status:" + strconv.Itoa(http.StatusCreated)))
	//if err != nil {
	//	return
	//}
	//id, _ := strconv.Atoi(u.Id)
	_ = json.NewEncoder(w).Encode(us.Id)
}

func HttpGetUserFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	rawDataIn, err := ioutil.ReadFile(userStorageFile)
	if err != nil {
		_, _ = w.Write([]byte("No users in Storage"))
	}
	var userStorage user.Storage
	_ = json.Unmarshal(rawDataIn, &userStorage)
	for _, u := range userStorage.Users {
		if u.Id == params {
			id, _ := strconv.Atoi(u.Id)
			_ = json.NewEncoder(w).Encode(&userStorage.Users[id-1].Friends)
			return
		}
	}
	_, err = w.Write([]byte("User not find"))
	if err != nil {
		return
	}
}

func HttpUpdateUserAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawDataIn, err := ioutil.ReadFile(userStorageFile)
	if err != nil {
		_, _ = w.Write([]byte("No users in Storage"))
	}
	var userStorage user.Storage
	_ = json.Unmarshal(rawDataIn, &userStorage)
	var updateAge updateUser.UpdateUser
	_ = json.NewDecoder(r.Body).Decode(&updateAge)
	params := chi.URLParam(r, "id")
	_, err = strconv.Atoi(updateAge.NewAge)

	if err != nil {
		_, err := w.Write([]byte("Age must be int type"))
		if err != nil {
			return
		}
		return
	}
	for index, item := range userStorage.Users {
		if item.Id == params {
			userStorage.Users[index].Age = updateAge.NewAge
			rawDataOut, err := json.MarshalIndent(&userStorage, "", "  ")
			if err != nil {
				log.Fatal("JSON marshaling failed:", err)
			}
			err = ioutil.WriteFile(userStorageFile, rawDataOut, 0644)
			if err != nil {
				log.Fatal("Cannot write updated settings file:", err)
			}
			_, err = w.Write([]byte("User " + item.Name + ". Age update successful! Status: " + strconv.Itoa(http.StatusOK)))
			if err != nil {
				return
			}
			return
		}
	}
	_, err = w.Write([]byte("User not found"))
	if err != nil {
		return
	}
}

func HttpDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawDataIn, err := ioutil.ReadFile(userStorageFile)
	if err != nil {
		_, _ = w.Write([]byte("No users in Storage"))
	}
	var userStorage user.Storage
	_ = json.Unmarshal(rawDataIn, &userStorage)

	var makeFriend makeFriends.MakeFriends
	_ = json.NewDecoder(r.Body).Decode(&makeFriend)

	for i, u := range userStorage.Users {
		for j, f := range u.Friends {
			if f == makeFriend.TargetId {
				userStorage.Users[i].Friends = append(u.Friends[:j], u.Friends[j+1:]...)
			}
		}
	}
	for index, u := range userStorage.Users {
		if u.Id == makeFriend.TargetId {
			userStorage.Users = append(userStorage.Users[:index], userStorage.Users[index+1:]...)
			rawDataOut, err := json.MarshalIndent(&userStorage, "", "  ")
			if err != nil {
				log.Fatal("JSON marshaling failed:", err)
			}
			err = ioutil.WriteFile(userStorageFile, rawDataOut, 0644)
			if err != nil {
				log.Fatal("Cannot write updated settings file:", err)
			}
			_, err = w.Write([]byte(u.Name + " was delete. Status: " + strconv.Itoa(http.StatusOK)))
			if err != nil {
				return
			}
			return
		}
	}
	_, err = w.Write([]byte("User not found"))
	if err != nil {
		return
	}
}
