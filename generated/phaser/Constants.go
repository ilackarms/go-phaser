package phaser

import "github.com/gopherjs/gopherjs/js"

// this file created manually!

const (
	/**
	* The Phaser version number.
	* @constant
	* @type {string}
	 */
	VERSION = "2.6.2"

	/**
	* AUTO renderer - picks between WebGL or Canvas based on device.
	* @constant
	* @type {integer}
	 */
	AUTO = 0

	/**
	* Canvas Renderer.
	* @constant
	* @type {integer}
	 */
	CANVAS = 1

	/**
	* WebGL Renderer.
	* @constant
	* @type {integer}
	 */
	WEBGL = 2

	/**
	* Headless renderer (not visual output)
	* @constant
	* @type {integer}
	 */
	HEADLESS = 3

	/**
	* Direction constant.
	* @constant
	* @type {integer}
	 */
	NONE = 0

	/**
	* Direction constant.
	* @constant
	* @type {integer}
	 */
	LEFT = 1

	/**
	* Direction constant.
	* @constant
	* @type {integer}
	 */
	RIGHT = 2

	/**
	* Direction constant.
	* @constant
	* @type {integer}
	 */
	UP = 3

	/**
	* Direction constant.
	* @constant
	* @type {integer}
	 */
	DOWN = 4

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	SPRITE = 0

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	BUTTON = 1

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	IMAGE = 2

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	GRAPHICS = 3

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	TEXT = 4

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	TILESPRITE = 5

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	BITMAPTEXT = 6

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	GROUP = 7

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	RENDERTEXTURE = 8

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	TILEMAP = 9

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	TILEMAPLAYER = 10

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	EMITTER = 11

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	POLYGON = 12

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	BITMAPDATA = 13

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	CANVAS_FILTER = 14

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	WEBGL_FILTER = 15

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	ELLIPSE = 16

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	SPRITEBATCH = 17

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	RETROFONT = 18

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	POINTER = 19

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	ROPE = 20

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	CIRCLE = 21

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	RECTANGLE = 22

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	LINE = 23

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	MATRIX = 24

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	POINT = 25

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	ROUNDEDRECTANGLE = 26

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	CREATURE = 27

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	VIDEO = 28

	/**
	* Game Object type.
	* @constant
	* @type {integer}
	 */
	PENDING_ATLAS = -1

	/**
	* A horizontal orientation
	* @constant
	* @type {integer}
	 */
	HORIZONTAL = 0

	/**
	* A vertical orientation
	* @constant
	* @type {integer}
	 */
	VERTICAL = 1

	/**
	* A landscape orientation
	* @constant
	* @type {integer}
	 */
	LANDSCAPE = 0

	/**
	* A portrait orientation
	* @constant
	* @type {integer}
	 */
	PORTRAIT = 1

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face up.
	* @constant
	* @type {integer}
	 */
	ANGLE_UP = 270

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face down.
	* @constant
	* @type {integer}
	 */
	ANGLE_DOWN = 90

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face left.
	* @constant
	* @type {integer}
	 */
	ANGLE_LEFT = 180

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face right.
	* @constant
	* @type {integer}
	 */
	ANGLE_RIGHT = 0

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face north east.
	* @constant
	* @type {integer}
	 */
	ANGLE_NORTH_EAST = 315

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face north west.
	* @constant
	* @type {integer}
	 */
	ANGLE_NORTH_WEST = 225

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face south east.
	* @constant
	* @type {integer}
	 */
	ANGLE_SOUTH_EAST = 45

	/**
	* The Angle (in degrees) a Game Object needs to be set to in order to face south west.
	* @constant
	* @type {integer}
	 */
	ANGLE_SOUTH_WEST = 135

	/**
	* A constant representing a top-left alignment or position.
	* @constant
	* @type {integer}
	 */
	TOP_LEFT = 0

	/**
	* A constant representing a top-center alignment or position.
	* @constant
	* @type {integer}
	 */
	TOP_CENTER = 1

	/**
	* A constant representing a top-right alignment or position.
	* @constant
	* @type {integer}
	 */
	TOP_RIGHT = 2

	/**
	* A constant representing a left-top alignment or position.
	* @constant
	* @type {integer}
	 */
	LEFT_TOP = 3

	/**
	* A constant representing a left-center alignment or position.
	* @constant
	* @type {integer}
	 */
	LEFT_CENTER = 4

	/**
	* A constant representing a left-bottom alignment or position.
	* @constant
	* @type {integer}
	 */
	LEFT_BOTTOM = 5

	/**
	* A constant representing a center alignment or position.
	* @constant
	* @type {integer}
	 */
	CENTER = 6

	/**
	* A constant representing a right-top alignment or position.
	* @constant
	* @type {integer}
	 */
	RIGHT_TOP = 7

	/**
	* A constant representing a right-center alignment or position.
	* @constant
	* @type {integer}
	 */
	RIGHT_CENTER = 8

	/**
	* A constant representing a right-bottom alignment or position.
	* @constant
	* @type {integer}
	 */
	RIGHT_BOTTOM = 9

	/**
	* A constant representing a bottom-left alignment or position.
	* @constant
	* @type {integer}
	 */
	BOTTOM_LEFT = 10

	/**
	* A constant representing a bottom-center alignment or position.
	* @constant
	* @type {integer}
	 */
	BOTTOM_CENTER = 11

	/**
	* A constant representing a bottom-right alignment or position.
	* @constant
	* @type {integer}
	 */
	BOTTOM_RIGHT = 12

	/**
	* @const
	* @type {number}
	 */
	Physics_ARCADE = 0

	/**
	* @const
	* @type {number}
	 */
	Physics_P2JS = 1

	/**
	* @const
	* @type {number}
	 */
	Physics_NINJA = 2

	/**
	* @const
	* @type {number}
	 */
	Physics_BOX2D = 3

	/**
	* @const
	* @type {number}
	 */
	Physics_CHIPMUNK = 4

	/**
	* @const
	* @type {number}
	 */
	Physics_MATTERJS = 5

	/**
	* A constant used for the sortDirection value.
	* Use this if you don't wish to perform any pre-collision sorting at all, or will manually sort your Groups.
	* @constant
	* @type {number}
	 */
	Physics_Arcade_SORT_NONE = 0

	/**
	* A constant used for the sortDirection value.
	* Use this if your game world is wide but short and scrolls from the left to the right (i.e. Mario)
	* @constant
	* @type {number}
	 */
	Physics_Arcade_LEFT_RIGHT = 1

	/**
	* A constant used for the sortDirection value.
	* Use this if your game world is wide but short and scrolls from the right to the left (i.e. Mario backwards)
	* @constant
	* @type {number}
	 */
	Physics_Arcade_RIGHT_LEFT = 2

	/**
	* A constant used for the sortDirection value.
	* Use this if your game world is narrow but tall and scrolls from the top to the bottom (i.e. Dig Dug)
	* @constant
	* @type {number}
	 */
	Physics_Arcade_TOP_BOTTOM = 3

	/**
	* A constant used for the sortDirection value.
	* Use this if your game world is narrow but tall and scrolls from the bottom to the top (i.e. Commando or a vertically scrolling shoot-em-up)
	* @constant
	* @type {number}
	 */
	Physics_Arcade_BOTTOM_TOP = 4

	/**
	* @constant
	* @type {number}
	 */
	Tilemap_CSV = 0

	/**
	 * @constant
	 * @type {number}
	 */
	Tilemap_TILED_JSON = 1

	/**
	 * @constant
	 * @type {number}
	 */
	Tilemap_NORTH = 0

	/**
	 * @constant
	 * @type {number}
	 */
	Tilemap_EAST = 1

	/**
	 * @constant
	 * @type {number}
	 */
	Tilemap_SOUTH = 2

	/**
	 * @constant
	 * @type {number}
	 */
	Tilemap_WEST = 3

	/**
	* A `bulletKillType` constant that stops the bullets from ever being destroyed automatically.
	* @constant
	* @type {integer}
	 */
	Weapon_KILL_NEVER = 0

	/**
	 * A `bulletKillType` constant that automatically kills the bullets when their `bulletLifespan` expires.
	 * @constant
	 * @type {integer}
	 */
	Weapon_KILL_LIFESPAN = 1

	/**
	 * A `bulletKillType` constant that automatically kills the bullets after they
	 * exceed the `bulletDistance` from their original firing position.
	 * @constant
	 * @type {integer}
	 */
	Weapon_KILL_DISTANCE = 2

	/**
	 * A `bulletKillType` constant that automatically kills the bullets when they leave the `Weapon.bounds` rectangle.
	 * @constant
	 * @type {integer}
	 */
	Weapon_KILL_WEAPON_BOUNDS = 3

	/**
	 * A `bulletKillType` constant that automatically kills the bullets when they leave the `Camera.bounds` rectangle.
	 * @constant
	 * @type {integer}
	 */
	Weapon_KILL_CAMERA_BOUNDS = 4

	/**
	 * A `bulletKillType` constant that automatically kills the bullets when they leave the `World.bounds` rectangle.
	 * @constant
	 * @type {integer}
	 */
	Weapon_KILL_WORLD_BOUNDS = 5

	/**
	 * A `bulletKillType` constant that automatically kills the bullets when they leave the `Weapon.bounds` rectangle.
	 * @constant
	 * @type {integer}
	 */
	Weapon_KILL_STATIC_BOUNDS = 6
)

var (
	/**
	* An array of Phaser game instances.
	* @constant
	* @type {array}
	 */
	GAMES = js.Global.Get("GAMES")

	/**
	 * Various blend modes supported by Pixi.
	 *
	 * IMPORTANT = The WebGL renderer only supports the NORMAL ADD MULTIPLY and SCREEN blend modes.
	 *
	 * @constant
	 * @property {Number} blendModes.NORMAL
	 * @property {Number} blendModes.ADD
	 * @property {Number} blendModes.MULTIPLY
	 * @property {Number} blendModes.SCREEN
	 * @property {Number} blendModes.OVERLAY
	 * @property {Number} blendModes.DARKEN
	 * @property {Number} blendModes.LIGHTEN
	 * @property {Number} blendModes.COLOR_DODGE
	 * @property {Number} blendModes.COLOR_BURN
	 * @property {Number} blendModes.HARD_LIGHT
	 * @property {Number} blendModes.SOFT_LIGHT
	 * @property {Number} blendModes.DIFFERENCE
	 * @property {Number} blendModes.EXCLUSION
	 * @property {Number} blendModes.HUE
	 * @property {Number} blendModes.SATURATION
	 * @property {Number} blendModes.COLOR
	 * @property {Number} blendModes.LUMINOSITY
	 * @static
	 */
	BlendModes = struct {
		NORMAL      int
		ADD         int
		MULTIPLY    int
		SCREEN      int
		OVERLAY     int
		DARKEN      int
		LIGHTEN     int
		COLOR_DODGE int
		COLOR_BURN  int
		HARD_LIGHT  int
		SOFT_LIGHT  int
		DIFFERENCE  int
		EXCLUSION   int
		HUE         int
		SATURATION  int
		COLOR       int
		LUMINOSITY  int
	}{
		NORMAL:      0,
		ADD:         1,
		MULTIPLY:    2,
		SCREEN:      3,
		OVERLAY:     4,
		DARKEN:      5,
		LIGHTEN:     6,
		COLOR_DODGE: 7,
		COLOR_BURN:  8,
		HARD_LIGHT:  9,
		SOFT_LIGHT:  10,
		DIFFERENCE:  11,
		EXCLUSION:   12,
		HUE:         13,
		SATURATION:  14,
		COLOR:       15,
		LUMINOSITY:  16,
	}

	/**
	 * The scale modes that are supported by Pixi.
	 *
	 * The DEFAULT scale mode affects the default scaling mode of future operations.
	 * It can be re-assigned to either LINEAR or NEAREST depending upon suitability.
	 *
	 * @constant
	 * @property {Object} scaleModes
	 * @property {Number} scaleModes.DEFAULT=LINEAR
	 * @property {Number} scaleModes.LINEAR Smooth scaling
	 * @property {Number} scaleModes.NEAREST Pixelating scaling
	 * @static
	 */
	ScaleModes = struct {
		DEFAULT int
		LINEAR  int
		NEAREST int
	}{
		DEFAULT: 0,
		LINEAR:  0,
		NEAREST: 1,
	}
)
