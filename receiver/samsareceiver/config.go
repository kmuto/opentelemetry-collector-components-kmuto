package samsareceiver

import (
	"github.com/kmuto/opentelemetry-collector-components-kmuto/receiver/samsareceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	DeviceName                     string `mapstructure:"device_name"`
}

func createDefaultConfig() component.Config {
	cc := scraperhelper.NewDefaultControllerConfig()
	mbc := metadata.DefaultMetricsBuilderConfig()
	return &config{
		ControllerConfig:     cc,
		MetricsBuilderConfig: mbc,
		DeviceName:           "",
	}
}

var _ component.CreateDefaultConfigFunc = createDefaultConfig
