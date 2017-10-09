// Package phaser Automatic generation for Phaser.Physics.Arcade.TilemapCollision
// generated file PhysicsArcadeTilemapCollision.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// PhysicsArcadeTilemapCollision The Arcade Physics Tile map collision methods.
type PhysicsArcadeTilemapCollision struct {
	*js.Object
}

// NewPhysicsArcadeTilemapCollision The Arcade Physics Tile map collision methods.
func NewPhysicsArcadeTilemapCollision() *PhysicsArcadeTilemapCollision {
	return &PhysicsArcadeTilemapCollision{js.Global.Get("Phaser").Get("Physics").Get("Arcade").Get("TilemapCollision").New()}
}

// NewPhysicsArcadeTilemapCollisionI The Arcade Physics Tile map collision methods.
func NewPhysicsArcadeTilemapCollisionI(args ...interface{}) *PhysicsArcadeTilemapCollision {
	return &PhysicsArcadeTilemapCollision{js.Global.Get("Phaser").Get("Physics").Get("Arcade").Get("TilemapCollision").New(args)}
}

// PhysicsArcadeTilemapCollision Binding conversion method to PhysicsArcadeTilemapCollision point
func ToPhysicsArcadeTilemapCollision(jsStruct interface{}) *PhysicsArcadeTilemapCollision {
	if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsArcadeTilemapCollision{Object: object}
	}
	return nil
}

// TILE_BIAS A value added to the delta values during collision with tiles. Adjust this if you get tunneling.
func (self *PhysicsArcadeTilemapCollision) TILE_BIAS() int {
	return self.Object.Get("TILE_BIAS").Int()
}

// SetTILE_BIASA A value added to the delta values during collision with tiles. Adjust this if you get tunneling.
func (self *PhysicsArcadeTilemapCollision) SetTILE_BIASA(member int) {
	self.Object.Set("TILE_BIAS", member)
}
