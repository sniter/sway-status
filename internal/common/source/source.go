package source

import (
	"os"
	"os/exec"
)

type Source interface {
	ReadBytes() ([]byte, error)
	ReadString() (string, error)
}

type Decorator[S any] interface {
	Render(input []byte) S
}

type File struct {
	FileName string
}

func (f File) ReadBytes() ([]byte, error) {
	return os.ReadFile(f.FileName)
}

func (f File) ReadString() (string, error) {
	bytes, err := os.ReadFile(f.FileName)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

type Run struct {
	Script string
}

func (r Run) ReadBytes() ([]byte, error) {
	cmd := exec.Command("bash", "-c", r.Script)
	return cmd.Output()
}

func (r Run) ReadString() (string, error) {
	bytes, err := r.ReadBytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
