package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Pilot holds the schema definition for the Pilot entity.
type Pilot struct {
	ent.Schema
}

// Fields of the Pilot.
func (Pilot) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Pilot.
func (Pilot) Edges() []ent.Edge {
	// join pilot_languages
	return []ent.Edge{
		edge.To("jets", Jet.Type).
			StorageKey(edge.Column("pilot_id")),
	}
}
