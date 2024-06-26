// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/FM1337/ASB/internal/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The CooldownQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CooldownQueryRuleFunc func(context.Context, *ent.CooldownQuery) error

// EvalQuery return f(ctx, q).
func (f CooldownQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CooldownQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CooldownQuery", q)
}

// The CooldownMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CooldownMutationRuleFunc func(context.Context, *ent.CooldownMutation) error

// EvalMutation calls f(ctx, m).
func (f CooldownMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CooldownMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CooldownMutation", m)
}

// The ServerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ServerQueryRuleFunc func(context.Context, *ent.ServerQuery) error

// EvalQuery return f(ctx, q).
func (f ServerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ServerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ServerQuery", q)
}

// The ServerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ServerMutationRuleFunc func(context.Context, *ent.ServerMutation) error

// EvalMutation calls f(ctx, m).
func (f ServerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ServerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ServerMutation", m)
}

// The ServerConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ServerConfigQueryRuleFunc func(context.Context, *ent.ServerConfigQuery) error

// EvalQuery return f(ctx, q).
func (f ServerConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ServerConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ServerConfigQuery", q)
}

// The ServerConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ServerConfigMutationRuleFunc func(context.Context, *ent.ServerConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f ServerConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ServerConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ServerConfigMutation", m)
}

// The SpammerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SpammerQueryRuleFunc func(context.Context, *ent.SpammerQuery) error

// EvalQuery return f(ctx, q).
func (f SpammerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SpammerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SpammerQuery", q)
}

// The SpammerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SpammerMutationRuleFunc func(context.Context, *ent.SpammerMutation) error

// EvalMutation calls f(ctx, m).
func (f SpammerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SpammerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SpammerMutation", m)
}

// The WordBlacklistQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type WordBlacklistQueryRuleFunc func(context.Context, *ent.WordBlacklistQuery) error

// EvalQuery return f(ctx, q).
func (f WordBlacklistQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WordBlacklistQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.WordBlacklistQuery", q)
}

// The WordBlacklistMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type WordBlacklistMutationRuleFunc func(context.Context, *ent.WordBlacklistMutation) error

// EvalMutation calls f(ctx, m).
func (f WordBlacklistMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.WordBlacklistMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.WordBlacklistMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.CooldownQuery:
		return q.Filter(), nil
	case *ent.ServerQuery:
		return q.Filter(), nil
	case *ent.ServerConfigQuery:
		return q.Filter(), nil
	case *ent.SpammerQuery:
		return q.Filter(), nil
	case *ent.WordBlacklistQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.CooldownMutation:
		return m.Filter(), nil
	case *ent.ServerMutation:
		return m.Filter(), nil
	case *ent.ServerConfigMutation:
		return m.Filter(), nil
	case *ent.SpammerMutation:
		return m.Filter(), nil
	case *ent.WordBlacklistMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
