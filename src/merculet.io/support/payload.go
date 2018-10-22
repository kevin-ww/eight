package support

import "time"

type Payload struct {
	TxId      string    `json:"txId,omitempty"`
	TxTs      int64     `json:"txTs,omitempty"`
	Memo      string    `json:"memo,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
}

type payload interface {
	KeyGen() string
	IsValid() bool
}
