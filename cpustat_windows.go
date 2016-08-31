package nux

import (
	"log"

	"github.com/shirou/gopsutil/cpu"
)

func CurrentProcStat() (*ProcStat, error) {
	stats, err := cpu.Times(false)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stat := stats[0]

	cu := &CpuUsage{
		User:    uint64(stat.User),
		Nice:    uint64(stat.Nice),
		System:  uint64(stat.System),
		Idle:    uint64(stat.Idle),
		Iowait:  uint64(stat.Iowait),
		Irq:     uint64(stat.Irq),
		SoftIrq: uint64(stat.Softirq),
		Steal:   uint64(stat.Steal),
		Guest:   uint64(stat.Guest),
		Total:   uint64(stat.User + stat.Nice + stat.System + stat.Idle + stat.Irq + stat.Softirq + stat.Steal + stat.Guest),
	}
	ps := &ProcStat{Cpus: make([]*CpuUsage, 1)}
	ps.Cpu = cu
	return ps, nil
}
