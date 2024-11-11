package data

type Store interface {
	GetPoints(id string) (int, error)
	CreatePointsEntry(points int) (id string, err error)
}
