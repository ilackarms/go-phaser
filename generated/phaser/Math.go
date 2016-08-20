// Automatic generation for Phaser.Math
// generated file Math.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A collection of useful mathematical functions.
// 
// These are normally accessed through `game.math`.
type Math struct {
    *js.Object
}


// Twice PI.
func (self *Math) GetPI2() interface{}{
    return self.Get("PI2")
}

// Twice PI.
func (self *Math) SetPI2(member interface{}) {
    self.Set("PI2", member)
}



// Returns a number between the `min` and `max` values.
func (self *Math) BetweenI(args ...interface{}) float64{
    return self.Call("between", args).Float()
}

// Two number are fuzzyEqual if their difference is less than epsilon.
func (self *Math) FuzzyEqualI(args ...interface{}) bool{
    return self.Call("fuzzyEqual", args).Bool()
}

// `a` is fuzzyLessThan `b` if it is less than b + epsilon.
func (self *Math) FuzzyLessThanI(args ...interface{}) bool{
    return self.Call("fuzzyLessThan", args).Bool()
}

// `a` is fuzzyGreaterThan `b` if it is more than b - epsilon.
func (self *Math) FuzzyGreaterThanI(args ...interface{}) bool{
    return self.Call("fuzzyGreaterThan", args).Bool()
}

// Applies a fuzzy ceil to the given value.
func (self *Math) FuzzyCeilI(args ...interface{}) float64{
    return self.Call("fuzzyCeil", args).Float()
}

// Applies a fuzzy floor to the given value.
func (self *Math) FuzzyFloorI(args ...interface{}) float64{
    return self.Call("fuzzyFloor", args).Float()
}

// Averages all values passed to the function and returns the result.
func (self *Math) AverageI(args ...interface{}) float64{
    return self.Call("average", args).Float()
}

// 
func (self *Math) ShearI(args ...interface{}) float64{
    return self.Call("shear", args).Float()
}

// Snap a value to nearest grid slice, using rounding.
// 
// Example: if you have an interval gap of 5 and a position of 12... you will snap to 10 whereas 14 will snap to 15.
func (self *Math) SnapToI(args ...interface{}) float64{
    return self.Call("snapTo", args).Float()
}

// Snap a value to nearest grid slice, using floor.
// 
// Example: if you have an interval gap of 5 and a position of 12... you will snap to 10.
// As will 14 snap to 10... but 16 will snap to 15.
func (self *Math) SnapToFloorI(args ...interface{}) float64{
    return self.Call("snapToFloor", args).Float()
}

// Snap a value to nearest grid slice, using ceil.
// 
// Example: if you have an interval gap of 5 and a position of 12... you will snap to 15.
// As will 14 will snap to 15... but 16 will snap to 20.
func (self *Math) SnapToCeilI(args ...interface{}) float64{
    return self.Call("snapToCeil", args).Float()
}

// Round to some place comparative to a `base`, default is 10 for decimal place.
// The `place` is represented by the power applied to `base` to get that place.
// 
//     e.g. 2000/7 ~= 285.714285714285714285714 ~= (bin)100011101.1011011011011011
// 
//     roundTo(2000/7,3) === 0
//     roundTo(2000/7,2) == 300
//     roundTo(2000/7,1) == 290
//     roundTo(2000/7,0) == 286
//     roundTo(2000/7,-1) == 285.7
//     roundTo(2000/7,-2) == 285.71
//     roundTo(2000/7,-3) == 285.714
//     roundTo(2000/7,-4) == 285.7143
//     roundTo(2000/7,-5) == 285.71429
// 
//     roundTo(2000/7,3,2)  == 288       -- 100100000
//     roundTo(2000/7,2,2)  == 284       -- 100011100
//     roundTo(2000/7,1,2)  == 286       -- 100011110
//     roundTo(2000/7,0,2)  == 286       -- 100011110
//     roundTo(2000/7,-1,2) == 285.5     -- 100011101.1
//     roundTo(2000/7,-2,2) == 285.75    -- 100011101.11
//     roundTo(2000/7,-3,2) == 285.75    -- 100011101.11
//     roundTo(2000/7,-4,2) == 285.6875  -- 100011101.1011
//     roundTo(2000/7,-5,2) == 285.71875 -- 100011101.10111
// 
// Note what occurs when we round to the 3rd space (8ths place), 100100000, this is to be assumed
// because we are rounding 100011.1011011011011011 which rounds up.
func (self *Math) RoundToI(args ...interface{}) float64{
    return self.Call("roundTo", args).Float()
}

// Floors to some place comparative to a `base`, default is 10 for decimal place.
// The `place` is represented by the power applied to `base` to get that place.
func (self *Math) FloorToI(args ...interface{}) float64{
    return self.Call("floorTo", args).Float()
}

// Ceils to some place comparative to a `base`, default is 10 for decimal place.
// The `place` is represented by the power applied to `base` to get that place.
func (self *Math) CeilToI(args ...interface{}) float64{
    return self.Call("ceilTo", args).Float()
}

// Find the angle of a segment from (x1, y1) -> (x2, y2).
func (self *Math) AngleBetweenI(args ...interface{}) float64{
    return self.Call("angleBetween", args).Float()
}

// Find the angle of a segment from (x1, y1) -> (x2, y2).
// 
// The difference between this method and Math.angleBetween is that this assumes the y coordinate travels
// down the screen.
func (self *Math) AngleBetweenYI(args ...interface{}) float64{
    return self.Call("angleBetweenY", args).Float()
}

// Find the angle of a segment from (point1.x, point1.y) -> (point2.x, point2.y).
func (self *Math) AngleBetweenPointsI(args ...interface{}) float64{
    return self.Call("angleBetweenPoints", args).Float()
}

// Find the angle of a segment from (point1.x, point1.y) -> (point2.x, point2.y).
func (self *Math) AngleBetweenPointsYI(args ...interface{}) float64{
    return self.Call("angleBetweenPointsY", args).Float()
}

// Reverses an angle.
func (self *Math) ReverseAngleI(args ...interface{}) float64{
    return self.Call("reverseAngle", args).Float()
}

// Normalizes an angle to the [0,2pi) range.
func (self *Math) NormalizeAngleI(args ...interface{}) float64{
    return self.Call("normalizeAngle", args).Float()
}

// Adds the given amount to the value, but never lets the value go over the specified maximum.
func (self *Math) MaxAddI(args ...interface{}) float64{
    return self.Call("maxAdd", args).Float()
}

// Subtracts the given amount from the value, but never lets the value go below the specified minimum.
func (self *Math) MinSubI(args ...interface{}) float64{
    return self.Call("minSub", args).Float()
}

// Ensures that the value always stays between min and max, by wrapping the value around.
// 
// If `max` is not larger than `min` the result is 0.
func (self *Math) WrapI(args ...interface{}) float64{
    return self.Call("wrap", args).Float()
}

// Adds value to amount and ensures that the result always stays between 0 and max, by wrapping the value around.
// 
// Values _must_ be positive integers, and are passed through Math.abs. See {@link Phaser.Math#wrap} for an alternative.
func (self *Math) WrapValueI(args ...interface{}) float64{
    return self.Call("wrapValue", args).Float()
}

// Returns true if the number given is odd.
func (self *Math) IsOddI(args ...interface{}) bool{
    return self.Call("isOdd", args).Bool()
}

// Returns true if the number given is even.
func (self *Math) IsEvenI(args ...interface{}) bool{
    return self.Call("isEven", args).Bool()
}

// Variation of Math.min that can be passed either an array of numbers or the numbers as parameters.
// 
// Prefer the standard `Math.min` function when appropriate.
func (self *Math) MinI(args ...interface{}) float64{
    return self.Call("min", args).Float()
}

// Variation of Math.max that can be passed either an array of numbers or the numbers as parameters.
// 
// Prefer the standard `Math.max` function when appropriate.
func (self *Math) MaxI(args ...interface{}) float64{
    return self.Call("max", args).Float()
}

// Variation of Math.min that can be passed a property and either an array of objects or the objects as parameters.
// It will find the lowest matching property value from the given objects.
func (self *Math) MinPropertyI(args ...interface{}) float64{
    return self.Call("minProperty", args).Float()
}

// Variation of Math.max that can be passed a property and either an array of objects or the objects as parameters.
// It will find the largest matching property value from the given objects.
func (self *Math) MaxPropertyI(args ...interface{}) float64{
    return self.Call("maxProperty", args).Float()
}

// Keeps an angle value between -180 and +180; or -PI and PI if radians.
func (self *Math) WrapAngleI(args ...interface{}) float64{
    return self.Call("wrapAngle", args).Float()
}

// A Linear Interpolation Method, mostly used by Phaser.Tween.
func (self *Math) LinearInterpolationI(args ...interface{}) float64{
    return self.Call("linearInterpolation", args).Float()
}

// A Bezier Interpolation Method, mostly used by Phaser.Tween.
func (self *Math) BezierInterpolationI(args ...interface{}) float64{
    return self.Call("bezierInterpolation", args).Float()
}

// A Catmull Rom Interpolation Method, mostly used by Phaser.Tween.
func (self *Math) CatmullRomInterpolationI(args ...interface{}) float64{
    return self.Call("catmullRomInterpolation", args).Float()
}

// Calculates a linear (interpolation) value over t.
func (self *Math) LinearI(args ...interface{}) float64{
    return self.Call("linear", args).Float()
}

// 
func (self *Math) BernsteinI(args ...interface{}) float64{
    return self.Call("bernstein", args).Float()
}

// 
func (self *Math) FactorialI(args ...interface{}) float64{
    return self.Call("factorial", args).Float()
}

// Calculates a catmum rom value.
func (self *Math) CatmullRomI(args ...interface{}) float64{
    return self.Call("catmullRom", args).Float()
}

// The absolute difference between two values.
func (self *Math) DifferenceI(args ...interface{}) float64{
    return self.Call("difference", args).Float()
}

// Round to the next whole number _away_ from zero.
func (self *Math) RoundAwayFromZeroI(args ...interface{}) int{
    return self.Call("roundAwayFromZero", args).Int()
}

// Generate a sine and cosine table simultaneously and extremely quickly.
// The parameters allow you to specify the length, amplitude and frequency of the wave.
// This generator is fast enough to be used in real-time.
// Code based on research by Franky of scene.at
func (self *Math) SinCosGeneratorI(args ...interface{}) interface{}{
    return self.Call("sinCosGenerator", args)
}

// Returns the euclidian distance between the two given set of coordinates.
func (self *Math) DistanceI(args ...interface{}) float64{
    return self.Call("distance", args).Float()
}

// Returns the euclidean distance squared between the two given set of
// coordinates (cuts out a square root operation before returning).
func (self *Math) DistanceSqI(args ...interface{}) float64{
    return self.Call("distanceSq", args).Float()
}

// Returns the distance between the two given set of coordinates at the power given.
func (self *Math) DistancePowI(args ...interface{}) float64{
    return self.Call("distancePow", args).Float()
}

// Force a value within the boundaries by clamping it to the range `min`, `max`.
func (self *Math) ClampI(args ...interface{}) float64{
    return self.Call("clamp", args).Float()
}

// Clamp `x` to the range `[a, Infinity)`.
// Roughly the same as `Math.max(x, a)`, except for NaN handling.
func (self *Math) ClampBottomI(args ...interface{}) float64{
    return self.Call("clampBottom", args).Float()
}

// Checks if two values are within the given tolerance of each other.
func (self *Math) WithinI(args ...interface{}) bool{
    return self.Call("within", args).Bool()
}

// Linear mapping from range <a1, a2> to range <b1, b2>
func (self *Math) MapLinearI(args ...interface{}) float64{
    return self.Call("mapLinear", args).Float()
}

// Smoothstep function as detailed at http://en.wikipedia.org/wiki/Smoothstep
func (self *Math) SmoothstepI(args ...interface{}) float64{
    return self.Call("smoothstep", args).Float()
}

// Smootherstep function as detailed at http://en.wikipedia.org/wiki/Smoothstep
func (self *Math) SmootherstepI(args ...interface{}) float64{
    return self.Call("smootherstep", args).Float()
}

// A value representing the sign of the value: -1 for negative, +1 for positive, 0 if value is 0.
// 
// This works differently from `Math.sign` for values of NaN and -0, etc.
func (self *Math) SignI(args ...interface{}) int{
    return self.Call("sign", args).Int()
}

// Work out what percentage value `a` is of value `b` using the given base.
func (self *Math) PercentI(args ...interface{}) float64{
    return self.Call("percent", args).Float()
}

// Convert degrees to radians.
func (self *Math) DegToRadI(args ...interface{}) float64{
    return self.Call("degToRad", args).Float()
}

// Convert radians to degrees.
func (self *Math) RadToDegI(args ...interface{}) float64{
    return self.Call("radToDeg", args).Float()
}
