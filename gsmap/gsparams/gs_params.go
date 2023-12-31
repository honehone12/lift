package gsparams

import (
	"fmt"
	"lift/brain/portman/port"
	"time"

	libuuid "github.com/google/uuid"
)

type GSParams struct {
	index   int
	process string
	uuid    [16]byte
	address string
	port    port.Port

	monitoringTimeout time.Duration
}

func NewGSParams(
	index int,
	process string,
	uuid [16]byte,
	address string,
	port port.Port,
	monitoringTimeout time.Duration,
) *GSParams {
	return &GSParams{
		index:             index,
		process:           process,
		uuid:              uuid,
		address:           address,
		port:              port,
		monitoringTimeout: monitoringTimeout,
	}
}

func (p *GSParams) Index() int {
	return p.index
}

func (p *GSParams) ProcessName() string {
	return p.process
}

func (p *GSParams) UuidString() string {
	return libuuid.UUID(p.uuid).String()
}

func (p *GSParams) UuidRaw() []byte {
	return p.uuid[:]
}

func (p *GSParams) Port() port.Port {
	return p.port
}

func (p *GSParams) ToArgs() []string {
	return []string{
		"-a", p.address,
		"-p", p.port.String(),
		"-u", p.UuidString(),
	}
}

func (p *GSParams) NextMonitoringTimeout() time.Time {
	return time.Now().Add(p.monitoringTimeout)
}

func (p *GSParams) LogWithId(msg string) string {
	return fmt.Sprintf("GS PROCESS [%s] ", p.UuidString()) + msg
}
