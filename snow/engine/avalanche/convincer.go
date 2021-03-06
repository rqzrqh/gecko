// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/ava-labs/gecko/ids"
	"github.com/ava-labs/gecko/snow/consensus/avalanche"
	"github.com/ava-labs/gecko/snow/engine/common"
	"github.com/ava-labs/gecko/utils/wrappers"
)

type convincer struct {
	consensus avalanche.Consensus
	sender    common.Sender
	vdr       ids.ShortID
	requestID uint32
	abandoned bool
	deps      ids.Set
	errs      *wrappers.Errs
}

func (c *convincer) Dependencies() ids.Set { return c.deps }

func (c *convincer) Fulfill(id ids.ID) {
	c.deps.Remove(id)
	c.Update()
}

func (c *convincer) Abandon(ids.ID) { c.abandoned = true }

func (c *convincer) Update() {
	if c.abandoned || c.deps.Len() != 0 || c.errs.Errored() {
		return
	}

	c.sender.Chits(c.vdr, c.requestID, c.consensus.Preferences())
}
