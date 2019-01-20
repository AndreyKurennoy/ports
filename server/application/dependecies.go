package application

import (
	"ports/server/infrastructure/persistence"
)

type Dependencies struct {
	PortService
}

func InitDependencies() Dependencies {
	d := Dependencies{}
	d.PortService = NewPortService(persistence.NewMapRepository())

	return d
}
