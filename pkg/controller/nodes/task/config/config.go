package config

import (
	"strings"
	"time"

	"github.com/lyft/flytestdlib/config"
	"k8s.io/apimachinery/pkg/util/sets"
)

//go:generate pflags Config --default-var defaultConfig

const SectionKey = "tasks"

var (
	defaultConfig = &Config{
		TaskPlugins:            TaskPluginConfig{EnabledPlugins: []string{}},
		MaxPluginPhaseVersions: 100000,
		BarrierConfig: BarrierConfig{
			Enabled:   true,
			CacheSize: 10000,
			CacheTTL:  config.Duration{Duration: time.Minute * 30},
		},
		BackOffConfig: BackOffConfig{
			BaseSecond:  2,
			MaxDuration: config.Duration{Duration: time.Minute * 10},
		},
		MaxErrorMessageLength: 2048,
	}

	section = config.MustRegisterSection(SectionKey, defaultConfig)
)

type Config struct {
	TaskPlugins            TaskPluginConfig `json:"task-plugins" pflag:",Task plugin configuration"`
	MaxPluginPhaseVersions int32            `json:"max-plugin-phase-versions" pflag:",Maximum number of plugin phase versions allowed for one phase."`
	BarrierConfig          BarrierConfig    `json:"barrier" pflag:",Config for Barrier implementation"`
	BackOffConfig          BackOffConfig    `json:"backoff" pflag:",Config for Exponential BackOff implementation"`
	MaxErrorMessageLength  int              `json:"maxLogMessageLength" pflag:",Max length of error message."`
}

type BarrierConfig struct {
	Enabled   bool            `json:"enabled" pflag:",Enable Barrier transitions using inmemory context"`
	CacheSize int             `json:"cache-size" pflag:",Max number of barrier to preserve in memory"`
	CacheTTL  config.Duration `json:"cache-ttl" pflag:", Max duration that a barrier would be respected if the process is not restarted. This should account for time required to store the record into persistent storage (across multiple rounds."`
}

type TaskPluginConfig struct {
	EnabledPlugins []string `json:"enabled-plugins" pflag:",Plugins enabled currently"`
}

type BackOffConfig struct {
	BaseSecond  int             `json:"base-second" pflag:",The number of seconds representing the base duration of the exponential backoff"`
	MaxDuration config.Duration `json:"max-duration" pflag:",The cap of the backoff duration"`
}

func (p TaskPluginConfig) GetEnabledPluginsSet() sets.String {
	s := sets.NewString()
	for _, e := range p.EnabledPlugins {
		cleanedPluginName := strings.Trim(e, " ")
		cleanedPluginName = strings.ToLower(cleanedPluginName)
		s.Insert(cleanedPluginName)
	}
	return s
}

func GetConfig() *Config {
	return section.GetConfig().(*Config)
}
