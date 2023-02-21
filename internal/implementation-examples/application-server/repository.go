package main

import (
	"go/types"
)

type repository struct {
	dbPool *types.Nil // stub for demo purposes
}

func NewRepository() repository {
	return repository{
		dbPool: new(types.Nil),
	}
}

func (r repository) GetItem() Item {
	return Item{Name: "Item 01"}
}
