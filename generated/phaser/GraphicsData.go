// Package phaser Automatic generation for PIXI.GraphicsData
// generated file GraphicsData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// GraphicsData A GraphicsData object.
type GraphicsData struct {
    *js.Object
}

// NewGraphicsData A GraphicsData object.
func NewGraphicsData() *GraphicsData {
    return &GraphicsData{js.Global.Get("PIXI").Get("GraphicsData").New()}
}
// NewGraphicsDataI A GraphicsData object.
func NewGraphicsDataI(args ...interface{}) *GraphicsData {
    return &GraphicsData{js.Global.Get("PIXI").Get("GraphicsData").New(args)}
}




