package internal

import (
	"context"
	"fmt"
	"runtime/metrics"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/semconv/v1.37.0/goconv"
)

// RegisterMemoryUsedMetric registers the metric for getting process memory
// utilisation.
func RegisterMemoryUsedMetric(meter metric.Meter) error {
	_, err := goconv.NewMemoryUsed(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				memoryTotalMetric := "/memory/classes/total:bytes"
				memoryReleasedMetric := "/memory/classes/heap/released:bytes"

				sample := make([]metrics.Sample, 2)

				sample[0].Name = memoryTotalMetric
				sample[1].Name = memoryReleasedMetric

				metrics.Read(sample)

				for _, s := range sample {
					if s.Value.Kind() == metrics.KindBad {
						return fmt.Errorf("bad sample metric: %s", s.Name)
					}
				}

				memoryTotal := sample[0].Value.Uint64()
				memoryReleased := sample[1].Value.Uint64()

				io.Observe(int64(memoryTotal - memoryReleased))

				return nil
			},
		),
	)

	return err
}

// RegisterMemoryLimitMetric registers the metric for collecting memory limit
// from the Go runtime.
func RegisterMemoryLimitMetric(meter metric.Meter) error {
	_, err := goconv.NewMemoryLimit(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				memLimitMetric := "/gc/gomemlimit:bytes"

				memLimit, err := readSingleMetricUint(memLimitMetric)
				if err != nil {
					return err
				}

				io.Observe(int64(memLimit))

				return nil
			},
		),
	)

	return err
}

// RegisterMemoryAllocatedMetric registers the metric to collect memory
// allocations made to the process.
func RegisterMemoryAllocatedMetric(meter metric.Meter) error {
	_, err := goconv.NewMemoryAllocated(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				memAllocMetric := "/gc/heap/allocs:bytes"

				memAlloc, err := readSingleMetricUint(memAllocMetric)
				if err != nil {
					return err
				}

				io.Observe(int64(memAlloc))

				return nil
			},
		),
	)

	return err
}

// RegisterMemoryAllocationsMetric registers the metric to count the number of
// memory allocations made to the process.
func RegisterMemoryAllocationsMetric(meter metric.Meter) error {
	_, err := goconv.NewMemoryAllocations(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				memAllocsMetric := "/gc/heap/allocs:objects"

				memAllocs, err := readSingleMetricUint(memAllocsMetric)
				if err != nil {
					return err
				}

				io.Observe(int64(memAllocs))

				return nil
			},
		),
	)

	return err
}

// RegisterGCGoalMetric registers the metric for measuring the GC goal.
func RegisterGCGoalMetric(meter metric.Meter) error {
	_, err := goconv.NewMemoryGCGoal(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				gcGoalMetric := "/gc/heap/goal:bytes"

				gcGoal, err := readSingleMetricUint(gcGoalMetric)
				if err != nil {
					return err
				}

				io.Observe(int64(gcGoal))

				return nil
			},
		),
	)

	return err
}

// RegisterGoRoutineCountMetric registers the metric to report the number of
// goroutines.
func RegisterGoRoutineCountMetric(meter metric.Meter) error {
	_, err := goconv.NewGoroutineCount(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				goRoutineMetric := "/sched/goroutines:goroutines"

				goRoutine, err := readSingleMetricUint(goRoutineMetric)
				if err != nil {
					return err
				}

				io.Observe(int64(goRoutine))

				return nil
			},
		),
	)

	return err
}

// RegisterProcessorLimitMetric registers the metric to report the maximum
// number of OS threads that can be spawned.
func RegisterProcessorLimitMetric(meter metric.Meter) error {
	_, err := goconv.NewProcessorLimit(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				processorLimitMetric := "/sched/gomaxprocs:threads"

				processorLimit, err := readSingleMetricUint(
					processorLimitMetric,
				)
				if err != nil {
					return err
				}

				io.Observe(int64(processorLimit))

				return nil
			},
		),
	)

	return err
}

// RegisterGoGCConfigMetric exports the heap size target metric.
func RegisterGoGCConfigMetric(meter metric.Meter) error {
	_, err := goconv.NewConfigGogc(
		meter,
		metric.WithInt64Callback(
			func(ctx context.Context, io metric.Int64Observer) error {
				goGCConfigMetric := "/gc/gogc:percent"

				goGCConfig, err := readSingleMetricUint(goGCConfigMetric)
				if err != nil {
					return nil
				}

				io.Observe(int64(goGCConfig))

				return nil
			},
		),
	)

	return err
}

// readSingleMetricUint reads the metric with the provided name and returns its
// value as a uint64 type.
func readSingleMetricUint(name string) (uint64, error) {
	sample := make([]metrics.Sample, 1)
	sample[0].Name = name

	metrics.Read(sample)

	if sample[0].Value.Kind() == metrics.KindBad {
		return 0, fmt.Errorf("bad sample metric: %s", name)
	}

	return sample[0].Value.Uint64(), nil
}
