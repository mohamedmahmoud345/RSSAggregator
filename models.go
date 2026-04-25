package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mohamedmahmoud345/RssAggregator/internal/database"
)

type User struct {
	ID        uuid.UUID 	`json:"id"`
	CreatedAt time.Time  	`json:"created_At"`
	UpdatedAt time.Time   	`json:"updated_At"`
	Name      string        `json:"name"`
}

func databaseUserToUser(dbUesr database.User) User {
	return User{
		ID: dbUesr.ID,
		CreatedAt: dbUesr.CreatedAt,
		UpdatedAt: dbUesr.UpdatedAt,
		Name: dbUesr.Name,	
	}
}