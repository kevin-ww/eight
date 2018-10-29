package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"merculet.io/support"
	"strings"
)

//const AddressLength = 20
//
//type Address [AddressLength]byte
//
//func (a Address) String() string {
//	return a.Hex()
//}
//func (a Address) Hex() string {
//	unchecksummed := hex.EncodeToString(a[:])
//	sha := sha3.NewKeccak256()
//	sha.Write([]byte(unchecksummed))
//	hash := sha.Sum(nil)
//
//	result := []byte(unchecksummed)
//	for i := 0; i < len(result); i++ {
//		hashByte := hash[i/2]
//		if i%2 == 0 {
//			hashByte = hashByte >> 4
//		} else {
//			hashByte &= 0xf
//		}
//		if result[i] > '9' && hashByte > 7 {
//			result[i] -= 32
//		}
//	}
//	return "0x" + string(result)
//}

type Account struct {
	Name                string   `json:"name,omitempty"`
	Symbol              string   `json:"symbol,omitempty"`
	Balance             *big.Int `json:"balance,omitempty"`
	LastTransactionTime int64    `json:"last,omitempty"`
}

func AsAccount(name string) Account {
	//TODO validation
	return Account{Name: name}
}

func AsAccountWithSymbol(name string, symbol string) Account {
	return Account{Name: name, Symbol: symbol}
}

func (a *Account) String() string {
	return a.Name
}

type Transaction struct {
	support.Payload
	From       Account  `json:"from,omitempty"`
	To         Account  `json:"to,omitempty"`
	Symbol     string   `json:"symbol,omitempty"`
	Amount     *big.Int `json:"amt,omitempty"`
	Ts         int64    `json:"ts,omitempty"`
	Balance    *big.Int `json:"balance,omitempty"`
	PreviousTx string   `json:"previousTx,omitempty"`
}

type TokenTransferService struct {
	ledgerStub *support.LedgerStub
	//keyGen     func(t *Transaction) string
	//isValid    func(t *Transaction) bool
}

func (t *Transaction) isValid() bool {
	return true
}

func (a *Account) KeyGen() string {
	return strings.Join([]string{a.Name, a.Symbol}, "|")
}

func (t *Transaction) KeyGen() string {
	return strings.Join([]string{t.From.String(), t.Symbol}, "|")
}

func (ts *TokenTransferService) New(s *support.LedgerStub) interface{} {
	ts.ledgerStub = s
	return ts
}

func (ts *TokenTransferService) Transfer(tx *Transaction) error {
	if !tx.isValid() {
		return errors.New(`entity is invalid`)
	}
	//compensateTx := buildCompensateTx(tx)
	var err error
	err = ts.transferInternal(tx)
	if err != nil {
		return nil
	}

	return ts.transferInternal(buildCompensateTx(tx))

	//for _, x := range []*Transaction{tx, buildCompensateTx(tx)} {
	//	fmt.Printf("handle transaction of %v \n",x)
	//	err = ts.transferInternal(x)
	//	if err != nil {
	//		return err
	//	}
	//}
	//return nil
}

//TODO
func (ts *TokenTransferService) BatchTransfer(transactions []Transaction) error {
	for _, transaction := range transactions {
		error := ts.Transfer(&transaction)
		if error != nil {
			return error
		}
	}
	return nil
}

func (ts *TokenTransferService) GetBalance(account *Account) (*Account, error) {

	res, e := ts.ledgerStub.Get(account.KeyGen(), &Transaction{})

	if e != nil || res == nil {
		return nil, e
	}
	account.Balance = res.(*Transaction).Balance

	return account, nil
}

func (ts *TokenTransferService) getInternal(tx *Transaction) (*Transaction, error) {

	res, e := ts.ledgerStub.Get(tx.KeyGen(), &Transaction{})
	if e != nil || res == nil {
		return nil, e
	}

	return res.(*Transaction), nil
}

func (ts *TokenTransferService) transferInternal(tx *Transaction) error {

	fmt.Printf("with tx %v \n", tx)

	previousTx, e := ts.getInternal(tx)

	if e != nil {
		return e
	}

	//TODO

	var previousBalance *big.Int
	if previousTx == nil {
		previousBalance = big.NewInt(0)
	} else {
		previousBalance = previousTx.Balance
		tx.PreviousTx = previousTx.TxId
	}

	previousBalance.Add(previousBalance, tx.Amount)
	tx.Balance = previousBalance
	//
	return ts.ledgerStub.Put(tx.KeyGen(), tx)
}

func buildCompensateTx(tx *Transaction) *Transaction {
	//TODO use clone
	tx.From, tx.To = tx.To, tx.From
	tx.Amount = tx.Amount.Neg(tx.Amount)
	return tx
}
