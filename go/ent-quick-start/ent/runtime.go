// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entqs/ent/car"
	"entqs/ent/group"
	"entqs/ent/schema"
	"entqs/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescModel is the schema descriptor for model field.
	carDescModel := carFields[1].Descriptor()
	// car.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	car.ModelValidator = carDescModel.Validators[0].(func(string) error)
	// carDescRegisteredAt is the schema descriptor for registered_at field.
	carDescRegisteredAt := carFields[2].Descriptor()
	// car.DefaultRegisteredAt holds the default value on creation for the registered_at field.
	car.DefaultRegisteredAt = carDescRegisteredAt.Default.(func() time.Time)
	// carDescID is the schema descriptor for id field.
	carDescID := carFields[0].Descriptor()
	// car.DefaultID holds the default value on creation for the id field.
	car.DefaultID = carDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
