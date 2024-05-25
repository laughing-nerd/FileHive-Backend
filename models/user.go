package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Picture   string             `json:"picture"`
	TotalSize int                `json:"totalSize"`
	Files     []File             `json:"files"`
}

type File struct {
	Name     string   `json:"name"`
	Size     int64    `json:"size"`
	Type     string   `json:"type"`
	Category string   `json:"category"`
	Link     string   `json:"link"`
	Shares   []string `json:"shares"`
}
