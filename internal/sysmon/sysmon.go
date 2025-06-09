package sysmon

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type SysMon struct {
	DiskName     string
	CpuStatsFile string
	LabelFormat  string
	Name         string
	Instance     string
}

func getTemperature() string {
	cmd := exec.Command("sensors")
	out, err := cmd.Output()
	if err != nil {
		return "N/A"
	}

	lines := strings.SplitSeq(string(out), "\n")
	for line := range lines {
		if strings.Contains(line, "CPU") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				return fields[1]
			}
		}
	}
	return "N/A"
}

func (s SysMon) getDiskUsage() string {
	cmd := exec.Command("df", "-h")
	out, err := cmd.Output()
	if err != nil {
		return "N/A"
	}

	lines := strings.SplitSeq(string(out), "\n")
	for line := range lines {
		if strings.Contains(line, s.DiskName) {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				if len(fields) >= 5 {
					return fields[4]
				} else {
					return fmt.Sprintf("??? (%d fields)", len(fields))
				}
			}
		}
	}
	return "N/A"
}

func readCPUStat() (idle, total uint64, ok bool) {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	fields := strings.Fields(strings.Split(string(data), "\n")[0])
	if len(fields) < 5 {
		return
	}
	for i, v := range fields[1:] {
		val, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		total += val
		if i == 3 {
			idle = val
		}
	}
	ok = true
	return
}

func (s SysMon) getCPUUsage() int {
	idleNow, totalNow, ok := readCPUStat()
	if !ok {
		return 0
	}

	lastData, err := os.ReadFile(s.CpuStatsFile)
	if err != nil {
		// No previous data, save current and return 0
		os.WriteFile(s.CpuStatsFile, fmt.Appendf(nil, "%d %d", idleNow, totalNow), 0644)
		return 0
	}

	parts := strings.Fields(string(lastData))
	if len(parts) != 2 {
		os.WriteFile(s.CpuStatsFile, fmt.Appendf(nil, "%d %d", idleNow, totalNow), 0644)
		return 0
	}

	idlePrev, _ := strconv.ParseUint(parts[0], 10, 64)
	totalPrev, _ := strconv.ParseUint(parts[1], 10, 64)

	os.WriteFile(s.CpuStatsFile, fmt.Appendf(nil, "%d %d", idleNow, totalNow), 0644)

	totalDelta := float64(totalNow - totalPrev)
	idleDelta := float64(idleNow - idlePrev)

	if totalDelta == 0 {
		return 0
	}

	usage := 100 * (1.0 - idleDelta/totalDelta)
	return int(usage + 0.5)
}

func getMemoryUsage() int {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0
	}

	var memTotal, memFree int
	for line := range strings.SplitSeq(string(data), "\n") {
		if strings.HasPrefix(line, "MemTotal:") {
			fmt.Sscanf(line, "MemTotal: %d kB", &memTotal)
		} else if strings.HasPrefix(line, "MemFree:") {
			fmt.Sscanf(line, "MemFree: %d kB", &memFree)
		}
	}
	if memTotal == 0 {
		return 0
	}
	used := int(100 * float64(memTotal-memFree) / float64(memTotal))
	return used
}

func getSwapUsage() int {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0
	}

	var swapTotal, swapFree int
	for line := range strings.SplitSeq(string(data), "\n") {
		if strings.HasPrefix(line, "SwapTotal:") {
			fmt.Sscanf(line, "SwapTotal: %d kB", &swapTotal)
		} else if strings.HasPrefix(line, "SwapFree:") {
			fmt.Sscanf(line, "SwapFree: %d kB", &swapFree)
		}
	}
	if swapTotal == 0 {
		return 0
	}
	used := int(100 * float64(swapTotal-swapFree) / float64(swapTotal))
	return used
}

func (s SysMon) GetName() string     { return s.Name }
func (s SysMon) GetInstance() string { return s.Instance }
func (s SysMon) GetFullText(_ []byte) (string, error) {
	return fmt.Sprintf(s.LabelFormat, getTemperature(), s.getCPUUsage(), s.getDiskUsage(), getMemoryUsage(), getSwapUsage()), nil
}
