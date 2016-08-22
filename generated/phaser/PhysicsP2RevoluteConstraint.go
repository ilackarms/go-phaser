// Automatic generation for Phaser.Physics.P2.RevoluteConstraint
// generated file PhysicsP2RevoluteConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
type PhysicsP2RevoluteConstraint struct {
    *js.Object
}


// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func NewPhysicsP2RevoluteConstraint(world *PhysicsP2, bodyA *P2Body, pivotA *Float32Array, bodyB *P2Body, pivotB *Float32Array) *PhysicsP2RevoluteConstraint {
    return &PhysicsP2RevoluteConstraint{js.Global.Call("Phaser.Physics.P2.RevoluteConstraint", world, bodyA, pivotA, bodyB, pivotB)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func NewPhysicsP2RevoluteConstraint1O(world *PhysicsP2, bodyA *P2Body, pivotA *Float32Array, bodyB *P2Body, pivotB *Float32Array, maxForce int) *PhysicsP2RevoluteConstraint {
    return &PhysicsP2RevoluteConstraint{js.Global.Call("Phaser.Physics.P2.RevoluteConstraint", world, bodyA, pivotA, bodyB, pivotB, maxForce)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func NewPhysicsP2RevoluteConstraint2O(world *PhysicsP2, bodyA *P2Body, pivotA *Float32Array, bodyB *P2Body, pivotB *Float32Array, maxForce int, worldPivot *Float32Array) *PhysicsP2RevoluteConstraint {
    return &PhysicsP2RevoluteConstraint{js.Global.Call("Phaser.Physics.P2.RevoluteConstraint", world, bodyA, pivotA, bodyB, pivotB, maxForce, worldPivot)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func NewPhysicsP2RevoluteConstraintI(args ...interface{}) *PhysicsP2RevoluteConstraint {
    return &PhysicsP2RevoluteConstraint{js.Global.Call("Phaser.Physics.P2.RevoluteConstraint", args)}
}



// Local reference to game.
func (self *PhysicsP2RevoluteConstraint) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2RevoluteConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2RevoluteConstraint) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2RevoluteConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


