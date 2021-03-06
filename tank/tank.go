package tank

import (
	tl "github.com/JoelOtter/termloop"
)

type Tank struct {
	*tl.Entity
	direction int
}

// Tank directions
const (
	UP    int = 1
	DOWN  int = 2
	LEFT  int = 3
	RIGHT int = 4
)

//var (
//	canvasUp     tl.Canvas
//	canvasRight  tl.Canvas
//	canvasDown   tl.Canvas
//	canvasLeft   tl.Canvas
//	tankBodyCell tl.Cell
//)

func (tank *Tank) Draw(screen *tl.Screen) {

	tank.Entity.Draw(screen)
}

func TankUpCanvas(tankBodyCell tl.Cell) tl.Canvas {

	canvasUp := tl.NewCanvas(9, 9)

	// Tank canvas up
	canvasUp[4][0] = tankBodyCell
	canvasUp[4][1] = tankBodyCell
	canvasUp[4][2] = tankBodyCell
	canvasUp[4][3] = tankBodyCell
	canvasUp[4][4] = tankBodyCell
	canvasUp[4][5] = tankBodyCell

	canvasUp[1][2] = tankBodyCell
	canvasUp[1][3] = tankBodyCell
	canvasUp[1][4] = tankBodyCell
	canvasUp[1][5] = tankBodyCell
	canvasUp[1][6] = tankBodyCell
	canvasUp[1][7] = tankBodyCell

	canvasUp[2][2] = tankBodyCell
	canvasUp[2][3] = tankBodyCell
	canvasUp[2][4] = tankBodyCell
	canvasUp[2][5] = tankBodyCell
	canvasUp[2][6] = tankBodyCell
	canvasUp[2][7] = tankBodyCell

	canvasUp[6][2] = tankBodyCell
	canvasUp[6][3] = tankBodyCell
	canvasUp[6][4] = tankBodyCell
	canvasUp[6][5] = tankBodyCell
	canvasUp[6][6] = tankBodyCell
	canvasUp[6][7] = tankBodyCell

	canvasUp[7][2] = tankBodyCell
	canvasUp[7][3] = tankBodyCell
	canvasUp[7][4] = tankBodyCell
	canvasUp[7][5] = tankBodyCell
	canvasUp[7][6] = tankBodyCell
	canvasUp[7][7] = tankBodyCell

	canvasUp[3][4] = tankBodyCell
	canvasUp[3][5] = tankBodyCell
	canvasUp[5][4] = tankBodyCell
	canvasUp[5][5] = tankBodyCell

	return canvasUp
}

func TankDownCanvas(tankBodyCell tl.Cell) tl.Canvas {

	canvasDown := tl.NewCanvas(9, 9)

	// Tank canvas down
	canvasDown[1][1] = tankBodyCell
	canvasDown[1][2] = tankBodyCell
	canvasDown[1][3] = tankBodyCell
	canvasDown[1][4] = tankBodyCell
	canvasDown[1][5] = tankBodyCell
	canvasDown[1][6] = tankBodyCell

	canvasDown[2][1] = tankBodyCell
	canvasDown[2][2] = tankBodyCell
	canvasDown[2][3] = tankBodyCell
	canvasDown[2][4] = tankBodyCell
	canvasDown[2][5] = tankBodyCell
	canvasDown[2][6] = tankBodyCell

	canvasDown[6][1] = tankBodyCell
	canvasDown[6][2] = tankBodyCell
	canvasDown[6][3] = tankBodyCell
	canvasDown[6][4] = tankBodyCell
	canvasDown[6][5] = tankBodyCell
	canvasDown[6][6] = tankBodyCell

	canvasDown[7][1] = tankBodyCell
	canvasDown[7][2] = tankBodyCell
	canvasDown[7][3] = tankBodyCell
	canvasDown[7][4] = tankBodyCell
	canvasDown[7][5] = tankBodyCell
	canvasDown[7][6] = tankBodyCell

	canvasDown[4][3] = tankBodyCell
	canvasDown[4][4] = tankBodyCell
	canvasDown[4][5] = tankBodyCell
	canvasDown[4][6] = tankBodyCell
	canvasDown[4][7] = tankBodyCell
	canvasDown[4][8] = tankBodyCell

	canvasDown[3][3] = tankBodyCell
	canvasDown[3][4] = tankBodyCell
	canvasDown[5][3] = tankBodyCell
	canvasDown[5][4] = tankBodyCell

	return canvasDown
}

func TankLeftCanvas(tankBodyCell tl.Cell) tl.Canvas {

	canvasLeft := tl.NewCanvas(9, 9)

	// Tank canvas left
	canvasLeft[2][1] = tankBodyCell
	canvasLeft[3][1] = tankBodyCell
	canvasLeft[4][1] = tankBodyCell
	canvasLeft[5][1] = tankBodyCell
	canvasLeft[6][1] = tankBodyCell
	canvasLeft[7][1] = tankBodyCell

	canvasLeft[2][2] = tankBodyCell
	canvasLeft[3][2] = tankBodyCell
	canvasLeft[4][2] = tankBodyCell
	canvasLeft[5][2] = tankBodyCell
	canvasLeft[6][2] = tankBodyCell
	canvasLeft[7][2] = tankBodyCell

	canvasLeft[2][6] = tankBodyCell
	canvasLeft[3][6] = tankBodyCell
	canvasLeft[4][6] = tankBodyCell
	canvasLeft[5][6] = tankBodyCell
	canvasLeft[6][6] = tankBodyCell
	canvasLeft[7][6] = tankBodyCell

	canvasLeft[2][7] = tankBodyCell
	canvasLeft[3][7] = tankBodyCell
	canvasLeft[4][7] = tankBodyCell
	canvasLeft[5][7] = tankBodyCell
	canvasLeft[6][7] = tankBodyCell
	canvasLeft[7][7] = tankBodyCell

	canvasLeft[0][4] = tankBodyCell
	canvasLeft[1][4] = tankBodyCell
	canvasLeft[2][4] = tankBodyCell
	canvasLeft[3][4] = tankBodyCell
	canvasLeft[4][4] = tankBodyCell
	canvasLeft[5][4] = tankBodyCell

	canvasLeft[4][3] = tankBodyCell
	canvasLeft[5][3] = tankBodyCell
	canvasLeft[4][5] = tankBodyCell
	canvasLeft[5][5] = tankBodyCell

	return canvasLeft
}

func TankRightCanvas(tankBodyCell tl.Cell) tl.Canvas {

	canvasRight := tl.NewCanvas(9, 9)

	// Tank canvas right
	canvasRight[2][1] = tankBodyCell
	canvasRight[3][1] = tankBodyCell
	canvasRight[4][1] = tankBodyCell
	canvasRight[5][1] = tankBodyCell
	canvasRight[6][1] = tankBodyCell
	canvasRight[1][1] = tankBodyCell

	canvasRight[2][2] = tankBodyCell
	canvasRight[3][2] = tankBodyCell
	canvasRight[4][2] = tankBodyCell
	canvasRight[5][2] = tankBodyCell
	canvasRight[6][2] = tankBodyCell
	canvasRight[1][2] = tankBodyCell

	canvasRight[2][6] = tankBodyCell
	canvasRight[3][6] = tankBodyCell
	canvasRight[4][6] = tankBodyCell
	canvasRight[5][6] = tankBodyCell
	canvasRight[6][6] = tankBodyCell
	canvasRight[1][6] = tankBodyCell

	canvasRight[2][7] = tankBodyCell
	canvasRight[3][7] = tankBodyCell
	canvasRight[4][7] = tankBodyCell
	canvasRight[5][7] = tankBodyCell
	canvasRight[6][7] = tankBodyCell
	canvasRight[1][7] = tankBodyCell

	canvasRight[7][4] = tankBodyCell
	canvasRight[3][4] = tankBodyCell
	canvasRight[4][4] = tankBodyCell
	canvasRight[5][4] = tankBodyCell
	canvasRight[6][4] = tankBodyCell
	canvasRight[8][4] = tankBodyCell

	canvasRight[3][3] = tankBodyCell
	canvasRight[4][3] = tankBodyCell
	canvasRight[3][5] = tankBodyCell
	canvasRight[4][5] = tankBodyCell

	return canvasRight
}

func init() {

	//	tankBodyCell = tl.Cell{Fg: tl.ColorRed, Bg: tl.ColorRed}

	// new blank canvas

}

// Initial a new tank
func NewTank(cell tl.Cell) *Tank {

	tank := Tank{
		Entity: tl.NewEntity(0, 0, 9, 9),
	}

	TankUp(&tank, cell)

	return &tank

}

// Initial a tank with position
func NewTankXY(x, y int, cell tl.Cell) *Tank {

	tank := Tank{
		Entity: tl.NewEntity(x, y, 9, 9),
	}

	TankUp(&tank, cell)

	return &tank

}

func TankUp(tank *Tank, cell tl.Cell) {

	// Refresh tank direction
	canvas := TankUpCanvas(cell)
	tank.SetCanvas(&canvas)
	tank.direction = UP
}

func TankDown(tank *Tank, cell tl.Cell) {

	// Refresh tank direction
	canvas := TankDownCanvas(cell)
	tank.SetCanvas(&canvas)
	tank.direction = DOWN
}

func TankLeft(tank *Tank, cell tl.Cell) {

	// Refresh tank direction
	canvas := TankLeftCanvas(cell)
	tank.SetCanvas(&canvas)
	tank.direction = LEFT
}

func TankRight(tank *Tank, cell tl.Cell) {

	// Refresh tank direction
	canvas := TankRightCanvas(cell)
	tank.SetCanvas(&canvas)
	tank.direction = RIGHT
}

func (tank *Tank) GetDirection() int {
	return tank.direction
}
