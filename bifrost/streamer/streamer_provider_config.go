package streamer

import "github.com/bj-wangjia/mtggokit/bifrost/log"

type StreamerProviderCfg struct {
	Name       string
	ExpireTime int64
	Logger     log.BiLogger
}
