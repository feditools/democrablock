package models

import (
	"fmt"
	"strings"
)

const (
	TransactionTypeCouncilInit       TransactionType = "COUNCIL_INIT"
	TransactionTypeCouncilInitString string          = "Council initialized with members: %s"
)

type TransactionCouncilInit struct {
	Members []TransactionCouncilInitMember `json:"members"`
}

type TransactionCouncilInitMember struct {
	DBID int64  `json:"db_id"`
	Name string `json:"name"`
}

func (t *TransactionCouncilInit) String() string {
	memberNames := make([]string, len(t.Members))
	for i, member := range t.Members {
		memberNames[i] = member.Name
	}

	return fmt.Sprintf(TransactionTypeCouncilInitString, strings.Join(memberNames, ", "))
}
