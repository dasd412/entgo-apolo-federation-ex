// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphqlmodel

import (
	"order/pkg/ent"
)

type ExampleTime struct {
	UnixTime  int    `json:"unixTime"`
	TimeStamp string `json:"timeStamp"`
}

type Subscription struct {
}

type User struct {
	ID     int          `json:"id"`
	Orders []*ent.Order `json:"orders,omitempty"`
}

func (User) IsEntity() {}
