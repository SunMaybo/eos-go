package token

import "github.com/SunMaybo/eos-go"

func NewTransfer(from, to eos.AccountName, quantity eos.Asset, memo string) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.token"),
		Name:    ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}
func NewTransferAndPermissionName(from, to eos.AccountName, quantity eos.Asset, memo string, name eos.PermissionName) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.token"),
		Name:    ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: PN(string(name))},
		},
		ActionData: eos.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `eosio.token` contract.
type Transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
