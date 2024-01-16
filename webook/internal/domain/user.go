package domain

import "time"

// User 领域对象，是DDD中的entity
// BO(business object)
type User struct {
	Id       int64
	Email    string
	Password string
	Name     string
	Phone    string
	//UnionID  string
	//OpenID      string
	Birthday    string
	Description string
	//不用组合，万一将来可能还有DingDingInfo,里面有同名字段 UnionID
	WechatInfo WechatInfo
	Ctime      time.Time
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
