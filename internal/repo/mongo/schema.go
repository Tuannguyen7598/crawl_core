package DatabaseMongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseHttpSchema struct {
	Id                primitive.ObjectID  `bson:"_id,omitempty"`
	Sns               string              `bson:"sns"`
	Url               string              `bson:"url"`
	Body              string              `bson:"body"`
	Created_at        time.Time           `bson:"created_at"`
	Latest_updated_at *primitive.DateTime `bson:"latest_updated_at"`
	Deleted_at        *primitive.DateTime `bson:"deleted_at"`
}

type DomSchema struct {
	Id                primitive.ObjectID  `bson:"_id,omitempty"`
	Sns               *string             `bson:"sns"`
	ActionType        *string             `bson:"action_type"`
	ElementId         *string             `bson:"element_id"`
	ElementName       *string             `bson:"element_name"`
	ElementValue      *string             `bson:"element_value"`
	ElementInnerText  *string             `bson:"element_inner_text"`
	Created_at        time.Time           `bson:"created_at"`
	Latest_updated_at *primitive.DateTime `bson:"latest_updated_at"`
	Deleted_at        *primitive.DateTime `bson:"deleted_at"`
}
