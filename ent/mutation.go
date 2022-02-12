// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"ent_test/ent/account"
	"ent_test/ent/predicate"
	"ent_test/ent/schema"
	"ent_test/ent/token"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount = "Account"
	TypeToken   = "Token"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op            Op
	typ           string
	id            *schema.ID
	email         *string
	clearedFields map[string]struct{}
	token         map[schema.ID]struct{}
	removedtoken  map[schema.ID]struct{}
	clearedtoken  bool
	done          bool
	oldValue      func(context.Context) (*Account, error)
	predicates    []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id schema.ID) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Account entities.
func (m *AccountMutation) SetID(id schema.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountMutation) ID() (id schema.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccountMutation) IDs(ctx context.Context) ([]schema.ID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []schema.ID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Account.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *AccountMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *AccountMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *AccountMutation) ResetEmail() {
	m.email = nil
}

// AddTokenIDs adds the "token" edge to the Token entity by ids.
func (m *AccountMutation) AddTokenIDs(ids ...schema.ID) {
	if m.token == nil {
		m.token = make(map[schema.ID]struct{})
	}
	for i := range ids {
		m.token[ids[i]] = struct{}{}
	}
}

// ClearToken clears the "token" edge to the Token entity.
func (m *AccountMutation) ClearToken() {
	m.clearedtoken = true
}

// TokenCleared reports if the "token" edge to the Token entity was cleared.
func (m *AccountMutation) TokenCleared() bool {
	return m.clearedtoken
}

// RemoveTokenIDs removes the "token" edge to the Token entity by IDs.
func (m *AccountMutation) RemoveTokenIDs(ids ...schema.ID) {
	if m.removedtoken == nil {
		m.removedtoken = make(map[schema.ID]struct{})
	}
	for i := range ids {
		delete(m.token, ids[i])
		m.removedtoken[ids[i]] = struct{}{}
	}
}

// RemovedToken returns the removed IDs of the "token" edge to the Token entity.
func (m *AccountMutation) RemovedTokenIDs() (ids []schema.ID) {
	for id := range m.removedtoken {
		ids = append(ids, id)
	}
	return
}

// TokenIDs returns the "token" edge IDs in the mutation.
func (m *AccountMutation) TokenIDs() (ids []schema.ID) {
	for id := range m.token {
		ids = append(ids, id)
	}
	return
}

// ResetToken resets all changes to the "token" edge.
func (m *AccountMutation) ResetToken() {
	m.token = nil
	m.clearedtoken = false
	m.removedtoken = nil
}

// Where appends a list predicates to the AccountMutation builder.
func (m *AccountMutation) Where(ps ...predicate.Account) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.email != nil {
		fields = append(fields, account.FieldEmail)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldEmail:
		return m.Email()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldEmail:
		return m.OldEmail(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldEmail:
		m.ResetEmail()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.token != nil {
		edges = append(edges, account.EdgeToken)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case account.EdgeToken:
		ids := make([]ent.Value, 0, len(m.token))
		for id := range m.token {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtoken != nil {
		edges = append(edges, account.EdgeToken)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case account.EdgeToken:
		ids := make([]ent.Value, 0, len(m.removedtoken))
		for id := range m.removedtoken {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtoken {
		edges = append(edges, account.EdgeToken)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	switch name {
	case account.EdgeToken:
		return m.clearedtoken
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	switch name {
	case account.EdgeToken:
		m.ResetToken()
		return nil
	}
	return fmt.Errorf("unknown Account edge %s", name)
}

// TokenMutation represents an operation that mutates the Token nodes in the graph.
type TokenMutation struct {
	config
	op             Op
	typ            string
	id             *schema.ID
	body           *string
	clearedFields  map[string]struct{}
	account        *schema.ID
	clearedaccount bool
	done           bool
	oldValue       func(context.Context) (*Token, error)
	predicates     []predicate.Token
}

var _ ent.Mutation = (*TokenMutation)(nil)

// tokenOption allows management of the mutation configuration using functional options.
type tokenOption func(*TokenMutation)

// newTokenMutation creates new mutation for the Token entity.
func newTokenMutation(c config, op Op, opts ...tokenOption) *TokenMutation {
	m := &TokenMutation{
		config:        c,
		op:            op,
		typ:           TypeToken,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTokenID sets the ID field of the mutation.
func withTokenID(id schema.ID) tokenOption {
	return func(m *TokenMutation) {
		var (
			err   error
			once  sync.Once
			value *Token
		)
		m.oldValue = func(ctx context.Context) (*Token, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Token.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withToken sets the old Token of the mutation.
func withToken(node *Token) tokenOption {
	return func(m *TokenMutation) {
		m.oldValue = func(context.Context) (*Token, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TokenMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TokenMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Token entities.
func (m *TokenMutation) SetID(id schema.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TokenMutation) ID() (id schema.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TokenMutation) IDs(ctx context.Context) ([]schema.ID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []schema.ID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Token.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetBody sets the "body" field.
func (m *TokenMutation) SetBody(s string) {
	m.body = &s
}

// Body returns the value of the "body" field in the mutation.
func (m *TokenMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "body" field's value of the Token entity.
// If the Token object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TokenMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody resets all changes to the "body" field.
func (m *TokenMutation) ResetBody() {
	m.body = nil
}

// SetAccountID sets the "account" edge to the Account entity by id.
func (m *TokenMutation) SetAccountID(id schema.ID) {
	m.account = &id
}

// ClearAccount clears the "account" edge to the Account entity.
func (m *TokenMutation) ClearAccount() {
	m.clearedaccount = true
}

// AccountCleared reports if the "account" edge to the Account entity was cleared.
func (m *TokenMutation) AccountCleared() bool {
	return m.clearedaccount
}

// AccountID returns the "account" edge ID in the mutation.
func (m *TokenMutation) AccountID() (id schema.ID, exists bool) {
	if m.account != nil {
		return *m.account, true
	}
	return
}

// AccountIDs returns the "account" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AccountID instead. It exists only for internal usage by the builders.
func (m *TokenMutation) AccountIDs() (ids []schema.ID) {
	if id := m.account; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAccount resets all changes to the "account" edge.
func (m *TokenMutation) ResetAccount() {
	m.account = nil
	m.clearedaccount = false
}

// Where appends a list predicates to the TokenMutation builder.
func (m *TokenMutation) Where(ps ...predicate.Token) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TokenMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Token).
func (m *TokenMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TokenMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.body != nil {
		fields = append(fields, token.FieldBody)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TokenMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case token.FieldBody:
		return m.Body()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TokenMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case token.FieldBody:
		return m.OldBody(ctx)
	}
	return nil, fmt.Errorf("unknown Token field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TokenMutation) SetField(name string, value ent.Value) error {
	switch name {
	case token.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	}
	return fmt.Errorf("unknown Token field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TokenMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TokenMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TokenMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Token numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TokenMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TokenMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TokenMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Token nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TokenMutation) ResetField(name string) error {
	switch name {
	case token.FieldBody:
		m.ResetBody()
		return nil
	}
	return fmt.Errorf("unknown Token field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TokenMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.account != nil {
		edges = append(edges, token.EdgeAccount)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TokenMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case token.EdgeAccount:
		if id := m.account; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TokenMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TokenMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TokenMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedaccount {
		edges = append(edges, token.EdgeAccount)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TokenMutation) EdgeCleared(name string) bool {
	switch name {
	case token.EdgeAccount:
		return m.clearedaccount
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TokenMutation) ClearEdge(name string) error {
	switch name {
	case token.EdgeAccount:
		m.ClearAccount()
		return nil
	}
	return fmt.Errorf("unknown Token unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TokenMutation) ResetEdge(name string) error {
	switch name {
	case token.EdgeAccount:
		m.ResetAccount()
		return nil
	}
	return fmt.Errorf("unknown Token edge %s", name)
}
