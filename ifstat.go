package nux

import "fmt"

type NetIf struct {
    Iface          string
    InBytes        int64
    InPackages     int64
    InErrors       int64
    InDropped      int64
    InFifoErrs     int64
    InFrameErrs    int64
    InCompressed   int64
    InMulticast    int64
    OutBytes       int64
    OutPackages    int64
    OutErrors      int64
    OutDropped     int64
    OutFifoErrs    int64
    OutCollisions  int64
    OutCarrierErrs int64
    OutCompressed  int64
    TotalBytes     int64
    TotalPackages  int64
    TotalErrors    int64
    TotalDropped   int64
}

func (this *NetIf) String() string {
    return fmt.Sprintf("<Iface:%s,InBytes:%d,InPackages:%d...>", this.Iface, this.InBytes, this.InPackages)
}
