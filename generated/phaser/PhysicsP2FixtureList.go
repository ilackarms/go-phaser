// Package phaser Automatic generation for Phaser.Physics.P2.FixtureList
// generated file PhysicsP2FixtureList.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// PhysicsP2FixtureList Allow to access a list of created fixture (coming from Body#addPhaserPolygon)
// which itself parse the input from PhysicsEditor with the custom phaser exporter.
// You can access fixtures of a Body by a group index or even by providing a fixture Key.
// You can set the fixture key and also the group index for a fixture in PhysicsEditor.
// This gives you the power to create a complex body built of many fixtures and modify them
// during runtime (to remove parts, set masks, categories & sensor properties)
type PhysicsP2FixtureList struct {
	*js.Object
}

// NewPhysicsP2FixtureList Allow to access a list of created fixture (coming from Body#addPhaserPolygon)
// which itself parse the input from PhysicsEditor with the custom phaser exporter.
// You can access fixtures of a Body by a group index or even by providing a fixture Key.
// You can set the fixture key and also the group index for a fixture in PhysicsEditor.
// This gives you the power to create a complex body built of many fixtures and modify them
// during runtime (to remove parts, set masks, categories & sensor properties)
func NewPhysicsP2FixtureList(list []interface{}) *PhysicsP2FixtureList {
	return &PhysicsP2FixtureList{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("FixtureList").New(list)}
}

// NewPhysicsP2FixtureListI Allow to access a list of created fixture (coming from Body#addPhaserPolygon)
// which itself parse the input from PhysicsEditor with the custom phaser exporter.
// You can access fixtures of a Body by a group index or even by providing a fixture Key.
// You can set the fixture key and also the group index for a fixture in PhysicsEditor.
// This gives you the power to create a complex body built of many fixtures and modify them
// during runtime (to remove parts, set masks, categories & sensor properties)
func NewPhysicsP2FixtureListI(args ...interface{}) *PhysicsP2FixtureList {
	return &PhysicsP2FixtureList{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("FixtureList").New(args...)}
}

// PhysicsP2FixtureList Binding conversion method to PhysicsP2FixtureList point
func ToPhysicsP2FixtureList(jsStruct interface{}) *PhysicsP2FixtureList {
	if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2FixtureList{Object: object}
	}
	return nil
}

// Init empty description
func (self *PhysicsP2FixtureList) Init() {
	self.Object.Call("init")
}

// InitI empty description
func (self *PhysicsP2FixtureList) InitI(args ...interface{}) {
	self.Object.Call("init", args...)
}

// SetCategory empty description
func (self *PhysicsP2FixtureList) SetCategory(bit int, fixtureKey string) {
	self.Object.Call("setCategory", bit, fixtureKey)
}

// SetCategoryI empty description
func (self *PhysicsP2FixtureList) SetCategoryI(args ...interface{}) {
	self.Object.Call("setCategory", args...)
}

// SetMask empty description
func (self *PhysicsP2FixtureList) SetMask(bit int, fixtureKey string) {
	self.Object.Call("setMask", bit, fixtureKey)
}

// SetMaskI empty description
func (self *PhysicsP2FixtureList) SetMaskI(args ...interface{}) {
	self.Object.Call("setMask", args...)
}

// SetSensor empty description
func (self *PhysicsP2FixtureList) SetSensor(value bool, fixtureKey string) {
	self.Object.Call("setSensor", value, fixtureKey)
}

// SetSensorI empty description
func (self *PhysicsP2FixtureList) SetSensorI(args ...interface{}) {
	self.Object.Call("setSensor", args...)
}

// SetMaterial empty description
func (self *PhysicsP2FixtureList) SetMaterial(material interface{}, fixtureKey string) {
	self.Object.Call("setMaterial", material, fixtureKey)
}

// SetMaterialI empty description
func (self *PhysicsP2FixtureList) SetMaterialI(args ...interface{}) {
	self.Object.Call("setMaterial", args...)
}

// GetFixtures Accessor to get either a list of specified fixtures by key or the whole fixture list
func (self *PhysicsP2FixtureList) GetFixtures(keys []interface{}) {
	self.Object.Call("getFixtures", keys)
}

// GetFixturesI Accessor to get either a list of specified fixtures by key or the whole fixture list
func (self *PhysicsP2FixtureList) GetFixturesI(args ...interface{}) {
	self.Object.Call("getFixtures", args...)
}

// GetFixtureByKey Accessor to get either a single fixture by its key.
func (self *PhysicsP2FixtureList) GetFixtureByKey(key string) {
	self.Object.Call("getFixtureByKey", key)
}

// GetFixtureByKeyI Accessor to get either a single fixture by its key.
func (self *PhysicsP2FixtureList) GetFixtureByKeyI(args ...interface{}) {
	self.Object.Call("getFixtureByKey", args...)
}

// GetGroup Accessor to get a group of fixtures by its group index.
func (self *PhysicsP2FixtureList) GetGroup(groupID int) {
	self.Object.Call("getGroup", groupID)
}

// GetGroupI Accessor to get a group of fixtures by its group index.
func (self *PhysicsP2FixtureList) GetGroupI(args ...interface{}) {
	self.Object.Call("getGroup", args...)
}

// Parse Parser for the output of Phaser.Physics.P2.Body#addPhaserPolygon
func (self *PhysicsP2FixtureList) Parse() {
	self.Object.Call("parse")
}

// ParseI Parser for the output of Phaser.Physics.P2.Body#addPhaserPolygon
func (self *PhysicsP2FixtureList) ParseI(args ...interface{}) {
	self.Object.Call("parse", args...)
}

// Flatten A helper to flatten arrays. This is very useful as the fixtures are nested from time to time due to the way P2 creates and splits polygons.
func (self *PhysicsP2FixtureList) Flatten(array []interface{}) {
	self.Object.Call("flatten", array)
}

// FlattenI A helper to flatten arrays. This is very useful as the fixtures are nested from time to time due to the way P2 creates and splits polygons.
func (self *PhysicsP2FixtureList) FlattenI(args ...interface{}) {
	self.Object.Call("flatten", args...)
}
