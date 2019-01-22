package persistence

import (
	"errors"
	"ports/server/domain/repository"
	ports "ports/server/interfaces"
)

// MapRepositoryImpl is abstraction over port model data storage.
type MapRepositoryImpl struct {
	src map[string]*ports.Port
}

// NewMapRepository creates map repository for ports
func NewMapRepository() repository.PortRepository {
	return &MapRepositoryImpl{src: map[string]*ports.Port{}}
}

// Get searches for port by id
func (r *MapRepositoryImpl) Get(id string) (*ports.Port, error) {
	if val, ok := r.src[id]; ok {
		return val, nil
	}

	return nil, errors.New("not found")
}

// Save adds new ports to storage and updates existed
func (r *MapRepositoryImpl) Save(port *ports.Ports) (*ports.Empty, error) {
	for _, v := range port.Port {
		r.src[v.Id] = v
	}

	return &ports.Empty{}, nil
}
