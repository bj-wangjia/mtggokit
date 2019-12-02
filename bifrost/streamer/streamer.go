package streamer

import (
	"context"
	"github.com/bj-wangjia/mtggokit/bifrost/container"
)

type Streamer interface {
	SetContainer(container.Container)
	GetContainer() container.Container
	GetSchedInfo() *SchedInfo
	UpdateData(ctx context.Context) error
}
