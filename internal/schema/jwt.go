package schema

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/opentreehole/backend/internal/model"
)

type UserClaims struct {
	jwt.RegisteredClaims

	// user id: all of `id`, `user_id`, `uid` is valid
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	UID    int `json:"uid"`

	// jwt type: access or refresh
	Type string `json:"type"`

	// user nickname
	Nickname string `json:"nickname"`

	// user joined time
	JoinedTime time.Time `json:"joined_time"`

	// whether user is admin
	IsAdmin bool `json:"is_admin"`

	// whether user has answered questions
	HasAnsweredQuestions bool `json:"has_answered_questions"`
}

func (u *UserClaims) FromModel(user *model.User) *UserClaims {
	*u = UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		ID:                   user.ID,
		UserID:               user.ID,
		UID:                  user.ID,
		Nickname:             user.Nickname,
		JoinedTime:           user.JoinedTime,
		IsAdmin:              user.IsAdmin,
		HasAnsweredQuestions: user.HasCompletedRegistrationTest,
	}
	return u
}

const (
	JWTTypeAccess  = "access"
	JWTTypeRefresh = "refresh"
)

func (u *UserClaims) ToModel() *model.User {
	var id = u.ID
	if id == 0 {
		id = u.UserID
	}
	if id == 0 {
		id = u.UID
	}
	return &model.User{
		ID:                           id,
		Nickname:                     u.Nickname,
		JoinedTime:                   u.JoinedTime,
		IsAdmin:                      u.IsAdmin,
		HasCompletedRegistrationTest: u.HasAnsweredQuestions,
	}
}
