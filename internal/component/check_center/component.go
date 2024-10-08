package checkCenter

import (
	"superman/internal/rpc"
	"time"

	facade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
)

// Component 启动时,检查center节点是否存活
type Component struct {
	facade.Component
}

func New() *Component {
	return &Component{}
}

func (c *Component) Name() string {
	return "check_heart"
}

func (c *Component) OnAfterInit() {
	for {
		if rpc.Ping(c.App()) {
			break
		}
		time.Sleep(2 * time.Second)
		clog.Warn("center node connect fail. retrying in 2 seconds.")
	}
}
