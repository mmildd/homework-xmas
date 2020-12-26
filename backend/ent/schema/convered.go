package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// covered holds the schema definition for the covered entity.
type covered struct {
	ent.Schema
}

// Fields of the covered
func (covered) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the covered
func (covered) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playlist", Playlist.Type).Ref("covered").Unique(),
		edge.From("video", Video.Type).Ref("covered").Unique(),
		edge.From("resolution", Resolution.Type).Ref("covered").Unique(),
	}
}
