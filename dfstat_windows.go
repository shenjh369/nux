package nux

import (
	"github.com/shirou/gopsutil/disk"
)

func ListMountPoint() ([][3]string, error) {
	result := make([][3]string, 0)
	list, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		result = append(result, [3]string{v.Device, v.Mountpoint, v.Fstype})
	}
	return result, nil
}

func BuildDeviceUsage(_fsSpec, _fsFile, _fsVfstype string) (*DeviceUsage, error) {
	var deviceUsage DeviceUsage
	stat, err := disk.Usage(_fsSpec)
	if err != nil {
		return nil, err
	}
	deviceUsage = DeviceUsage{_fsSpec,
		_fsFile,
		_fsVfstype,
		stat.Total,
		stat.Used,
		stat.Free,
		stat.UsedPercent,
		100 - stat.UsedPercent,
		stat.InodesTotal,
		stat.InodesUsed,
		stat.InodesFree,
		stat.InodesUsedPercent,
		100 - stat.InodesUsedPercent,
	}
	return &deviceUsage, nil
}
