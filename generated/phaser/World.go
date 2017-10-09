// Package phaser Automatic generation for Phaser.World
// generated file World.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// World "This world is but a canvas to our imagination." - Henry David Thoreau
//
// A game has only one world. The world is an abstract place in which all game objects live. It is not bound
// by stage limits and can be any size. You look into the world via cameras. All game objects live within
// the world at world-based coordinates. By default a world is created the same size as your Stage.
type World struct {
	*js.Object
}

// NewWorld "This world is but a canvas to our imagination." - Henry David Thoreau
//
// A game has only one world. The world is an abstract place in which all game objects live. It is not bound
// by stage limits and can be any size. You look into the world via cameras. All game objects live within
// the world at world-based coordinates. By default a world is created the same size as your Stage.
func NewWorld(game *Game) *World {
	return &World{js.Global.Get("Phaser").Get("World").New(game)}
}

// NewWorldI "This world is but a canvas to our imagination." - Henry David Thoreau
//
// A game has only one world. The world is an abstract place in which all game objects live. It is not bound
// by stage limits and can be any size. You look into the world via cameras. All game objects live within
// the world at world-based coordinates. By default a world is created the same size as your Stage.
func NewWorldI(args ...interface{}) *World {
	return &World{js.Global.Get("Phaser").Get("World").New(args)}
}

// World Binding conversion method to World point
func ToWorld(jsStruct interface{}) *World {
	if object, ok := jsStruct.(*js.Object); ok {
		return &World{Object: object}
	}
	return nil
}

// Bounds The World has no fixed size, but it does have a bounds outside of which objects are no longer considered as being "in world" and you should use this to clean-up the display list and purge dead objects.
// By default we set the Bounds to be from 0,0 to Game.width,Game.height. I.e. it will match the size given to the game constructor with 0,0 representing the top-left of the display.
// However 0,0 is actually the center of the world, and if you rotate or scale the world all of that will happen from 0,0.
// So if you want to make a game in which the world itself will rotate you should adjust the bounds so that 0,0 is the center point, i.e. set them to -1000,-1000,2000,2000 for a 2000x2000 sized world centered around 0,0. Bound of this world that objects can not escape from.
func (self *World) Bounds() *Rectangle {
	return &Rectangle{self.Object.Get("bounds")}
}

// SetBoundsA The World has no fixed size, but it does have a bounds outside of which objects are no longer considered as being "in world" and you should use this to clean-up the display list and purge dead objects.
// By default we set the Bounds to be from 0,0 to Game.width,Game.height. I.e. it will match the size given to the game constructor with 0,0 representing the top-left of the display.
// However 0,0 is actually the center of the world, and if you rotate or scale the world all of that will happen from 0,0.
// So if you want to make a game in which the world itself will rotate you should adjust the bounds so that 0,0 is the center point, i.e. set them to -1000,-1000,2000,2000 for a 2000x2000 sized world centered around 0,0. Bound of this world that objects can not escape from.
func (self *World) SetBoundsA(member *Rectangle) {
	self.Object.Set("bounds", member)
}

// Camera Camera instance.
func (self *World) Camera() *Camera {
	return &Camera{self.Object.Get("camera")}
}

// SetCameraA Camera instance.
func (self *World) SetCameraA(member *Camera) {
	self.Object.Set("camera", member)
}

// _definedSize True if the World has been given a specifically defined size (i.e. from a Tilemap or direct in code) or false if it's just matched to the Game dimensions.
func (self *World) _definedSize() bool {
	return self.Object.Get("_definedSize").Bool()
}

// Set_definedSizeA True if the World has been given a specifically defined size (i.e. from a Tilemap or direct in code) or false if it's just matched to the Game dimensions.
func (self *World) Set_definedSizeA(member bool) {
	self.Object.Set("_definedSize", member)
}

// _width The defined width of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) _width() interface{} {
	return self.Object.Get("_width")
}

// Set_widthA The defined width of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) Set_widthA(member interface{}) {
	self.Object.Set("_width", member)
}

// _height The defined height of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) _height() interface{} {
	return self.Object.Get("_height")
}

// Set_heightA The defined height of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) Set_heightA(member interface{}) {
	self.Object.Set("_height", member)
}

// Width Gets or sets the current width of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) Width() int {
	return self.Object.Get("width").Int()
}

// SetWidthA Gets or sets the current width of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) SetWidthA(member int) {
	self.Object.Set("width", member)
}

// Height Gets or sets the current height of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) Height() int {
	return self.Object.Get("height").Int()
}

// SetHeightA Gets or sets the current height of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) SetHeightA(member int) {
	self.Object.Set("height", member)
}

// CenterX Gets the X position corresponding to the center point of the world.
func (self *World) CenterX() int {
	return self.Object.Get("centerX").Int()
}

// SetCenterXA Gets the X position corresponding to the center point of the world.
func (self *World) SetCenterXA(member int) {
	self.Object.Set("centerX", member)
}

// CenterY Gets the Y position corresponding to the center point of the world.
func (self *World) CenterY() int {
	return self.Object.Get("centerY").Int()
}

// SetCenterYA Gets the Y position corresponding to the center point of the world.
func (self *World) SetCenterYA(member int) {
	self.Object.Set("centerY", member)
}

// RandomX Gets a random integer which is lesser than or equal to the current width of the game world.
func (self *World) RandomX() int {
	return self.Object.Get("randomX").Int()
}

// SetRandomXA Gets a random integer which is lesser than or equal to the current width of the game world.
func (self *World) SetRandomXA(member int) {
	self.Object.Set("randomX", member)
}

// RandomY Gets a random integer which is lesser than or equal to the current height of the game world.
func (self *World) RandomY() int {
	return self.Object.Get("randomY").Int()
}

// SetRandomYA Gets a random integer which is lesser than or equal to the current height of the game world.
func (self *World) SetRandomYA(member int) {
	self.Object.Set("randomY", member)
}

// X The x coordinate of the group container.
//
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) X() int {
	return self.Object.Get("x").Int()
}

// SetXA The x coordinate of the group container.
//
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) SetXA(member int) {
	self.Object.Set("x", member)
}

// Y The y coordinate of the group container.
//
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) Y() int {
	return self.Object.Get("y").Int()
}

// SetYA The y coordinate of the group container.
//
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) SetYA(member int) {
	self.Object.Set("y", member)
}

// Rotation The angle of rotation of the group container, in radians.
//
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *World) Rotation() int {
	return self.Object.Get("rotation").Int()
}

// SetRotationA The angle of rotation of the group container, in radians.
//
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *World) SetRotationA(member int) {
	self.Object.Set("rotation", member)
}

// Visible The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *World) Visible() bool {
	return self.Object.Get("visible").Bool()
}

// SetVisibleA The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *World) SetVisibleA(member bool) {
	self.Object.Set("visible", member)
}

// Alpha The alpha value of the group container.
func (self *World) Alpha() int {
	return self.Object.Get("alpha").Int()
}

// SetAlphaA The alpha value of the group container.
func (self *World) SetAlphaA(member int) {
	self.Object.Set("alpha", member)
}

// Game A reference to the currently running Game.
func (self *World) Game() *Game {
	return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *World) SetGameA(member *Game) {
	self.Object.Set("game", member)
}

// Name A name for this group. Not used internally but useful for debugging.
func (self *World) Name() string {
	return self.Object.Get("name").String()
}

// SetNameA A name for this group. Not used internally but useful for debugging.
func (self *World) SetNameA(member string) {
	self.Object.Set("name", member)
}

// Z The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *World) Z() int {
	return self.Object.Get("z").Int()
}

// SetZA The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *World) SetZA(member int) {
	self.Object.Set("z", member)
}

// Type Internal Phaser Type value.
func (self *World) Type() int {
	return self.Object.Get("type").Int()
}

// SetTypeA Internal Phaser Type value.
func (self *World) SetTypeA(member int) {
	self.Object.Set("type", member)
}

// PhysicsType The const physics body type of this object.
func (self *World) PhysicsType() int {
	return self.Object.Get("physicsType").Int()
}

// SetPhysicsTypeA The const physics body type of this object.
func (self *World) SetPhysicsTypeA(member int) {
	self.Object.Set("physicsType", member)
}

// Alive The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *World) Alive() bool {
	return self.Object.Get("alive").Bool()
}

// SetAliveA The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *World) SetAliveA(member bool) {
	self.Object.Set("alive", member)
}

// Exists If exists is true the group is updated, otherwise it is skipped.
func (self *World) Exists() bool {
	return self.Object.Get("exists").Bool()
}

// SetExistsA If exists is true the group is updated, otherwise it is skipped.
func (self *World) SetExistsA(member bool) {
	self.Object.Set("exists", member)
}

// IgnoreDestroy A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *World) IgnoreDestroy() bool {
	return self.Object.Get("ignoreDestroy").Bool()
}

// SetIgnoreDestroyA A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *World) SetIgnoreDestroyA(member bool) {
	self.Object.Set("ignoreDestroy", member)
}

// PendingDestroy A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
//
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *World) PendingDestroy() bool {
	return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
//
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *World) SetPendingDestroyA(member bool) {
	self.Object.Set("pendingDestroy", member)
}

// ClassType The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
//
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *World) ClassType() interface{} {
	return self.Object.Get("classType")
}

// SetClassTypeA The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
//
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *World) SetClassTypeA(member interface{}) {
	self.Object.Set("classType", member)
}

// Cursor The current display object that the group cursor is pointing to, if any. (Can be set manually.)
//
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *World) Cursor() *DisplayObject {
	return &DisplayObject{self.Object.Get("cursor")}
}

// SetCursorA The current display object that the group cursor is pointing to, if any. (Can be set manually.)
//
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *World) SetCursorA(member *DisplayObject) {
	self.Object.Set("cursor", member)
}

// InputEnableChildren A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true`
// on any children _added_ to, or _created by_, this Group.
//
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *World) InputEnableChildren() bool {
	return self.Object.Get("inputEnableChildren").Bool()
}

// SetInputEnableChildrenA A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true`
// on any children _added_ to, or _created by_, this Group.
//
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *World) SetInputEnableChildrenA(member bool) {
	self.Object.Set("inputEnableChildren", member)
}

// OnChildInputDown This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) OnChildInputDown() *Signal {
	return &Signal{self.Object.Get("onChildInputDown")}
}

// SetOnChildInputDownA This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) SetOnChildInputDownA(member *Signal) {
	self.Object.Set("onChildInputDown", member)
}

// OnChildInputUp This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal,
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *World) OnChildInputUp() *Signal {
	return &Signal{self.Object.Get("onChildInputUp")}
}

// SetOnChildInputUpA This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal,
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *World) SetOnChildInputUpA(member *Signal) {
	self.Object.Set("onChildInputUp", member)
}

// OnChildInputOver This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) OnChildInputOver() *Signal {
	return &Signal{self.Object.Get("onChildInputOver")}
}

// SetOnChildInputOverA This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) SetOnChildInputOverA(member *Signal) {
	self.Object.Set("onChildInputOver", member)
}

// OnChildInputOut This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) OnChildInputOut() *Signal {
	return &Signal{self.Object.Get("onChildInputOut")}
}

// SetOnChildInputOutA This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
//
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) SetOnChildInputOutA(member *Signal) {
	self.Object.Set("onChildInputOut", member)
}

// EnableBody If true all Sprites created by, or added to this group, will have a physics body enabled on them.
//
// If there are children already in the Group at the time you set this property, they are not changed.
//
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *World) EnableBody() bool {
	return self.Object.Get("enableBody").Bool()
}

// SetEnableBodyA If true all Sprites created by, or added to this group, will have a physics body enabled on them.
//
// If there are children already in the Group at the time you set this property, they are not changed.
//
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *World) SetEnableBodyA(member bool) {
	self.Object.Set("enableBody", member)
}

// EnableBodyDebug If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
//
// This only works for P2 bodies.
func (self *World) EnableBodyDebug() bool {
	return self.Object.Get("enableBodyDebug").Bool()
}

// SetEnableBodyDebugA If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
//
// This only works for P2 bodies.
func (self *World) SetEnableBodyDebugA(member bool) {
	self.Object.Set("enableBodyDebug", member)
}

// PhysicsBodyType If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
//
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *World) PhysicsBodyType() int {
	return self.Object.Get("physicsBodyType").Int()
}

// SetPhysicsBodyTypeA If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
//
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *World) SetPhysicsBodyTypeA(member int) {
	self.Object.Set("physicsBodyType", member)
}

// PhysicsSortDirection If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
//
// It should be set to one of the Phaser.Physics.Arcade sort direction constants:
//
// Phaser.Physics.Arcade.SORT_NONE
// Phaser.Physics.Arcade.LEFT_RIGHT
// Phaser.Physics.Arcade.RIGHT_LEFT
// Phaser.Physics.Arcade.TOP_BOTTOM
// Phaser.Physics.Arcade.BOTTOM_TOP
//
// If set to `null` the Group will use whatever Phaser.Physics.Arcade.sortDirection is set to. This is the default behavior.
func (self *World) PhysicsSortDirection() int {
	return self.Object.Get("physicsSortDirection").Int()
}

// SetPhysicsSortDirectionA If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
//
// It should be set to one of the Phaser.Physics.Arcade sort direction constants:
//
// Phaser.Physics.Arcade.SORT_NONE
// Phaser.Physics.Arcade.LEFT_RIGHT
// Phaser.Physics.Arcade.RIGHT_LEFT
// Phaser.Physics.Arcade.TOP_BOTTOM
// Phaser.Physics.Arcade.BOTTOM_TOP
//
// If set to `null` the Group will use whatever Phaser.Physics.Arcade.sortDirection is set to. This is the default behavior.
func (self *World) SetPhysicsSortDirectionA(member int) {
	self.Object.Set("physicsSortDirection", member)
}

// OnDestroy This signal is dispatched when the group is destroyed.
func (self *World) OnDestroy() *Signal {
	return &Signal{self.Object.Get("onDestroy")}
}

// SetOnDestroyA This signal is dispatched when the group is destroyed.
func (self *World) SetOnDestroyA(member *Signal) {
	self.Object.Set("onDestroy", member)
}

// CursorIndex The current index of the Group cursor. Advance it with Group.next.
func (self *World) CursorIndex() int {
	return self.Object.Get("cursorIndex").Int()
}

// SetCursorIndexA The current index of the Group cursor. Advance it with Group.next.
func (self *World) SetCursorIndexA(member int) {
	self.Object.Set("cursorIndex", member)
}

// FixedToCamera A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
//
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *World) FixedToCamera() bool {
	return self.Object.Get("fixedToCamera").Bool()
}

// SetFixedToCameraA A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
//
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *World) SetFixedToCameraA(member bool) {
	self.Object.Set("fixedToCamera", member)
}

// CameraOffset If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *World) CameraOffset() *Point {
	return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *World) SetCameraOffsetA(member *Point) {
	self.Object.Set("cameraOffset", member)
}

// Hash The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
//
// Only children of this Group can be added to and removed from the hash.
//
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *World) Hash() []interface{} {
	array00 := self.Object.Get("hash")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetHashA The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
//
// Only children of this Group can be added to and removed from the hash.
//
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *World) SetHashA(member []interface{}) {
	self.Object.Set("hash", member)
}

// Total Total number of existing children in the group.
func (self *World) Total() int {
	return self.Object.Get("total").Int()
}

// SetTotalA Total number of existing children in the group.
func (self *World) SetTotalA(member int) {
	self.Object.Set("total", member)
}

// Length Total number of children in this group, regardless of exists/alive status.
func (self *World) Length() int {
	return self.Object.Get("length").Int()
}

// SetLengthA Total number of children in this group, regardless of exists/alive status.
func (self *World) SetLengthA(member int) {
	self.Object.Set("length", member)
}

// Angle The angle of rotation of the group container, in degrees.
//
// This adjusts the group itself by modifying its local rotation transform.
//
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *World) Angle() int {
	return self.Object.Get("angle").Int()
}

// SetAngleA The angle of rotation of the group container, in degrees.
//
// This adjusts the group itself by modifying its local rotation transform.
//
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *World) SetAngleA(member int) {
	self.Object.Set("angle", member)
}

// Left The left coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) Left() int {
	return self.Object.Get("left").Int()
}

// SetLeftA The left coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) SetLeftA(member int) {
	self.Object.Set("left", member)
}

// Right The right coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) Right() int {
	return self.Object.Get("right").Int()
}

// SetRightA The right coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) SetRightA(member int) {
	self.Object.Set("right", member)
}

// Top The top coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) Top() int {
	return self.Object.Get("top").Int()
}

// SetTopA The top coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) SetTopA(member int) {
	self.Object.Set("top", member)
}

// Bottom The bottom coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) Bottom() int {
	return self.Object.Get("bottom").Int()
}

// SetBottomA The bottom coordinate of this Group.
//
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
func (self *World) SetBottomA(member int) {
	self.Object.Set("bottom", member)
}

// Children [read-only] The array of children of this container.
func (self *World) Children() []DisplayObject {
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *World) SetChildrenA(member []DisplayObject) {
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
func (self *World) IgnoreChildInput() bool {
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
func (self *World) SetIgnoreChildInputA(member bool) {
	self.Object.Set("ignoreChildInput", member)
}

// Boot Initialises the game world.
func (self *World) Boot() {
	self.Object.Call("boot")
}

// BootI Initialises the game world.
func (self *World) BootI(args ...interface{}) {
	self.Object.Call("boot", args)
}

// StateChange Called whenever the State changes or resets.
//
// It resets the world.x and world.y coordinates back to zero,
// then resets the Camera.
func (self *World) StateChange() {
	self.Object.Call("stateChange")
}

// StateChangeI Called whenever the State changes or resets.
//
// It resets the world.x and world.y coordinates back to zero,
// then resets the Camera.
func (self *World) StateChangeI(args ...interface{}) {
	self.Object.Call("stateChange", args)
}

// SetBounds Updates the size of this world and sets World.x/y to the given values
// The Camera bounds and Physics bounds (if set) are also updated to match the new World bounds.
func (self *World) SetBounds(x int, y int, width int, height int) {
	self.Object.Call("setBounds", x, y, width, height)
}

// SetBoundsI Updates the size of this world and sets World.x/y to the given values
// The Camera bounds and Physics bounds (if set) are also updated to match the new World bounds.
func (self *World) SetBoundsI(args ...interface{}) {
	self.Object.Call("setBounds", args)
}

// Resize Updates the size of this world. Note that this doesn't modify the world x/y coordinates, just the width and height.
func (self *World) Resize(width int, height int) {
	self.Object.Call("resize", width, height)
}

// ResizeI Updates the size of this world. Note that this doesn't modify the world x/y coordinates, just the width and height.
func (self *World) ResizeI(args ...interface{}) {
	self.Object.Call("resize", args)
}

// Shutdown Destroyer of worlds.
func (self *World) Shutdown() {
	self.Object.Call("shutdown")
}

// ShutdownI Destroyer of worlds.
func (self *World) ShutdownI(args ...interface{}) {
	self.Object.Call("shutdown", args)
}

// Wrap This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
//
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) Wrap(sprite interface{}) {
	self.Object.Call("wrap", sprite)
}

// Wrap1O This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
//
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) Wrap1O(sprite interface{}, padding int) {
	self.Object.Call("wrap", sprite, padding)
}

// Wrap2O This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
//
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) Wrap2O(sprite interface{}, padding int, useBounds bool) {
	self.Object.Call("wrap", sprite, padding, useBounds)
}

// Wrap3O This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
//
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) Wrap3O(sprite interface{}, padding int, useBounds bool, horizontal bool) {
	self.Object.Call("wrap", sprite, padding, useBounds, horizontal)
}

// Wrap4O This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
//
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) Wrap4O(sprite interface{}, padding int, useBounds bool, horizontal bool, vertical bool) {
	self.Object.Call("wrap", sprite, padding, useBounds, horizontal, vertical)
}

// WrapI This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
//
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) WrapI(args ...interface{}) {
	self.Object.Call("wrap", args)
}

// AlignIn Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignIn(container interface{}) *Group {
	return &Group{self.Object.Call("alignIn", container)}
}

// AlignIn1O Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignIn1O(container interface{}, position int) *Group {
	return &Group{self.Object.Call("alignIn", container, position)}
}

// AlignIn2O Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignIn2O(container interface{}, position int, offsetX int) *Group {
	return &Group{self.Object.Call("alignIn", container, position, offsetX)}
}

// AlignIn3O Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) *Group {
	return &Group{self.Object.Call("alignIn", container, position, offsetX, offsetY)}
}

// AlignInI Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
//
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`,
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignInI(args ...interface{}) *Group {
	return &Group{self.Object.Call("alignIn", args)}
}

// AlignTo Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignTo(parent interface{}) *Group {
	return &Group{self.Object.Call("alignTo", parent)}
}

// AlignTo1O Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignTo1O(parent interface{}, position int) *Group {
	return &Group{self.Object.Call("alignTo", parent, position)}
}

// AlignTo2O Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignTo2O(parent interface{}, position int, offsetX int) *Group {
	return &Group{self.Object.Call("alignTo", parent, position, offsetX)}
}

// AlignTo3O Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) *Group {
	return &Group{self.Object.Call("alignTo", parent, position, offsetX, offsetY)}
}

// AlignToI Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
//
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
//
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
//
// The position constants you can use are:
//
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`,
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`,
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER`
// and `Phaser.BOTTOM_RIGHT`.
//
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
//
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
//
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
//
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *World) AlignToI(args ...interface{}) *Group {
	return &Group{self.Object.Call("alignTo", args)}
}

// Add Adds an existing object as the top child in this group.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
//
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *World) Add(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("add", child)}
}

// Add1O Adds an existing object as the top child in this group.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
//
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *World) Add1O(child *DisplayObject, silent bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("add", child, silent)}
}

// Add2O Adds an existing object as the top child in this group.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
//
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *World) Add2O(child *DisplayObject, silent bool, index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("add", child, silent, index)}
}

// AddI Adds an existing object as the top child in this group.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
//
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *World) AddI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("add", args)}
}

// AddAt Adds an existing object to this group.
//
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) AddAt(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("addAt", child)}
}

// AddAt1O Adds an existing object to this group.
//
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) AddAt1O(child *DisplayObject, index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("addAt", child, index)}
}

// AddAt2O Adds an existing object to this group.
//
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) AddAt2O(child *DisplayObject, index int, silent bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("addAt", child, index, silent)}
}

// AddAtI Adds an existing object to this group.
//
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) AddAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("addAt", args)}
}

// AddToHash Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *World) AddToHash(child *DisplayObject) bool {
	return self.Object.Call("addToHash", child).Bool()
}

// AddToHashI Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *World) AddToHashI(args ...interface{}) bool {
	return self.Object.Call("addToHash", args).Bool()
}

// RemoveFromHash Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *World) RemoveFromHash(child *DisplayObject) bool {
	return self.Object.Call("removeFromHash", child).Bool()
}

// RemoveFromHashI Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *World) RemoveFromHashI(args ...interface{}) bool {
	return self.Object.Call("removeFromHash", args).Bool()
}

// AddMultiple Adds an array of existing Display Objects to this Group.
//
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
//
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) AddMultiple(children interface{}) interface{} {
	return self.Object.Call("addMultiple", children)
}

// AddMultiple1O Adds an array of existing Display Objects to this Group.
//
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
//
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) AddMultiple1O(children interface{}, silent bool) interface{} {
	return self.Object.Call("addMultiple", children, silent)
}

// AddMultipleI Adds an array of existing Display Objects to this Group.
//
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
//
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) AddMultipleI(args ...interface{}) interface{} {
	return self.Object.Call("addMultiple", args)
}

// GetAt Returns the child found at the given index within this group.
func (self *World) GetAt(index int) interface{} {
	return self.Object.Call("getAt", index)
}

// GetAtI Returns the child found at the given index within this group.
func (self *World) GetAtI(args ...interface{}) interface{} {
	return self.Object.Call("getAt", args)
}

// Create Creates a new Phaser.Sprite object and adds it to the top of this group.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) Create(x int, y int) *DisplayObject {
	return &DisplayObject{self.Object.Call("create", x, y)}
}

// Create1O Creates a new Phaser.Sprite object and adds it to the top of this group.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) Create1O(x int, y int, key interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("create", x, y, key)}
}

// Create2O Creates a new Phaser.Sprite object and adds it to the top of this group.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) Create2O(x int, y int, key interface{}, frame interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("create", x, y, key, frame)}
}

// Create3O Creates a new Phaser.Sprite object and adds it to the top of this group.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) Create3O(x int, y int, key interface{}, frame interface{}, exists bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("create", x, y, key, frame, exists)}
}

// Create4O Creates a new Phaser.Sprite object and adds it to the top of this group.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) Create4O(x int, y int, key interface{}, frame interface{}, exists bool, index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("create", x, y, key, frame, exists, index)}
}

// CreateI Creates a new Phaser.Sprite object and adds it to the top of this group.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// The child is automatically added to the top of the group, and is displayed above every previous child.
//
// Or if the _optional_ index is specified, the child is added at the location specified by the index value,
// this allows you to control child ordering.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) CreateI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("create", args)}
}

// CreateMultiple Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
//
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
//
// For example:
//
// `createMultiple(25, ['ball', 'carrot'])`
//
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
//
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
//
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
//
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
//
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
//
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
//
// By default the Sprites will have their `exists` property set to `false`, and they will be
// positioned at 0x0, relative to the `Group.x / y` values.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) CreateMultiple(quantity int, key interface{}) []interface{} {
	array00 := self.Object.Call("createMultiple", quantity, key)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CreateMultiple1O Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
//
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
//
// For example:
//
// `createMultiple(25, ['ball', 'carrot'])`
//
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
//
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
//
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
//
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
//
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
//
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
//
// By default the Sprites will have their `exists` property set to `false`, and they will be
// positioned at 0x0, relative to the `Group.x / y` values.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) CreateMultiple1O(quantity int, key interface{}, frame interface{}) []interface{} {
	array00 := self.Object.Call("createMultiple", quantity, key, frame)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CreateMultiple2O Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
//
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
//
// For example:
//
// `createMultiple(25, ['ball', 'carrot'])`
//
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
//
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
//
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
//
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
//
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
//
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
//
// By default the Sprites will have their `exists` property set to `false`, and they will be
// positioned at 0x0, relative to the `Group.x / y` values.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) CreateMultiple2O(quantity int, key interface{}, frame interface{}, exists bool) []interface{} {
	array00 := self.Object.Call("createMultiple", quantity, key, frame, exists)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CreateMultipleI Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
//
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
//
// Use {@link Phaser.Group#classType classType} to change the type of object created.
//
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
//
// For example:
//
// `createMultiple(25, ['ball', 'carrot'])`
//
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
//
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
//
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
//
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
//
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
//
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
//
// By default the Sprites will have their `exists` property set to `false`, and they will be
// positioned at 0x0, relative to the `Group.x / y` values.
//
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *World) CreateMultipleI(args ...interface{}) []interface{} {
	array00 := self.Object.Call("createMultiple", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// UpdateZ Internal method that re-applies all of the children's Z values.
//
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *World) UpdateZ() {
	self.Object.Call("updateZ")
}

// UpdateZI Internal method that re-applies all of the children's Z values.
//
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *World) UpdateZI(args ...interface{}) {
	self.Object.Call("updateZ", args)
}

// Align This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
//
// The grid dimensions are determined by the first four arguments. The `width` and `height` arguments
// relate to the width and height of the grid respectively.
//
// For example if the Group had 100 children in it:
//
// `Group.align(10, 10, 32, 32)`
//
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
//
// `Group.align(25, 4, 32, 32)`
//
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
//
// You can choose to set _either_ the `width` or `height` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
//
// `Group.align(-1, 8, 32, 32)`
//
// ... will align the children so that there are 8 children vertically (the second argument),
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
//
// You can also do:
//
// `Group.align(10, -1, 32, 32)`
//
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
//
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`,
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`,
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
//
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *World) Align(width int, height int, cellWidth int, cellHeight int) bool {
	return self.Object.Call("align", width, height, cellWidth, cellHeight).Bool()
}

// Align1O This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
//
// The grid dimensions are determined by the first four arguments. The `width` and `height` arguments
// relate to the width and height of the grid respectively.
//
// For example if the Group had 100 children in it:
//
// `Group.align(10, 10, 32, 32)`
//
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
//
// `Group.align(25, 4, 32, 32)`
//
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
//
// You can choose to set _either_ the `width` or `height` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
//
// `Group.align(-1, 8, 32, 32)`
//
// ... will align the children so that there are 8 children vertically (the second argument),
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
//
// You can also do:
//
// `Group.align(10, -1, 32, 32)`
//
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
//
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`,
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`,
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
//
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *World) Align1O(width int, height int, cellWidth int, cellHeight int, position int) bool {
	return self.Object.Call("align", width, height, cellWidth, cellHeight, position).Bool()
}

// Align2O This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
//
// The grid dimensions are determined by the first four arguments. The `width` and `height` arguments
// relate to the width and height of the grid respectively.
//
// For example if the Group had 100 children in it:
//
// `Group.align(10, 10, 32, 32)`
//
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
//
// `Group.align(25, 4, 32, 32)`
//
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
//
// You can choose to set _either_ the `width` or `height` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
//
// `Group.align(-1, 8, 32, 32)`
//
// ... will align the children so that there are 8 children vertically (the second argument),
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
//
// You can also do:
//
// `Group.align(10, -1, 32, 32)`
//
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
//
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`,
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`,
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
//
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *World) Align2O(width int, height int, cellWidth int, cellHeight int, position int, offset int) bool {
	return self.Object.Call("align", width, height, cellWidth, cellHeight, position, offset).Bool()
}

// AlignI This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
//
// The grid dimensions are determined by the first four arguments. The `width` and `height` arguments
// relate to the width and height of the grid respectively.
//
// For example if the Group had 100 children in it:
//
// `Group.align(10, 10, 32, 32)`
//
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
//
// `Group.align(25, 4, 32, 32)`
//
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
//
// You can choose to set _either_ the `width` or `height` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
//
// `Group.align(-1, 8, 32, 32)`
//
// ... will align the children so that there are 8 children vertically (the second argument),
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
//
// You can also do:
//
// `Group.align(10, -1, 32, 32)`
//
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
//
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`,
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`,
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
//
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *World) AlignI(args ...interface{}) bool {
	return self.Object.Call("align", args).Bool()
}

// ResetCursor Sets the group cursor to the first child in the group.
//
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *World) ResetCursor() interface{} {
	return self.Object.Call("resetCursor")
}

// ResetCursor1O Sets the group cursor to the first child in the group.
//
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *World) ResetCursor1O(index int) interface{} {
	return self.Object.Call("resetCursor", index)
}

// ResetCursorI Sets the group cursor to the first child in the group.
//
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *World) ResetCursorI(args ...interface{}) interface{} {
	return self.Object.Call("resetCursor", args)
}

// Next Advances the group cursor to the next (higher) object in the group.
//
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *World) Next() interface{} {
	return self.Object.Call("next")
}

// NextI Advances the group cursor to the next (higher) object in the group.
//
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *World) NextI(args ...interface{}) interface{} {
	return self.Object.Call("next", args)
}

// Previous Moves the group cursor to the previous (lower) child in the group.
//
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *World) Previous() interface{} {
	return self.Object.Call("previous")
}

// PreviousI Moves the group cursor to the previous (lower) child in the group.
//
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *World) PreviousI(args ...interface{}) interface{} {
	return self.Object.Call("previous", args)
}

// Swap Swaps the position of two children in this group.
//
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *World) Swap(child1 interface{}, child2 interface{}) {
	self.Object.Call("swap", child1, child2)
}

// SwapI Swaps the position of two children in this group.
//
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *World) SwapI(args ...interface{}) {
	self.Object.Call("swap", args)
}

// BringToTop Brings the given child to the top of this group so it renders above all other children.
func (self *World) BringToTop(child interface{}) interface{} {
	return self.Object.Call("bringToTop", child)
}

// BringToTopI Brings the given child to the top of this group so it renders above all other children.
func (self *World) BringToTopI(args ...interface{}) interface{} {
	return self.Object.Call("bringToTop", args)
}

// SendToBack Sends the given child to the bottom of this group so it renders below all other children.
func (self *World) SendToBack(child interface{}) interface{} {
	return self.Object.Call("sendToBack", child)
}

// SendToBackI Sends the given child to the bottom of this group so it renders below all other children.
func (self *World) SendToBackI(args ...interface{}) interface{} {
	return self.Object.Call("sendToBack", args)
}

// MoveUp Moves the given child up one place in this group unless it's already at the top.
func (self *World) MoveUp(child interface{}) interface{} {
	return self.Object.Call("moveUp", child)
}

// MoveUpI Moves the given child up one place in this group unless it's already at the top.
func (self *World) MoveUpI(args ...interface{}) interface{} {
	return self.Object.Call("moveUp", args)
}

// MoveDown Moves the given child down one place in this group unless it's already at the bottom.
func (self *World) MoveDown(child interface{}) interface{} {
	return self.Object.Call("moveDown", child)
}

// MoveDownI Moves the given child down one place in this group unless it's already at the bottom.
func (self *World) MoveDownI(args ...interface{}) interface{} {
	return self.Object.Call("moveDown", args)
}

// Xy Positions the child found at the given index within this group to the given x and y coordinates.
func (self *World) Xy(index int, x int, y int) {
	self.Object.Call("xy", index, x, y)
}

// XyI Positions the child found at the given index within this group to the given x and y coordinates.
func (self *World) XyI(args ...interface{}) {
	self.Object.Call("xy", args)
}

// Reverse Reverses all children in this group.
//
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *World) Reverse() {
	self.Object.Call("reverse")
}

// ReverseI Reverses all children in this group.
//
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *World) ReverseI(args ...interface{}) {
	self.Object.Call("reverse", args)
}

// GetIndex Get the index position of the given child in this group, which should match the child's `z` property.
func (self *World) GetIndex(child interface{}) int {
	return self.Object.Call("getIndex", child).Int()
}

// GetIndexI Get the index position of the given child in this group, which should match the child's `z` property.
func (self *World) GetIndexI(args ...interface{}) int {
	return self.Object.Call("getIndex", args).Int()
}

// GetByName Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *World) GetByName(name string) interface{} {
	return self.Object.Call("getByName", name)
}

// GetByNameI Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *World) GetByNameI(args ...interface{}) interface{} {
	return self.Object.Call("getByName", args)
}

// Replace Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) Replace(oldChild interface{}, newChild interface{}) interface{} {
	return self.Object.Call("replace", oldChild, newChild)
}

// ReplaceI Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
//
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
//
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) ReplaceI(args ...interface{}) interface{} {
	return self.Object.Call("replace", args)
}

// HasProperty Checks if the child has the given property.
//
// Will scan up to 4 levels deep only.
func (self *World) HasProperty(child interface{}, key []string) bool {
	return self.Object.Call("hasProperty", child, key).Bool()
}

// HasPropertyI Checks if the child has the given property.
//
// Will scan up to 4 levels deep only.
func (self *World) HasPropertyI(args ...interface{}) bool {
	return self.Object.Call("hasProperty", args).Bool()
}

// SetProperty Sets a property to the given value on the child. The operation parameter controls how the value is set.
//
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *World) SetProperty(child interface{}, key []interface{}, value interface{}) bool {
	return self.Object.Call("setProperty", child, key, value).Bool()
}

// SetProperty1O Sets a property to the given value on the child. The operation parameter controls how the value is set.
//
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *World) SetProperty1O(child interface{}, key []interface{}, value interface{}, operation int) bool {
	return self.Object.Call("setProperty", child, key, value, operation).Bool()
}

// SetProperty2O Sets a property to the given value on the child. The operation parameter controls how the value is set.
//
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *World) SetProperty2O(child interface{}, key []interface{}, value interface{}, operation int, force bool) bool {
	return self.Object.Call("setProperty", child, key, value, operation, force).Bool()
}

// SetPropertyI Sets a property to the given value on the child. The operation parameter controls how the value is set.
//
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *World) SetPropertyI(args ...interface{}) bool {
	return self.Object.Call("setProperty", args).Bool()
}

// CheckProperty Checks a property for the given value on the child.
func (self *World) CheckProperty(child interface{}, key []interface{}, value interface{}) bool {
	return self.Object.Call("checkProperty", child, key, value).Bool()
}

// CheckProperty1O Checks a property for the given value on the child.
func (self *World) CheckProperty1O(child interface{}, key []interface{}, value interface{}, force bool) bool {
	return self.Object.Call("checkProperty", child, key, value, force).Bool()
}

// CheckPropertyI Checks a property for the given value on the child.
func (self *World) CheckPropertyI(args ...interface{}) bool {
	return self.Object.Call("checkProperty", args).Bool()
}

// Set Quickly set a property on a single child of this group to a new value.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) Set(child *Sprite, key string, value interface{}) bool {
	return self.Object.Call("set", child, key, value).Bool()
}

// Set1O Quickly set a property on a single child of this group to a new value.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) Set1O(child *Sprite, key string, value interface{}, checkAlive bool) bool {
	return self.Object.Call("set", child, key, value, checkAlive).Bool()
}

// Set2O Quickly set a property on a single child of this group to a new value.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) Set2O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool) bool {
	return self.Object.Call("set", child, key, value, checkAlive, checkVisible).Bool()
}

// Set3O Quickly set a property on a single child of this group to a new value.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) Set3O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool, operation int) bool {
	return self.Object.Call("set", child, key, value, checkAlive, checkVisible, operation).Bool()
}

// Set4O Quickly set a property on a single child of this group to a new value.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) Set4O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) bool {
	return self.Object.Call("set", child, key, value, checkAlive, checkVisible, operation, force).Bool()
}

// SetI Quickly set a property on a single child of this group to a new value.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetI(args ...interface{}) bool {
	return self.Object.Call("set", args).Bool()
}

// SetAll Quickly set the same property across all children of this group to a new value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAll(key string, value interface{}) {
	self.Object.Call("setAll", key, value)
}

// SetAll1O Quickly set the same property across all children of this group to a new value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAll1O(key string, value interface{}, checkAlive bool) {
	self.Object.Call("setAll", key, value, checkAlive)
}

// SetAll2O Quickly set the same property across all children of this group to a new value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAll2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
	self.Object.Call("setAll", key, value, checkAlive, checkVisible)
}

// SetAll3O Quickly set the same property across all children of this group to a new value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAll3O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int) {
	self.Object.Call("setAll", key, value, checkAlive, checkVisible, operation)
}

// SetAll4O Quickly set the same property across all children of this group to a new value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAll4O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) {
	self.Object.Call("setAll", key, value, checkAlive, checkVisible, operation, force)
}

// SetAllI Quickly set the same property across all children of this group to a new value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllI(args ...interface{}) {
	self.Object.Call("setAll", args)
}

// SetAllChildren Quickly set the same property across all children of this group, and any child Groups, to a new value.
//
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildren(key string, value interface{}) {
	self.Object.Call("setAllChildren", key, value)
}

// SetAllChildren1O Quickly set the same property across all children of this group, and any child Groups, to a new value.
//
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildren1O(key string, value interface{}, checkAlive bool) {
	self.Object.Call("setAllChildren", key, value, checkAlive)
}

// SetAllChildren2O Quickly set the same property across all children of this group, and any child Groups, to a new value.
//
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildren2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
	self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible)
}

// SetAllChildren3O Quickly set the same property across all children of this group, and any child Groups, to a new value.
//
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildren3O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int) {
	self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible, operation)
}

// SetAllChildren4O Quickly set the same property across all children of this group, and any child Groups, to a new value.
//
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildren4O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) {
	self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible, operation, force)
}

// SetAllChildrenI Quickly set the same property across all children of this group, and any child Groups, to a new value.
//
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
//
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildrenI(args ...interface{}) {
	self.Object.Call("setAllChildren", args)
}

// CheckAll Quickly check that the same property across all children of this group is equal to the given value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *World) CheckAll(key string, value interface{}) {
	self.Object.Call("checkAll", key, value)
}

// CheckAll1O Quickly check that the same property across all children of this group is equal to the given value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *World) CheckAll1O(key string, value interface{}, checkAlive bool) {
	self.Object.Call("checkAll", key, value, checkAlive)
}

// CheckAll2O Quickly check that the same property across all children of this group is equal to the given value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *World) CheckAll2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
	self.Object.Call("checkAll", key, value, checkAlive, checkVisible)
}

// CheckAll3O Quickly check that the same property across all children of this group is equal to the given value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *World) CheckAll3O(key string, value interface{}, checkAlive bool, checkVisible bool, force bool) {
	self.Object.Call("checkAll", key, value, checkAlive, checkVisible, force)
}

// CheckAllI Quickly check that the same property across all children of this group is equal to the given value.
//
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *World) CheckAllI(args ...interface{}) {
	self.Object.Call("checkAll", args)
}

// AddAll Adds the amount to the given property on all children in this group.
//
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *World) AddAll(property string, amount int, checkAlive bool, checkVisible bool) {
	self.Object.Call("addAll", property, amount, checkAlive, checkVisible)
}

// AddAllI Adds the amount to the given property on all children in this group.
//
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *World) AddAllI(args ...interface{}) {
	self.Object.Call("addAll", args)
}

// SubAll Subtracts the amount from the given property on all children in this group.
//
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *World) SubAll(property string, amount int, checkAlive bool, checkVisible bool) {
	self.Object.Call("subAll", property, amount, checkAlive, checkVisible)
}

// SubAllI Subtracts the amount from the given property on all children in this group.
//
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *World) SubAllI(args ...interface{}) {
	self.Object.Call("subAll", args)
}

// MultiplyAll Multiplies the given property by the amount on all children in this group.
//
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *World) MultiplyAll(property string, amount int, checkAlive bool, checkVisible bool) {
	self.Object.Call("multiplyAll", property, amount, checkAlive, checkVisible)
}

// MultiplyAllI Multiplies the given property by the amount on all children in this group.
//
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *World) MultiplyAllI(args ...interface{}) {
	self.Object.Call("multiplyAll", args)
}

// DivideAll Divides the given property by the amount on all children in this group.
//
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *World) DivideAll(property string, amount int, checkAlive bool, checkVisible bool) {
	self.Object.Call("divideAll", property, amount, checkAlive, checkVisible)
}

// DivideAllI Divides the given property by the amount on all children in this group.
//
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *World) DivideAllI(args ...interface{}) {
	self.Object.Call("divideAll", args)
}

// CallAllExists Calls a function, specified by name, on all children in the group who exist (or do not exist).
//
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *World) CallAllExists(callback string, existsValue bool, parameter interface{}) {
	self.Object.Call("callAllExists", callback, existsValue, parameter)
}

// CallAllExistsI Calls a function, specified by name, on all children in the group who exist (or do not exist).
//
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *World) CallAllExistsI(args ...interface{}) {
	self.Object.Call("callAllExists", args)
}

// CallbackFromArray Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *World) CallbackFromArray(child interface{}, callback []interface{}, length int) {
	self.Object.Call("callbackFromArray", child, callback, length)
}

// CallbackFromArrayI Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *World) CallbackFromArrayI(args ...interface{}) {
	self.Object.Call("callbackFromArray", args)
}

// CallAll Calls a function, specified by name, on all on children.
//
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *World) CallAll(method string, context string, args interface{}) {
	self.Object.Call("callAll", method, context, args)
}

// CallAllI Calls a function, specified by name, on all on children.
//
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *World) CallAllI(args ...interface{}) {
	self.Object.Call("callAll", args)
}

// PreUpdate The core preUpdate - as called by World.
func (self *World) PreUpdate() {
	self.Object.Call("preUpdate")
}

// PreUpdateI The core preUpdate - as called by World.
func (self *World) PreUpdateI(args ...interface{}) {
	self.Object.Call("preUpdate", args)
}

// Update The core update - as called by World.
func (self *World) Update() {
	self.Object.Call("update")
}

// UpdateI The core update - as called by World.
func (self *World) UpdateI(args ...interface{}) {
	self.Object.Call("update", args)
}

// PostUpdate The core postUpdate - as called by World.
func (self *World) PostUpdate() {
	self.Object.Call("postUpdate")
}

// PostUpdateI The core postUpdate - as called by World.
func (self *World) PostUpdateI(args ...interface{}) {
	self.Object.Call("postUpdate", args)
}

// Filter Find children matching a certain predicate.
//
// For example:
//
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
//
// Note: Currently this will skip any children which are Groups themselves.
func (self *World) Filter(predicate interface{}) *ArraySet {
	return &ArraySet{self.Object.Call("filter", predicate)}
}

// Filter1O Find children matching a certain predicate.
//
// For example:
//
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
//
// Note: Currently this will skip any children which are Groups themselves.
func (self *World) Filter1O(predicate interface{}, checkExists bool) *ArraySet {
	return &ArraySet{self.Object.Call("filter", predicate, checkExists)}
}

// FilterI Find children matching a certain predicate.
//
// For example:
//
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
//
// Note: Currently this will skip any children which are Groups themselves.
func (self *World) FilterI(args ...interface{}) *ArraySet {
	return &ArraySet{self.Object.Call("filter", args)}
}

// ForEach Call a function on each child in this group.
//
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
//
//     Group.forEach(awardBonusGold, this, true, 100, 500)
//
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
//
// Note: This check will skip any children which are Groups themselves.
func (self *World) ForEach(callback interface{}, callbackContext interface{}) {
	self.Object.Call("forEach", callback, callbackContext)
}

// ForEach1O Call a function on each child in this group.
//
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
//
//     Group.forEach(awardBonusGold, this, true, 100, 500)
//
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
//
// Note: This check will skip any children which are Groups themselves.
func (self *World) ForEach1O(callback interface{}, callbackContext interface{}, checkExists bool) {
	self.Object.Call("forEach", callback, callbackContext, checkExists)
}

// ForEach2O Call a function on each child in this group.
//
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
//
//     Group.forEach(awardBonusGold, this, true, 100, 500)
//
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
//
// Note: This check will skip any children which are Groups themselves.
func (self *World) ForEach2O(callback interface{}, callbackContext interface{}, checkExists bool, args interface{}) {
	self.Object.Call("forEach", callback, callbackContext, checkExists, args)
}

// ForEachI Call a function on each child in this group.
//
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
//
//     Group.forEach(awardBonusGold, this, true, 100, 500)
//
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
//
// Note: This check will skip any children which are Groups themselves.
func (self *World) ForEachI(args ...interface{}) {
	self.Object.Call("forEach", args)
}

// ForEachExists Call a function on each existing child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachExists(callback interface{}, callbackContext interface{}) {
	self.Object.Call("forEachExists", callback, callbackContext)
}

// ForEachExists1O Call a function on each existing child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachExists1O(callback interface{}, callbackContext interface{}, args interface{}) {
	self.Object.Call("forEachExists", callback, callbackContext, args)
}

// ForEachExistsI Call a function on each existing child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachExistsI(args ...interface{}) {
	self.Object.Call("forEachExists", args)
}

// ForEachAlive Call a function on each alive child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachAlive(callback interface{}, callbackContext interface{}) {
	self.Object.Call("forEachAlive", callback, callbackContext)
}

// ForEachAlive1O Call a function on each alive child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachAlive1O(callback interface{}, callbackContext interface{}, args interface{}) {
	self.Object.Call("forEachAlive", callback, callbackContext, args)
}

// ForEachAliveI Call a function on each alive child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachAliveI(args ...interface{}) {
	self.Object.Call("forEachAlive", args)
}

// ForEachDead Call a function on each dead child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachDead(callback interface{}, callbackContext interface{}) {
	self.Object.Call("forEachDead", callback, callbackContext)
}

// ForEachDead1O Call a function on each dead child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachDead1O(callback interface{}, callbackContext interface{}, args interface{}) {
	self.Object.Call("forEachDead", callback, callbackContext, args)
}

// ForEachDeadI Call a function on each dead child in this group.
//
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachDeadI(args ...interface{}) {
	self.Object.Call("forEachDead", args)
}

// Sort Sort the children in the group according to a particular key and ordering.
//
// Call this function to sort the group according to a particular key value and order.
//
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
//
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *World) Sort() {
	self.Object.Call("sort")
}

// Sort1O Sort the children in the group according to a particular key and ordering.
//
// Call this function to sort the group according to a particular key value and order.
//
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
//
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *World) Sort1O(key string) {
	self.Object.Call("sort", key)
}

// Sort2O Sort the children in the group according to a particular key and ordering.
//
// Call this function to sort the group according to a particular key value and order.
//
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
//
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *World) Sort2O(key string, order int) {
	self.Object.Call("sort", key, order)
}

// SortI Sort the children in the group according to a particular key and ordering.
//
// Call this function to sort the group according to a particular key value and order.
//
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
//
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *World) SortI(args ...interface{}) {
	self.Object.Call("sort", args)
}

// CustomSort Sort the children in the group according to custom sort function.
//
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *World) CustomSort(sortHandler interface{}) {
	self.Object.Call("customSort", sortHandler)
}

// CustomSort1O Sort the children in the group according to custom sort function.
//
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *World) CustomSort1O(sortHandler interface{}, context interface{}) {
	self.Object.Call("customSort", sortHandler, context)
}

// CustomSortI Sort the children in the group according to custom sort function.
//
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *World) CustomSortI(args ...interface{}) {
	self.Object.Call("customSort", args)
}

// AscendingSortHandler An internal helper function for the sort process.
func (self *World) AscendingSortHandler(a interface{}, b interface{}) {
	self.Object.Call("ascendingSortHandler", a, b)
}

// AscendingSortHandlerI An internal helper function for the sort process.
func (self *World) AscendingSortHandlerI(args ...interface{}) {
	self.Object.Call("ascendingSortHandler", args)
}

// DescendingSortHandler An internal helper function for the sort process.
func (self *World) DescendingSortHandler(a interface{}, b interface{}) {
	self.Object.Call("descendingSortHandler", a, b)
}

// DescendingSortHandlerI An internal helper function for the sort process.
func (self *World) DescendingSortHandlerI(args ...interface{}) {
	self.Object.Call("descendingSortHandler", args)
}

// Iterate Iterates over the children of the group performing one of several actions for matched children.
//
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
//
// The result depends on the `returnType`:
//
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
//
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *World) Iterate(key string, value interface{}, returnType int) interface{} {
	return self.Object.Call("iterate", key, value, returnType)
}

// Iterate1O Iterates over the children of the group performing one of several actions for matched children.
//
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
//
// The result depends on the `returnType`:
//
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
//
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *World) Iterate1O(key string, value interface{}, returnType int, callback interface{}) interface{} {
	return self.Object.Call("iterate", key, value, returnType, callback)
}

// Iterate2O Iterates over the children of the group performing one of several actions for matched children.
//
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
//
// The result depends on the `returnType`:
//
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
//
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *World) Iterate2O(key string, value interface{}, returnType int, callback interface{}, callbackContext interface{}) interface{} {
	return self.Object.Call("iterate", key, value, returnType, callback, callbackContext)
}

// Iterate3O Iterates over the children of the group performing one of several actions for matched children.
//
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
//
// The result depends on the `returnType`:
//
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
//
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *World) Iterate3O(key string, value interface{}, returnType int, callback interface{}, callbackContext interface{}, args []interface{}) interface{} {
	return self.Object.Call("iterate", key, value, returnType, callback, callbackContext, args)
}

// IterateI Iterates over the children of the group performing one of several actions for matched children.
//
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
//
// The result depends on the `returnType`:
//
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
//
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *World) IterateI(args ...interface{}) interface{} {
	return self.Object.Call("iterate", args)
}

// GetFirstExists Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists() *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists")}
}

// GetFirstExists1O Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists1O(exists bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", exists)}
}

// GetFirstExists2O Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists2O(exists bool, createIfNull bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull)}
}

// GetFirstExists3O Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists3O(exists bool, createIfNull bool, x int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x)}
}

// GetFirstExists4O Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists4O(exists bool, createIfNull bool, x int, y int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y)}
}

// GetFirstExists5O Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists5O(exists bool, createIfNull bool, x int, y int, key interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y, key)}
}

// GetFirstExists6O Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExists6O(exists bool, createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y, key, frame)}
}

// GetFirstExistsI Get the first display object that exists, or doesn't exist.
//
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExistsI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstExists", args)}
}

// GetFirstAlive Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAlive() *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive")}
}

// GetFirstAlive1O Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAlive1O(createIfNull bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull)}
}

// GetFirstAlive2O Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAlive2O(createIfNull bool, x int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x)}
}

// GetFirstAlive3O Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAlive3O(createIfNull bool, x int, y int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y)}
}

// GetFirstAlive4O Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAlive4O(createIfNull bool, x int, y int, key interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y, key)}
}

// GetFirstAlive5O Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAlive5O(createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y, key, frame)}
}

// GetFirstAliveI Get the first child that is alive (`child.alive === true`).
//
// This is handy for choosing a squad leader, etc.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstAliveI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstAlive", args)}
}

// GetFirstDead Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDead() *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead")}
}

// GetFirstDead1O Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDead1O(createIfNull bool) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead", createIfNull)}
}

// GetFirstDead2O Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDead2O(createIfNull bool, x int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x)}
}

// GetFirstDead3O Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDead3O(createIfNull bool, x int, y int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y)}
}

// GetFirstDead4O Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDead4O(createIfNull bool, x int, y int, key interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y, key)}
}

// GetFirstDead5O Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDead5O(createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y, key, frame)}
}

// GetFirstDeadI Get the first child that is dead (`child.alive === false`).
//
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
//
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
//
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
//
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstDeadI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getFirstDead", args)}
}

// ResetChild Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
//
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
//
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChild(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("resetChild", child)}
}

// ResetChild1O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
//
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
//
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChild1O(child *DisplayObject, x int) *DisplayObject {
	return &DisplayObject{self.Object.Call("resetChild", child, x)}
}

// ResetChild2O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
//
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
//
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChild2O(child *DisplayObject, x int, y int) *DisplayObject {
	return &DisplayObject{self.Object.Call("resetChild", child, x, y)}
}

// ResetChild3O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
//
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
//
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChild3O(child *DisplayObject, x int, y int, key interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("resetChild", child, x, y, key)}
}

// ResetChild4O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
//
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
//
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChild4O(child *DisplayObject, x int, y int, key interface{}, frame interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("resetChild", child, x, y, key, frame)}
}

// ResetChildI Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
//
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
//
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChildI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("resetChild", args)}
}

// GetTop Return the child at the top of this group.
//
// The top child is the child displayed (rendered) above every other child.
func (self *World) GetTop() interface{} {
	return self.Object.Call("getTop")
}

// GetTopI Return the child at the top of this group.
//
// The top child is the child displayed (rendered) above every other child.
func (self *World) GetTopI(args ...interface{}) interface{} {
	return self.Object.Call("getTop", args)
}

// GetBottom Returns the child at the bottom of this group.
//
// The bottom child the child being displayed (rendered) below every other child.
func (self *World) GetBottom() interface{} {
	return self.Object.Call("getBottom")
}

// GetBottomI Returns the child at the bottom of this group.
//
// The bottom child the child being displayed (rendered) below every other child.
func (self *World) GetBottomI(args ...interface{}) interface{} {
	return self.Object.Call("getBottom", args)
}

// GetClosestTo Get the closest child to given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetClosestTo(object interface{}) interface{} {
	return self.Object.Call("getClosestTo", object)
}

// GetClosestTo1O Get the closest child to given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetClosestTo1O(object interface{}, callback interface{}) interface{} {
	return self.Object.Call("getClosestTo", object, callback)
}

// GetClosestTo2O Get the closest child to given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetClosestTo2O(object interface{}, callback interface{}, callbackContext interface{}) interface{} {
	return self.Object.Call("getClosestTo", object, callback, callbackContext)
}

// GetClosestToI Get the closest child to given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetClosestToI(args ...interface{}) interface{} {
	return self.Object.Call("getClosestTo", args)
}

// GetFurthestFrom Get the child furthest away from the given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetFurthestFrom(object interface{}) interface{} {
	return self.Object.Call("getFurthestFrom", object)
}

// GetFurthestFrom1O Get the child furthest away from the given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetFurthestFrom1O(object interface{}, callback interface{}) interface{} {
	return self.Object.Call("getFurthestFrom", object, callback)
}

// GetFurthestFrom2O Get the child furthest away from the given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetFurthestFrom2O(object interface{}, callback interface{}, callbackContext interface{}) interface{} {
	return self.Object.Call("getFurthestFrom", object, callback, callbackContext)
}

// GetFurthestFromI Get the child furthest away from the given Object, with optional callback to filter children.
//
// This can be a Sprite, Group, Image or any object with public x and y properties.
//
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
//
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your
// filtering criteria, otherwise it should return `false`.
func (self *World) GetFurthestFromI(args ...interface{}) interface{} {
	return self.Object.Call("getFurthestFrom", args)
}

// CountLiving Get the number of living children in this group.
func (self *World) CountLiving() int {
	return self.Object.Call("countLiving").Int()
}

// CountLivingI Get the number of living children in this group.
func (self *World) CountLivingI(args ...interface{}) int {
	return self.Object.Call("countLiving", args).Int()
}

// CountDead Get the number of dead children in this group.
func (self *World) CountDead() int {
	return self.Object.Call("countDead").Int()
}

// CountDeadI Get the number of dead children in this group.
func (self *World) CountDeadI(args ...interface{}) int {
	return self.Object.Call("countDead", args).Int()
}

// GetRandom Returns a random child from the group.
func (self *World) GetRandom() interface{} {
	return self.Object.Call("getRandom")
}

// GetRandom1O Returns a random child from the group.
func (self *World) GetRandom1O(startIndex int) interface{} {
	return self.Object.Call("getRandom", startIndex)
}

// GetRandom2O Returns a random child from the group.
func (self *World) GetRandom2O(startIndex int, length int) interface{} {
	return self.Object.Call("getRandom", startIndex, length)
}

// GetRandomI Returns a random child from the group.
func (self *World) GetRandomI(args ...interface{}) interface{} {
	return self.Object.Call("getRandom", args)
}

// GetRandomExists Returns a random child from the Group that has `exists` set to `true`.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetRandomExists() interface{} {
	return self.Object.Call("getRandomExists")
}

// GetRandomExists1O Returns a random child from the Group that has `exists` set to `true`.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetRandomExists1O(startIndex int) interface{} {
	return self.Object.Call("getRandomExists", startIndex)
}

// GetRandomExists2O Returns a random child from the Group that has `exists` set to `true`.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetRandomExists2O(startIndex int, endIndex int) interface{} {
	return self.Object.Call("getRandomExists", startIndex, endIndex)
}

// GetRandomExistsI Returns a random child from the Group that has `exists` set to `true`.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetRandomExistsI(args ...interface{}) interface{} {
	return self.Object.Call("getRandomExists", args)
}

// GetAll Returns all children in this Group.
//
// You can optionally specify a matching criteria using the `property` and `value` arguments.
//
// For example: `getAll('exists', true)` would return only children that have their exists property set.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetAll() interface{} {
	return self.Object.Call("getAll")
}

// GetAll1O Returns all children in this Group.
//
// You can optionally specify a matching criteria using the `property` and `value` arguments.
//
// For example: `getAll('exists', true)` would return only children that have their exists property set.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetAll1O(property string) interface{} {
	return self.Object.Call("getAll", property)
}

// GetAll2O Returns all children in this Group.
//
// You can optionally specify a matching criteria using the `property` and `value` arguments.
//
// For example: `getAll('exists', true)` would return only children that have their exists property set.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetAll2O(property string, value interface{}) interface{} {
	return self.Object.Call("getAll", property, value)
}

// GetAll3O Returns all children in this Group.
//
// You can optionally specify a matching criteria using the `property` and `value` arguments.
//
// For example: `getAll('exists', true)` would return only children that have their exists property set.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetAll3O(property string, value interface{}, startIndex int) interface{} {
	return self.Object.Call("getAll", property, value, startIndex)
}

// GetAll4O Returns all children in this Group.
//
// You can optionally specify a matching criteria using the `property` and `value` arguments.
//
// For example: `getAll('exists', true)` would return only children that have their exists property set.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetAll4O(property string, value interface{}, startIndex int, endIndex int) interface{} {
	return self.Object.Call("getAll", property, value, startIndex, endIndex)
}

// GetAllI Returns all children in this Group.
//
// You can optionally specify a matching criteria using the `property` and `value` arguments.
//
// For example: `getAll('exists', true)` would return only children that have their exists property set.
//
// Optionally you can specify a start and end index. For example if this Group had 100 children,
// and you set `startIndex` to 0 and `endIndex` to 50, it would return a random child from only
// the first 50 children in the Group.
func (self *World) GetAllI(args ...interface{}) interface{} {
	return self.Object.Call("getAll", args)
}

// Remove Removes the given child from this group.
//
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
//
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *World) Remove(child interface{}) bool {
	return self.Object.Call("remove", child).Bool()
}

// Remove1O Removes the given child from this group.
//
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
//
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *World) Remove1O(child interface{}, destroy bool) bool {
	return self.Object.Call("remove", child, destroy).Bool()
}

// Remove2O Removes the given child from this group.
//
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
//
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *World) Remove2O(child interface{}, destroy bool, silent bool) bool {
	return self.Object.Call("remove", child, destroy, silent).Bool()
}

// RemoveI Removes the given child from this group.
//
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
//
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *World) RemoveI(args ...interface{}) bool {
	return self.Object.Call("remove", args).Bool()
}

// MoveAll Moves all children from this Group to the Group given.
func (self *World) MoveAll(group *Group) *Group {
	return &Group{self.Object.Call("moveAll", group)}
}

// MoveAll1O Moves all children from this Group to the Group given.
func (self *World) MoveAll1O(group *Group, silent bool) *Group {
	return &Group{self.Object.Call("moveAll", group, silent)}
}

// MoveAllI Moves all children from this Group to the Group given.
func (self *World) MoveAllI(args ...interface{}) *Group {
	return &Group{self.Object.Call("moveAll", args)}
}

// RemoveAll Removes all children from this Group, but does not remove the group from its parent.
//
// The children can be optionally destroyed as they are removed.
//
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *World) RemoveAll() {
	self.Object.Call("removeAll")
}

// RemoveAll1O Removes all children from this Group, but does not remove the group from its parent.
//
// The children can be optionally destroyed as they are removed.
//
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *World) RemoveAll1O(destroy bool) {
	self.Object.Call("removeAll", destroy)
}

// RemoveAll2O Removes all children from this Group, but does not remove the group from its parent.
//
// The children can be optionally destroyed as they are removed.
//
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *World) RemoveAll2O(destroy bool, silent bool) {
	self.Object.Call("removeAll", destroy, silent)
}

// RemoveAll3O Removes all children from this Group, but does not remove the group from its parent.
//
// The children can be optionally destroyed as they are removed.
//
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *World) RemoveAll3O(destroy bool, silent bool, destroyTexture bool) {
	self.Object.Call("removeAll", destroy, silent, destroyTexture)
}

// RemoveAllI Removes all children from this Group, but does not remove the group from its parent.
//
// The children can be optionally destroyed as they are removed.
//
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *World) RemoveAllI(args ...interface{}) {
	self.Object.Call("removeAll", args)
}

// RemoveBetween Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *World) RemoveBetween(startIndex int) {
	self.Object.Call("removeBetween", startIndex)
}

// RemoveBetween1O Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *World) RemoveBetween1O(startIndex int, endIndex int) {
	self.Object.Call("removeBetween", startIndex, endIndex)
}

// RemoveBetween2O Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *World) RemoveBetween2O(startIndex int, endIndex int, destroy bool) {
	self.Object.Call("removeBetween", startIndex, endIndex, destroy)
}

// RemoveBetween3O Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *World) RemoveBetween3O(startIndex int, endIndex int, destroy bool, silent bool) {
	self.Object.Call("removeBetween", startIndex, endIndex, destroy, silent)
}

// RemoveBetweenI Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *World) RemoveBetweenI(args ...interface{}) {
	self.Object.Call("removeBetween", args)
}

// Destroy Destroys this group.
//
// Removes all children, then removes this group from its parent and nulls references.
func (self *World) Destroy() {
	self.Object.Call("destroy")
}

// Destroy1O Destroys this group.
//
// Removes all children, then removes this group from its parent and nulls references.
func (self *World) Destroy1O(destroyChildren bool) {
	self.Object.Call("destroy", destroyChildren)
}

// Destroy2O Destroys this group.
//
// Removes all children, then removes this group from its parent and nulls references.
func (self *World) Destroy2O(destroyChildren bool, soft bool) {
	self.Object.Call("destroy", destroyChildren, soft)
}

// DestroyI Destroys this group.
//
// Removes all children, then removes this group from its parent and nulls references.
func (self *World) DestroyI(args ...interface{}) {
	self.Object.Call("destroy", args)
}

// AddChild Adds a child to the container.
func (self *World) AddChild(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *World) AddChildI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *World) AddChildAt(child *DisplayObject, index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *World) AddChildAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *World) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
	self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *World) SwapChildrenI(args ...interface{}) {
	self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *World) GetChildIndex(child *DisplayObject) int {
	return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *World) GetChildIndexI(args ...interface{}) int {
	return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *World) SetChildIndex(child *DisplayObject, index int) {
	self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *World) SetChildIndexI(args ...interface{}) {
	self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *World) GetChildAt(index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *World) GetChildAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *World) RemoveChild(child *DisplayObject) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *World) RemoveChildI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *World) RemoveChildAt(index int) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *World) RemoveChildAtI(args ...interface{}) *DisplayObject {
	return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *World) RemoveChildren(beginIndex int, endIndex int) {
	self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *World) RemoveChildrenI(args ...interface{}) {
	self.Object.Call("removeChildren", args)
}

// GetBounds Retrieves the global bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *World) GetBounds() *Rectangle {
	return &Rectangle{self.Object.Call("getBounds")}
}

// GetBounds1O Retrieves the global bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *World) GetBounds1O(targetCoordinateSpace interface{}) *Rectangle {
	return &Rectangle{self.Object.Call("getBounds", targetCoordinateSpace)}
}

// GetBoundsI Retrieves the global bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *World) GetBoundsI(args ...interface{}) *Rectangle {
	return &Rectangle{self.Object.Call("getBounds", args)}
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle without any transformations. The calculation takes all visible children into consideration.
func (self *World) GetLocalBounds() *Rectangle {
	return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle without any transformations. The calculation takes all visible children into consideration.
func (self *World) GetLocalBoundsI(args ...interface{}) *Rectangle {
	return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Contains Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *World) Contains(child *DisplayObject) bool {
	return self.Object.Call("contains", child).Bool()
}

// ContainsI Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *World) ContainsI(args ...interface{}) bool {
	return self.Object.Call("contains", args).Bool()
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *World) _renderWebGL(renderSession *RenderSession) {
	self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *World) _renderWebGLI(args ...interface{}) {
	self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *World) _renderCanvas(renderSession *RenderSession) {
	self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *World) _renderCanvasI(args ...interface{}) {
	self.Object.Call("_renderCanvas", args)
}
