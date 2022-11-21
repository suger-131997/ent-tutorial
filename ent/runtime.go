// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent-tutorial/ent/group"
	"ent-tutorial/ent/pet"
	"ent-tutorial/ent/schema"
	"ent-tutorial/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.DefaultName holds the default value on creation for the name field.
	group.DefaultName = groupDescName.Default.(string)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[0].Descriptor()
	// pet.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	pet.AgeValidator = petDescAge.Validators[0].(func(int) error)
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[1].Descriptor()
	// pet.DefaultName holds the default value on creation for the name field.
	pet.DefaultName = petDescName.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
