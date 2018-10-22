package main

import (
	"merculet.io/support"
)

type Token struct {
	*support.Payload
	Symbol      string
	Name        string
	TotalSupply uint64
	Decimals    uint8
}

type TokenService struct {
	ledgerStub *support.LedgerStub
	keyGen     func(t *Token) string
}

func (ts *TokenService) New(s *support.LedgerStub) interface{} {
	ts.ledgerStub = s
	ts.keyGen = func(t *Token) string {
		return t.Symbol
	}
	return ts
}

//
func (ts *TokenService) create(token *Token) (*Token, error) {
	return token, ts.ledgerStub.Put(ts.keyGen(token), token)
}

func (ts *TokenService) has(token *Token) (bool, error) {
	return ts.ledgerStub.Exists(ts.keyGen(token))
}

func (ts *TokenService) update(token *Token) (*Token, error) {
	return ts.create(token)
}

func (ts *TokenService) get(token *Token) (*Token, error) {
	res, e := ts.ledgerStub.Get(ts.keyGen(token), &Token{})
	if e != nil {
		return nil, e
	}
	return res.(*Token), nil
}
