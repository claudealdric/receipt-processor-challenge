package main

type Store interface {
	GetPoints(id string) (int, error)
}
