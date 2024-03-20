package memory

import (
	"context"
	"fmt"
	"os"
)

const basePath = "framework/output"

type Memory struct{}

type Result struct {
	Link []string
	Body string
}

type SaveRequest struct {
	Body     string
	Domain   string
	FileName string
}

type IMemory interface {
	Save(context.Context, SaveRequest) error
}

func NewMemory() IMemory {
	return &Memory{}
}

func (m *Memory) Save(ctx context.Context, req SaveRequest) error {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", basePath, req.Domain), 0755); err != nil {
		return fmt.Errorf("entity.http.Save: %v", err)
	}

	if err := os.WriteFile(m.generatePath(req), []byte(req.Body), 0644); err != nil {
		return fmt.Errorf("entity.http.Save: %v", err)
	}
	return nil
}

func (m *Memory) generatePath(req SaveRequest) string {
	return fmt.Sprintf("%s/%s/%s.html", basePath, req.Domain, req.FileName)
}
