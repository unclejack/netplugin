package mastercfg

import (
	"encoding/json"
	"fmt"

	"github.com/contiv/netplugin/core"
)

const (
	netProfileConfigPathPrefix = StateConfigPath + "netProfile/"
	netProfileConfigPath       = netProfileConfigPathPrefix + "%s"
)

// EpgNetProfile has an instance of policy attached to an endpoint group
type EpgNetProfile struct {
	core.CommonState
	EpgNetProfileKey string // Key for this epg netProfile
	EndpointGroupID  int    // Endpoint group where this policy is attached to
}

// Epg NetProfile database
var epgNetProfileDb = make(map[string]*EpgNetProfile)

// Write the state.
func (p *EpgNetProfile) Write() error {
	key := fmt.Sprintf(netProfileConfigPath, p.ID)
	return p.StateDriver.WriteState(key, p, json.Marshal)
}

// Read the state for a given identifier
func (p *EpgNetProfile) Read(id string) error {
	key := fmt.Sprintf(netProfileConfigPath, id)
	return p.StateDriver.ReadState(key, p, json.Unmarshal)
}

// ReadAll state and return the collection.
func (p *EpgNetProfile) ReadAll() ([]core.State, error) {
	return p.StateDriver.ReadAllState(netProfileConfigPathPrefix, p, json.Unmarshal)
}

// WatchAll state transitions and send them through the channel.
func (p *EpgNetProfile) WatchAll(rsps chan core.WatchState) error {
	return p.StateDriver.WatchAllState(netProfileConfigPathPrefix, p, json.Unmarshal,
		rsps)
}

// Clear removes the state.
func (p *EpgNetProfile) Clear() error {
	key := fmt.Sprintf(netProfileConfigPath, p.ID)
	return p.StateDriver.ClearState(key)
}
