package persistence

import (
	ports "ports/server/interfaces"
	"errors"
	"ports/server/domain/repository"
)

type MapRepositoryImpl struct {
	src map[string]*ports.Port
}

func NewMapRepository() repository.PortRepository {
	return &MapRepositoryImpl{src: map[string]*ports.Port{}}
}

func (r *MapRepositoryImpl) Get(id string) (*ports.Port, error) {
	if val, ok := r.src[id]; ok {
		return val, nil
	}

	return nil, errors.New("not found")
}

func (r *MapRepositoryImpl) Save(port *ports.Ports) (*ports.Empty, error) {
	for _, v := range port.Port {
		r.src[v.Id] = v
	}

	return &ports.Empty{}, nil
}
