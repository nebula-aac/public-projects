package graph

import "github.com/nebula-aac/public-projects/graph/model"

func FindFoo(id string) (*model.Foo, error) {
	if id == "1" {
		var nameValue = "Name"
		return &model.Foo{
			ID:   "1",
			Name: &nameValue,
		}, nil
	} else {
		return nil, nil
	}
}
