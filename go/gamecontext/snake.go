package gamecontext

type Direction int

const (
	UP Direction = 1 + iota
	DOWN
	LEFT
	RIGHT
)

type Vector2 struct {
	X int
	Y int
}

type Snake struct {
	headPosition Vector2
	dir          Direction
	body         []Vector2
}

func (s *Snake) addBodypart() {
	s.body = append(s.body, s.headPosition)
}

func (s *Snake) removeTail() {
	if len(s.body) != 0 {
		s.body = s.body[1:]
	}
}

func (s *Snake) move() {
	s.addBodypart()
	s.removeTail()
}

func newSnake() (snake Snake) {
	snake.dir = LEFT
	snake.headPosition = Vector2{
		X: FIELD_SIZE_X / 2,
		Y: FIELD_SIZE_Y / 2,
	}
	snake.addBodypart()
	snake.addBodypart()
	snake.addBodypart()
	return
}

func (s *Snake) propogate() {
	switch s.dir {
	case UP:
		s.headPosition.Y -= 1
	case DOWN:
		s.headPosition.Y += 1
	case LEFT:
		s.headPosition.X -= 1
	case RIGHT:
		s.headPosition.X += 1
	}
	s.move()
}
