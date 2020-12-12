package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// License holds the schema definition for the License entity.
type License struct {
	ent.Schema
}

// Fields of the License.
func (License) Fields() []ent.Field {
	return nil
}

// Edges of the License.
func (License) Edges() []ent.Edge {
	// TODO: pilot
	return []ent.Edge{
		edge.To("pilot", Pilot.Type).
			Unique().StorageKey(edge.Column("pilot_id")),
	}
}
