package models

type QpayRequest struct {
	SenderInvoiceNo     string `json:"sender_invoice_no"`
	InvoiceReceiverCode string `json:"invoice_receiver_code"`
	InvoiceDescription  string `json:"invoice_description"`
	Amount              string `json:"amount"`
}

type QpayCheck struct {
	ObjectType        string `json:"object_type"`
	ObjectId          string `json:"object_id"`
	OrderNumber       string `json:"order_number"`
	EbarimtType       int    `json:"ebarimt_type"`
	OrgRegisterNumber int    `json:"org_register_number"`
}

type QpayResonpse struct {
	Count      int `json:"count"`
	PaidAmount int `json:"paid_amount"`
	Rows       []struct {
		PaymentID           string        `json:"payment_id"`
		PaymentStatus       string        `json:"payment_status"`
		PaymentAmount       string        `json:"payment_amount"`
		TrxFee              string        `json:"trx_fee"`
		PaymentCurrency     string        `json:"payment_currency"`
		PaymentWallet       string        `json:"payment_wallet"`
		PaymentType         string        `json:"payment_type"`
		NextPaymentDate     interface{}   `json:"next_payment_date"`
		NextPaymentDatetime interface{}   `json:"next_payment_datetime"`
		CardTransactions    []interface{} `json:"card_transactions"`
		P2PTransactions     []struct {
			TransactionBankCode string `json:"transaction_bank_code"`
			AccountBankCode     string `json:"account_bank_code"`
			AccountBankName     string `json:"account_bank_name"`
			AccountNumber       string `json:"account_number"`
			Status              string `json:"status"`
			Amount              string `json:"amount"`
			Currency            string `json:"currency"`
			SettlementStatus    string `json:"settlement_status"`
		} `json:"p2p_transactions"`
	} `json:"rows"`
}
