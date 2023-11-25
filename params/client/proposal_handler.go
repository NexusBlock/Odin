package client

import (
	govclient "github.com/nexusblock/heimdall/gov/client"
	"github.com/nexusblock/heimdall/params/client/cli"
	"github.com/nexusblock/heimdall/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
