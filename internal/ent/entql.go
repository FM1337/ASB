// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/FM1337/ASB/internal/ent/predicate"
	"github.com/FM1337/ASB/internal/ent/server"
	"github.com/FM1337/ASB/internal/ent/serverconfig"
	"github.com/FM1337/ASB/internal/ent/wordblacklist"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   server.Table,
			Columns: server.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: server.FieldID,
			},
		},
		Type: "Server",
		Fields: map[string]*sqlgraph.FieldSpec{
			server.FieldCreateTime: {Type: field.TypeTime, Column: server.FieldCreateTime},
			server.FieldUpdateTime: {Type: field.TypeTime, Column: server.FieldUpdateTime},
			server.FieldOwnerID:    {Type: field.TypeString, Column: server.FieldOwnerID},
			server.FieldEnabled:    {Type: field.TypeBool, Column: server.FieldEnabled},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   serverconfig.Table,
			Columns: serverconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serverconfig.FieldID,
			},
		},
		Type: "ServerConfig",
		Fields: map[string]*sqlgraph.FieldSpec{
			serverconfig.FieldCreateTime:           {Type: field.TypeTime, Column: serverconfig.FieldCreateTime},
			serverconfig.FieldUpdateTime:           {Type: field.TypeTime, Column: serverconfig.FieldUpdateTime},
			serverconfig.FieldRemoveRoles:          {Type: field.TypeBool, Column: serverconfig.FieldRemoveRoles},
			serverconfig.FieldGiveRole:             {Type: field.TypeBool, Column: serverconfig.FieldGiveRole},
			serverconfig.FieldTimeout:              {Type: field.TypeBool, Column: serverconfig.FieldTimeout},
			serverconfig.FieldKick:                 {Type: field.TypeBool, Column: serverconfig.FieldKick},
			serverconfig.FieldBan:                  {Type: field.TypeBool, Column: serverconfig.FieldBan},
			serverconfig.FieldCheckInvites:         {Type: field.TypeBool, Column: serverconfig.FieldCheckInvites},
			serverconfig.FieldCheckLinks:           {Type: field.TypeBool, Column: serverconfig.FieldCheckLinks},
			serverconfig.FieldRatelimit:            {Type: field.TypeBool, Column: serverconfig.FieldRatelimit},
			serverconfig.FieldAlerts:               {Type: field.TypeBool, Column: serverconfig.FieldAlerts},
			serverconfig.FieldFlagLinks:            {Type: field.TypeBool, Column: serverconfig.FieldFlagLinks},
			serverconfig.FieldLogChannel:           {Type: field.TypeString, Column: serverconfig.FieldLogChannel},
			serverconfig.FieldExcludedChannels:     {Type: field.TypeJSON, Column: serverconfig.FieldExcludedChannels},
			serverconfig.FieldExcludedRoles:        {Type: field.TypeJSON, Column: serverconfig.FieldExcludedRoles},
			serverconfig.FieldExcludedUsers:        {Type: field.TypeJSON, Column: serverconfig.FieldExcludedUsers},
			serverconfig.FieldGivenRole:            {Type: field.TypeString, Column: serverconfig.FieldGivenRole},
			serverconfig.FieldRatelimitMessage:     {Type: field.TypeInt, Column: serverconfig.FieldRatelimitMessage},
			serverconfig.FieldRatelimitTime:        {Type: field.TypeEnum, Column: serverconfig.FieldRatelimitTime},
			serverconfig.FieldTimeoutTime:          {Type: field.TypeEnum, Column: serverconfig.FieldTimeoutTime},
			serverconfig.FieldBanDeleteMessageTime: {Type: field.TypeEnum, Column: serverconfig.FieldBanDeleteMessageTime},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   wordblacklist.Table,
			Columns: wordblacklist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wordblacklist.FieldID,
			},
		},
		Type: "WordBlacklist",
		Fields: map[string]*sqlgraph.FieldSpec{
			wordblacklist.FieldCreateTime: {Type: field.TypeTime, Column: wordblacklist.FieldCreateTime},
			wordblacklist.FieldUpdateTime: {Type: field.TypeTime, Column: wordblacklist.FieldUpdateTime},
			wordblacklist.FieldWord:       {Type: field.TypeString, Column: wordblacklist.FieldWord},
		},
	}
	graph.MustAddE(
		"configuration",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   server.ConfigurationTable,
			Columns: []string{server.ConfigurationColumn},
			Bidi:    false,
		},
		"Server",
		"ServerConfig",
	)
	graph.MustAddE(
		"word_blacklist",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   server.WordBlacklistTable,
			Columns: server.WordBlacklistPrimaryKey,
			Bidi:    false,
		},
		"Server",
		"WordBlacklist",
	)
	graph.MustAddE(
		"server",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   serverconfig.ServerTable,
			Columns: []string{serverconfig.ServerColumn},
			Bidi:    false,
		},
		"ServerConfig",
		"Server",
	)
	graph.MustAddE(
		"server",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wordblacklist.ServerTable,
			Columns: wordblacklist.ServerPrimaryKey,
			Bidi:    false,
		},
		"WordBlacklist",
		"Server",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (sq *ServerQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ServerQuery builder.
func (sq *ServerQuery) Filter() *ServerFilter {
	return &ServerFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *ServerMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ServerMutation builder.
func (m *ServerMutation) Filter() *ServerFilter {
	return &ServerFilter{config: m.config, predicateAdder: m}
}

// ServerFilter provides a generic filtering capability at runtime for ServerQuery.
type ServerFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ServerFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *ServerFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(server.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *ServerFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(server.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *ServerFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(server.FieldUpdateTime))
}

// WhereOwnerID applies the entql string predicate on the owner_id field.
func (f *ServerFilter) WhereOwnerID(p entql.StringP) {
	f.Where(p.Field(server.FieldOwnerID))
}

// WhereEnabled applies the entql bool predicate on the enabled field.
func (f *ServerFilter) WhereEnabled(p entql.BoolP) {
	f.Where(p.Field(server.FieldEnabled))
}

// WhereHasConfiguration applies a predicate to check if query has an edge configuration.
func (f *ServerFilter) WhereHasConfiguration() {
	f.Where(entql.HasEdge("configuration"))
}

// WhereHasConfigurationWith applies a predicate to check if query has an edge configuration with a given conditions (other predicates).
func (f *ServerFilter) WhereHasConfigurationWith(preds ...predicate.ServerConfig) {
	f.Where(entql.HasEdgeWith("configuration", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasWordBlacklist applies a predicate to check if query has an edge word_blacklist.
func (f *ServerFilter) WhereHasWordBlacklist() {
	f.Where(entql.HasEdge("word_blacklist"))
}

// WhereHasWordBlacklistWith applies a predicate to check if query has an edge word_blacklist with a given conditions (other predicates).
func (f *ServerFilter) WhereHasWordBlacklistWith(preds ...predicate.WordBlacklist) {
	f.Where(entql.HasEdgeWith("word_blacklist", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (scq *ServerConfigQuery) addPredicate(pred func(s *sql.Selector)) {
	scq.predicates = append(scq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ServerConfigQuery builder.
func (scq *ServerConfigQuery) Filter() *ServerConfigFilter {
	return &ServerConfigFilter{config: scq.config, predicateAdder: scq}
}

// addPredicate implements the predicateAdder interface.
func (m *ServerConfigMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ServerConfigMutation builder.
func (m *ServerConfigMutation) Filter() *ServerConfigFilter {
	return &ServerConfigFilter{config: m.config, predicateAdder: m}
}

// ServerConfigFilter provides a generic filtering capability at runtime for ServerConfigQuery.
type ServerConfigFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ServerConfigFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ServerConfigFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(serverconfig.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *ServerConfigFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(serverconfig.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *ServerConfigFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(serverconfig.FieldUpdateTime))
}

// WhereRemoveRoles applies the entql bool predicate on the remove_roles field.
func (f *ServerConfigFilter) WhereRemoveRoles(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldRemoveRoles))
}

// WhereGiveRole applies the entql bool predicate on the give_role field.
func (f *ServerConfigFilter) WhereGiveRole(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldGiveRole))
}

// WhereTimeout applies the entql bool predicate on the timeout field.
func (f *ServerConfigFilter) WhereTimeout(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldTimeout))
}

// WhereKick applies the entql bool predicate on the kick field.
func (f *ServerConfigFilter) WhereKick(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldKick))
}

// WhereBan applies the entql bool predicate on the ban field.
func (f *ServerConfigFilter) WhereBan(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldBan))
}

// WhereCheckInvites applies the entql bool predicate on the check_invites field.
func (f *ServerConfigFilter) WhereCheckInvites(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldCheckInvites))
}

// WhereCheckLinks applies the entql bool predicate on the check_links field.
func (f *ServerConfigFilter) WhereCheckLinks(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldCheckLinks))
}

// WhereRatelimit applies the entql bool predicate on the ratelimit field.
func (f *ServerConfigFilter) WhereRatelimit(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldRatelimit))
}

// WhereAlerts applies the entql bool predicate on the alerts field.
func (f *ServerConfigFilter) WhereAlerts(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldAlerts))
}

// WhereFlagLinks applies the entql bool predicate on the flag_links field.
func (f *ServerConfigFilter) WhereFlagLinks(p entql.BoolP) {
	f.Where(p.Field(serverconfig.FieldFlagLinks))
}

// WhereLogChannel applies the entql string predicate on the log_channel field.
func (f *ServerConfigFilter) WhereLogChannel(p entql.StringP) {
	f.Where(p.Field(serverconfig.FieldLogChannel))
}

// WhereExcludedChannels applies the entql json.RawMessage predicate on the excluded_channels field.
func (f *ServerConfigFilter) WhereExcludedChannels(p entql.BytesP) {
	f.Where(p.Field(serverconfig.FieldExcludedChannels))
}

// WhereExcludedRoles applies the entql json.RawMessage predicate on the excluded_roles field.
func (f *ServerConfigFilter) WhereExcludedRoles(p entql.BytesP) {
	f.Where(p.Field(serverconfig.FieldExcludedRoles))
}

// WhereExcludedUsers applies the entql json.RawMessage predicate on the excluded_users field.
func (f *ServerConfigFilter) WhereExcludedUsers(p entql.BytesP) {
	f.Where(p.Field(serverconfig.FieldExcludedUsers))
}

// WhereGivenRole applies the entql string predicate on the given_role field.
func (f *ServerConfigFilter) WhereGivenRole(p entql.StringP) {
	f.Where(p.Field(serverconfig.FieldGivenRole))
}

// WhereRatelimitMessage applies the entql int predicate on the ratelimit_message field.
func (f *ServerConfigFilter) WhereRatelimitMessage(p entql.IntP) {
	f.Where(p.Field(serverconfig.FieldRatelimitMessage))
}

// WhereRatelimitTime applies the entql string predicate on the ratelimit_time field.
func (f *ServerConfigFilter) WhereRatelimitTime(p entql.StringP) {
	f.Where(p.Field(serverconfig.FieldRatelimitTime))
}

// WhereTimeoutTime applies the entql string predicate on the timeout_time field.
func (f *ServerConfigFilter) WhereTimeoutTime(p entql.StringP) {
	f.Where(p.Field(serverconfig.FieldTimeoutTime))
}

// WhereBanDeleteMessageTime applies the entql string predicate on the ban_delete_message_time field.
func (f *ServerConfigFilter) WhereBanDeleteMessageTime(p entql.StringP) {
	f.Where(p.Field(serverconfig.FieldBanDeleteMessageTime))
}

// WhereHasServer applies a predicate to check if query has an edge server.
func (f *ServerConfigFilter) WhereHasServer() {
	f.Where(entql.HasEdge("server"))
}

// WhereHasServerWith applies a predicate to check if query has an edge server with a given conditions (other predicates).
func (f *ServerConfigFilter) WhereHasServerWith(preds ...predicate.Server) {
	f.Where(entql.HasEdgeWith("server", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (wbq *WordBlacklistQuery) addPredicate(pred func(s *sql.Selector)) {
	wbq.predicates = append(wbq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the WordBlacklistQuery builder.
func (wbq *WordBlacklistQuery) Filter() *WordBlacklistFilter {
	return &WordBlacklistFilter{config: wbq.config, predicateAdder: wbq}
}

// addPredicate implements the predicateAdder interface.
func (m *WordBlacklistMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the WordBlacklistMutation builder.
func (m *WordBlacklistMutation) Filter() *WordBlacklistFilter {
	return &WordBlacklistFilter{config: m.config, predicateAdder: m}
}

// WordBlacklistFilter provides a generic filtering capability at runtime for WordBlacklistQuery.
type WordBlacklistFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *WordBlacklistFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *WordBlacklistFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(wordblacklist.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *WordBlacklistFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(wordblacklist.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *WordBlacklistFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(wordblacklist.FieldUpdateTime))
}

// WhereWord applies the entql string predicate on the word field.
func (f *WordBlacklistFilter) WhereWord(p entql.StringP) {
	f.Where(p.Field(wordblacklist.FieldWord))
}

// WhereHasServer applies a predicate to check if query has an edge server.
func (f *WordBlacklistFilter) WhereHasServer() {
	f.Where(entql.HasEdge("server"))
}

// WhereHasServerWith applies a predicate to check if query has an edge server with a given conditions (other predicates).
func (f *WordBlacklistFilter) WhereHasServerWith(preds ...predicate.Server) {
	f.Where(entql.HasEdgeWith("server", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
