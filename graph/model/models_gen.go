// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cat struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Meow string `json:"meow"`
}

type Dog struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Bark  string `json:"bark"`
}

type Query struct {
}
