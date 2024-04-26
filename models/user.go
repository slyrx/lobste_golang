package models

import (
	"time"
)

type User struct {
	ID              int64      `db:"id"`
	Username        string     `db:"username"`
	Email           string     `db:"email"`
	PasswordDigest  string     `db:"password_digest"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
	IsAdmin         bool       `db:"is_admin"`
	IsModerator     bool       `db:"is_moderator"`
	Karma           int        `db:"karma"`
	InvitedByUserID *int64     `db:"invited_by_user_id"`
	BannedAt        *time.Time `db:"banned_at"`
	BannedByUserID  *int64     `db:"banned_by_user_id"`
	BannedReason    *string    `db:"banned_reason"`
}
