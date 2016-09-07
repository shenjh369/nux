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
		User:    stat.User,
		Nice:    stat.Nice,
		System:  stat.System,
		Idle:    stat.Idle,
		Iowait:  stat.Iowait,
		Irq:     stat.Irq,
		SoftIrq: stat.Softirq,
		Steal:   stat.Steal,
		Guest:   stat.Guest,
		Total:   stat.User + stat.Nice + stat.System + stat.Idle + stat.Irq + stat.Softirq + stat.Steal + stat.Guest,
	}
	ps := &ProcStat{Cpus: make([]*CpuUsage, 1)}
	ps.Cpu = cu
	return ps, nil
}
