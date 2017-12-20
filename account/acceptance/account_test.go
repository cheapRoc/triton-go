package account_acc

import (
	"context"
	"testing"

	triton "github.com/joyent/triton-go"
	"github.com/joyent/triton-go/account"
	"github.com/joyent/triton-go/testutils"
)

func TestAccAccount_Get(t *testing.T) {
	testutils.AccTest(t, testutils.TestCase{
		Steps: []testutils.Step{

			&testutils.StepClient{
				StateBagKey: "account",
				CallFunc: func(config *triton.ClientConfig) (interface{}, error) {
					return account.NewClient(config)
				},
			},

			&testutils.StepAPICall{
				StateBagKey: "account",
				CallFunc: func(client interface{}) (interface{}, error) {
					c := client.(*account.AccountClient)
					ctx := context.Background()
					input := &account.GetInput{}
					return c.Get(ctx, input)
				},
			},

			&testutils.StepAssertSet{
				StateBagKey: "account",
				Keys:        []string{"ID", "Login", "Email"},
			},
		},
	})
}
