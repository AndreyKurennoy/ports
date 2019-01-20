package application

import (
	"ports/client/domain/repository"
	ports "ports/client/interfaces"
	"mime/multipart"
	"github.com/bcicen/jstream"
	"encoding/json"
	"fmt"
)

const numPorts = 10

type PortService interface {
	FindPort(id string) (*ports.Port, error)
	AddPorts(file *multipart.FileHeader) (*ports.Empty, error)
}

func CreatePortService(repository repository.PortRepository) PortService {
	return &PortServiceImpl{portRepository: repository}
}

type PortServiceImpl struct {
	portRepository repository.PortRepository
}

func (p *PortServiceImpl) FindPort(id string) (*ports.Port, error) {
	return p.portRepository.Get(id)
}

func (p *PortServiceImpl) AddPorts(file *multipart.FileHeader) (*ports.Empty, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	p.readJsonFile(src)

	return &ports.Empty{}, nil
}

//TODO: Better to change json file struct to use standard json stream package
func (p *PortServiceImpl) readJsonFile(file multipart.File) {
	slice := []*ports.Port{}
	decoder := jstream.NewDecoder(file, 1).EmitKV() // extract JSON values at a depth level of 1

	for mv := range decoder.Stream() {
		value := mv.Value.(jstream.KV)
		result := serializePort(value)

		slice = append(slice, result)
		if len(slice) == numPorts {
			//TODO:handle errors
			p.portRepository.Save(&ports.Ports{Port: slice})
			slice = []*ports.Port{}
		}

	}
	p.portRepository.Save(&ports.Ports{Port: slice})
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
