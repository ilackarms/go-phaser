// Package phaser Automatic generation for Phaser.RandomDataGenerator
// generated file RandomDataGenerator.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

// RandomDataGenerator An extremely useful repeatable random data generator.
//
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
//
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
type RandomDataGenerator struct {
	*js.Object
}

// NewRandomDataGenerator An extremely useful repeatable random data generator.
//
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
//
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
func NewRandomDataGenerator() *RandomDataGenerator {
	return &RandomDataGenerator{js.Global.Get("Phaser").Get("RandomDataGenerator").New()}
}

// NewRandomDataGenerator1O An extremely useful repeatable random data generator.
//
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
//
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
func NewRandomDataGenerator1O(seeds interface{}) *RandomDataGenerator {
	return &RandomDataGenerator{js.Global.Get("Phaser").Get("RandomDataGenerator").New(seeds)}
}

// NewRandomDataGeneratorI An extremely useful repeatable random data generator.
//
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
//
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
func NewRandomDataGeneratorI(args ...interface{}) *RandomDataGenerator {
	return &RandomDataGenerator{js.Global.Get("Phaser").Get("RandomDataGenerator").New(args...)}
}

// RandomDataGenerator Binding conversion method to RandomDataGenerator point
func ToRandomDataGenerator(jsStruct interface{}) *RandomDataGenerator {
	if object, ok := jsStruct.(*js.Object); ok {
		return &RandomDataGenerator{Object: object}
	}
	return nil
}

// Rnd Private random helper.
func (self *RandomDataGenerator) Rnd() int {
	return self.Object.Call("rnd").Int()
}

// RndI Private random helper.
func (self *RandomDataGenerator) RndI(args ...interface{}) int {
	return self.Object.Call("rnd", args...).Int()
}

// Sow Reset the seed of the random data generator.
//
// _Note_: the seed array is only processed up to the first `undefined` (or `null`) value, should such be present.
func (self *RandomDataGenerator) Sow(seeds []interface{}) {
	self.Object.Call("sow", seeds)
}

// SowI Reset the seed of the random data generator.
//
// _Note_: the seed array is only processed up to the first `undefined` (or `null`) value, should such be present.
func (self *RandomDataGenerator) SowI(args ...interface{}) {
	self.Object.Call("sow", args...)
}

// Hash Internal method that creates a seed hash.
func (self *RandomDataGenerator) Hash(data interface{}) int {
	return self.Object.Call("hash", data).Int()
}

// HashI Internal method that creates a seed hash.
func (self *RandomDataGenerator) HashI(args ...interface{}) int {
	return self.Object.Call("hash", args...).Int()
}

// Integer Returns a random integer between 0 and 2^32.
func (self *RandomDataGenerator) Integer() int {
	return self.Object.Call("integer").Int()
}

// IntegerI Returns a random integer between 0 and 2^32.
func (self *RandomDataGenerator) IntegerI(args ...interface{}) int {
	return self.Object.Call("integer", args...).Int()
}

// Frac Returns a random real number between 0 and 1.
func (self *RandomDataGenerator) Frac() int {
	return self.Object.Call("frac").Int()
}

// FracI Returns a random real number between 0 and 1.
func (self *RandomDataGenerator) FracI(args ...interface{}) int {
	return self.Object.Call("frac", args...).Int()
}

// Real Returns a random real number between 0 and 2^32.
func (self *RandomDataGenerator) Real() int {
	return self.Object.Call("real").Int()
}

// RealI Returns a random real number between 0 and 2^32.
func (self *RandomDataGenerator) RealI(args ...interface{}) int {
	return self.Object.Call("real", args...).Int()
}

// IntegerInRange Returns a random integer between and including min and max.
func (self *RandomDataGenerator) IntegerInRange(min int, max int) int {
	return self.Object.Call("integerInRange", min, max).Int()
}

// IntegerInRangeI Returns a random integer between and including min and max.
func (self *RandomDataGenerator) IntegerInRangeI(args ...interface{}) int {
	return self.Object.Call("integerInRange", args...).Int()
}

// Between Returns a random integer between and including min and max.
// This method is an alias for RandomDataGenerator.integerInRange.
func (self *RandomDataGenerator) Between(min int, max int) int {
	return self.Object.Call("between", min, max).Int()
}

// BetweenI Returns a random integer between and including min and max.
// This method is an alias for RandomDataGenerator.integerInRange.
func (self *RandomDataGenerator) BetweenI(args ...interface{}) int {
	return self.Object.Call("between", args...).Int()
}

// RealInRange Returns a random real number between min and max.
func (self *RandomDataGenerator) RealInRange(min int, max int) int {
	return self.Object.Call("realInRange", min, max).Int()
}

// RealInRangeI Returns a random real number between min and max.
func (self *RandomDataGenerator) RealInRangeI(args ...interface{}) int {
	return self.Object.Call("realInRange", args...).Int()
}

// Normal Returns a random real number between -1 and 1.
func (self *RandomDataGenerator) Normal() int {
	return self.Object.Call("normal").Int()
}

// NormalI Returns a random real number between -1 and 1.
func (self *RandomDataGenerator) NormalI(args ...interface{}) int {
	return self.Object.Call("normal", args...).Int()
}

// Uuid Returns a valid RFC4122 version4 ID hex string from https://gist.github.com/1308368
func (self *RandomDataGenerator) Uuid() string {
	return self.Object.Call("uuid").String()
}

// UuidI Returns a valid RFC4122 version4 ID hex string from https://gist.github.com/1308368
func (self *RandomDataGenerator) UuidI(args ...interface{}) string {
	return self.Object.Call("uuid", args...).String()
}

// Pick Returns a random member of `array`.
func (self *RandomDataGenerator) Pick(ary []interface{}) interface{} {
	return self.Object.Call("pick", ary)
}

// PickI Returns a random member of `array`.
func (self *RandomDataGenerator) PickI(args ...interface{}) interface{} {
	return self.Object.Call("pick", args...)
}

// Sign Returns a sign to be used with multiplication operator.
func (self *RandomDataGenerator) Sign() int {
	return self.Object.Call("sign").Int()
}

// SignI Returns a sign to be used with multiplication operator.
func (self *RandomDataGenerator) SignI(args ...interface{}) int {
	return self.Object.Call("sign", args...).Int()
}

// WeightedPick Returns a random member of `array`, favoring the earlier entries.
func (self *RandomDataGenerator) WeightedPick(ary []interface{}) interface{} {
	return self.Object.Call("weightedPick", ary)
}

// WeightedPickI Returns a random member of `array`, favoring the earlier entries.
func (self *RandomDataGenerator) WeightedPickI(args ...interface{}) interface{} {
	return self.Object.Call("weightedPick", args...)
}

// Timestamp Returns a random timestamp between min and max, or between the beginning of 2000 and the end of 2020 if min and max aren't specified.
func (self *RandomDataGenerator) Timestamp(min int, max int) int {
	return self.Object.Call("timestamp", min, max).Int()
}

// TimestampI Returns a random timestamp between min and max, or between the beginning of 2000 and the end of 2020 if min and max aren't specified.
func (self *RandomDataGenerator) TimestampI(args ...interface{}) int {
	return self.Object.Call("timestamp", args...).Int()
}

// Angle Returns a random angle between -180 and 180.
func (self *RandomDataGenerator) Angle() int {
	return self.Object.Call("angle").Int()
}

// AngleI Returns a random angle between -180 and 180.
func (self *RandomDataGenerator) AngleI(args ...interface{}) int {
	return self.Object.Call("angle", args...).Int()
}

// State Gets or Sets the state of the generator. This allows you to retain the values
// that the generator is using between games, i.e. in a game save file.
//
// To seed this generator with a previously saved state you can pass it as the
// `seed` value in your game config, or call this method directly after Phaser has booted.
//
// Call this method with no parameters to return the current state.
//
// If providing a state it should match the same format that this method
// returns, which is a string with a header `!rnd` followed by the `c`,
// `s0`, `s1` and `s2` values respectively, each comma-delimited.
func (self *RandomDataGenerator) State() string {
	return self.Object.Call("state").String()
}

// State1O Gets or Sets the state of the generator. This allows you to retain the values
// that the generator is using between games, i.e. in a game save file.
//
// To seed this generator with a previously saved state you can pass it as the
// `seed` value in your game config, or call this method directly after Phaser has booted.
//
// Call this method with no parameters to return the current state.
//
// If providing a state it should match the same format that this method
// returns, which is a string with a header `!rnd` followed by the `c`,
// `s0`, `s1` and `s2` values respectively, each comma-delimited.
func (self *RandomDataGenerator) State1O(state string) string {
	return self.Object.Call("state", state).String()
}

// StateI Gets or Sets the state of the generator. This allows you to retain the values
// that the generator is using between games, i.e. in a game save file.
//
// To seed this generator with a previously saved state you can pass it as the
// `seed` value in your game config, or call this method directly after Phaser has booted.
//
// Call this method with no parameters to return the current state.
//
// If providing a state it should match the same format that this method
// returns, which is a string with a header `!rnd` followed by the `c`,
// `s0`, `s1` and `s2` values respectively, each comma-delimited.
func (self *RandomDataGenerator) StateI(args ...interface{}) string {
	return self.Object.Call("state", args...).String()
}
