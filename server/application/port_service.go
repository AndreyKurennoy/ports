package application

import (
	"ports/server/domain/repository"
	ports "ports/server/interfaces"
)

type PortService interface {
	FindPort(id string) (*ports.Port, error)
	AddPorts(ports *ports.Ports) (*ports.Empty, error)
}

func NewPortService(repository repository.PortRepository) PortService {
	return &PortServiceImpl{repository: repository}
}

type PortServiceImpl struct {
	repository repository.PortRepository
}

func (p *PortServiceImpl) FindPort(id string) (*ports.Port, error) {
	return p.repository.Get(id)
}

func (p *PortServiceImpl) AddPorts(ports *ports.Ports) (*ports.Empty, error) {
	return p.repository.Save(ports)
}
