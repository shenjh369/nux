package nux

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/toolkits/file"
)

/*
Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
  eth0: 1990350    2838    0    0    0     0          0         0   401351    2218    0    0    0     0       0          0
    lo:   26105     286    0    0    0     0          0         0    26105     286    0    0    0     0       0          0
*/
func NetIfs(onlyPrefix []string) ([]*NetIf, error) {
	contents, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return nil, err
	}

	ret := []*NetIf{}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		lineBytes, err := file.ReadLine(reader)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil, err
		}

		line := string(lineBytes)
		idx := strings.Index(line, ":")
		if idx < 0 {
			continue
		}

		netIf := NetIf{}

		eth := strings.TrimSpace(line[0:idx])
		if len(onlyPrefix) > 0 {
			found := false
			for _, prefix := range onlyPrefix {
				if strings.HasPrefix(eth, prefix) {
					found = true
					break
				}
			}

			if !found {
				continue
			}
		}

		netIf.Iface = eth

		fields := strings.Fields(line[idx+1:])

		if len(fields) != 16 {
			continue
		}

		netIf.InBytes, _ = strconv.ParseInt(fields[0], 10, 64)
		netIf.InPackages, _ = strconv.ParseInt(fields[1], 10, 64)
		netIf.InErrors, _ = strconv.ParseInt(fields[2], 10, 64)
		netIf.InDropped, _ = strconv.ParseInt(fields[3], 10, 64)
		netIf.InFifoErrs, _ = strconv.ParseInt(fields[4], 10, 64)
		netIf.InFrameErrs, _ = strconv.ParseInt(fields[5], 10, 64)
		netIf.InCompressed, _ = strconv.ParseInt(fields[6], 10, 64)
		netIf.InMulticast, _ = strconv.ParseInt(fields[7], 10, 64)

		netIf.OutBytes, _ = strconv.ParseInt(fields[8], 10, 64)
		netIf.OutPackages, _ = strconv.ParseInt(fields[9], 10, 64)
		netIf.OutErrors, _ = strconv.ParseInt(fields[10], 10, 64)
		netIf.OutDropped, _ = strconv.ParseInt(fields[11], 10, 64)
		netIf.OutFifoErrs, _ = strconv.ParseInt(fields[12], 10, 64)
		netIf.OutCollisions, _ = strconv.ParseInt(fields[13], 10, 64)
		netIf.OutCarrierErrs, _ = strconv.ParseInt(fields[14], 10, 64)
		netIf.OutCompressed, _ = strconv.ParseInt(fields[15], 10, 64)

		netIf.TotalBytes = netIf.InBytes + netIf.OutBytes
		netIf.TotalPackages = netIf.InPackages + netIf.OutPackages
		netIf.TotalErrors = netIf.InErrors + netIf.OutErrors
		netIf.TotalDropped = netIf.InDropped + netIf.OutDropped

		ret = append(ret, &netIf)
	}

	return ret, nil
}
