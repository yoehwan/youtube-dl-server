package core

import (
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core/src/melon"
	"github.com/youtube-dl-server/core/src/youtube"
	"github.com/youtube-dl-server/core/src/youtube_dl"
)

type Core struct {
	config       *config.Config
	youtubeDl    *youtube_dl.YoutubeDL
	melon        *melon.Melon
	youtubeChart *youtube.Youtube
}

func InitCore(c *config.Config) *Core {
	dl := youtube_dl.NewYoutubeDL(c.YoutubeDlConfig)
	m := melon.NewMelon(c.MelonConfig)
	y := youtube.NewYoutube(c.YoutubeConfig)
	return &Core{
		config:       c,
		youtubeDl:    dl,
		melon:        m,
		youtubeChart: y,
	}
}

func (c *Core) LoadAudioURL(url string) ([]byte, error) {
	return c.youtubeDl.LoadAudio(url)
}

func (c *Core) LoadConfig() interface{} {
	return c.config
}

func (c *Core) LoadMelonChart() (interface{}, error) {
	return c.melon.LoadChartList()
}

func (c *Core) LoadYoutubeChart() (interface{}, error) {
	return c.youtubeChart.LoadYoutubeChart()
}
