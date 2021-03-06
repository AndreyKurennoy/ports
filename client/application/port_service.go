package application

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"ports/client/domain/repository"
	ports "ports/client/interfaces"

	"github.com/bcicen/jstream"
)

const numPorts = 10

// PortService contains logic of ports searching and adding
type PortService interface {
	FindPort(id string) (*ports.Port, error)
	AddPorts(file *multipart.FileHeader) (*ports.Empty, error)
}

// CreatePortService creates new instance of port service
func CreatePortService(repository repository.PortRepository) PortService {
	return &PortServiceImpl{portRepository: repository}
}

// PortServiceImpl contains logic of ports searching and adding
type PortServiceImpl struct {
	portRepository repository.PortRepository
}

// FindPort searches port by id
func (p *PortServiceImpl) FindPort(id string) (*ports.Port, error) {
	return p.portRepository.Get(id)
}

// AddPorts add ports
func (p *PortServiceImpl) AddPorts(file *multipart.FileHeader) (empty *ports.Empty, err error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := src.Close(); cerr != nil && err != nil {
			err = cerr
		}
	}()

	readJSONFile(func(i *ports.Ports) {
		_, _ = p.portRepository.Save(i)
	}, src)

	return &ports.Empty{}, nil
}

//TODO: Better to change json file struct to use standard json stream package
func readJSONFile(data func(*ports.Ports), file multipart.File) {
	slice := &ports.Ports{}
	decoder := jstream.NewDecoder(file, 1).EmitKV() // extract JSON values at a depth level of 1

	for mv := range decoder.Stream() {
		value := mv.Value.(jstream.KV)
		result := serializePort(value)

		slice.Port = append(slice.Port, result)
		if len(slice.Port) == numPorts {
			data(slice)
			slice = &ports.Ports{}
		}

	}

	data(slice)
}

func serializePort(value jstream.KV) *ports.Port {
	result := &ports.Port{}

	bytes, err := json.Marshal(value.Value)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(bytes, result)
	if err != nil {
		fmt.Print(err)
	}

	result.Id = value.Key

	return result
}
