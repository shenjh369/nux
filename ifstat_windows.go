package nux

import (
    "errors"
    "github.com/shirou/gopsutil/net"
    snet "net"
)

func NetIfs(ifacePrefix []string) ([]*NetIf, error) {
    netifs := make([]*NetIf, 0)
    list, err := net.NetIOCounters(true)
    if err != nil {
        return nil, err
    }
    for _, v := range list {
        iface, err := GetNetCardMac(v.Addr)
        if err != nil {
            continue
        }
        N := NetIf{
            iface, int64(v.BytesRecv), int64(v.PacketsRecv),
            int64(v.Errin), int64(v.Dropin), 0, 0, 0, 0, int64(v.BytesSent),
            int64(v.PacketsSent), int64(v.Errout), int64(v.Dropout), 0, 0, 0, 0,
            int64(v.BytesRecv + v.BytesSent), int64(v.PacketsRecv + v.PacketsSent),
            int64(v.Errin + v.Errout), int64(v.Dropin + v.Dropout),
        }
        netifs = append(netifs, &N)
    }
    return netifs, nil
}

func GetNetCardMac(mac [8]byte) (string, error) {
    list, err := net.NetInterfaces()
    if err != nil {
        return "", err
    }
    for _, v := range list {
        if v.HardwareAddr == snet.HardwareAddr(mac[:8]).String() {
            return v.Name, nil
        }
    }
    return "", errors.New("Have no this mac")
}
