package resHttpSchema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseHttpSchema struct {
	
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Sns             string `bson:"sns,omitempty"`
	Url             string 	`bson:"url,omitempty"`
	Body            string  `bson:"body,omitempty"`
	Created_at      time.Time `bson:"created_at"`
	Latest_updated_at  time.Time `bson:"latest_updated_at"`
	Deleted_at      *time.Time 	`bson:"deleted_at"`
}

