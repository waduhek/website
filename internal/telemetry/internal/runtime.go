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
