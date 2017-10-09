// Package phaser Automatic generation for Phaser.Component.Health
// generated file ComponentHealth.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// ComponentHealth The Health component provides the ability for Game Objects to have a `health` property
// that can be damaged and reset through game code.
// Requires the LifeSpan component.
type ComponentHealth struct {
	*js.Object
}

// NewComponentHealth The Health component provides the ability for Game Objects to have a `health` property
// that can be damaged and reset through game code.
// Requires the LifeSpan component.
func NewComponentHealth() *ComponentHealth {
	return &ComponentHealth{js.Global.Get("Phaser").Get("Component").Get("Health").New()}
}

// NewComponentHealthI The Health component provides the ability for Game Objects to have a `health` property
// that can be damaged and reset through game code.
// Requires the LifeSpan component.
func NewComponentHealthI(args ...interface{}) *ComponentHealth {
	return &ComponentHealth{js.Global.Get("Phaser").Get("Component").Get("Health").New(args)}
}

// ComponentHealth Binding conversion method to ComponentHealth point
func ToComponentHealth(jsStruct interface{}) *ComponentHealth {
	if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentHealth{Object: object}
	}
	return nil
}

// Health The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
//
// It can be used in combination with the `damage` method or modified directly.
func (self *ComponentHealth) Health() int {
	return self.Object.Get("health").Int()
}

// SetHealthA The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
//
// It can be used in combination with the `damage` method or modified directly.
func (self *ComponentHealth) SetHealthA(member int) {
	self.Object.Set("health", member)
}

// MaxHealth The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *ComponentHealth) MaxHealth() int {
	return self.Object.Get("maxHealth").Int()
}

// SetMaxHealthA The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *ComponentHealth) SetMaxHealthA(member int) {
	self.Object.Set("maxHealth", member)
}

// Damage Damages the Game Object. This removes the given amount of health from the `health` property.
//
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *ComponentHealth) Damage() interface{} {
	return self.Object.Get("damage")
}

// SetDamageA Damages the Game Object. This removes the given amount of health from the `health` property.
//
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *ComponentHealth) SetDamageA(member interface{}) {
	self.Object.Set("damage", member)
}

// SetHealth Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *ComponentHealth) SetHealth() interface{} {
	return self.Object.Get("setHealth")
}

// SetSetHealthA Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *ComponentHealth) SetSetHealthA(member interface{}) {
	self.Object.Set("setHealth", member)
}

// Heal Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *ComponentHealth) Heal() interface{} {
	return self.Object.Get("heal")
}

// SetHealA Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *ComponentHealth) SetHealA(member interface{}) {
	self.Object.Set("heal", member)
}
