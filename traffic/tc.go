package traffic

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func deleteQdisc(netIf string) error {
	intIf, err := netlink.LinkByName(netIf)
	if err != nil {
		return fmt.Errorf("failed to get the interface: %v", err)
	}

	qdiscs, err := netlink.QdiscList(intIf)
	if err != nil {
		return fmt.Errorf("failed to list qdiscs on %v: %v", netIf, err)
	}

	if len(qdiscs) == 0 {
		return nil
	}

	err = netlink.QdiscDel(qdiscs[0])
	if err != nil {
		return fmt.Errorf("failed to delete qdisc on %v: %v", netIf, err)
	}

	qdiscs, err = netlink.QdiscList(intIf)
	if err != nil {
		return fmt.Errorf("failed to list qdiscs on %v: %v", netIf, err)
	}

	if len(qdiscs) != 0 {
		return fmt.Errorf("found qdisc on %v after removing qdisc", netIf)
	}
	return nil
}

// SetUpBandwidthShaping sets up TBF on the network interface
func SetUpBandwidthShaping(netIf string, bandwidth uint64) error {
	intIf, err := netlink.LinkByName(netIf)
	if err != nil {
		return fmt.Errorf("failed to get the interface: %v", err)
	}

	// we don't care about this error
	_ = deleteQdisc(netIf)

	qdisc := &netlink.Tbf{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: intIf.Attrs().Index,
			Handle:    netlink.MakeHandle(1, 0),
			Parent:    netlink.HANDLE_ROOT,
		},
		Rate:   bandwidth,
		Limit:  1500000,
		Buffer: 32768,
	}
	if err := netlink.QdiscAdd(qdisc); err != nil {
		return err
	}
	return nil
}
