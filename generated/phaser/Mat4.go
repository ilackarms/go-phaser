// Automatic generation for mat4
// generated file Mat4.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Mat4 struct {
    *js.Object
}


// 
func NewMat4() *Mat4 {
    return &Mat4{js.Global.Call("mat4")}
}

// 
func NewMat4I(args ...interface{}) *Mat4 {
    return &Mat4{js.Global.Call("mat4", args)}
}




