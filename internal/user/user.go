package user

//var Users []User

type User struct {
	Id      string   `json:"id" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"Name"`
	Age     string   `json:"age" bson:"Age"`
	Friends []string `json:"friends" bson:"Friends"`
}
type Storage struct {
	Users []User
}

//var Database *mongo.Database

//func (u *User) FriendsToString() string {
//	var friend []string
//	var name string
//	for _, fr := range u.Friends {
//		for _, user := range Users {
//			if user.Id == fr {
//				name = user.Name
//			}
//		}
//		friend = append(friend, name)
//	}
//
//	friends := strings.Join(friend, ", ")
//	return friends
//}
//
//func (u *User) ToString() string {
//	return fmt.Sprintf("id %s name %s age %s friends %s", u.Id, u.Name, u.Age, u.FriendsToString())
//}
