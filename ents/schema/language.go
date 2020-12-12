package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.Text("language"),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
}
