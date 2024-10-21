package graph

import "gql-gen-errorhandling/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type CatInternal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DogInternal struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

var Cats = []CatInternal{
	{
		ID:   "1",
		Name: "Cat 1",
		Age:  1,
	},
	{
		ID:   "2",
		Name: "Cat 2",
		Age:  2,
	},
}

var Dogs = []DogInternal{
	{
		ID:    "1",
		Name:  "Dog 1",
		Breed: "Breed 1",
	},
	{
		ID:    "2",
		Name:  "Dog 2",
		Breed: "Breed 2",
	},
}

func TransformDog(dog DogInternal) *model.Dog {
	return &model.Dog{
		ID:    dog.ID,
		Name:  dog.Name,
		Breed: dog.Breed,
	}
}

func TransformCat(cat CatInternal) *model.Cat {
	return &model.Cat{
		ID:   cat.ID,
		Name: cat.Name,
		Age:  cat.Age,
	}
}

func TransformDogs(dogs []DogInternal) []*model.Dog {
	var result []*model.Dog
	for _, dog := range dogs {
		result = append(result, TransformDog(dog))
	}
	return result
}

func TransformCats(cats []CatInternal) []*model.Cat {
	var result []*model.Cat
	for _, cat := range cats {
		result = append(result, TransformCat(cat))
	}
	return result
}

type Resolver struct {
	CatsData []CatInternal
	DogsData []DogInternal
}
