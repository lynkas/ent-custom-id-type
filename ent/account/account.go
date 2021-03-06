// Code generated by entc, DO NOT EDIT.

package account

import (
	"ent_test/ent/schema"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "tokens"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "account_token"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldEmail,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() schema.ID
)
