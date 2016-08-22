// Automatic generation for Phaser.ImageCollection
// generated file ImageCollection.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
type ImageCollection struct {
    *js.Object
}


// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection(name string, firstgid int) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", name, firstgid)}
}

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection1O(name string, firstgid int, width int) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", name, firstgid, width)}
}

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection2O(name string, firstgid int, width int, height int) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", name, firstgid, width, height)}
}

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection3O(name string, firstgid int, width int, height int, margin int) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", name, firstgid, width, height, margin)}
}

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection4O(name string, firstgid int, width int, height int, margin int, spacing int) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", name, firstgid, width, height, margin, spacing)}
}

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection5O(name string, firstgid int, width int, height int, margin int, spacing int, properties interface{}) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", name, firstgid, width, height, margin, spacing, properties)}
}

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollectionI(args ...interface{}) *ImageCollection {
    return &ImageCollection{js.Global.Call("Phaser.ImageCollection", args)}
}



// The name of the Image Collection.
func (self *ImageCollection) GetNameA() string{
    return self.Object.Get("name").String()
}

// The name of the Image Collection.
func (self *ImageCollection) SetNameA(member string) {
    self.Object.Set("name", member)
}

// The Tiled firstgid value.
// This is the starting index of the first image index this Image Collection contains.
func (self *ImageCollection) GetFirstgidA() int{
    return self.Object.Get("firstgid").Int()
}

// The Tiled firstgid value.
// This is the starting index of the first image index this Image Collection contains.
func (self *ImageCollection) SetFirstgidA(member int) {
    self.Object.Set("firstgid", member)
}

// The width of the widest image (in pixels).
func (self *ImageCollection) GetImageWidthA() int{
    return self.Object.Get("imageWidth").Int()
}

// The width of the widest image (in pixels).
func (self *ImageCollection) SetImageWidthA(member int) {
    self.Object.Set("imageWidth", member)
}

// The height of the tallest image (in pixels).
func (self *ImageCollection) GetImageHeightA() int{
    return self.Object.Get("imageHeight").Int()
}

// The height of the tallest image (in pixels).
func (self *ImageCollection) SetImageHeightA(member int) {
    self.Object.Set("imageHeight", member)
}

// The margin around the images in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) GetImageMarginA() interface{}{
    return self.Object.Get("imageMargin")
}

// The margin around the images in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) SetImageMarginA(member interface{}) {
    self.Object.Set("imageMargin", member)
}

// The spacing between each image in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) GetImageSpacingA() int{
    return self.Object.Get("imageSpacing").Int()
}

// The spacing between each image in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) SetImageSpacingA(member int) {
    self.Object.Set("imageSpacing", member)
}

// Image Collection-specific properties that are typically defined in the Tiled editor.
func (self *ImageCollection) GetPropertiesA() interface{}{
    return self.Object.Get("properties")
}

// Image Collection-specific properties that are typically defined in the Tiled editor.
func (self *ImageCollection) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// The cached images that are a part of this collection.
func (self *ImageCollection) GetImagesA() []interface{}{
	array00 := self.Object.Get("images")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The cached images that are a part of this collection.
func (self *ImageCollection) SetImagesA(member []interface{}) {
    self.Object.Set("images", member)
}

// The total number of images in the image collection.
func (self *ImageCollection) GetTotalA() int{
    return self.Object.Get("total").Int()
}

// The total number of images in the image collection.
func (self *ImageCollection) SetTotalA(member int) {
    self.Object.Set("total", member)
}



// Returns true if and only if this image collection contains the given image index.
func (self *ImageCollection) ContainsImageIndex(imageIndex int) bool{
    return self.Object.Call("containsImageIndex", imageIndex).Bool()
}

// Returns true if and only if this image collection contains the given image index.
func (self *ImageCollection) ContainsImageIndexI(args ...interface{}) bool{
    return self.Object.Call("containsImageIndex", args).Bool()
}

// Add an image to this Image Collection.
func (self *ImageCollection) AddImage(gid int, image string) {
    self.Object.Call("addImage", gid, image)
}

// Add an image to this Image Collection.
func (self *ImageCollection) AddImageI(args ...interface{}) {
    self.Object.Call("addImage", args)
}
