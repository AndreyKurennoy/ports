package repository

import ports "ports/server/interfaces"

// PortRepository is abstraction over port model data storage.
type PortRepository interface {
	Get(id string) (*ports.Port, error)
	Save(ports *ports.Ports) (*ports.Empty, error)
}
