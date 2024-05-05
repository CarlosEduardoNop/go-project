package model

type Order struct {
	ID    string `bson:"_id,omiempty" json:"id"`
	Name  string `bson:"name" json:"name"`
	Price int    `bson:"price" json:"price"`
}