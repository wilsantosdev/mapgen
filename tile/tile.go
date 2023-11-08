package tile

import (
	"math/rand"
)

type Direction int

const (
	UP Direction = 0 + iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) Oposite() Direction {
	switch d {
	case UP:
		return DOWN
	case RIGHT:
		return LEFT
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	default:
		return UP
	}
}

type TileType string

const (
	NonCollapsedTile    TileType = "X"
	EmptyTile           TileType = " "
	UpDownLeftTile      TileType = "╣"
	UpRightDownTile     TileType = "╠"
	UpDownTile          TileType = "║"
	RightLeftTile       TileType = "═"
	UpLeftTile          TileType = "╝"
	UpRightTile         TileType = "╚"
	DownLeftTile        TileType = "╗"
	RightDownTile       TileType = "╔"
	UpRightLeftTile     TileType = "╩"
	RightDownLeftTile   TileType = "╦"
	UpRightDownLeftTile TileType = "╬"
)

func (t TileType) String() string {
	return string(t)
}

func (t TileType) getConnectors() []int {
	switch t {
	case UpDownLeftTile:
		return []int{1, 0, 1, 1}
	case UpRightDownTile:
		return []int{1, 1, 1, 0}
	case UpDownTile:
		return []int{1, 0, 1, 0}
	case RightLeftTile:
		return []int{0, 1, 0, 1}
	case UpLeftTile:
		return []int{1, 0, 0, 1}
	case UpRightTile:
		return []int{1, 1, 0, 0}
	case DownLeftTile:
		return []int{0, 0, 1, 1}
	case RightDownTile:
		return []int{0, 1, 1, 0}
	case UpRightLeftTile:
		return []int{1, 1, 0, 1}
	case RightDownLeftTile:
		return []int{0, 1, 1, 1}
	case UpRightDownLeftTile:
		return []int{1, 1, 1, 1}
	default:
		return []int{0, 0, 0, 0}
	}
}

func (t TileType) hasConnector(connector, direction int) bool {
	connectors := t.getConnectors()
	return connectors[direction] == connector
}

type Tile struct {
	tileType  TileType
	collapsed bool
	xIndex    int
	yIndex    int
	options   []TileType
}

func NewTile(xIndex, yIndex int) *Tile {
	return &Tile{
		tileType:  NonCollapsedTile,
		collapsed: false,
		xIndex:    xIndex,
		yIndex:    yIndex,
		options: []TileType{
			UpDownLeftTile,
			UpRightDownLeftTile,
			UpDownTile,
			RightLeftTile,
			UpLeftTile,
			UpRightTile,
			DownLeftTile,
			RightDownTile,
			UpRightLeftTile,
			RightDownLeftTile,
			UpRightDownTile,
		},
	}
}

func (t Tile) GetTileType() TileType {
	return t.tileType
}

func (t Tile) GetOptions() []TileType {
	return t.options
}

func (t Tile) IsCollapsed() bool {
	return t.collapsed
}

func (t Tile) GetXIndex() int {
	return t.xIndex
}

func (t Tile) GetYIndex() int {
	return t.yIndex
}

func (t *Tile) Collapse() {
	t.collapsed = true
	t.tileType = t.getRandomOption()
}

func (t Tile) getRandomOption() TileType {
	randomIndex := rand.Intn(len(t.options))
	return t.options[randomIndex]
}

func (t *Tile) UpdateOptions(tt TileType, direction Direction) {
	var newOptions []TileType
	connector := tt.getConnectors()[direction]
	for _, option := range t.options {
		if option.hasConnector(connector, int(direction.Oposite())) {
			newOptions = append(newOptions, option)
		}
	}
	t.options = newOptions
}
