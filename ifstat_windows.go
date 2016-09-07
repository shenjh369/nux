package nux

import (
	"github.com/shirou/gopsutil/net"
)

func NetIfs(ifacePrefix []string) ([]*NetIf, error) {
	netifs := make([]*NetIf, 0)
	list, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		N := NetIf{
			v.Name, int64(v.BytesRecv), int64(v.PacketsRecv),
			int64(v.Errin), int64(v.Dropin), 0, 0, 0, 0, int64(v.BytesSent),
			int64(v.PacketsSent), int64(v.Errout), int64(v.Dropout), 0, 0, 0, 0,
			int64(v.BytesRecv + v.BytesSent), int64(v.PacketsRecv + v.PacketsSent),
			int64(v.Errin + v.Errout), int64(v.Dropin + v.Dropout), 0, 0, 0,
		}
		netifs = append(netifs, &N)
	}
	return netifs, nil
}
