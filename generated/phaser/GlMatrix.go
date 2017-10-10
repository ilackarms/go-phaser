// Package phaser Automatic generation for glMatrix
// generated file GlMatrix.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// GlMatrix empty description
type GlMatrix struct {
	*js.Object
}

// NewGlMatrix empty description
func NewGlMatrix() *GlMatrix {
	return &GlMatrix{js.Global.Get("glMatrix").New()}
}

// NewGlMatrixI empty description
func NewGlMatrixI(args ...interface{}) *GlMatrix {
	return &GlMatrix{js.Global.Get("glMatrix").New(args...)}
}

// GlMatrix Binding conversion method to GlMatrix point
func ToGlMatrix(jsStruct interface{}) *GlMatrix {
	if object, ok := jsStruct.(*js.Object); ok {
		return &GlMatrix{Object: object}
	}
	return nil
}
