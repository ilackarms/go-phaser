// Package phaser Automatic generation for Phaser.Rope
// generated file Rope.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// Rope A Rope is a Sprite that has a repeating texture.
//
// The texture will automatically wrap on the edges as it moves.
//
// Please note that Ropes cannot have an input handler.
type Rope struct {
	*js.Object
}

// NewRope A Rope is a Sprite that has a repeating texture.
//
// The texture will automatically wrap on the edges as it moves.
//
// Please note that Ropes cannot have an input handler.
func NewRope(game *Game, x int, y int, key interface{}, frame interface{}, points []interface{}) *Rope {
	return &Rope{js.Global.Get("Phaser").Get("Rope").New(game, x, y, key, frame, points)}
}

// NewRopeI A Rope is a Sprite that has a repeating texture.
//
// The texture will automatically wrap on the edges as it moves.
//
// Please note that Ropes cannot have an input handler.
func NewRopeI(args ...interface{}) *Rope {
	return &Rope{js.Global.Get("Phaser").Get("Rope").New(args...)}
}

// Rope Binding conversion method to Rope point
func ToRope(jsStruct interface{}) *Rope {
	if object, ok := jsStruct.(*js.Object); ok {
		return &Rope{Object: object}
	}
	return nil
}

// Type The const type of this object.
func (self *Rope) Type() int {
	return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Rope) SetTypeA(member int) {
	self.Object.Set("type", member)
}

// UpdateAnimation A Rope will call its updateAnimation function on each update loop if it has one. Set to a function if you'd like the rope to animate during the update phase. Set to false or null to remove it.
func (self *Rope) UpdateAnimation() interface{} {
	return self.Object.Get("updateAnimation")
}

// SetUpdateAnimationA A Rope will call its updateAnimation function on each update loop if it has one. Set to a function if you'd like the rope to animate during the update phase. Set to false or null to remove it.
func (self *Rope) SetUpdateAnimationA(member interface{}) {
	self.Object.Set("updateAnimation", member)
}

// Segments The segments that make up the rope body as an array of Phaser.Rectangles
func (self *Rope) Segments() interface{} {
	return self.Object.Get("segments")
}

// SetSegmentsA The segments that make up the rope body as an array of Phaser.Rectangles
func (self *Rope) SetSegmentsA(member interface{}) {
	self.Object.Set("segments", member)
}

// Texture The texture of the strip
func (self *Rope) Texture() *Texture {
	return &Texture{self.Object.Get("texture")}
}

// SetTextureA The texture of the strip
func (self *Rope) SetTextureA(member *Texture) {
	self.Object.Set("texture", member)
}

// Dirty Whether the strip is dirty or not
func (self *Rope) Dirty() bool {
	return self.Object.Get("dirty").Bool()
}

// SetDirtyA Whether the strip is dirty or not
func (self *Rope) SetDirtyA(member bool) {
	self.Object.Set("dirty", member)
}

// BlendMode The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
func (self *Rope) BlendMode() int {
	return self.Object.Get("blendMode").Int()
}

// SetBlendModeA The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
func (self *Rope) SetBlendModeA(member int) {
	self.Object.Set("blendMode", member)
}

// CanvasPadding Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
func (self *Rope) CanvasPadding() int {
	return self.Object.Get("canvasPadding").Int()
}

// SetCanvasPaddingA Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
func (self *Rope) SetCanvasPaddingA(member int) {
	self.Object.Set("canvasPadding", member)
}

// Children [read-only] The array of children of this container.
func (self *Rope) Children() []DisplayObject {
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *Rope) SetChildrenA(member []DisplayObject) {
	self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
//
//
//
// If this property is `true` then the children will _not_ be considered as valid for Input events.
//
//
//
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Rope) IgnoreChildInput() bool {
	return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
//
//
//
// If this property is `true` then the children will _not_ be considered as valid for Input events.
//
//
//
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Rope) SetIgnoreChildInputA(member bool) {
	self.Object.Set("ignoreChildInput", member)
}

// Width The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Rope) Width() int {
	return self.Object.Get("width").Int()
}

// SetWidthA The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Rope) SetWidthA(member int) {
	self.Object.Set("width", member)
}

// Height The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Rope) Height() int {
	return self.Object.Get("height").Int()
}

// SetHeightA The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Rope) SetHeightA(member int) {
	self.Object.Set("height", member)
}

// Game A reference to the currently running Game.
func (self *Rope) Game() *Game {
	return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Rope) SetGameA(member *Game) {
	self.Object.Set("game", member)
}

// Name A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Rope) Name() string {
	return self.Object.Get("name").String()
}

// SetNameA A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Rope) SetNameA(member string) {
	self.Object.Set("name", member)
}

// Data An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Rope) Data() interface{} {
	return self.Object.Get("data")
}

// SetDataA An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Rope) SetDataA(member interface{}) {
	self.Object.Set("data", member)
}

// Components The components this Game Object has installed.
func (self *Rope) Components() interface{} {
	return self.Object.Get("components")
}

// SetComponentsA The components this Game Object has installed.
func (self *Rope) SetComponentsA(member interface{}) {
	self.Object.Set("components", member)
}

// Z The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Rope) Z() int {
	return self.Object.Get("z").Int()
}

// SetZA The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Rope) SetZA(member int) {
	self.Object.Set("z", member)
}

// Events All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Rope) Events() *Events {
	return &Events{self.Object.Get("events")}
}

// SetEventsA All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Rope) SetEventsA(member *Events) {
	self.Object.Set("events", member)
}

// Animations If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Rope) Animations() *AnimationManager {
	return &AnimationManager{self.Object.Get("animations")}
}

// SetAnimationsA If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Rope) SetAnimationsA(member *AnimationManager) {
	self.Object.Set("animations", member)
}

// Key The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Rope) Key() interface{} {
	return self.Object.Get("key")
}

// SetKeyA The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Rope) SetKeyA(member interface{}) {
	self.Object.Set("key", member)
}

// World The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`,
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Rope) World() *Point {
	return &Point{self.Object.Get("world")}
}

// SetWorldA The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`,
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Rope) SetWorldA(member *Point) {
	self.Object.Set("world", member)
}

// Debug A debug flag designed for use with `Game.enableStep`.
func (self *Rope) Debug() bool {
	return self.Object.Get("debug").Bool()
}

// SetDebugA A debug flag designed for use with `Game.enableStep`.
func (self *Rope) SetDebugA(member bool) {
	self.Object.Set("debug", member)
}

// PreviousPosition The position the Game Object was located in the previous frame.
func (self *Rope) PreviousPosition() *Point {
	return &Point{self.Object.Get("previousPosition")}
}

// SetPreviousPositionA The position the Game Object was located in the previous frame.
func (self *Rope) SetPreviousPositionA(member *Point) {
	self.Object.Set("previousPosition", member)
}

// PreviousRotation The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Rope) PreviousRotation() int {
	return self.Object.Get("previousRotation").Int()
}

// SetPreviousRotationA The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Rope) SetPreviousRotationA(member int) {
	self.Object.Set("previousRotation", member)
}

// RenderOrderID The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Rope) RenderOrderID() int {
	return self.Object.Get("renderOrderID").Int()
}

// SetRenderOrderIDA The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Rope) SetRenderOrderIDA(member int) {
	self.Object.Set("renderOrderID", member)
}

// Fresh A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Rope) Fresh() bool {
	return self.Object.Get("fresh").Bool()
}

// SetFreshA A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Rope) SetFreshA(member bool) {
	self.Object.Set("fresh", member)
}

// PendingDestroy A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
//
// This is extremely useful if you wish to destroy an object from within one of its own callbacks
// such as with Buttons or other Input events.
func (self *Rope) PendingDestroy() bool {
	return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
//
// This is extremely useful if you wish to destroy an object from within one of its own callbacks
// such as with Buttons or other Input events.
func (self *Rope) SetPendingDestroyA(member bool) {
	self.Object.Set("pendingDestroy", member)
}

// Exists Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
//
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Rope) Exists() bool {
	return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
//
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Rope) SetExistsA(member bool) {
	self.Object.Set("exists", member)
}

// Angle The angle property is the rotation of the Game Object in *degrees* from its original orientation.
//
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
//
// Values outside this range are added to or subtracted from 360 to obtain a value within the range.
// For example, the statement player.angle = 450 is the same as player.angle = 90.
//
// If you wish to work in radians instead of degrees you can use the property `rotation` instead.
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *Rope) Angle() int {
	return self.Object.Get("angle").Int()
}

// SetAngleA The angle property is the rotation of the Game Object in *degrees* from its original orientation.
//
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
//
// Values outside this range are added to or subtracted from 360 to obtain a value within the range.
// For example, the statement player.angle = 450 is the same as player.angle = 90.
//
// If you wish to work in radians instead of degrees you can use the property `rotation` instead.
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *Rope) SetAngleA(member int) {
	self.Object.Set("angle", member)
}

// AutoCull A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
//
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Rope) AutoCull() bool {
	return self.Object.Get("autoCull").Bool()
}

// SetAutoCullA A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
//
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Rope) SetAutoCullA(member bool) {
	self.Object.Set("autoCull", member)
}

// InCamera Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Rope) InCamera() bool {
	return self.Object.Get("inCamera").Bool()
}

// SetInCameraA Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Rope) SetInCameraA(member bool) {
	self.Object.Set("inCamera", member)
}

// OffsetX The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Rope) OffsetX() int {
	return self.Object.Get("offsetX").Int()
}

// SetOffsetXA The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Rope) SetOffsetXA(member int) {
	self.Object.Set("offsetX", member)
}

// OffsetY The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Rope) OffsetY() int {
	return self.Object.Get("offsetY").Int()
}

// SetOffsetYA The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Rope) SetOffsetYA(member int) {
	self.Object.Set("offsetY", member)
}

// CenterX The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Rope) CenterX() int {
	return self.Object.Get("centerX").Int()
}

// SetCenterXA The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Rope) SetCenterXA(member int) {
	self.Object.Set("centerX", member)
}

// CenterY The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Rope) CenterY() int {
	return self.Object.Get("centerY").Int()
}

// SetCenterYA The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Rope) SetCenterYA(member int) {
	self.Object.Set("centerY", member)
}

// Left The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Rope) Left() int {
	return self.Object.Get("left").Int()
}

// SetLeftA The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Rope) SetLeftA(member int) {
	self.Object.Set("left", member)
}

// Right The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Rope) Right() int {
	return self.Object.Get("right").Int()
}

// SetRightA The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Rope) SetRightA(member int) {
	self.Object.Set("right", member)
}

// Top The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Rope) Top() int {
	return self.Object.Get("top").Int()
}

// SetTopA The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Rope) SetTopA(member int) {
	self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Rope) Bottom() int {
	return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Rope) SetBottomA(member int) {
	self.Object.Set("bottom", member)
}

// CropRect The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`.
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Rope) CropRect() *Rectangle {
	return &Rectangle{self.Object.Get("cropRect")}
}

// SetCropRectA The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`.
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Rope) SetCropRectA(member *Rectangle) {
	self.Object.Set("cropRect", member)
}

// DeltaX Returns the delta x value. The difference between world.x now and in the previous frame.
//
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *Rope) DeltaX() int {
	return self.Object.Get("deltaX").Int()
}

// SetDeltaXA Returns the delta x value. The difference between world.x now and in the previous frame.
//
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *Rope) SetDeltaXA(member int) {
	self.Object.Set("deltaX", member)
}

// DeltaY Returns the delta y value. The difference between world.y now and in the previous frame.
//
// The value will be positive if the Game Object has moved down or negative if up.
func (self *Rope) DeltaY() int {
	return self.Object.Get("deltaY").Int()
}

// SetDeltaYA Returns the delta y value. The difference between world.y now and in the previous frame.
//
// The value will be positive if the Game Object has moved down or negative if up.
func (self *Rope) SetDeltaYA(member int) {
	self.Object.Set("deltaY", member)
}

// DeltaZ Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *Rope) DeltaZ() int {
	return self.Object.Get("deltaZ").Int()
}

// SetDeltaZA Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *Rope) SetDeltaZA(member int) {
	self.Object.Set("deltaZ", member)
}

// DestroyPhase As a Game Object runs through its destroy method this flag is set to true,
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Rope) DestroyPhase() bool {
	return self.Object.Get("destroyPhase").Bool()
}

// SetDestroyPhaseA As a Game Object runs through its destroy method this flag is set to true,
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Rope) SetDestroyPhaseA(member bool) {
	self.Object.Set("destroyPhase", member)
}

// FixedToCamera A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
//
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
//
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times
// regardless where in the world the camera is.
//
// The offsets are stored in the `cameraOffset` property.
//
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
//
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *Rope) FixedToCamera() bool {
	return self.Object.Get("fixedToCamera").Bool()
}

// SetFixedToCameraA A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
//
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
//
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times
// regardless where in the world the camera is.
//
// The offsets are stored in the `cameraOffset` property.
//
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
//
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *Rope) SetFixedToCameraA(member bool) {
	self.Object.Set("fixedToCamera", member)
}

// CameraOffset The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
//
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Rope) CameraOffset() *Point {
	return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
//
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Rope) SetCameraOffsetA(member *Point) {
	self.Object.Set("cameraOffset", member)
}

// CheckWorldBounds If this is set to `true` the Game Object checks if it is within the World bounds each frame.
//
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
//
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
//
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
//
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
//
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Rope) CheckWorldBounds() bool {
	return self.Object.Get("checkWorldBounds").Bool()
}

// SetCheckWorldBoundsA If this is set to `true` the Game Object checks if it is within the World bounds each frame.
//
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
//
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
//
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
//
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
//
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Rope) SetCheckWorldBoundsA(member bool) {
	self.Object.Set("checkWorldBounds", member)
}

// OutOfBoundsKill If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Rope) OutOfBoundsKill() bool {
	return self.Object.Get("outOfBoundsKill").Bool()
}

// SetOutOfBoundsKillA If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Rope) SetOutOfBoundsKillA(member bool) {
	self.Object.Set("outOfBoundsKill", member)
}

// OutOfCameraBoundsKill If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Rope) OutOfCameraBoundsKill() bool {
	return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// SetOutOfCameraBoundsKillA If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Rope) SetOutOfCameraBoundsKillA(member bool) {
	self.Object.Set("outOfCameraBoundsKill", member)
}

// InWorld Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Rope) InWorld() bool {
	return self.Object.Get("inWorld").Bool()
}

// SetInWorldA Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Rope) SetInWorldA(member bool) {
	self.Object.Set("inWorld", member)
}

// Alive A useful flag to control if the Game Object is alive or dead.
//
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
//
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Rope) Alive() bool {
	return self.Object.Get("alive").Bool()
}

// SetAliveA A useful flag to control if the Game Object is alive or dead.
//
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
//
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Rope) SetAliveA(member bool) {
	self.Object.Set("alive", member)
}

// Lifespan The lifespan allows you to give a Game Object a lifespan in milliseconds.
//
// Once the Game Object is 'born' you can set this to a positive value.
//
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
//
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Rope) Lifespan() int {
	return self.Object.Get("lifespan").Int()
}

// SetLifespanA The lifespan allows you to give a Game Object a lifespan in milliseconds.
//
// Once the Game Object is 'born' you can set this to a positive value.
//
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
//
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Rope) SetLifespanA(member int) {
	self.Object.Set("lifespan", member)
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
func (self *Rope) Frame() int {
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
func (self *Rope) SetFrameA(member int) {
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
func (self *Rope) FrameName() string {
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
func (self *Rope) SetFrameNameA(member string) {
	self.Object.Set("frameName", member)
}

// Body `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated
// properties and methods via it.
//
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
//
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
//
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
//
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5,
// so the physics body is centered on the Game Object.
//
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *Rope) Body() interface{} {
	return self.Object.Get("body")
}

// SetBodyA `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated
// properties and methods via it.
//
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
//
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
//
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
//
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5,
// so the physics body is centered on the Game Object.
//
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *Rope) SetBodyA(member interface{}) {
	self.Object.Set("body", member)
}

// X The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Rope) X() int {
	return self.Object.Get("x").Int()
}

// SetXA The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Rope) SetXA(member int) {
	self.Object.Set("x", member)
}

// Y The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Rope) Y() int {
	return self.Object.Get("y").Int()
}

// SetYA The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Rope) SetYA(member int) {
	self.Object.Set("y", member)
}

// TransformCallback The callback that will apply any scale limiting to the worldTransform.
func (self *Rope) TransformCallback() interface{} {
	return self.Object.Get("transformCallback")
}

// SetTransformCallbackA The callback that will apply any scale limiting to the worldTransform.
func (self *Rope) SetTransformCallbackA(member interface{}) {
	self.Object.Set("transformCallback", member)
}

// TransformCallbackContext The context under which `transformCallback` is called.
func (self *Rope) TransformCallbackContext() interface{} {
	return self.Object.Get("transformCallbackContext")
}

// SetTransformCallbackContextA The context under which `transformCallback` is called.
func (self *Rope) SetTransformCallbackContextA(member interface{}) {
	self.Object.Set("transformCallbackContext", member)
}

// ScaleMin The minimum scale this Game Object will scale down to.
//
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
//
// Set it to `null` to remove the limit.
func (self *Rope) ScaleMin() *Point {
	return &Point{self.Object.Get("scaleMin")}
}

// SetScaleMinA The minimum scale this Game Object will scale down to.
//
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
//
// Set it to `null` to remove the limit.
func (self *Rope) SetScaleMinA(member *Point) {
	self.Object.Set("scaleMin", member)
}

// ScaleMax The maximum scale this Game Object will scale up to.
//
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
//
// Set it to `null` to remove the limit.
func (self *Rope) ScaleMax() *Point {
	return &Point{self.Object.Get("scaleMax")}
}

// SetScaleMaxA The maximum scale this Game Object will scale up to.
//
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
//
// Set it to `null` to remove the limit.
func (self *Rope) SetScaleMaxA(member *Point) {
	self.Object.Set("scaleMax", member)
}

// Smoothed Enable or disable texture smoothing for this Game Object.
//
// It only takes effect if the Game Object is using an image based texture.
//
// Smoothing is enabled by default.
func (self *Rope) Smoothed() bool {
	return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Enable or disable texture smoothing for this Game Object.
//
// It only takes effect if the Game Object is using an image based texture.
//
// Smoothing is enabled by default.
func (self *Rope) SetSmoothedA(member bool) {
	self.Object.Set("smoothed", member)
}

// PreUpdate Automatically called by World.preUpdate.
func (self *Rope) PreUpdate() {
	self.Object.Call("preUpdate")
}

// PreUpdateI Automatically called by World.preUpdate.
func (self *Rope) PreUpdateI(args ...interface{}) {
	self.Object.Call("preUpdate", args...)
}

// Update Override and use this function in your own custom objects to handle any update requirements you may have.
func (self *Rope) Update() {
	self.Object.Call("update")
}

// UpdateI Override and use this function in your own custom objects to handle any update requirements you may have.
func (self *Rope) UpdateI(args ...interface{}) {
	self.Object.Call("update", args...)
}

// Reset Resets the Rope. This places the Rope at the given x/y world coordinates and then
// sets alive, exists, visible and renderable all to true. Also resets the outOfBounds state.
// If the Rope has a physics body that too is reset.
func (self *Rope) Reset(x int, y int) *Rope {
	return &Rope{self.Object.Call("reset", x, y)}
}

// ResetI Resets the Rope. This places the Rope at the given x/y world coordinates and then
// sets alive, exists, visible and renderable all to true. Also resets the outOfBounds state.
// If the Rope has a physics body that too is reset.
func (self *Rope) ResetI(args ...interface{}) *Rope {
	return &Rope{self.Object.Call("reset", args...)}
}

// RenderStripFlat Renders a flat strip
func (self *Rope) RenderStripFlat(strip *Strip) {
	self.Object.Call("renderStripFlat", strip)
}

// RenderStripFlatI Renders a flat strip
func (self *Rope) RenderStripFlatI(args ...interface{}) {
	self.Object.Call("renderStripFlat", args...)
}

// OnTextureUpdate When the texture is updated, this event will fire to update the scale and frame
func (self *Rope) OnTextureUpdate(event interface{}) {
	self.Object.Call("onTextureUpdate", event)
}

// OnTextureUpdateI When the texture is updated, this event will fire to update the scale and frame
func (self *Rope) OnTextureUpdateI(args ...interface{}) {
	self.Object.Call("onTextureUpdate", args...)
}

// GetBounds Returns the bounds of the mesh as a rectangle. The bounds calculation takes the worldTransform into account.
func (self *Rope) GetBounds(matrix *Matrix) *Rectangle {
	return &Rectangle{self.Object.Call("getBounds", matrix)}
}

// GetBoundsI Returns the bounds of the mesh as a rectangle. The bounds calculation takes the worldTransform into account.
func (self *Rope) GetBoundsI(args ...interface{}) *Rectangle {
	return &Rectangle{self.Object.Call("getBounds", args...)}
}

// AddChild Adds a child to the container.
func (self *Rope) AddChild(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *Rope) AddChildI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChild", args...)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Rope) AddChildAt(child *DisplayObject, index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Rope) AddChildAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChildAt", args...)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *Rope) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
	self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *Rope) SwapChildrenI(args ...interface{}) {
	self.Object.Call("swapChildren", args...)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *Rope) GetChildIndex(child *DisplayObject) int {
	return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *Rope) GetChildIndexI(args ...interface{}) int {
	return self.Object.Call("getChildIndex", args...).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *Rope) SetChildIndex(child *DisplayObject, index int) {
	self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *Rope) SetChildIndexI(args ...interface{}) {
	self.Object.Call("setChildIndex", args...)
}

// GetChildAt Returns the child at the specified index
func (self *Rope) GetChildAt(index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *Rope) GetChildAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getChildAt", args...)}
}

// RemoveChild Removes a child from the container.
func (self *Rope) RemoveChild(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *Rope) RemoveChildI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChild", args...)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *Rope) RemoveChildAt(index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *Rope) RemoveChildAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChildAt", args...)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *Rope) RemoveChildren(beginIndex int, endIndex int) {
	self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *Rope) RemoveChildrenI(args ...interface{}) {
	self.Object.Call("removeChildren", args...)
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle without any transformations. The calculation takes all visible children into consideration.
func (self *Rope) GetLocalBounds() *Rectangle {
	return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle without any transformations. The calculation takes all visible children into consideration.
func (self *Rope) GetLocalBoundsI(args ...interface{}) *Rectangle {
	return &Rectangle{self.Object.Call("getLocalBounds", args...)}
}

// Contains Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *Rope) Contains(child *DisplayObject) bool {
	return self.Object.Call("contains", child).Bool()
}

// ContainsI Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *Rope) ContainsI(args ...interface{}) bool {
	return self.Object.Call("contains", args...).Bool()
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *Rope) _renderWebGL(renderSession *RenderSession) {
	self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *Rope) _renderWebGLI(args ...interface{}) {
	self.Object.Call("_renderWebGL", args...)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *Rope) _renderCanvas(renderSession *RenderSession) {
	self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *Rope) _renderCanvasI(args ...interface{}) {
	self.Object.Call("_renderCanvas", args...)
}

// PostUpdate Internal method called by the World postUpdate cycle.
func (self *Rope) PostUpdate() {
	self.Object.Call("postUpdate")
}

// PostUpdateI Internal method called by the World postUpdate cycle.
func (self *Rope) PostUpdateI(args ...interface{}) {
	self.Object.Call("postUpdate", args...)
}

// Play Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Rope) Play(name string) *Animation {
	return &Animation{self.Object.Call("play", name)}
}

// Play1O Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Rope) Play1O(name string, frameRate int) *Animation {
	return &Animation{self.Object.Call("play", name, frameRate)}
}

// Play2O Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Rope) Play2O(name string, frameRate int, loop bool) *Animation {
	return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Play3O Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Rope) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation {
	return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// PlayI Plays an Animation.
//
// The animation should have previously been created via `animations.add`.
//
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Rope) PlayI(args ...interface{}) *Animation {
	return &Animation{self.Object.Call("play", args...)}
}

// AlignIn Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignIn(container interface{}) interface{} {
	return self.Object.Call("alignIn", container)
}

// AlignIn1O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignIn1O(container interface{}, position int) interface{} {
	return self.Object.Call("alignIn", container, position)
}

// AlignIn2O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignIn2O(container interface{}, position int, offsetX int) interface{} {
	return self.Object.Call("alignIn", container, position, offsetX)
}

// AlignIn3O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{} {
	return self.Object.Call("alignIn", container, position, offsetX, offsetY)
}

// AlignInI Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignInI(args ...interface{}) interface{} {
	return self.Object.Call("alignIn", args...)
}

// AlignTo Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignTo(parent interface{}) interface{} {
	return self.Object.Call("alignTo", parent)
}

// AlignTo1O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignTo1O(parent interface{}, position int) interface{} {
	return self.Object.Call("alignTo", parent, position)
}

// AlignTo2O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignTo2O(parent interface{}, position int, offsetX int) interface{} {
	return self.Object.Call("alignTo", parent, position, offsetX)
}

// AlignTo3O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{} {
	return self.Object.Call("alignTo", parent, position, offsetX, offsetY)
}

// AlignToI Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
//
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Rope) AlignToI(args ...interface{}) interface{} {
	return self.Object.Call("alignTo", args...)
}

// BringToTop Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
//
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) BringToTop() *DisplayObject {
	return &DisplayObject{self.Object.Call("bringToTop")}
}

// BringToTopI Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
//
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) BringToTopI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("bringToTop", args...)}
}

// SendToBack Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
//
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) SendToBack() *DisplayObject {
	return &DisplayObject{self.Object.Call("sendToBack")}
}

// SendToBackI Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
//
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) SendToBackI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("sendToBack", args...)}
}

// MoveUp Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
//
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) MoveUp() *DisplayObject {
	return &DisplayObject{self.Object.Call("moveUp")}
}

// MoveUpI Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
//
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) MoveUpI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("moveUp", args...)}
}

// MoveDown Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
//
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) MoveDown() *DisplayObject {
	return &DisplayObject{self.Object.Call("moveDown")}
}

// MoveDownI Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
//
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World,
// because the World is the root Group from which all Game Objects descend.
func (self *Rope) MoveDownI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("moveDown", args...)}
}

// Crop Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
//
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
//
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
//
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`,
// in which case the values are duplicated to a local object.
func (self *Rope) Crop(rect *Rectangle) {
	self.Object.Call("crop", rect)
}

// Crop1O Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
//
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
//
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
//
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`,
// in which case the values are duplicated to a local object.
func (self *Rope) Crop1O(rect *Rectangle, copy bool) {
	self.Object.Call("crop", rect, copy)
}

// CropI Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
//
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
//
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
//
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`,
// in which case the values are duplicated to a local object.
func (self *Rope) CropI(args ...interface{}) {
	self.Object.Call("crop", args...)
}

// UpdateCrop If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Rope) UpdateCrop() {
	self.Object.Call("updateCrop")
}

// UpdateCropI If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Rope) UpdateCropI(args ...interface{}) {
	self.Object.Call("updateCrop", args...)
}

// Destroy Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
//
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
//
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Rope) Destroy() {
	self.Object.Call("destroy")
}

// Destroy1O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
//
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
//
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Rope) Destroy1O(destroyChildren bool) {
	self.Object.Call("destroy", destroyChildren)
}

// Destroy2O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
//
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
//
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Rope) Destroy2O(destroyChildren bool, destroyTexture bool) {
	self.Object.Call("destroy", destroyChildren, destroyTexture)
}

// DestroyI Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
//
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
//
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Rope) DestroyI(args ...interface{}) {
	self.Object.Call("destroy", args...)
}

// Revive Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
//
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
//
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Rope) Revive() *DisplayObject {
	return &DisplayObject{self.Object.Call("revive")}
}

// Revive1O Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
//
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
//
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Rope) Revive1O(health int) *DisplayObject {
	return &DisplayObject{self.Object.Call("revive", health)}
}

// ReviveI Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
//
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
//
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Rope) ReviveI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("revive", args...)}
}

// Kill Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
//
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
//
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
//
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Rope) Kill() *DisplayObject {
	return &DisplayObject{self.Object.Call("kill")}
}

// KillI Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
//
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
//
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
//
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Rope) KillI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("kill", args...)}
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
func (self *Rope) LoadTexture(key interface{}) {
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
func (self *Rope) LoadTexture1O(key interface{}, frame interface{}) {
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
func (self *Rope) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
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
func (self *Rope) LoadTextureI(args ...interface{}) {
	self.Object.Call("loadTexture", args...)
}

// SetFrame Sets the texture frame the Game Object uses for rendering.
//
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Rope) SetFrame(frame *Frame) {
	self.Object.Call("setFrame", frame)
}

// SetFrameI Sets the texture frame the Game Object uses for rendering.
//
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Rope) SetFrameI(args ...interface{}) {
	self.Object.Call("setFrame", args...)
}

// ResizeFrame Resizes the Frame dimensions that the Game Object uses for rendering.
//
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Rope) ResizeFrame(parent interface{}, width int, height int) {
	self.Object.Call("resizeFrame", parent, width, height)
}

// ResizeFrameI Resizes the Frame dimensions that the Game Object uses for rendering.
//
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Rope) ResizeFrameI(args ...interface{}) {
	self.Object.Call("resizeFrame", args...)
}

// ResetFrame Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Rope) ResetFrame() {
	self.Object.Call("resetFrame")
}

// ResetFrameI Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Rope) ResetFrameI(args ...interface{}) {
	self.Object.Call("resetFrame", args...)
}

// Overlap Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object,
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
//
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
//
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Rope) Overlap(displayObject interface{}) bool {
	return self.Object.Call("overlap", displayObject).Bool()
}

// OverlapI Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object,
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
//
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
//
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Rope) OverlapI(args ...interface{}) bool {
	return self.Object.Call("overlap", args...).Bool()
}

// CheckTransform Adjust scaling limits, if set, to this Game Object.
func (self *Rope) CheckTransform(wt *Matrix) {
	self.Object.Call("checkTransform", wt)
}

// CheckTransformI Adjust scaling limits, if set, to this Game Object.
func (self *Rope) CheckTransformI(args ...interface{}) {
	self.Object.Call("checkTransform", args...)
}

// SetScaleMinMax Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
//
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
//
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
//
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
//
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
//
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly,
// or pass `null` for the `maxX` and `maxY` parameters.
//
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *Rope) SetScaleMinMax(minX interface{}, minY interface{}, maxX interface{}, maxY interface{}) {
	self.Object.Call("setScaleMinMax", minX, minY, maxX, maxY)
}

// SetScaleMinMaxI Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
//
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
//
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
//
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
//
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
//
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly,
// or pass `null` for the `maxX` and `maxY` parameters.
//
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *Rope) SetScaleMinMaxI(args ...interface{}) {
	self.Object.Call("setScaleMinMax", args...)
}
