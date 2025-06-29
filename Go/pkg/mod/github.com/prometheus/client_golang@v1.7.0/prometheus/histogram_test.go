// Copyright 2015 The Prometheus Authors
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
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"sync"
	"testing"
	"testing/quick"
	"time"

	//lint:ignore SA1019 Need to keep deprecated package for compatibility.
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	dto "github.com/prometheus/client_model/go"
)

func benchmarkHistogramObserve(w int, b *testing.B) {
	b.StopTimer()

	wg := new(sync.WaitGroup)
	wg.Add(w)

	g := new(sync.WaitGroup)
	g.Add(1)

	s := NewHistogram(HistogramOpts{})

	for i := 0; i < w; i++ {
		go func() {
			g.Wait()

			for i := 0; i < b.N; i++ {
				s.Observe(float64(i))
			}

			wg.Done()
		}()
	}

	b.StartTimer()
	g.Done()
	wg.Wait()
}

func BenchmarkHistogramObserve1(b *testing.B) {
	benchmarkHistogramObserve(1, b)
}

func BenchmarkHistogramObserve2(b *testing.B) {
	benchmarkHistogramObserve(2, b)
}

func BenchmarkHistogramObserve4(b *testing.B) {
	benchmarkHistogramObserve(4, b)
}

func BenchmarkHistogramObserve8(b *testing.B) {
	benchmarkHistogramObserve(8, b)
}

func benchmarkHistogramWrite(w int, b *testing.B) {
	b.StopTimer()

	wg := new(sync.WaitGroup)
	wg.Add(w)

	g := new(sync.WaitGroup)
	g.Add(1)

	s := NewHistogram(HistogramOpts{})

	for i := 0; i < 1000000; i++ {
		s.Observe(float64(i))
	}

	for j := 0; j < w; j++ {
		outs := make([]dto.Metric, b.N)

		go func(o []dto.Metric) {
			g.Wait()

			for i := 0; i < b.N; i++ {
				s.Write(&o[i])
			}

			wg.Done()
		}(outs)
	}

	b.StartTimer()
	g.Done()
	wg.Wait()
}

func BenchmarkHistogramWrite1(b *testing.B) {
	benchmarkHistogramWrite(1, b)
}

func BenchmarkHistogramWrite2(b *testing.B) {
	benchmarkHistogramWrite(2, b)
}

func BenchmarkHistogramWrite4(b *testing.B) {
	benchmarkHistogramWrite(4, b)
}

func BenchmarkHistogramWrite8(b *testing.B) {
	benchmarkHistogramWrite(8, b)
}

func TestHistogramNonMonotonicBuckets(t *testing.T) {
	testCases := map[string][]float64{
		"not strictly monotonic":  {1, 2, 2, 3},
		"not monotonic at all":    {1, 2, 4, 3, 5},
		"have +Inf in the middle": {1, 2, math.Inf(+1), 3},
	}
	for name, buckets := range testCases {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Buckets %v are %s but NewHistogram did not panic.", buckets, name)
				}
			}()
			_ = NewHistogram(HistogramOpts{
				Name:    "test_histogram",
				Help:    "helpless",
				Buckets: buckets,
			})
		}()
	}
}

// Intentionally adding +Inf here to test if that case is handled correctly.
// Also, getCumulativeCounts depends on it.
var testBuckets = []float64{-2, -1, -0.5, 0, 0.5, 1, 2, math.Inf(+1)}

func TestHistogramConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	rand.Seed(42)

	it := func(n uint32) bool {
		mutations := int(n%1e4 + 1e4)
		concLevel := int(n%5 + 1)
		total := mutations * concLevel

		var start, end sync.WaitGroup
		start.Add(1)
		end.Add(concLevel)

		sum := NewHistogram(HistogramOpts{
			Name:    "test_histogram",
			Help:    "helpless",
			Buckets: testBuckets,
		})

		allVars := make([]float64, total)
		var sampleSum float64
		for i := 0; i < concLevel; i++ {
			vals := make([]float64, mutations)
			for j := 0; j < mutations; j++ {
				v := rand.NormFloat64()
				vals[j] = v
				allVars[i*mutations+j] = v
				sampleSum += v
			}

			go func(vals []float64) {
				start.Wait()
				for _, v := range vals {
					if n%2 == 0 {
						sum.Observe(v)
					} else {
						sum.(ExemplarObserver).ObserveWithExemplar(v, Labels{"foo": "bar"})
					}
				}
				end.Done()
			}(vals)
		}
		sort.Float64s(allVars)
		start.Done()
		end.Wait()

		m := &dto.Metric{}
		sum.Write(m)
		if got, want := int(*m.Histogram.SampleCount), total; got != want {
			t.Errorf("got sample count %d, want %d", got, want)
		}
		if got, want := *m.Histogram.SampleSum, sampleSum; math.Abs((got-want)/want) > 0.001 {
			t.Errorf("got sample sum %f, want %f", got, want)
		}

		wantCounts := getCumulativeCounts(allVars)
		wantBuckets := len(testBuckets)
		if !math.IsInf(m.Histogram.Bucket[len(m.Histogram.Bucket)-1].GetUpperBound(), +1) {
			wantBuckets--
		}

		if got := len(m.Histogram.Bucket); got != wantBuckets {
			t.Errorf("got %d buckets in protobuf, want %d", got, wantBuckets)
		}
		for i, wantBound := range testBuckets {
			if i == len(testBuckets)-1 {
				break // No +Inf bucket in protobuf.
			}
			if gotBound := *m.Histogram.Bucket[i].UpperBound; gotBound != wantBound {
				t.Errorf("got bound %f, want %f", gotBound, wantBound)
			}
			if gotCount, wantCount := *m.Histogram.Bucket[i].CumulativeCount, wantCounts[i]; gotCount != wantCount {
				t.Errorf("got count %d, want %d", gotCount, wantCount)
			}
		}
		return true
	}

	if err := quick.Check(it, nil); err != nil {
		t.Error(err)
	}
}

func TestHistogramVecConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	rand.Seed(42)

	it := func(n uint32) bool {
		mutations := int(n%1e4 + 1e4)
		concLevel := int(n%7 + 1)
		vecLength := int(n%3 + 1)

		var start, end sync.WaitGroup
		start.Add(1)
		end.Add(concLevel)

		his := NewHistogramVec(
			HistogramOpts{
				Name:    "test_histogram",
				Help:    "helpless",
				Buckets: []float64{-2, -1, -0.5, 0, 0.5, 1, 2, math.Inf(+1)},
			},
			[]string{"label"},
		)

		allVars := make([][]float64, vecLength)
		sampleSums := make([]float64, vecLength)
		for i := 0; i < concLevel; i++ {
			vals := make([]float64, mutations)
			picks := make([]int, mutations)
			for j := 0; j < mutations; j++ {
				v := rand.NormFloat64()
				vals[j] = v
				pick := rand.Intn(vecLength)
				picks[j] = pick
				allVars[pick] = append(allVars[pick], v)
				sampleSums[pick] += v
			}

			go func(vals []float64) {
				start.Wait()
				for i, v := range vals {
					his.WithLabelValues(string('A' + picks[i])).Observe(v)
				}
				end.Done()
			}(vals)
		}
		for _, vars := range allVars {
			sort.Float64s(vars)
		}
		start.Done()
		end.Wait()

		for i := 0; i < vecLength; i++ {
			m := &dto.Metric{}
			s := his.WithLabelValues(string('A' + i))
			s.(Histogram).Write(m)

			if got, want := len(m.Histogram.Bucket), len(testBuckets)-1; got != want {
				t.Errorf("got %d buckets in protobuf, want %d", got, want)
			}
			if got, want := int(*m.Histogram.SampleCount), len(allVars[i]); got != want {
				t.Errorf("got sample count %d, want %d", got, want)
			}
			if got, want := *m.Histogram.SampleSum, sampleSums[i]; math.Abs((got-want)/want) > 0.001 {
				t.Errorf("got sample sum %f, want %f", got, want)
			}

			wantCounts := getCumulativeCounts(allVars[i])

			for j, wantBound := range testBuckets {
				if j == len(testBuckets)-1 {
					break // No +Inf bucket in protobuf.
				}
				if gotBound := *m.Histogram.Bucket[j].UpperBound; gotBound != wantBound {
					t.Errorf("got bound %f, want %f", gotBound, wantBound)
				}
				if gotCount, wantCount := *m.Histogram.Bucket[j].CumulativeCount, wantCounts[j]; gotCount != wantCount {
					t.Errorf("got count %d, want %d", gotCount, wantCount)
				}
			}
		}
		return true
	}

	if err := quick.Check(it, nil); err != nil {
		t.Error(err)
	}
}

func getCumulativeCounts(vars []float64) []uint64 {
	counts := make([]uint64, len(testBuckets))
	for _, v := range vars {
		for i := len(testBuckets) - 1; i >= 0; i-- {
			if v > testBuckets[i] {
				break
			}
			counts[i]++
		}
	}
	return counts
}

func TestBuckets(t *testing.T) {
	got := LinearBuckets(-15, 5, 6)
	want := []float64{-15, -10, -5, 0, 5, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("linear buckets: got %v, want %v", got, want)
	}

	got = ExponentialBuckets(100, 1.2, 3)
	want = []float64{100, 120, 144}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("exponential buckets: got %v, want %v", got, want)
	}
}

func TestHistogramAtomicObserve(t *testing.T) {
	var (
		quit = make(chan struct{})
		his  = NewHistogram(HistogramOpts{
			Buckets: []float64{0.5, 10, 20},
		})
	)

	defer func() { close(quit) }()

	observe := func() {
		for {
			select {
			case <-quit:
				return
			default:
				his.Observe(1)
			}
		}
	}

	go observe()
	go observe()
	go observe()

	for i := 0; i < 100; i++ {
		m := &dto.Metric{}
		if err := his.Write(m); err != nil {
			t.Fatal("unexpected error writing histogram:", err)
		}
		h := m.GetHistogram()
		if h.GetSampleCount() != uint64(h.GetSampleSum()) ||
			h.GetSampleCount() != h.GetBucket()[1].GetCumulativeCount() ||
			h.GetSampleCount() != h.GetBucket()[2].GetCumulativeCount() {
			t.Fatalf(
				"inconsistent counts in histogram: count=%d sum=%f buckets=[%d, %d]",
				h.GetSampleCount(), h.GetSampleSum(),
				h.GetBucket()[1].GetCumulativeCount(), h.GetBucket()[2].GetCumulativeCount(),
			)
		}
		runtime.Gosched()
	}
}

func TestHistogramExemplar(t *testing.T) {
	now := time.Now()

	histogram := NewHistogram(HistogramOpts{
		Name:    "test",
		Help:    "test help",
		Buckets: []float64{1, 2, 3, 4},
	}).(*histogram)
	histogram.now = func() time.Time { return now }

	ts, err := ptypes.TimestampProto(now)
	if err != nil {
		t.Fatal(err)
	}
	expectedExemplars := []*dto.Exemplar{
		nil,
		&dto.Exemplar{
			Label: []*dto.LabelPair{
				&dto.LabelPair{Name: proto.String("id"), Value: proto.String("control")},
			},
			Value:     proto.Float64(1.6),
			Timestamp: ts,
		},
		nil,
		&dto.Exemplar{
			Label: []*dto.LabelPair{
				&dto.LabelPair{Name: proto.String("id"), Value: proto.String("3")},
			},
			Value:     proto.Float64(4),
			Timestamp: ts,
		},
		&dto.Exemplar{
			Label: []*dto.LabelPair{
				&dto.LabelPair{Name: proto.String("id"), Value: proto.String("memory")},
			},
			Value:     proto.Float64(4.5),
			Timestamp: ts,
		},
	}

	histogram.ObserveWithExemplar(1.5, Labels{"id": "1"})
	histogram.ObserveWithExemplar(1.6, Labels{"id": "control"}) // To replace exemplar in bucket 0.
	histogram.ObserveWithExemplar(4, Labels{"id": "3"})
	histogram.ObserveWithExemplar(4.5, Labels{"id": "memory"}) // Should go to +Inf bucket.

	for i, ex := range histogram.exemplars {
		var got, expected string
		if val := ex.Load(); val != nil {
			got = val.(*dto.Exemplar).String()
		}
		if expectedExemplars[i] != nil {
			expected = expectedExemplars[i].String()
		}
		if got != expected {
			t.Errorf("expected exemplar %s, got %s.", expected, got)
		}
	}
}
