/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package drivers

import (
	"encoding/json"
	"fmt"

	"github.com/contiv/netplugin/core"
	"github.com/jainvipin/bitset"
)

// implements the State interface for a network implemented using
// vlans with ovs. The state is stored as Json objects.
const (
	BASE_PATH           = "/contiv/"
	CFG_PATH            = BASE_PATH + "config/"
	NW_CFG_PATH_PREFIX  = CFG_PATH + "nets/"
	NW_CFG_PATH         = NW_CFG_PATH_PREFIX + "%s"
	EP_CFG_PATH_PREFIX  = CFG_PATH + "eps/"
	EP_CFG_PATH         = EP_CFG_PATH_PREFIX + "%s"
	OPER_PATH           = BASE_PATH + "oper/"
	NW_OPER_PATH_PREFIX = OPER_PATH + "nets/"
	NW_OPER_PATH        = NW_OPER_PATH_PREFIX + "%s"
	EP_OPER_PATH_PREFIX = OPER_PATH + "eps/"
	EP_OPER_PATH        = EP_OPER_PATH_PREFIX + "%s"
)

type OvsCfgNetworkState struct {
	core.CommonState
	Tenant     string        `json:"tenant"`
	PktTagType string        `json:"pktTagType"`
	PktTag     int           `json:"pktTag"`
	ExtPktTag  int           `json:"extPktTag"`
	SubnetIp   string        `json:"subnetIp"`
	SubnetLen  uint          `json:"subnetLen"`
	DefaultGw  string        `json:"defaultGw"`
	EpCount    int           `json:"epCount"`
	IpAllocMap bitset.BitSet `json:"ipAllocMap"`
}

func (s *OvsCfgNetworkState) Write() error {
	key := fmt.Sprintf(NW_CFG_PATH, s.Id)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

func (s *OvsCfgNetworkState) Read(id string) error {
	key := fmt.Sprintf(NW_CFG_PATH, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

func (s *OvsCfgNetworkState) ReadAll() ([]core.State, error) {
	return s.StateDriver.ReadAllState(NW_CFG_PATH_PREFIX, s, json.Unmarshal)
}

func (s *OvsCfgNetworkState) WatchAll(rsps chan core.WatchState) error {
	return s.StateDriver.WatchAllState(NW_CFG_PATH_PREFIX, s, json.Unmarshal,
		rsps)
}

func (s *OvsCfgNetworkState) Clear() error {
	key := fmt.Sprintf(NW_CFG_PATH, s.Id)
	return s.StateDriver.ClearState(key)
}
