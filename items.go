package main

import "fmt"

func LoadItems() []Item {
	var items = []Item{
		{
			"-",
			0,
			0,
		},
		{
			"Iron Hoe",
			HOE,
			IRON_HOE,
		},
		{
			"Grass Seeds",
			SOIL_SEED,
			GRASS_SEEDS,
		},
		{
			"Tall Grass Starter",
			GRASS_PLANT,
			TALL_GRASS_STARTER,
		},
		{
			"Sapling",
			GRASS_PLANT,
			SAPLING,
		},
		{
			"Delete",
			OVERWRITE_PLACEABLE,
			DELETE,
		},
		{
			"Grass Tile",
			OVERWRITE_PLACEABLE,
			GRASS_TILE,
		},
		{
			"Soil Tile",
			OVERWRITE_PLACEABLE,
			SOIL_TILE,
		},
		{
			"Water Tile",
			OVERWRITE_PLACEABLE,
			WATER_TILE,
		},
		{
			"Stone Tile",
			OVERWRITE_PLACEABLE,
			STONE_TILE,
		},
		{
			"Bridge Tile",
			OVERWRITE_PLACEABLE,
			BRIDGE_TILE,
		},
	}

	fmt.Printf("Loaded items!\n")

	return items
}

func LoadInventory() {
	for i := range player.Inventory {
		player.Inventory[i] = items[NONE]
	}

	player.Inventory[0] = items[IRON_HOE]
	player.Inventory[1] = items[GRASS_SEEDS]
	player.Inventory[2] = items[TALL_GRASS_STARTER]
	player.Inventory[3] = items[SAPLING]
	player.Inventory[4] = items[DELETE]
	player.Inventory[9] = items[GRASS_TILE]
	player.Inventory[10] = items[SOIL_TILE]
	player.Inventory[11] = items[WATER_TILE]
	player.Inventory[12] = items[STONE_TILE]
	player.Inventory[13] = items[BRIDGE_TILE]
}
