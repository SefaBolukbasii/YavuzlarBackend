package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Challange struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Questions []Quest `json:"questions"`
}
type Quest struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
type Message struct {
	Message string `json:"message"`
}
type User struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
type USER_CHALLANGE struct {
	UserId      int `json:"user_id"`
	ChallangeId int `json:"challange_id"`
}
type ScoreTable struct {
	User      User      `json:"user"`
	Challange Challange `json:"challange"`
	Score     int       `json:"score"`
}
