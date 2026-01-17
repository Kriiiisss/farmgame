package main

// Basic stuff
const (
	NONE                        int     = 0
	TILE_SIZE                   float32 = 32
	TILES_RENDER_TOLERANCE      float32 = 3
	PLACEABLES_RENDER_TOLERANCE float32 = 6
	MAX_INVENTORY_SIZE          int     = 27
	EPSILON                     float32 = 0.125
	GRASS_TINT                          = 0x91bd59ff
	SQRT2                       float32 = 1.41421356
)

// Client states
const (
	MAIN_MENU = iota
	LOADING_TO_WORLD
	IN_A_WORLD
)

// Main menu sections
const (
	WELCOME = iota
	MANAGE_WORLDS
	OPTIONS
)

// Manage worlds sections
const (
	MNGWORLD_WELCOME = iota
	CREATE_NEW_WORLD
	DELETE_WORLD
)

// Option sections
const (
	OPT_WELCOME = iota
	VOLUME
)

// Placeables
const (
	TALL_GRASS int = 1
	TREE       int = 2
)

// HUD Textures
const (
	INDEX_TEXTURE_ID = iota
	INDEX_SEL_TEXTURE_ID
	SLOT_TEXTURE_ID
	SLOT_SEL_TEXTURE_ID
	SLOT_UNAV_TEXTURE_ID
	HIGHLIGHT_TEXTURE_ID
)

// Hexes of tiles
var hexToTileId = map[uint32]int{
	0x5c4300ff: SOIL,
	0x99e550ff: GRASS,
	0x959491ff: STONE,
	0x5cfff8ff: WATER,
	0x7a5a00ff: BRIDGE,
}

// Tiles
const (
	WATER  = 1
	SOIL   = 2
	GRASS  = 3
	STONE  = 4
	BRIDGE = 5
)

// Items
const (
	NULL = iota
	IRON_HOE
	IRON_SHOVEL
	GRASS_SEEDS
	TALL_GRASS_STARTER
	SAPLING
	DELETE
	GRASS_TILE
	SOIL_TILE
	WATER_TILE
	STONE_TILE
	BRIDGE_TILE
)

// Item categories
const (
	SOIL_SEED           = 1
	TILLED_SOIL_SEED    = 2
	GRASS_PLANT         = 3
	HOE                 = 4
	SHOVEL              = 5
	OVERWRITE_TILE      = 6
	OVERWRITE_PLACEABLE = 7
)

// Mesh tile maps
const MESH_COUNT = 4
const (
	SOIL_MESH = iota
	STONE_MESH
	GRASS_MESH
	BRIDGE_MESH
)

// Debug dimensions
const (
	DEBUG_FONTSIZE_SCREEN int32 = 20
	DEBUG_FONTSIZE_TILE   int32 = int32(TILE_SIZE) / 6
	DEBUG_PADDING               = 10
	DEBUG_WIDTH                 = 300
)

// Welcome buttons
const (
	BTN_MANAGEWORLDS = iota
	BTN_OPTIONS
)

// Manage worlds buttons
const (
	BTN_CREATEWORLD = iota
	BTN_LOADWORLD
	BTN_DELETEWORLD
	BTN_REFRESH
)

// Create world buttons
const (
	BTN_CREATE = iota
	BTN_CREWORLD_CANCEL
)

// Delete world buttons
const (
	BTN_DELETE = iota
	BTN_DELWORLD_CANCEL
)

// Options buttons
const (
	BTN_VOLUME = iota
)

// Inventory dimensions (percentage of render width)
const (
	HOTBAR_SLOT_PADDING float32 = 0.010
	HOTBAR_SLOT_SIZE    float32 = 0.0375
	HOTBAR_INDEX_SIZE   float32 = HOTBAR_SLOT_SIZE / 7
)
