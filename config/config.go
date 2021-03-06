package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	YoutubeDlConfig *YoutubeDlConfig `json:"youtube_dl_config"`
	ApiConfig       *ApiConfig       `json:"api_config"`
	ViewConfig      *ViewConfig      `json:"view_config"`
	MelonConfig     *MelonConfig     `json:"melon_config"`
	LoggerConfig    *LoggerConfig    `json:"logger_config"`
	YoutubeConfig   *YoutubeConfig   `json:"youtube_config"`
}

type YoutubeDlConfig struct {
	AudioFormat  string `json:"audio_format"`
	AudioQuality int    `json:"audio_quality"`
}

type ApiConfig struct {
	Port       string `json:"port"`
	Version    string `json:"version"`
	ConfigApi  string `json:"config_api"`
	AudioApi   string `json:"audio_api"`
	MelonApi   string `json:"melon_api"`
	YoutubeApi string `json:"youtube_api"`
}

type ViewConfig struct {
	Path string `json:"path"`
}

type MelonConfig struct {
	Top     int `json:"top"`
	Ballade int `json:"ballade"`
	Dance   int `json:"dance"`
	Hiphop  int `json:"hiphop"`
	Rnb     int `json:"rnb"`
	Indie   int `json:"indie"`
	Rock    int `json:"rock"`
	Trot    int `json:"trot"`
	Folk    int `json:"folk"`
}

type YoutubeConfig struct {
	Top int `json:"top"`
}

type LoggerConfig struct {
	Path string `json:"path"`
}

func NewConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	viper.SetConfigType("yaml")

	err = viper.ReadConfig(file)
	if err != nil {
		log.Panicln(err)
	}
	return &Config{
		YoutubeDlConfig: &YoutubeDlConfig{
			AudioFormat:  viper.GetString("youtube_dl.audio_format"),
			AudioQuality: viper.GetInt("youtube_dl.audio_quality"),
		},
		ApiConfig: &ApiConfig{
			Port:       viper.GetString("api.port"),
			Version:    viper.GetString("api.version"),
			ConfigApi:  viper.GetString("api.config_api"),
			AudioApi:   viper.GetString("api.audio_api"),
			MelonApi:   viper.GetString("api.melon_api"),
			YoutubeApi: viper.GetString("api.youtube_api"),
		},
		ViewConfig: &ViewConfig{
			Path: viper.GetString("view.path"),
		},
		MelonConfig: &MelonConfig{
			Top:     viper.GetInt("melon_chart.top"),
			Ballade: viper.GetInt("melon_chart.ballade"),
			Dance:   viper.GetInt("melon_chart.dance"),
			Hiphop:  viper.GetInt("melon_chart.hiphop"),
			Rnb:     viper.GetInt("melon_chart.rnb"),
			Indie:   viper.GetInt("melon_chart.indie"),
			Rock:    viper.GetInt("melon_chart.rock"),
			Trot:    viper.GetInt("melon_chart.trot"),
			Folk:    viper.GetInt("melon_chart.folk"),
		},
		LoggerConfig: &LoggerConfig{
			Path: viper.GetString("logger.path"),
		},
		YoutubeConfig: &YoutubeConfig{
			Top: viper.GetInt("youtube_chart.top"),
		},
	}
}
