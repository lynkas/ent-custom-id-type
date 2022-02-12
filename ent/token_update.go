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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetBody sets the "body" field.
func (tu *TokenUpdate) SetBody(s string) *TokenUpdate {
	tu.mutation.SetBody(s)
	return tu
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (tu *TokenUpdate) SetAccountID(id schema.ID) *TokenUpdate {
	tu.mutation.SetAccountID(id)
	return tu
}

// SetAccount sets the "account" edge to the Account entity.
func (tu *TokenUpdate) SetAccount(a *Account) *TokenUpdate {
	return tu.SetAccountID(a.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (tu *TokenUpdate) ClearAccount() *TokenUpdate {
	tu.mutation.ClearAccount()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokenUpdate) check() error {
	if v, ok := tu.mutation.Body(); ok {
		if err := token.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Token.body": %w`, err)}
		}
	}
	if _, ok := tu.mutation.AccountID(); tu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Token.account"`)
	}
	return nil
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: token.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldBody,
		})
	}
	if tu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.AccountTable,
			Columns: []string{token.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeOther,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.AccountTable,
			Columns: []string{token.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeOther,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenMutation
}

// SetBody sets the "body" field.
func (tuo *TokenUpdateOne) SetBody(s string) *TokenUpdateOne {
	tuo.mutation.SetBody(s)
	return tuo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (tuo *TokenUpdateOne) SetAccountID(id schema.ID) *TokenUpdateOne {
	tuo.mutation.SetAccountID(id)
	return tuo
}

// SetAccount sets the "account" edge to the Account entity.
func (tuo *TokenUpdateOne) SetAccount(a *Account) *TokenUpdateOne {
	return tuo.SetAccountID(a.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (tuo *TokenUpdateOne) ClearAccount() *TokenUpdateOne {
	tuo.mutation.ClearAccount()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	var (
		err  error
		node *Token
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokenUpdateOne) check() error {
	if v, ok := tuo.mutation.Body(); ok {
		if err := token.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Token.body": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.AccountID(); tuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Token.account"`)
	}
	return nil
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: token.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldBody,
		})
	}
	if tuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.AccountTable,
			Columns: []string{token.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeOther,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.AccountTable,
			Columns: []string{token.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeOther,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
