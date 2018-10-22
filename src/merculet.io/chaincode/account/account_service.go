package main

import "merculet.io/support"

type Account struct {
	*support.Payload
	Name         string
	Organization string
}

type AccountService struct {
	ledgerStub *support.LedgerStub
	keyGen     func(t *Account) string
}

func (as *AccountService) New(s *support.LedgerStub) interface{} {
	as.ledgerStub = s
	as.keyGen = func(t *Account) string {
		//name only ?
		return t.Name
	}
	return as
}

//
func (as *AccountService) create(acc *Account) (*Account, error) {
	return acc, as.ledgerStub.Put(as.keyGen(acc), acc)
}

func (as *AccountService) has(acc *Account) (bool, error) {
	return as.ledgerStub.Exists(as.keyGen(acc))
}

func (as *AccountService) update(acc *Account) (*Account, error) {
	return as.create(acc)
}

func (as *AccountService) get(acc *Account) (*Account, error) {
	res, e := as.ledgerStub.Get(as.keyGen(acc), &Account{})
	if e != nil {
		return nil, e
	}
	return res.(*Account), nil
}
