// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"net/url"
	"os"
	"time"

	"github.com/gardener/logging/fluent-bit-to-loki/pkg/config"

	"github.com/cortexproject/cortex/pkg/util"
	"github.com/cortexproject/cortex/pkg/util/flagext"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/grafana/loki/pkg/promtail/client"
	lokiflag "github.com/grafana/loki/pkg/util/flagext"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/prometheus/common/model"
	"github.com/weaveworks/common/logging"
)

var _ = Describe("Client", func() {
	defaultURL, _ := parseURL("http://localhost:3100/loki/api/v1/push")
	var infoLogLevel logging.Level
	_ = infoLogLevel.Set("info")
	conf := &config.Config{
		LineFormat: config.KvPairFormat,
		ClientConfig: client.Config{
			URL:            defaultURL,
			TenantID:       "", // empty as not set in fluent-bit plugin config map
			BatchSize:      100,
			BatchWait:      30 * time.Second,
			ExternalLabels: lokiflag.LabelSet{LabelSet: model.LabelSet{"app": "foo"}},
			BackoffConfig: util.BackoffConfig{
				MinBackoff: (1 * time.Second),
				MaxBackoff: 300 * time.Second,
				MaxRetries: 10,
			},
			Timeout: 10 * time.Second,
		},
		BufferConfig: config.BufferConfig{
			Buffer:     false,
			BufferType: config.DefaultBufferConfig.BufferType,
			DqueConfig: config.DqueConfig{
				QueueDir:         config.DefaultDqueConfig.QueueDir,
				QueueSegmentSize: config.DefaultDqueConfig.QueueSegmentSize,
				QueueSync:        config.DefaultDqueConfig.QueueSync,
				QueueName:        config.DefaultDqueConfig.QueueName,
			},
		},
		LogLevel:      infoLogLevel,
		LabelKeys:     []string{"foo", "bar"},
		RemoveKeys:    []string{"buzz", "fuzz"},
		DropSingleKey: false,
		DynamicHostPath: map[string]interface{}{
			"kubernetes": map[string]interface{}{
				"namespace_name": "namespace",
			},
		},
		DynamicHostPrefix: "http://loki.",
		DynamicHostSuffix: ".svc:3100/loki/api/v1/push",
		DynamicHostRegex:  "shoot--",
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, infoLogLevel.Gokit)

	Describe("NewClient", func() {
		It("should create a client", func() {
			c, err := NewClient(conf, logger)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).ToNot(BeNil())
		})
	})

	Describe("newTimestampOrderingClient", func() {
		It("should create a client", func() {
			c, err := newTimestampOrderingClient(conf.ClientConfig, logger)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).ToNot(BeNil())
		})
	})
})

func parseURL(u string) (flagext.URLValue, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return flagext.URLValue{}, err
	}
	return flagext.URLValue{URL: parsed}, nil
}
