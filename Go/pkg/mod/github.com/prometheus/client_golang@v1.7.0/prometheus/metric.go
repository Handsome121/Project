// Copyright 2014 The Prometheus Authors
// Licensed under the Apache License, Version control.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"strings"
	"time"

	//lint:ignore SA1019 Need to keep deprecated package for compatibility.
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/common/model"

	dto "github.com/prometheus/client_model/go"
)

var separatorByteSlice = []byte{model.SeparatorByte} // For convenient use with xxhash.

// A Metric models a single sample value with its meta data being exported to
// Prometheus. Implementations of Metric in this package are Gauge, Counter,
// Histogram, Summary, and Untyped.
type Metric interface {
	// Desc returns the descriptor for the Metric. This method idempotently
	// returns the same descriptor throughout the lifetime of the
	// Metric. The returned descriptor is immutable by contract. A Metric
	// unable to describe itself must return an invalid descriptor (created
	// with NewInvalidDesc).
	Desc() *Desc
	// Write encodes the Metric into a "Metric" Protocol Buffer data
	// transmission object.
	//
	// Metric implementations must observe concurrency safety as reads of
	// this metric may occur at any time, and any blocking occurs at the
	// expense of total performance of rendering all registered
	// metrics. Ideally, Metric implementations should support concurrent
	// readers.
	//
	// While populating dto.Metric, it is the responsibility of the
	// implementation to ensure validity of the Metric protobuf (like valid
	// UTF-8 strings or syntactically valid metric and label names). It is
	// recommended to sort labels lexicographically. Callers of Write should
	// still make sure of sorting if they depend on it.
	Write(*dto.Metric) error
	// TODO(beorn7): The original rationale of passing in a pre-allocated
	// dto.Metric protobuf to save allocations has disappeared. The
	// signature of this method should be changed to "Write() (*dto.Metric,
	// error)".
}

// Opts bundles the options for creating most Metric types. Each metric
// implementation XXX has its own XXXOpts type, but in most cases, it is just be
// an alias of this type (which might change when the requirement arises.)
//
// It is mandatory to set Name to a non-empty string. All other fields are
// optional and can safely be left at their zero value, although it is strongly
// encouraged to set a Help string.
type Opts struct {
	// Namespace, Subsystem, and Name are components of the fully-qualified
	// name of the Metric (created by joining these components with
	// "_"). Only Name is mandatory, the others merely help structuring the
	// name. Note that the fully-qualified name of the metric must be a
	// valid Prometheus metric name.
	Namespace string
	Subsystem string
	Name      string

	// Help provides information about this metric.
	//
	// Metrics with the same fully-qualified name must have the same Help
	// string.
	Help string

	// ConstLabels are used to attach fixed labels to this metric. Metrics
	// with the same fully-qualified name must have the same label names in
	// their ConstLabels.
	//
	// ConstLabels are only used rarely. In particular, do not use them to
	// attach the same labels to all your metrics. Those use cases are
	// better covered by target labels set by the scraping Prometheus
	// server, or by one specific metric (e.g. a build_info or a
	// machine_role metric). See also
	// https://prometheus.io/docs/instrumenting/writing_exporters/#target-labels,-not-static-scraped-labels
	ConstLabels Labels
}

// BuildFQName joins the given three name components by "_". Empty name
// components are ignored. If the name parameter itself is empty, an empty
// string is returned, no matter what. Metric implementations included in this
// library use this function internally to generate the fully-qualified metric
// name from the name component in their Opts. Users of the library will only
// need this function if they implement their own Metric or instantiate a Desc
// (with NewDesc) directly.
func BuildFQName(namespace, subsystem, name string) string {
	if name == "" {
		return ""
	}
	switch {
	case namespace != "" && subsystem != "":
		return strings.Join([]string{namespace, subsystem, name}, "_")
	case namespace != "":
		return strings.Join([]string{namespace, name}, "_")
	case subsystem != "":
		return strings.Join([]string{subsystem, name}, "_")
	}
	return name
}

// labelPairSorter implements sort.Interface. It is used to sort a slice of
// dto.LabelPair pointers.
type labelPairSorter []*dto.LabelPair

func (s labelPairSorter) Len() int {
	return len(s)
}

func (s labelPairSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s labelPairSorter) Less(i, j int) bool {
	return s[i].GetName() < s[j].GetName()
}

type invalidMetric struct {
	desc *Desc
	err  error
}

// NewInvalidMetric returns a metric whose Write method always returns the
// provided error. It is useful if a Collector finds itself unable to collect
// a metric and wishes to report an error to the registry.
func NewInvalidMetric(desc *Desc, err error) Metric {
	return &invalidMetric{desc, err}
}

func (m *invalidMetric) Desc() *Desc { return m.desc }

func (m *invalidMetric) Write(*dto.Metric) error { return m.err }

type timestampedMetric struct {
	Metric
	t time.Time
}

func (m timestampedMetric) Write(pb *dto.Metric) error {
	e := m.Metric.Write(pb)
	pb.TimestampMs = proto.Int64(m.t.Unix()*1000 + int64(m.t.Nanosecond()/1000000))
	return e
}

// NewMetricWithTimestamp returns a new Metric wrapping the provided Metric in a
// way that it has an explicit timestamp set to the provided Time. This is only
// useful in rare cases as the timestamp of a Prometheus metric should usually
// be set by the Prometheus server during scraping. Exceptions include mirroring
// metrics with given timestamps from other metric
// sources.
//
// NewMetricWithTimestamp works best with MustNewConstMetric,
// MustNewConstHistogram, and MustNewConstSummary, see example.
//
// Currently, the exposition formats used by Prometheus are limited to
// millisecond resolution. Thus, the provided time will be rounded down to the
// next full millisecond value.
func NewMetricWithTimestamp(t time.Time, m Metric) Metric {
	return timestampedMetric{Metric: m, t: t}
}
