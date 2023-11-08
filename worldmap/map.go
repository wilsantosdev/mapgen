package worldmap

import (
	"bufio"
	"mapgen/tile"
	"math/rand"
	"os"
	"sort"
)

type Map struct {
	width  int
	height int
	tiles  [][]tile.Tile
}

func NewMap(width int, height int) *Map {
	m := &Map{
		width:  width,
		height: height,
		tiles:  make([][]tile.Tile, height),
	}
	for i := 0; i < height; i++ {
		var tiles []tile.Tile
		for j := 0; j < width; j++ {
			tiles = append(tiles, *tile.NewTile(i, j))
		}
		m.tiles[i] = tiles
	}

	m.CollapseTiles()

	return m
}

func (m *Map) CollapseTiles() {
	// get all tiles that are not collapsed
	nonCollapsedTiles := m.getNonCollapsed()
	empty := true
	for _, row := range nonCollapsedTiles {
		if len(row) > 0 {
			empty = false
		}
	}
	if empty {
		return
	}

	// search lowest number of options
	possibleTiles := m.getLowestNumberOfOptions(nonCollapsedTiles)
	// get random tile from lowest number of options
	randomTile := m.getRandomTile(possibleTiles)
	// collapse tile
	m.tiles[randomTile.GetXIndex()][randomTile.GetYIndex()].Collapse()
	// update surrounding tiles
	m.updateSurroundingTiles(randomTile.GetXIndex(), randomTile.GetYIndex())
	// repeat until all tiles are collapsed
	m.CollapseTiles()
}

func (m *Map) updateSurroundingTiles(xIndex, yIndex int) {
	// update left tile
	if yIndex > 0 {
		m.tiles[xIndex][yIndex-1].UpdateOptions(
			m.tiles[xIndex][yIndex].GetTileType(),
			tile.LEFT,
		)
	}

	// update right tile
	if yIndex < m.width-1 {
		m.tiles[xIndex][yIndex+1].UpdateOptions(
			m.tiles[xIndex][yIndex].GetTileType(),
			tile.RIGHT,
		)
	}

	// update up tile
	if xIndex > 0 {
		m.tiles[xIndex-1][yIndex].UpdateOptions(
			m.tiles[xIndex][yIndex].GetTileType(),
			tile.UP,
		)
	}

	// update down tile
	if xIndex < m.height-1 {
		m.tiles[xIndex+1][yIndex].UpdateOptions(
			m.tiles[xIndex][yIndex].GetTileType(),
			tile.DOWN,
		)
	}

}

func (m Map) getRandomTile(possibleTiles []tile.Tile) tile.Tile {
	randomIndex := rand.Intn(len(possibleTiles))
	return possibleTiles[randomIndex]
}

func (m Map) getLowestNumberOfOptions(nonCollapsedTiles [][]tile.Tile) []tile.Tile {
	var allTiles []tile.Tile
	for _, row := range nonCollapsedTiles {
		allTiles = append(allTiles, row...)
	}
	sort.Slice(allTiles, func(i, j int) bool {
		return len(allTiles[i].GetOptions()) < len(allTiles[j].GetOptions())
	})

	numOptions := len(allTiles[0].GetOptions())
	var possibleTiles []tile.Tile

	for _, tile := range allTiles {
		if len(tile.GetOptions()) == numOptions {
			possibleTiles = append(possibleTiles, tile)
		}
	}

	return possibleTiles

}

func (m Map) getNonCollapsed() [][]tile.Tile {
	var result [][]tile.Tile

	for _, row := range m.tiles {
		var nonCollapsed []tile.Tile
		for _, tile := range row {
			if !tile.IsCollapsed() {
				nonCollapsed = append(nonCollapsed, tile)
			}
		}
		result = append(result, nonCollapsed)
	}
	return result
}

func (m Map) SaveToFile() {

	file, err := os.Create("map.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufioWriter := bufio.NewWriter(file)
	defer bufioWriter.Flush()

	for _, row := range m.tiles {
		for _, tile := range row {
			bufioWriter.WriteString(tile.GetTileType().String())
		}
		bufioWriter.WriteString("\n")
	}
}

func (m Map) GetMap() string {
	var result string
	for _, row := range m.tiles {
		for _, tile := range row {
			result += tile.GetTileType().String()
		}
		result += "\n"
	}
	return result
}
