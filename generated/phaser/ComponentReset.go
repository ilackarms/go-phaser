// Package phaser Automatic generation for Phaser.Component.Reset
// generated file ComponentReset.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// ComponentReset The Reset component allows a Game Object to be reset and repositioned to a new location.
type ComponentReset struct {
	*js.Object
}

// NewComponentReset The Reset component allows a Game Object to be reset and repositioned to a new location.
func NewComponentReset() *ComponentReset {
	return &ComponentReset{js.Global.Get("Phaser").Get("Component").Get("Reset").New()}
}

// NewComponentResetI The Reset component allows a Game Object to be reset and repositioned to a new location.
func NewComponentResetI(args ...interface{}) *ComponentReset {
	return &ComponentReset{js.Global.Get("Phaser").Get("Component").Get("Reset").New(args...)}
}

// ComponentReset Binding conversion method to ComponentReset point
func ToComponentReset(jsStruct interface{}) *ComponentReset {
	if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentReset{Object: object}
	}
	return nil
}

// Reset Resets the Game Object.
//
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`,
// `visible` and `renderable` to true.
//
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
//
// If this Game Object has a Physics Body it will reset the Body.
func (self *ComponentReset) Reset(x int, y int) *DisplayObject {
	return &DisplayObject{self.Object.Call("reset", x, y)}
}

// Reset1O Resets the Game Object.
//
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`,
// `visible` and `renderable` to true.
//
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
//
// If this Game Object has a Physics Body it will reset the Body.
func (self *ComponentReset) Reset1O(x int, y int, health int) *DisplayObject {
	return &DisplayObject{self.Object.Call("reset", x, y, health)}
}

// ResetI Resets the Game Object.
//
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`,
// `visible` and `renderable` to true.
//
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
//
// If this Game Object has a Physics Body it will reset the Body.
func (self *ComponentReset) ResetI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("reset", args...)}
}
