// Package phaser Automatic generation for Phaser.Physics.P2.RotationalSpring
// generated file PhysicsP2RotationalSpring.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// PhysicsP2RotationalSpring Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
type PhysicsP2RotationalSpring struct {
	*js.Object
}

// NewPhysicsP2RotationalSpring Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2RotationalSpring(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2RotationalSpring {
	return &PhysicsP2RotationalSpring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("RotationalSpring").New(world, bodyA, bodyB)}
}

// NewPhysicsP2RotationalSpring1O Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2RotationalSpring1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restAngle int) *PhysicsP2RotationalSpring {
	return &PhysicsP2RotationalSpring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("RotationalSpring").New(world, bodyA, bodyB, restAngle)}
}

// NewPhysicsP2RotationalSpring2O Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2RotationalSpring2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restAngle int, stiffness int) *PhysicsP2RotationalSpring {
	return &PhysicsP2RotationalSpring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("RotationalSpring").New(world, bodyA, bodyB, restAngle, stiffness)}
}

// NewPhysicsP2RotationalSpring3O Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2RotationalSpring3O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restAngle int, stiffness int, damping int) *PhysicsP2RotationalSpring {
	return &PhysicsP2RotationalSpring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("RotationalSpring").New(world, bodyA, bodyB, restAngle, stiffness, damping)}
}

// NewPhysicsP2RotationalSpringI Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2RotationalSpringI(args ...interface{}) *PhysicsP2RotationalSpring {
	return &PhysicsP2RotationalSpring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("RotationalSpring").New(args...)}
}

// PhysicsP2RotationalSpring Binding conversion method to PhysicsP2RotationalSpring point
func ToPhysicsP2RotationalSpring(jsStruct interface{}) *PhysicsP2RotationalSpring {
	if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2RotationalSpring{Object: object}
	}
	return nil
}

// Game Local reference to game.
func (self *PhysicsP2RotationalSpring) Game() *Game {
	return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsP2RotationalSpring) SetGameA(member *Game) {
	self.Object.Set("game", member)
}

// World Local reference to P2 World.
func (self *PhysicsP2RotationalSpring) World() *PhysicsP2 {
	return &PhysicsP2{self.Object.Get("world")}
}

// SetWorldA Local reference to P2 World.
func (self *PhysicsP2RotationalSpring) SetWorldA(member *PhysicsP2) {
	self.Object.Set("world", member)
}

// Data The actual p2 spring object.
func (self *PhysicsP2RotationalSpring) Data() *P2RotationalSpring {
	return &P2RotationalSpring{self.Object.Get("data")}
}

// SetDataA The actual p2 spring object.
func (self *PhysicsP2RotationalSpring) SetDataA(member *P2RotationalSpring) {
	self.Object.Set("data", member)
}
