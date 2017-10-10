// Package phaser Automatic generation for Phaser.Component.LoadTexture
// generated file ComponentLoadTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// ComponentLoadTexture The LoadTexture component manages the loading of a texture into the Game Object and the changing of frames.
type ComponentLoadTexture struct {
	*js.Object
}

// NewComponentLoadTexture The LoadTexture component manages the loading of a texture into the Game Object and the changing of frames.
func NewComponentLoadTexture() *ComponentLoadTexture {
	return &ComponentLoadTexture{js.Global.Get("Phaser").Get("Component").Get("LoadTexture").New()}
}

// NewComponentLoadTextureI The LoadTexture component manages the loading of a texture into the Game Object and the changing of frames.
func NewComponentLoadTextureI(args ...interface{}) *ComponentLoadTexture {
	return &ComponentLoadTexture{js.Global.Get("Phaser").Get("Component").Get("LoadTexture").New(args...)}
}

// ComponentLoadTexture Binding conversion method to ComponentLoadTexture point
func ToComponentLoadTexture(jsStruct interface{}) *ComponentLoadTexture {
	if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentLoadTexture{Object: object}
	}
	return nil
}

// Frame Gets or sets the current frame index of the texture being used to render this Game Object.
//
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
//
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
//
// If you are using a texture atlas then you should use the `frameName` property instead.
//
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *ComponentLoadTexture) Frame() int {
	return self.Object.Get("frame").Int()
}

// SetFrameA Gets or sets the current frame index of the texture being used to render this Game Object.
//
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
//
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
//
// If you are using a texture atlas then you should use the `frameName` property instead.
//
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *ComponentLoadTexture) SetFrameA(member int) {
	self.Object.Set("frame", member)
}

// FrameName Gets or sets the current frame name of the texture being used to render this Game Object.
//
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use,
// for example: `player.frameName = "idle"`.
//
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
//
// If you are using a sprite sheet then you should use the `frame` property instead.
//
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *ComponentLoadTexture) FrameName() string {
	return self.Object.Get("frameName").String()
}

// SetFrameNameA Gets or sets the current frame name of the texture being used to render this Game Object.
//
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use,
// for example: `player.frameName = "idle"`.
//
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
//
// If you are using a sprite sheet then you should use the `frame` property instead.
//
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *ComponentLoadTexture) SetFrameNameA(member string) {
	self.Object.Set("frameName", member)
}

// LoadTexture Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
//
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
//
// You should only use `loadTexture` if you want to replace the base texture entirely.
//
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
//
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite.
// Doing this then sets the key to be the `frame` argument (the frame is set to zero).
//
// This allows you to create sprites using `load.image` during development, and then change them
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS'
// and swapping it to be the key of the atlas data.
//
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *ComponentLoadTexture) LoadTexture(key interface{}) {
	self.Object.Call("loadTexture", key)
}

// LoadTexture1O Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
//
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
//
// You should only use `loadTexture` if you want to replace the base texture entirely.
//
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
//
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite.
// Doing this then sets the key to be the `frame` argument (the frame is set to zero).
//
// This allows you to create sprites using `load.image` during development, and then change them
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS'
// and swapping it to be the key of the atlas data.
//
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *ComponentLoadTexture) LoadTexture1O(key interface{}, frame interface{}) {
	self.Object.Call("loadTexture", key, frame)
}

// LoadTexture2O Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
//
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
//
// You should only use `loadTexture` if you want to replace the base texture entirely.
//
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
//
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite.
// Doing this then sets the key to be the `frame` argument (the frame is set to zero).
//
// This allows you to create sprites using `load.image` during development, and then change them
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS'
// and swapping it to be the key of the atlas data.
//
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *ComponentLoadTexture) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
	self.Object.Call("loadTexture", key, frame, stopAnimation)
}

// LoadTextureI Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
//
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
//
// You should only use `loadTexture` if you want to replace the base texture entirely.
//
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
//
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite.
// Doing this then sets the key to be the `frame` argument (the frame is set to zero).
//
// This allows you to create sprites using `load.image` during development, and then change them
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS'
// and swapping it to be the key of the atlas data.
//
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *ComponentLoadTexture) LoadTextureI(args ...interface{}) {
	self.Object.Call("loadTexture", args...)
}

// SetFrame Sets the texture frame the Game Object uses for rendering.
//
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *ComponentLoadTexture) SetFrame(frame *Frame) {
	self.Object.Call("setFrame", frame)
}

// SetFrameI Sets the texture frame the Game Object uses for rendering.
//
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *ComponentLoadTexture) SetFrameI(args ...interface{}) {
	self.Object.Call("setFrame", args...)
}

// ResizeFrame Resizes the Frame dimensions that the Game Object uses for rendering.
//
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *ComponentLoadTexture) ResizeFrame(parent interface{}, width int, height int) {
	self.Object.Call("resizeFrame", parent, width, height)
}

// ResizeFrameI Resizes the Frame dimensions that the Game Object uses for rendering.
//
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *ComponentLoadTexture) ResizeFrameI(args ...interface{}) {
	self.Object.Call("resizeFrame", args...)
}

// ResetFrame Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *ComponentLoadTexture) ResetFrame() {
	self.Object.Call("resetFrame")
}

// ResetFrameI Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *ComponentLoadTexture) ResetFrameI(args ...interface{}) {
	self.Object.Call("resetFrame", args...)
}
