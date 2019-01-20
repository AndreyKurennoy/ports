package repository

import ports "ports/server/interfaces"

type PortRepository interface {
	Get(id string) (*ports.Port, error)
	Save(ports *ports.Ports) (*ports.Empty, error)
}
