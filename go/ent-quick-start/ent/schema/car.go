package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("model").NotEmpty().Comment("モデル名"),
		field.Time("registered_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return nil
}
