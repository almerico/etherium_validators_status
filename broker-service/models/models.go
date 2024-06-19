package models

type Info struct {
	Status string   `json:"status" validate:"required"`
	Data   InfoData `json:"data" validate:"required"`
}

type InfoData struct {
	Activationeligibilityepoch int64  `json:"activationeligibilityepoch"`
	Activationepoch            int64  `json:"activationepoch"`
	Balance                    int64  `json:"balance"`
	Effectivebalance           int64  `json:"effectivebalance"`
	Exitepoch                  int64  `json:"exitepoch"`
	Lastattestationslot        int64  `json:"lastattestationslot"`
	Name                       string `json:"name"`
	Pubkey                     string `json:"pubkey"`
	Slashed                    bool   `json:"slashed"`
	Status                     string `json:"status"`
	Validatorindex             int64  `json:"validatorindex"`
	Withdrawableepoch          int64  `json:"withdrawableepoch"`
	Withdrawalcredentials      string `json:"withdrawalcredentials"`
	Total_withdrawals          int64  `json:"total_withdrawals"`
}
