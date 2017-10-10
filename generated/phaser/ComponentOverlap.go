// Package phaser Automatic generation for Phaser.Component.Overlap
// generated file ComponentOverlap.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// ComponentOverlap The Overlap component allows a Game Object to check if it overlaps with the bounds of another Game Object.
type ComponentOverlap struct {
	*js.Object
}

// NewComponentOverlap The Overlap component allows a Game Object to check if it overlaps with the bounds of another Game Object.
func NewComponentOverlap() *ComponentOverlap {
	return &ComponentOverlap{js.Global.Get("Phaser").Get("Component").Get("Overlap").New()}
}

// NewComponentOverlapI The Overlap component allows a Game Object to check if it overlaps with the bounds of another Game Object.
func NewComponentOverlapI(args ...interface{}) *ComponentOverlap {
	return &ComponentOverlap{js.Global.Get("Phaser").Get("Component").Get("Overlap").New(args...)}
}

// ComponentOverlap Binding conversion method to ComponentOverlap point
func ToComponentOverlap(jsStruct interface{}) *ComponentOverlap {
	if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentOverlap{Object: object}
	}
	return nil
}

// Overlap Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object,
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
//
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
//
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *ComponentOverlap) Overlap(displayObject interface{}) bool {
	return self.Object.Call("overlap", displayObject).Bool()
}

// OverlapI Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object,
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
//
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
//
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *ComponentOverlap) OverlapI(args ...interface{}) bool {
	return self.Object.Call("overlap", args...).Bool()
}
