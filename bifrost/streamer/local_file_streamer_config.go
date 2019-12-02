package streamer

import "github.com/bj-wangjia/mtggokit/bifrost/log"

type LocalFileStreamerCfg struct {
	Name       string
	Path       string
	UpdatMode  UpdatMode
	Interval   int
	IsSync     bool
	DataParser DataParser
	UserData   interface{}
	Logger     log.BiLogger
}
