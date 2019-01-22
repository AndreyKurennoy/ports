package application

import (
	"ports/server/domain/repository"
	ports "ports/server/interfaces"
)

// PortService contains logic of ports searching and adding
type PortService interface {
	FindPort(id string) (*ports.Port, error)
	AddPorts(ports *ports.Ports) (*ports.Empty, error)
}

// NewPortService creates new instance of port service
func NewPortService(repository repository.PortRepository) PortService {
	return &PortServiceImpl{repository: repository}
}

// PortServiceImpl contains logic of ports searching and adding
type PortServiceImpl struct {
	repository repository.PortRepository
}

// FindPort searches port by id
func (p *PortServiceImpl) FindPort(id string) (*ports.Port, error) {
	return p.repository.Get(id)
}

// AddPorts add ports
func (p *PortServiceImpl) AddPorts(ports *ports.Ports) (*ports.Empty, error) {
	return p.repository.Save(ports)
}
