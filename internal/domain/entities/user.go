package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" pg:",pk,type:uuid,default:gen_random_uuid()"`
	Email    string    `json:"email" pg:",unique,notnull"`
	Password string    `json:"-" pg:",notnull"`
	BaseModel
}
