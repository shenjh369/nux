package nux

import (
    "github.com/shirou/gopsutil/mem"
)

func MemInfo() (*Mem, error) {
    info, err := mem.VirtualMemory()
    if err != nil {
        return nil, err
    }
    return &Mem{info.Buffers, info.Cached, info.Total, info.Available, 0, 0, 0}, nil
}
