package samsareceiver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/kmuto/opentelemetry-collector-components-kmuto/receiver/samsareceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type samsaScraper struct {
	cfg *config
	mb  *metadata.MetricsBuilder
}

type SamsaState struct {
	BatPct int64 `json:"batPct"`
}

func (s *samsaScraper) scrape(context.Context) (pmetric.Metrics, error) {
	ts := pcommon.NewTimestampFromTime(time.Now())

	samsaIP := os.Getenv("SAMSA_IP")
	samsaPassword := os.Getenv("SAMSA_PASSWORD")
	url := fmt.Sprintf("http://%s/api/local/info/state", samsaIP)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return s.mb.Emit(), err
	}
	req.Header.Set("Authorization", "Basic "+basicAuth("user", samsaPassword))

	resp, err := client.Do(req)
	if err != nil {
		return s.mb.Emit(), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s.mb.Emit(), err
	}

	var state SamsaState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return s.mb.Emit(), err
	}

	s.mb.RecordSamsaBatteryRemainingDataPoint(ts, state.BatPct, s.cfg.DeviceName)
	return s.mb.Emit(), nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64Encode([]byte(auth))
}

func base64Encode(b []byte) string {
	return string(bytes.TrimSpace([]byte(fmt.Sprintf("%x", b))))
}

func newSamsaScraper(cfg *config, settings receiver.Settings) *samsaScraper {
	mb := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings)
	return &samsaScraper{
		cfg: cfg,
		mb:  mb,
	}
}
