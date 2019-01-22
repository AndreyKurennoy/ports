package application

import "ports/server/infrastructure/persistence"

// Dependencies stores all application dependencies.
type Dependencies struct {
	PortService
}

// InitDependencies initializes application dependencies.
func InitDependencies() Dependencies {
	d := Dependencies{}
	d.PortService = NewPortService(persistence.NewMapRepository())

	return d
}
