package main

import (
	"github.com/pkg/errors"
	"merculet.io/support"
	"strings"
)

type Transaction struct {
	*support.Payload
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	Symbol     string `json:"symbol,omitempty"`
	Amount     int64  `json:"amt,omitempty"`
	Ts         int64  `json:"ts,omitempty"`
	Balance    int64  `json:"balance,omitempty"`
	PreviousTx string `json:"previousTx,omitempty"`
}

type TokenTransferService struct {
	ledgerStub *support.LedgerStub
	keyGen     func(t *Transaction) string
	isValid    func(t *Transaction) bool
}

func (ts *TokenTransferService) New(s *support.LedgerStub) interface{} {
	ts.ledgerStub = s
	ts.keyGen = func(t *Transaction) string {
		//TODO
		return strings.Join([]string{t.From, t.Symbol}, `|`)
	}
	ts.isValid = func(t *Transaction) (valid bool) {
		return true
	}
	return ts
}

func (ts *TokenTransferService) Transfer(tx *Transaction) error {
	if !ts.isValid(tx) {
		return errors.New(`entity is invalid`)
	}
	compensateTx := buildCompensateTx(tx)
	var err error
	for _, tx := range []*Transaction{tx, compensateTx} {
		err = ts.transferInternal(tx)
		if err != nil {
			return err
		}
	}
	return nil
}

//TODO
func (ts *TokenTransferService) BatchTransfer(transactions ...*Transaction) error {
	for _, transaction := range transactions {
		error := ts.Transfer(transaction)
		if error != nil {
			return error
		}
	}
	return nil
}

func (ts *TokenTransferService) Get(tx *Transaction) (*Transaction, error) {
	res, e := ts.ledgerStub.Get(ts.keyGen(tx), &Transaction{})
	if e != nil {
		return nil, e
	}
	return res.(*Transaction), nil
}

func (ts *TokenTransferService) transferInternal(tx *Transaction) error {

	previousTx, e := ts.Get(tx)

	if e != nil {
		return nil
	}

	//TODO
	tx.Balance = previousTx.Balance + tx.Amount
	tx.PreviousTx = previousTx.TxId
	//
	return ts.ledgerStub.Put(ts.keyGen(tx), tx)
}

func buildCompensateTx(tx *Transaction) *Transaction {
	tx.From, tx.To = tx.To, tx.From
	tx.Amount = (-1) * tx.Amount
	return tx
}
