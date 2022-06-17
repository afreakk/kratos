package login

import (
	"context"

	"github.com/ory/kratos/ui/node"
)

func sortNodes(ctx context.Context, n node.Nodes) error {
	return n.SortBySchema(ctx,
		node.SortByGroups([]node.UiNodeGroup{
			node.DefaultGroup,
			node.WebAuthnGroup,
			node.PasswordGroup,
			node.OpenIDConnectGroup,
			node.TOTPGroup,
			node.LookupGroup,
		}),
		node.SortUseOrder([]string{
			"csrf_token",
			"identifier",
			"password",
		}),
	)
}
