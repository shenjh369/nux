package nux

import (
	"fmt"
)

type CpuUsage struct {
	User    float64 // time spent in user mode
	Nice    float64 // time spent in user mode with low priority (nice)
	System  float64 // time spent in system mode
	Idle    float64 // time spent in the idle task
	Iowait  float64 // time spent waiting for I/O to complete (since Linux 2.5.41)
	Irq     float64 // time spent servicing  interrupts  (since  2.6.0-test4)
	SoftIrq float64 // time spent servicing softirqs (since 2.6.0-test4)
	Steal   float64 // time spent in other OSes when running in a virtualized environment (since 2.6.11)
	Guest   float64 // time spent running a virtual CPU for guest operating systems under the control of the Linux kernel. (since 2.6.24)
	Total   float64 // total of all time fields
}

func (this *CpuUsage) String() string {
	return fmt.Sprintf("<User:%d, Nice:%d, System:%d, Idle:%d, Iowait:%d, Irq:%d, SoftIrq:%d, Steal:%d, Guest:%d, Total:%d>",
		this.User,
		this.Nice,
		this.System,
		this.Idle,
		this.Iowait,
		this.Irq,
		this.SoftIrq,
		this.Steal,
		this.Guest,
		this.Total)
}

type ProcStat struct {
	Cpu          *CpuUsage
	Cpus         []*CpuUsage
	Ctxt         uint64
	Processes    uint64
	ProcsRunning uint64
	ProcsBlocked uint64
}

func (this *ProcStat) String() string {
	return fmt.Sprintf("<Cpu:%v, Cpus:%v, Ctxt:%d, Processes:%d, ProcsRunning:%d, ProcsBlocking:%d>",
		this.Cpu,
		this.Cpus,
		this.Ctxt,
		this.Processes,
		this.ProcsRunning,
		this.ProcsBlocked)
}
