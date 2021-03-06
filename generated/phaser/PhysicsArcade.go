// Package phaser Automatic generation for Phaser.Physics.Arcade
// generated file PhysicsArcade.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// PhysicsArcade The Arcade Physics world. Contains Arcade Physics related collision, overlap and motion methods.
type PhysicsArcade struct {
	*js.Object
}

// NewPhysicsArcade The Arcade Physics world. Contains Arcade Physics related collision, overlap and motion methods.
func NewPhysicsArcade(game *Game) *PhysicsArcade {
	return &PhysicsArcade{js.Global.Get("Phaser").Get("Physics").Get("Arcade").New(game)}
}

// NewPhysicsArcadeI The Arcade Physics world. Contains Arcade Physics related collision, overlap and motion methods.
func NewPhysicsArcadeI(args ...interface{}) *PhysicsArcade {
	return &PhysicsArcade{js.Global.Get("Phaser").Get("Physics").Get("Arcade").New(args...)}
}

// PhysicsArcade Binding conversion method to PhysicsArcade point
func ToPhysicsArcade(jsStruct interface{}) *PhysicsArcade {
	if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsArcade{Object: object}
	}
	return nil
}

// Game Local reference to game.
func (self *PhysicsArcade) Game() *Game {
	return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsArcade) SetGameA(member *Game) {
	self.Object.Set("game", member)
}

// Gravity The World gravity setting. Defaults to x: 0, y: 0, or no gravity.
func (self *PhysicsArcade) Gravity() *Point {
	return &Point{self.Object.Get("gravity")}
}

// SetGravityA The World gravity setting. Defaults to x: 0, y: 0, or no gravity.
func (self *PhysicsArcade) SetGravityA(member *Point) {
	self.Object.Set("gravity", member)
}

// Bounds The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsArcade) Bounds() *Rectangle {
	return &Rectangle{self.Object.Get("bounds")}
}

// SetBoundsA The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsArcade) SetBoundsA(member *Rectangle) {
	self.Object.Set("bounds", member)
}

// CheckCollision Set the checkCollision properties to control for which bounds collision is processed.
// For example checkCollision.down = false means Bodies cannot collide with the World.bounds.bottom. An object containing allowed collision flags.
func (self *PhysicsArcade) CheckCollision() interface{} {
	return self.Object.Get("checkCollision")
}

// SetCheckCollisionA Set the checkCollision properties to control for which bounds collision is processed.
// For example checkCollision.down = false means Bodies cannot collide with the World.bounds.bottom. An object containing allowed collision flags.
func (self *PhysicsArcade) SetCheckCollisionA(member interface{}) {
	self.Object.Set("checkCollision", member)
}

// MaxObjects Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsArcade) MaxObjects() int {
	return self.Object.Get("maxObjects").Int()
}

// SetMaxObjectsA Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsArcade) SetMaxObjectsA(member int) {
	self.Object.Set("maxObjects", member)
}

// MaxLevels Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsArcade) MaxLevels() int {
	return self.Object.Get("maxLevels").Int()
}

// SetMaxLevelsA Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsArcade) SetMaxLevelsA(member int) {
	self.Object.Set("maxLevels", member)
}

// OVERLAP_BIAS A value added to the delta values during collision checks.
func (self *PhysicsArcade) OVERLAP_BIAS() int {
	return self.Object.Get("OVERLAP_BIAS").Int()
}

// SetOVERLAP_BIASA A value added to the delta values during collision checks.
func (self *PhysicsArcade) SetOVERLAP_BIASA(member int) {
	self.Object.Set("OVERLAP_BIAS", member)
}

// ForceX If true World.separate will always separate on the X axis before Y. Otherwise it will check gravity totals first.
func (self *PhysicsArcade) ForceX() bool {
	return self.Object.Get("forceX").Bool()
}

// SetForceXA If true World.separate will always separate on the X axis before Y. Otherwise it will check gravity totals first.
func (self *PhysicsArcade) SetForceXA(member bool) {
	self.Object.Set("forceX", member)
}

// SortDirection Used when colliding a Sprite vs. a Group, or a Group vs. a Group, this defines the direction the sort is based on. Default is Phaser.Physics.Arcade.LEFT_RIGHT.
func (self *PhysicsArcade) SortDirection() int {
	return self.Object.Get("sortDirection").Int()
}

// SetSortDirectionA Used when colliding a Sprite vs. a Group, or a Group vs. a Group, this defines the direction the sort is based on. Default is Phaser.Physics.Arcade.LEFT_RIGHT.
func (self *PhysicsArcade) SetSortDirectionA(member int) {
	self.Object.Set("sortDirection", member)
}

// SkipQuadTree If true the QuadTree will not be used for any collision. QuadTrees are great if objects are well spread out in your game, otherwise they are a performance hit. If you enable this you can disable on a per body basis via `Body.skipQuadTree`.
func (self *PhysicsArcade) SkipQuadTree() bool {
	return self.Object.Get("skipQuadTree").Bool()
}

// SetSkipQuadTreeA If true the QuadTree will not be used for any collision. QuadTrees are great if objects are well spread out in your game, otherwise they are a performance hit. If you enable this you can disable on a per body basis via `Body.skipQuadTree`.
func (self *PhysicsArcade) SetSkipQuadTreeA(member bool) {
	self.Object.Set("skipQuadTree", member)
}

// IsPaused If `true` the `Body.preUpdate` method will be skipped, halting all motion for all bodies. Note that other methods such as `collide` will still work, so be careful not to call them on paused bodies.
func (self *PhysicsArcade) IsPaused() bool {
	return self.Object.Get("isPaused").Bool()
}

// SetIsPausedA If `true` the `Body.preUpdate` method will be skipped, halting all motion for all bodies. Note that other methods such as `collide` will still work, so be careful not to call them on paused bodies.
func (self *PhysicsArcade) SetIsPausedA(member bool) {
	self.Object.Set("isPaused", member)
}

// QuadTree The world QuadTree.
func (self *PhysicsArcade) QuadTree() *QuadTree {
	return &QuadTree{self.Object.Get("quadTree")}
}

// SetQuadTreeA The world QuadTree.
func (self *PhysicsArcade) SetQuadTreeA(member *QuadTree) {
	self.Object.Set("quadTree", member)
}

// SORT_NONE A constant used for the sortDirection value.
// Use this if you don't wish to perform any pre-collision sorting at all, or will manually sort your Groups.
func (self *PhysicsArcade) SORT_NONE() int {
	return self.Object.Get("SORT_NONE").Int()
}

// SetSORT_NONEA A constant used for the sortDirection value.
// Use this if you don't wish to perform any pre-collision sorting at all, or will manually sort your Groups.
func (self *PhysicsArcade) SetSORT_NONEA(member int) {
	self.Object.Set("SORT_NONE", member)
}

// LEFT_RIGHT A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the left to the right (i.e. Mario)
func (self *PhysicsArcade) LEFT_RIGHT() int {
	return self.Object.Get("LEFT_RIGHT").Int()
}

// SetLEFT_RIGHTA A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the left to the right (i.e. Mario)
func (self *PhysicsArcade) SetLEFT_RIGHTA(member int) {
	self.Object.Set("LEFT_RIGHT", member)
}

// RIGHT_LEFT A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the right to the left (i.e. Mario backwards)
func (self *PhysicsArcade) RIGHT_LEFT() int {
	return self.Object.Get("RIGHT_LEFT").Int()
}

// SetRIGHT_LEFTA A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the right to the left (i.e. Mario backwards)
func (self *PhysicsArcade) SetRIGHT_LEFTA(member int) {
	self.Object.Set("RIGHT_LEFT", member)
}

// TOP_BOTTOM A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the top to the bottom (i.e. Dig Dug)
func (self *PhysicsArcade) TOP_BOTTOM() int {
	return self.Object.Get("TOP_BOTTOM").Int()
}

// SetTOP_BOTTOMA A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the top to the bottom (i.e. Dig Dug)
func (self *PhysicsArcade) SetTOP_BOTTOMA(member int) {
	self.Object.Set("TOP_BOTTOM", member)
}

// BOTTOM_TOP A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the bottom to the top (i.e. Commando or a vertically scrolling shoot-em-up)
func (self *PhysicsArcade) BOTTOM_TOP() int {
	return self.Object.Get("BOTTOM_TOP").Int()
}

// SetBOTTOM_TOPA A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the bottom to the top (i.e. Commando or a vertically scrolling shoot-em-up)
func (self *PhysicsArcade) SetBOTTOM_TOPA(member int) {
	self.Object.Set("BOTTOM_TOP", member)
}

// SetBounds Updates the size of this physics world.
func (self *PhysicsArcade) SetBounds(x int, y int, width int, height int) {
	self.Object.Call("setBounds", x, y, width, height)
}

// SetBoundsI Updates the size of this physics world.
func (self *PhysicsArcade) SetBoundsI(args ...interface{}) {
	self.Object.Call("setBounds", args...)
}

// SetBoundsToWorld Updates the size of this physics world to match the size of the game world.
func (self *PhysicsArcade) SetBoundsToWorld() {
	self.Object.Call("setBoundsToWorld")
}

// SetBoundsToWorldI Updates the size of this physics world to match the size of the game world.
func (self *PhysicsArcade) SetBoundsToWorldI(args ...interface{}) {
	self.Object.Call("setBoundsToWorld", args...)
}

// Enable This will create an Arcade Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsArcade) Enable(object interface{}) {
	self.Object.Call("enable", object)
}

// Enable1O This will create an Arcade Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsArcade) Enable1O(object interface{}, children bool) {
	self.Object.Call("enable", object, children)
}

// EnableI This will create an Arcade Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsArcade) EnableI(args ...interface{}) {
	self.Object.Call("enable", args...)
}

// EnableBody Creates an Arcade Physics body on the given game object.
//
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
//
// When you add an Arcade Physics body to an object it will automatically add the object into its parent Groups hash array.
func (self *PhysicsArcade) EnableBody(object interface{}) {
	self.Object.Call("enableBody", object)
}

// EnableBodyI Creates an Arcade Physics body on the given game object.
//
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
//
// When you add an Arcade Physics body to an object it will automatically add the object into its parent Groups hash array.
func (self *PhysicsArcade) EnableBodyI(args ...interface{}) {
	self.Object.Call("enableBody", args...)
}

// UpdateMotion Called automatically by a Physics body, it updates all motion related values on the Body unless `World.isPaused` is `true`.
func (self *PhysicsArcade) UpdateMotion(The *PhysicsArcadeBody) {
	self.Object.Call("updateMotion", The)
}

// UpdateMotionI Called automatically by a Physics body, it updates all motion related values on the Body unless `World.isPaused` is `true`.
func (self *PhysicsArcade) UpdateMotionI(args ...interface{}) {
	self.Object.Call("updateMotion", args...)
}

// ComputeVelocity A tween-like function that takes a starting velocity and some other factors and returns an altered velocity.
// Based on a function in Flixel by @ADAMATOMIC
func (self *PhysicsArcade) ComputeVelocity(axis int, body *PhysicsArcadeBody, velocity int, acceleration int, drag int) int {
	return self.Object.Call("computeVelocity", axis, body, velocity, acceleration, drag).Int()
}

// ComputeVelocity1O A tween-like function that takes a starting velocity and some other factors and returns an altered velocity.
// Based on a function in Flixel by @ADAMATOMIC
func (self *PhysicsArcade) ComputeVelocity1O(axis int, body *PhysicsArcadeBody, velocity int, acceleration int, drag int, max int) int {
	return self.Object.Call("computeVelocity", axis, body, velocity, acceleration, drag, max).Int()
}

// ComputeVelocityI A tween-like function that takes a starting velocity and some other factors and returns an altered velocity.
// Based on a function in Flixel by @ADAMATOMIC
func (self *PhysicsArcade) ComputeVelocityI(args ...interface{}) int {
	return self.Object.Call("computeVelocity", args...).Int()
}

// Overlap Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups within Groups).
func (self *PhysicsArcade) Overlap(object1 interface{}, object2 interface{}) bool {
	return self.Object.Call("overlap", object1, object2).Bool()
}

// Overlap1O Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups within Groups).
func (self *PhysicsArcade) Overlap1O(object1 interface{}, object2 interface{}, overlapCallback interface{}) bool {
	return self.Object.Call("overlap", object1, object2, overlapCallback).Bool()
}

// Overlap2O Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups within Groups).
func (self *PhysicsArcade) Overlap2O(object1 interface{}, object2 interface{}, overlapCallback interface{}, processCallback interface{}) bool {
	return self.Object.Call("overlap", object1, object2, overlapCallback, processCallback).Bool()
}

// Overlap3O Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups within Groups).
func (self *PhysicsArcade) Overlap3O(object1 interface{}, object2 interface{}, overlapCallback interface{}, processCallback interface{}, callbackContext interface{}) bool {
	return self.Object.Call("overlap", object1, object2, overlapCallback, processCallback, callbackContext).Bool()
}

// OverlapI Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups within Groups).
func (self *PhysicsArcade) OverlapI(args ...interface{}) bool {
	return self.Object.Call("overlap", args...).Bool()
}

// Collide Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups or Tilemaps within other Groups).
func (self *PhysicsArcade) Collide(object1 interface{}, object2 interface{}) bool {
	return self.Object.Call("collide", object1, object2).Bool()
}

// Collide1O Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups or Tilemaps within other Groups).
func (self *PhysicsArcade) Collide1O(object1 interface{}, object2 interface{}, collideCallback interface{}) bool {
	return self.Object.Call("collide", object1, object2, collideCallback).Bool()
}

// Collide2O Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups or Tilemaps within other Groups).
func (self *PhysicsArcade) Collide2O(object1 interface{}, object2 interface{}, collideCallback interface{}, processCallback interface{}) bool {
	return self.Object.Call("collide", object1, object2, collideCallback, processCallback).Bool()
}

// Collide3O Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups or Tilemaps within other Groups).
func (self *PhysicsArcade) Collide3O(object1 interface{}, object2 interface{}, collideCallback interface{}, processCallback interface{}, callbackContext interface{}) bool {
	return self.Object.Call("collide", object1, object2, collideCallback, processCallback, callbackContext).Bool()
}

// CollideI Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups or Tilemaps within other Groups).
func (self *PhysicsArcade) CollideI(args ...interface{}) bool {
	return self.Object.Call("collide", args...).Bool()
}

// SortLeftRight A Sort function for sorting two bodies based on a LEFT to RIGHT sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortLeftRight(a *Sprite, b *Sprite) int {
	return self.Object.Call("sortLeftRight", a, b).Int()
}

// SortLeftRightI A Sort function for sorting two bodies based on a LEFT to RIGHT sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortLeftRightI(args ...interface{}) int {
	return self.Object.Call("sortLeftRight", args...).Int()
}

// SortRightLeft A Sort function for sorting two bodies based on a RIGHT to LEFT sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortRightLeft(a *Sprite, b *Sprite) int {
	return self.Object.Call("sortRightLeft", a, b).Int()
}

// SortRightLeftI A Sort function for sorting two bodies based on a RIGHT to LEFT sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortRightLeftI(args ...interface{}) int {
	return self.Object.Call("sortRightLeft", args...).Int()
}

// SortTopBottom A Sort function for sorting two bodies based on a TOP to BOTTOM sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortTopBottom(a *Sprite, b *Sprite) int {
	return self.Object.Call("sortTopBottom", a, b).Int()
}

// SortTopBottomI A Sort function for sorting two bodies based on a TOP to BOTTOM sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortTopBottomI(args ...interface{}) int {
	return self.Object.Call("sortTopBottom", args...).Int()
}

// SortBottomTop A Sort function for sorting two bodies based on a BOTTOM to TOP sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortBottomTop(a *Sprite, b *Sprite) int {
	return self.Object.Call("sortBottomTop", a, b).Int()
}

// SortBottomTopI A Sort function for sorting two bodies based on a BOTTOM to TOP sort direction.
//
// This is called automatically by World.sort
func (self *PhysicsArcade) SortBottomTopI(args ...interface{}) int {
	return self.Object.Call("sortBottomTop", args...).Int()
}

// Sort This method will sort a Groups hash array.
//
// If the Group has `physicsSortDirection` set it will use the sort direction defined.
//
// Otherwise if the sortDirection parameter is undefined, or Group.physicsSortDirection is null, it will use Phaser.Physics.Arcade.sortDirection.
//
// By changing Group.physicsSortDirection you can customise each Group to sort in a different order.
func (self *PhysicsArcade) Sort(group *Group) {
	self.Object.Call("sort", group)
}

// Sort1O This method will sort a Groups hash array.
//
// If the Group has `physicsSortDirection` set it will use the sort direction defined.
//
// Otherwise if the sortDirection parameter is undefined, or Group.physicsSortDirection is null, it will use Phaser.Physics.Arcade.sortDirection.
//
// By changing Group.physicsSortDirection you can customise each Group to sort in a different order.
func (self *PhysicsArcade) Sort1O(group *Group, sortDirection int) {
	self.Object.Call("sort", group, sortDirection)
}

// SortI This method will sort a Groups hash array.
//
// If the Group has `physicsSortDirection` set it will use the sort direction defined.
//
// Otherwise if the sortDirection parameter is undefined, or Group.physicsSortDirection is null, it will use Phaser.Physics.Arcade.sortDirection.
//
// By changing Group.physicsSortDirection you can customise each Group to sort in a different order.
func (self *PhysicsArcade) SortI(args ...interface{}) {
	self.Object.Call("sort", args...)
}

// CollideHandler Internal collision handler.
func (self *PhysicsArcade) CollideHandler(object1 interface{}, object2 interface{}, collideCallback interface{}, processCallback interface{}, callbackContext interface{}, overlapOnly bool) {
	self.Object.Call("collideHandler", object1, object2, collideCallback, processCallback, callbackContext, overlapOnly)
}

// CollideHandlerI Internal collision handler.
func (self *PhysicsArcade) CollideHandlerI(args ...interface{}) {
	self.Object.Call("collideHandler", args...)
}

// CollideSpriteVsSprite An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideSpriteVsSprite(sprite1 *Sprite, sprite2 *Sprite, collideCallback interface{}, processCallback interface{}, callbackContext interface{}, overlapOnly bool) bool {
	return self.Object.Call("collideSpriteVsSprite", sprite1, sprite2, collideCallback, processCallback, callbackContext, overlapOnly).Bool()
}

// CollideSpriteVsSpriteI An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideSpriteVsSpriteI(args ...interface{}) bool {
	return self.Object.Call("collideSpriteVsSprite", args...).Bool()
}

// CollideSpriteVsGroup An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideSpriteVsGroup(sprite *Sprite, group *Group, collideCallback interface{}, processCallback interface{}, callbackContext interface{}, overlapOnly bool) {
	self.Object.Call("collideSpriteVsGroup", sprite, group, collideCallback, processCallback, callbackContext, overlapOnly)
}

// CollideSpriteVsGroupI An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideSpriteVsGroupI(args ...interface{}) {
	self.Object.Call("collideSpriteVsGroup", args...)
}

// CollideGroupVsSelf An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideGroupVsSelf(group *Group, collideCallback interface{}, processCallback interface{}, callbackContext interface{}, overlapOnly bool) bool {
	return self.Object.Call("collideGroupVsSelf", group, collideCallback, processCallback, callbackContext, overlapOnly).Bool()
}

// CollideGroupVsSelfI An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideGroupVsSelfI(args ...interface{}) bool {
	return self.Object.Call("collideGroupVsSelf", args...).Bool()
}

// CollideGroupVsGroup An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideGroupVsGroup(group1 *Group, group2 *Group, collideCallback interface{}, processCallback interface{}, callbackContext interface{}, overlapOnly bool) {
	self.Object.Call("collideGroupVsGroup", group1, group2, collideCallback, processCallback, callbackContext, overlapOnly)
}

// CollideGroupVsGroupI An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideGroupVsGroupI(args ...interface{}) {
	self.Object.Call("collideGroupVsGroup", args...)
}

// Separate The core separation function to separate two physics bodies.
func (self *PhysicsArcade) Separate(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody, processCallback interface{}, callbackContext interface{}, overlapOnly bool) bool {
	return self.Object.Call("separate", body1, body2, processCallback, callbackContext, overlapOnly).Bool()
}

// SeparateI The core separation function to separate two physics bodies.
func (self *PhysicsArcade) SeparateI(args ...interface{}) bool {
	return self.Object.Call("separate", args...).Bool()
}

// Intersects Check for intersection against two bodies.
func (self *PhysicsArcade) Intersects(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody) bool {
	return self.Object.Call("intersects", body1, body2).Bool()
}

// IntersectsI Check for intersection against two bodies.
func (self *PhysicsArcade) IntersectsI(args ...interface{}) bool {
	return self.Object.Call("intersects", args...).Bool()
}

// CircleBodyIntersects Checks to see if a circular Body intersects with a Rectangular Body.
func (self *PhysicsArcade) CircleBodyIntersects(circle *PhysicsArcadeBody, body *PhysicsArcadeBody) bool {
	return self.Object.Call("circleBodyIntersects", circle, body).Bool()
}

// CircleBodyIntersectsI Checks to see if a circular Body intersects with a Rectangular Body.
func (self *PhysicsArcade) CircleBodyIntersectsI(args ...interface{}) bool {
	return self.Object.Call("circleBodyIntersects", args...).Bool()
}

// SeparateCircle The core separation function to separate two circular physics bodies.
func (self *PhysicsArcade) SeparateCircle(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody, overlapOnly bool) bool {
	return self.Object.Call("separateCircle", body1, body2, overlapOnly).Bool()
}

// SeparateCircleI The core separation function to separate two circular physics bodies.
func (self *PhysicsArcade) SeparateCircleI(args ...interface{}) bool {
	return self.Object.Call("separateCircle", args...).Bool()
}

// GetOverlapX Calculates the horizontal overlap between two Bodies and sets their properties accordingly, including:
// `touching.left`, `touching.right` and `overlapX`.
func (self *PhysicsArcade) GetOverlapX(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody, overlapOnly bool) float64 {
	return self.Object.Call("getOverlapX", body1, body2, overlapOnly).Float()
}

// GetOverlapXI Calculates the horizontal overlap between two Bodies and sets their properties accordingly, including:
// `touching.left`, `touching.right` and `overlapX`.
func (self *PhysicsArcade) GetOverlapXI(args ...interface{}) float64 {
	return self.Object.Call("getOverlapX", args...).Float()
}

// GetOverlapY Calculates the vertical overlap between two Bodies and sets their properties accordingly, including:
// `touching.up`, `touching.down` and `overlapY`.
func (self *PhysicsArcade) GetOverlapY(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody, overlapOnly bool) float64 {
	return self.Object.Call("getOverlapY", body1, body2, overlapOnly).Float()
}

// GetOverlapYI Calculates the vertical overlap between two Bodies and sets their properties accordingly, including:
// `touching.up`, `touching.down` and `overlapY`.
func (self *PhysicsArcade) GetOverlapYI(args ...interface{}) float64 {
	return self.Object.Call("getOverlapY", args...).Float()
}

// SeparateX The core separation function to separate two physics bodies on the x axis.
func (self *PhysicsArcade) SeparateX(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody, overlapOnly bool) bool {
	return self.Object.Call("separateX", body1, body2, overlapOnly).Bool()
}

// SeparateXI The core separation function to separate two physics bodies on the x axis.
func (self *PhysicsArcade) SeparateXI(args ...interface{}) bool {
	return self.Object.Call("separateX", args...).Bool()
}

// SeparateY The core separation function to separate two physics bodies on the y axis.
func (self *PhysicsArcade) SeparateY(body1 *PhysicsArcadeBody, body2 *PhysicsArcadeBody, overlapOnly bool) bool {
	return self.Object.Call("separateY", body1, body2, overlapOnly).Bool()
}

// SeparateYI The core separation function to separate two physics bodies on the y axis.
func (self *PhysicsArcade) SeparateYI(args ...interface{}) bool {
	return self.Object.Call("separateY", args...).Bool()
}

// GetObjectsUnderPointer Given a Group and a Pointer this will check to see which Group children overlap with the Pointer coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the Pointer or not.
func (self *PhysicsArcade) GetObjectsUnderPointer(pointer *Pointer, group *Group) []DisplayObject {
	array00 := self.Object.Call("getObjectsUnderPointer", pointer, group)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsUnderPointer1O Given a Group and a Pointer this will check to see which Group children overlap with the Pointer coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the Pointer or not.
func (self *PhysicsArcade) GetObjectsUnderPointer1O(pointer *Pointer, group *Group, callback interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsUnderPointer", pointer, group, callback)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsUnderPointer2O Given a Group and a Pointer this will check to see which Group children overlap with the Pointer coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the Pointer or not.
func (self *PhysicsArcade) GetObjectsUnderPointer2O(pointer *Pointer, group *Group, callback interface{}, callbackContext interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsUnderPointer", pointer, group, callback, callbackContext)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsUnderPointerI Given a Group and a Pointer this will check to see which Group children overlap with the Pointer coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the Pointer or not.
func (self *PhysicsArcade) GetObjectsUnderPointerI(args ...interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsUnderPointer", args...)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsAtLocation Given a Group and a location this will check to see which Group children overlap with the coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the coordinate or not.
func (self *PhysicsArcade) GetObjectsAtLocation(x int, y int, group *Group) []DisplayObject {
	array00 := self.Object.Call("getObjectsAtLocation", x, y, group)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsAtLocation1O Given a Group and a location this will check to see which Group children overlap with the coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the coordinate or not.
func (self *PhysicsArcade) GetObjectsAtLocation1O(x int, y int, group *Group, callback interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsAtLocation", x, y, group, callback)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsAtLocation2O Given a Group and a location this will check to see which Group children overlap with the coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the coordinate or not.
func (self *PhysicsArcade) GetObjectsAtLocation2O(x int, y int, group *Group, callback interface{}, callbackContext interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsAtLocation", x, y, group, callback, callbackContext)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsAtLocation3O Given a Group and a location this will check to see which Group children overlap with the coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the coordinate or not.
func (self *PhysicsArcade) GetObjectsAtLocation3O(x int, y int, group *Group, callback interface{}, callbackContext interface{}, callbackArg interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsAtLocation", x, y, group, callback, callbackContext, callbackArg)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// GetObjectsAtLocationI Given a Group and a location this will check to see which Group children overlap with the coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the coordinate or not.
func (self *PhysicsArcade) GetObjectsAtLocationI(args ...interface{}) []DisplayObject {
	array00 := self.Object.Call("getObjectsAtLocation", args...)
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {

		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// MoveToObject Move the given display object towards the destination object at a steady velocity.
// If you specify a maxTime then it will adjust the speed (overwriting what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToObject(displayObject interface{}, destination interface{}) int {
	return self.Object.Call("moveToObject", displayObject, destination).Int()
}

// MoveToObject1O Move the given display object towards the destination object at a steady velocity.
// If you specify a maxTime then it will adjust the speed (overwriting what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToObject1O(displayObject interface{}, destination interface{}, speed int) int {
	return self.Object.Call("moveToObject", displayObject, destination, speed).Int()
}

// MoveToObject2O Move the given display object towards the destination object at a steady velocity.
// If you specify a maxTime then it will adjust the speed (overwriting what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToObject2O(displayObject interface{}, destination interface{}, speed int, maxTime int) int {
	return self.Object.Call("moveToObject", displayObject, destination, speed, maxTime).Int()
}

// MoveToObjectI Move the given display object towards the destination object at a steady velocity.
// If you specify a maxTime then it will adjust the speed (overwriting what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToObjectI(args ...interface{}) int {
	return self.Object.Call("moveToObject", args...).Int()
}

// MoveToPointer Move the given display object towards the pointer at a steady velocity. If no pointer is given it will use Phaser.Input.activePointer.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) MoveToPointer(displayObject interface{}) int {
	return self.Object.Call("moveToPointer", displayObject).Int()
}

// MoveToPointer1O Move the given display object towards the pointer at a steady velocity. If no pointer is given it will use Phaser.Input.activePointer.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) MoveToPointer1O(displayObject interface{}, speed int) int {
	return self.Object.Call("moveToPointer", displayObject, speed).Int()
}

// MoveToPointer2O Move the given display object towards the pointer at a steady velocity. If no pointer is given it will use Phaser.Input.activePointer.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) MoveToPointer2O(displayObject interface{}, speed int, pointer *Pointer) int {
	return self.Object.Call("moveToPointer", displayObject, speed, pointer).Int()
}

// MoveToPointer3O Move the given display object towards the pointer at a steady velocity. If no pointer is given it will use Phaser.Input.activePointer.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) MoveToPointer3O(displayObject interface{}, speed int, pointer *Pointer, maxTime int) int {
	return self.Object.Call("moveToPointer", displayObject, speed, pointer, maxTime).Int()
}

// MoveToPointerI Move the given display object towards the pointer at a steady velocity. If no pointer is given it will use Phaser.Input.activePointer.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) MoveToPointerI(args ...interface{}) int {
	return self.Object.Call("moveToPointer", args...).Int()
}

// MoveToXY Move the given display object towards the x/y coordinates at a steady velocity.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToXY(displayObject interface{}, x int, y int) int {
	return self.Object.Call("moveToXY", displayObject, x, y).Int()
}

// MoveToXY1O Move the given display object towards the x/y coordinates at a steady velocity.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToXY1O(displayObject interface{}, x int, y int, speed int) int {
	return self.Object.Call("moveToXY", displayObject, x, y, speed).Int()
}

// MoveToXY2O Move the given display object towards the x/y coordinates at a steady velocity.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToXY2O(displayObject interface{}, x int, y int, speed int, maxTime int) int {
	return self.Object.Call("moveToXY", displayObject, x, y, speed, maxTime).Int()
}

// MoveToXYI Move the given display object towards the x/y coordinates at a steady velocity.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToXYI(args ...interface{}) int {
	return self.Object.Call("moveToXY", args...).Int()
}

// VelocityFromAngle Given the angle (in degrees) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromAngle(angle, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromAngle(angle int) *Point {
	return &Point{self.Object.Call("velocityFromAngle", angle)}
}

// VelocityFromAngle1O Given the angle (in degrees) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromAngle(angle, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromAngle1O(angle int, speed int) *Point {
	return &Point{self.Object.Call("velocityFromAngle", angle, speed)}
}

// VelocityFromAngle2O Given the angle (in degrees) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromAngle(angle, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromAngle2O(angle int, speed int, point interface{}) *Point {
	return &Point{self.Object.Call("velocityFromAngle", angle, speed, point)}
}

// VelocityFromAngleI Given the angle (in degrees) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromAngle(angle, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromAngleI(args ...interface{}) *Point {
	return &Point{self.Object.Call("velocityFromAngle", args...)}
}

// VelocityFromRotation Given the rotation (in radians) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromRotation(rotation, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromRotation(rotation int) *Point {
	return &Point{self.Object.Call("velocityFromRotation", rotation)}
}

// VelocityFromRotation1O Given the rotation (in radians) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromRotation(rotation, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromRotation1O(rotation int, speed int) *Point {
	return &Point{self.Object.Call("velocityFromRotation", rotation, speed)}
}

// VelocityFromRotation2O Given the rotation (in radians) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromRotation(rotation, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromRotation2O(rotation int, speed int, point interface{}) *Point {
	return &Point{self.Object.Call("velocityFromRotation", rotation, speed, point)}
}

// VelocityFromRotationI Given the rotation (in radians) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromRotation(rotation, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromRotationI(args ...interface{}) *Point {
	return &Point{self.Object.Call("velocityFromRotation", args...)}
}

// AccelerationFromRotation Given the rotation (in radians) and speed calculate the acceleration and return it as a Point object, or set it to the given point object.
// One way to use this is: accelerationFromRotation(rotation, 200, sprite.acceleration) which will set the values directly to the sprites acceleration and not create a new Point object.
func (self *PhysicsArcade) AccelerationFromRotation(rotation int) *Point {
	return &Point{self.Object.Call("accelerationFromRotation", rotation)}
}

// AccelerationFromRotation1O Given the rotation (in radians) and speed calculate the acceleration and return it as a Point object, or set it to the given point object.
// One way to use this is: accelerationFromRotation(rotation, 200, sprite.acceleration) which will set the values directly to the sprites acceleration and not create a new Point object.
func (self *PhysicsArcade) AccelerationFromRotation1O(rotation int, speed int) *Point {
	return &Point{self.Object.Call("accelerationFromRotation", rotation, speed)}
}

// AccelerationFromRotation2O Given the rotation (in radians) and speed calculate the acceleration and return it as a Point object, or set it to the given point object.
// One way to use this is: accelerationFromRotation(rotation, 200, sprite.acceleration) which will set the values directly to the sprites acceleration and not create a new Point object.
func (self *PhysicsArcade) AccelerationFromRotation2O(rotation int, speed int, point interface{}) *Point {
	return &Point{self.Object.Call("accelerationFromRotation", rotation, speed, point)}
}

// AccelerationFromRotationI Given the rotation (in radians) and speed calculate the acceleration and return it as a Point object, or set it to the given point object.
// One way to use this is: accelerationFromRotation(rotation, 200, sprite.acceleration) which will set the values directly to the sprites acceleration and not create a new Point object.
func (self *PhysicsArcade) AccelerationFromRotationI(args ...interface{}) *Point {
	return &Point{self.Object.Call("accelerationFromRotation", args...)}
}

// AccelerateToObject Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToObject(displayObject interface{}, destination interface{}) int {
	return self.Object.Call("accelerateToObject", displayObject, destination).Int()
}

// AccelerateToObject1O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToObject1O(displayObject interface{}, destination interface{}, speed int) int {
	return self.Object.Call("accelerateToObject", displayObject, destination, speed).Int()
}

// AccelerateToObject2O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToObject2O(displayObject interface{}, destination interface{}, speed int, xSpeedMax int) int {
	return self.Object.Call("accelerateToObject", displayObject, destination, speed, xSpeedMax).Int()
}

// AccelerateToObject3O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToObject3O(displayObject interface{}, destination interface{}, speed int, xSpeedMax int, ySpeedMax int) int {
	return self.Object.Call("accelerateToObject", displayObject, destination, speed, xSpeedMax, ySpeedMax).Int()
}

// AccelerateToObjectI Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToObjectI(args ...interface{}) int {
	return self.Object.Call("accelerateToObject", args...).Int()
}

// AccelerateToPointer Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointer(displayObject interface{}) int {
	return self.Object.Call("accelerateToPointer", displayObject).Int()
}

// AccelerateToPointer1O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointer1O(displayObject interface{}, pointer *Pointer) int {
	return self.Object.Call("accelerateToPointer", displayObject, pointer).Int()
}

// AccelerateToPointer2O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointer2O(displayObject interface{}, pointer *Pointer, speed int) int {
	return self.Object.Call("accelerateToPointer", displayObject, pointer, speed).Int()
}

// AccelerateToPointer3O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointer3O(displayObject interface{}, pointer *Pointer, speed int, xSpeedMax int) int {
	return self.Object.Call("accelerateToPointer", displayObject, pointer, speed, xSpeedMax).Int()
}

// AccelerateToPointer4O Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointer4O(displayObject interface{}, pointer *Pointer, speed int, xSpeedMax int, ySpeedMax int) int {
	return self.Object.Call("accelerateToPointer", displayObject, pointer, speed, xSpeedMax, ySpeedMax).Int()
}

// AccelerateToPointerI Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointerI(args ...interface{}) int {
	return self.Object.Call("accelerateToPointer", args...).Int()
}

// AccelerateToXY Sets the acceleration.x/y property on the display object so it will move towards the x/y coordinates at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToXY(displayObject interface{}, x int, y int) int {
	return self.Object.Call("accelerateToXY", displayObject, x, y).Int()
}

// AccelerateToXY1O Sets the acceleration.x/y property on the display object so it will move towards the x/y coordinates at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToXY1O(displayObject interface{}, x int, y int, speed int) int {
	return self.Object.Call("accelerateToXY", displayObject, x, y, speed).Int()
}

// AccelerateToXY2O Sets the acceleration.x/y property on the display object so it will move towards the x/y coordinates at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToXY2O(displayObject interface{}, x int, y int, speed int, xSpeedMax int) int {
	return self.Object.Call("accelerateToXY", displayObject, x, y, speed, xSpeedMax).Int()
}

// AccelerateToXY3O Sets the acceleration.x/y property on the display object so it will move towards the x/y coordinates at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToXY3O(displayObject interface{}, x int, y int, speed int, xSpeedMax int, ySpeedMax int) int {
	return self.Object.Call("accelerateToXY", displayObject, x, y, speed, xSpeedMax, ySpeedMax).Int()
}

// AccelerateToXYI Sets the acceleration.x/y property on the display object so it will move towards the x/y coordinates at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToXYI(args ...interface{}) int {
	return self.Object.Call("accelerateToXY", args...).Int()
}

// DistanceBetween Find the distance between two display objects (like Sprites).
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceBetween(source interface{}, target interface{}) int {
	return self.Object.Call("distanceBetween", source, target).Int()
}

// DistanceBetween1O Find the distance between two display objects (like Sprites).
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceBetween1O(source interface{}, target interface{}, world bool) int {
	return self.Object.Call("distanceBetween", source, target, world).Int()
}

// DistanceBetweenI Find the distance between two display objects (like Sprites).
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceBetweenI(args ...interface{}) int {
	return self.Object.Call("distanceBetween", args...).Int()
}

// DistanceToXY Find the distance between a display object (like a Sprite) and the given x/y coordinates.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToXY(displayObject interface{}, x int, y int) int {
	return self.Object.Call("distanceToXY", displayObject, x, y).Int()
}

// DistanceToXY1O Find the distance between a display object (like a Sprite) and the given x/y coordinates.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToXY1O(displayObject interface{}, x int, y int, world bool) int {
	return self.Object.Call("distanceToXY", displayObject, x, y, world).Int()
}

// DistanceToXYI Find the distance between a display object (like a Sprite) and the given x/y coordinates.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToXYI(args ...interface{}) int {
	return self.Object.Call("distanceToXY", args...).Int()
}

// DistanceToPointer Find the distance between a display object (like a Sprite) and a Pointer. If no Pointer is given the Input.activePointer is used.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToPointer(displayObject interface{}) int {
	return self.Object.Call("distanceToPointer", displayObject).Int()
}

// DistanceToPointer1O Find the distance between a display object (like a Sprite) and a Pointer. If no Pointer is given the Input.activePointer is used.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToPointer1O(displayObject interface{}, pointer *Pointer) int {
	return self.Object.Call("distanceToPointer", displayObject, pointer).Int()
}

// DistanceToPointer2O Find the distance between a display object (like a Sprite) and a Pointer. If no Pointer is given the Input.activePointer is used.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToPointer2O(displayObject interface{}, pointer *Pointer, world bool) int {
	return self.Object.Call("distanceToPointer", displayObject, pointer, world).Int()
}

// DistanceToPointerI Find the distance between a display object (like a Sprite) and a Pointer. If no Pointer is given the Input.activePointer is used.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToPointerI(args ...interface{}) int {
	return self.Object.Call("distanceToPointer", args...).Int()
}

// AngleBetween Find the angle in radians between two display objects (like Sprites).
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleBetween(source interface{}, target interface{}) int {
	return self.Object.Call("angleBetween", source, target).Int()
}

// AngleBetween1O Find the angle in radians between two display objects (like Sprites).
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleBetween1O(source interface{}, target interface{}, world bool) int {
	return self.Object.Call("angleBetween", source, target, world).Int()
}

// AngleBetweenI Find the angle in radians between two display objects (like Sprites).
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleBetweenI(args ...interface{}) int {
	return self.Object.Call("angleBetween", args...).Int()
}

// AngleBetweenCenters Find the angle in radians between centers of two display objects (like Sprites).
func (self *PhysicsArcade) AngleBetweenCenters(source interface{}, target interface{}) int {
	return self.Object.Call("angleBetweenCenters", source, target).Int()
}

// AngleBetweenCentersI Find the angle in radians between centers of two display objects (like Sprites).
func (self *PhysicsArcade) AngleBetweenCentersI(args ...interface{}) int {
	return self.Object.Call("angleBetweenCenters", args...).Int()
}

// AngleToXY Find the angle in radians between a display object (like a Sprite) and the given x/y coordinate.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToXY(displayObject interface{}, x int, y int) int {
	return self.Object.Call("angleToXY", displayObject, x, y).Int()
}

// AngleToXY1O Find the angle in radians between a display object (like a Sprite) and the given x/y coordinate.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToXY1O(displayObject interface{}, x int, y int, world bool) int {
	return self.Object.Call("angleToXY", displayObject, x, y, world).Int()
}

// AngleToXYI Find the angle in radians between a display object (like a Sprite) and the given x/y coordinate.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToXYI(args ...interface{}) int {
	return self.Object.Call("angleToXY", args...).Int()
}

// AngleToPointer Find the angle in radians between a display object (like a Sprite) and a Pointer, taking their x/y and center into account.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToPointer(displayObject interface{}) int {
	return self.Object.Call("angleToPointer", displayObject).Int()
}

// AngleToPointer1O Find the angle in radians between a display object (like a Sprite) and a Pointer, taking their x/y and center into account.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToPointer1O(displayObject interface{}, pointer *Pointer) int {
	return self.Object.Call("angleToPointer", displayObject, pointer).Int()
}

// AngleToPointer2O Find the angle in radians between a display object (like a Sprite) and a Pointer, taking their x/y and center into account.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToPointer2O(displayObject interface{}, pointer *Pointer, world bool) int {
	return self.Object.Call("angleToPointer", displayObject, pointer, world).Int()
}

// AngleToPointerI Find the angle in radians between a display object (like a Sprite) and a Pointer, taking their x/y and center into account.
//
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToPointerI(args ...interface{}) int {
	return self.Object.Call("angleToPointer", args...).Int()
}

// WorldAngleToPointer Find the angle in radians between a display object (like a Sprite) and a Pointer,
// taking their x/y and center into account relative to the world.
func (self *PhysicsArcade) WorldAngleToPointer(displayObject interface{}) int {
	return self.Object.Call("worldAngleToPointer", displayObject).Int()
}

// WorldAngleToPointer1O Find the angle in radians between a display object (like a Sprite) and a Pointer,
// taking their x/y and center into account relative to the world.
func (self *PhysicsArcade) WorldAngleToPointer1O(displayObject interface{}, pointer *Pointer) int {
	return self.Object.Call("worldAngleToPointer", displayObject, pointer).Int()
}

// WorldAngleToPointerI Find the angle in radians between a display object (like a Sprite) and a Pointer,
// taking their x/y and center into account relative to the world.
func (self *PhysicsArcade) WorldAngleToPointerI(args ...interface{}) int {
	return self.Object.Call("worldAngleToPointer", args...).Int()
}
