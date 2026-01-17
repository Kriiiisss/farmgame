package main

func loadPlaceables() []Placeable {
	var placeables = []Placeable{
		{
			"-",
			0,
			0,
			0,
			0,
			0,
		},
		{
			"Tall Grass",
			TALL_GRASS,
			160,
			0,
			32,
			64,
		},
		{
			"Tree",
			TREE,
			0,
			0,
			160,
			192,
		},
	}

	return placeables
}
