package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct{}

func (m *CPUMonitor) Check(ctx context.Context) (string, error) {
	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil {
		return "N/A", err
	}

	return fmt.Sprintf("%.2f", percent[0]) + "%", nil
}
