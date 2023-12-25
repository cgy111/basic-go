package domain

import "time"

// User 领域对象，是DDD中的entity
// BO(business object)
type User struct {
	Id          int64
	Email       string
	Password    string
	Name        string
	Phone       string
	Birthday    string
	Description string
	Ctime       time.Time
}

//type Address struct {
//}

//func (u User) Encrypt() {
//
//}
//
//func (u User) ComparParssword(input string)  {
//
//}
