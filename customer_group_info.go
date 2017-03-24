package multivers

type CustomerGroupInfo struct {
	CustomerGroupID int    `json:"customerGroupId"`
	FiscalYear      int    `json:"fiscalYear"`
	LedgerAccountID string `json:"ledgerAccountId"`
}
