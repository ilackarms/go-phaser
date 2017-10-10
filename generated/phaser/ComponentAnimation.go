// Package phaser Automatic generation for Phaser.Component.Animation
// generated file ComponentAnimation.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// ComponentAnimation The Animation Component provides a `play` method, which is a proxy to the `AnimationManager.play` method.
type ComponentAnimation struct {
	*js.Object
}

// NewComponentAnimation The Animation Component provides a `play` method, which is a proxy to the `AnimationManager.play` method.
func NewComponentAnimation() *ComponentAnimation {
	return &ComponentAnimation{js.Global.Get("Phaser").Get("Component").Get("Animation").New()}
}

// NewComponentAnimationI The Animation Component provides a `play` method, which is a proxy to the `AnimationManager.play` method.
func NewComponentAnimationI(args ...interface{}) *ComponentAnimation {
	return &ComponentAnimation{js.Global.Get("Phaser").Get("Component").Get("Animation").New(args...)}
}

// ComponentAnimation Binding conversion method to ComponentAnimation point
func ToComponentAnimation(jsStruct interface{}) *ComponentAnimation {
	if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentAnimation{Object: object}
	}
	return nil
}

// Play Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play(name string) *Animation {
	return &Animation{self.Object.Call("play", name)}
}

// Play1O Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play1O(name string, frameRate int) *Animation {
	return &Animation{self.Object.Call("play", name, frameRate)}
}

// Play2O Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play2O(name string, frameRate int, loop bool) *Animation {
	return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Play3O Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation {
	return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// PlayI Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) PlayI(args ...interface{}) *Animation {
	return &Animation{self.Object.Call("play", args...)}
}
