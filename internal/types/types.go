// Code generated by goctl. DO NOT EDIT.
package types

type PrimeTime struct {
	Hour  uint8 `json:"hour"`
	Count int   `json:"count"`
}

type Message struct {
	Id        int    `json:"id"`
	Direction string `json:"direction"`
	Sent      string `json:"sent"`
	MailboxId int    `json:"mailboxId"`
	CrmRid    int    `json:"crmRid"`
}

type PrimeCache struct {
	Hour  uint8 `json:"hour"`
	Count int   `json:"count"`
}

type RangeQuery struct {
	After  string `query:"after"`
	Before string `query:"before"`
}
