package termgame

type Game struct {
	board   [][]int
	playerX int
	playerY int

	boardHeight int
	boardWidth  int
}

func NewGame() *Game {
	g := new(Game)
	g.boardHeight = 16
	g.boardWidth = 16

	//g.resetGame()

	return g
}

func (g *Game) resetGame() {
	g.board = make([][]int, g.boardHeight)

	for i := 0; i < g.boardHeight; i++ {
		g.board[i] = make([]int, g.boardWidth)
		for j := 0; j < g.boardWidth; j++ {
			g.board[i][j] = 0
		}
	}

	g.board[0][0] = 1
	g.playerX = 0
	g.playerY = 0

}
