package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Hangar holds the schema definition for the Hangar entity.
type Hangar struct {
	ent.Schema
}

// Fields of the Hangar.
func (Hangar) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Hangar.
func (Hangar) Edges() []ent.Edge {
	return nil
}
