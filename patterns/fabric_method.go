package patterns

import "fmt"

const (
	ServerType           = "Server"
	PersonalComputerType = "PC"
	NotebookType         = "Notebook"
)

type ComputerPC interface {
	GetType() string
	PrintDetails()
}

func New(TypeName string) ComputerPC {
	switch TypeName {
	default:
		fmt.Printf("несуществующий тип [%s]\n", TypeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}
}

type Server struct {
	Type   string
	Core   int
	Memory int
}

func (s Server) GetType() string {
	return s.Type
}
func (s Server) PrintDetails() {
	fmt.Printf("Type: [%s] Core: [%d] Memory: [%d]\n", s.Type, s.Core, s.Memory)
}
func NewServer() ComputerPC {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 512,
	}
}

type PersonalComputer struct {
	Type   string
	Core   int
	Memory int
}

func (p PersonalComputer) GetType() string {
	return p.Type
}
func (p PersonalComputer) PrintDetails() {
	fmt.Printf("Type: [%s] Core: [%d] Memory: [%d]\n", p.Type, p.Core, p.Memory)
}
func NewPersonalComputer() ComputerPC {
	return PersonalComputer{
		Type:   PersonalComputerType,
		Core:   32,
		Memory: 1024,
	}
}

type Notebook struct {
	Type   string
	Core   int
	Memory int
}

func (n Notebook) GetType() string {
	return n.Type
}
func (n Notebook) PrintDetails() {
	fmt.Printf("Type: [%s] Core: [%d] Memory: [%d]\n", n.Type, n.Core, n.Memory)
}
func NewNotebook() ComputerPC {
	return Notebook{
		Type:   NotebookType,
		Core:   8,
		Memory: 512,
	}
}

func main() {
	var types = []string{"NewType", PersonalComputerType, ServerType, NotebookType}
	for _, i := range types {
		comp := New(i)
		if comp == nil {
			continue
		}
		comp.PrintDetails()
	}
}
