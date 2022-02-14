package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	FirstName string             `json:"firstname,omitempty" validate:"required"`
	LastName  string             `json:"lastname,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"required"`
	Password  string             `json:"password,omitempty" validate:"required"`
}

type AdminControl struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	CompanyName       string             `json:"companyname,omitempty" validate:"required"`
	TIN               string             `json:"tin,omitempty" validate:"required"`
	NumberOfEmployees int64              `json:"numberofemployees,omitempty" validate:"required"`
	Subscription      string             `json:"Subscription,omitempty" validate:"required"`
	FreeTrail         string             `json:"freetrail ,omitempty"`
	Address           string             `json:"Address,omitempty" validate:"required"`
	ContactNUmber     int64              `json:"contactnumber,omitempty" validate:"required"`
}
