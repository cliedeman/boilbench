package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Jet holds the schema definition for the Jet entity.
type Jet struct {
	ent.Schema
}

// Fields of the Jet.
func (Jet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Text("name").Optional(),
		field.Text("color").Optional(),
		// TODO: change to uuid
		field.Text("uuid").Optional(),
		field.Text("identifier").Optional(),
		field.Bytes("cargo").Optional(),
		field.Bytes("manifest").Optional(),
	}
}

// Edges of the Jet.
func (Jet) Edges() []ent.Edge {
	// TODO: pilot
	// TODO: airport
	return []ent.Edge{
		edge.From("airport", Airport.Type).
			Ref("jets").
			Unique(),
		edge.From("pilot", Pilot.Type).
			Ref("jets").
			Unique(),
	}
}
