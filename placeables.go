package main

func loadPlaceables() []Placeable {
	var placeables = []Placeable{
		{
			Id:     0,
			AtlasX: 0,
			AtlasY: 0,
			Width:  0,
			Height: 0,
		},
		{
			Id:     TALL_GRASS,
			AtlasX: 160,
			AtlasY: 0,
			Width:  32,
			Height: 64,
		},
		{
			Id:     TREE,
			AtlasX: 0,
			AtlasY: 0,
			Width:  160,
			Height: 192,
		},
		{
			Id:     POTATO_CROP,
			AtlasX: 192,
			AtlasY: 0,
			Width:  32,
			Height: 32,
		},
		{
			Id:     CARROT_CROP,
			AtlasX: 224,
			AtlasY: 0,
			Width:  32,
			Height: 32,
		},
		{
			Id:     WHEAT_CROP,
			AtlasX: 256,
			AtlasY: 0,
			Width:  32,
			Height: 32,
		},
	}

	return placeables
}
