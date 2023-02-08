package snake

import (
	"math/rand"
)

type Board struct {
    gameOver bool
    rows int
    cols int
    snake *Snake
    food Position
    score int
}

func NewBoard(rows int, cols int) *Board {
    b := &Board{
        rows: rows,
        cols: cols,
        gameOver: false,
    }
    snake := NewSnake(
        []Position{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, Right)
    b.snake = snake
    b.spawnFood()
    return b
}


func (b *Board) Update(input *Input) error {
    if b.gameOver {
        return nil
    }

    if direction, hasInput := input.GetDirection(); hasInput {
        b.snake.Turn(direction)
    }

    b.snake.Move()
    head := b.snake.Head()
    if b.isOutOfBounds(&head) || b.snake.HitsSelf() {
        b.gameOver = true
        return nil
    }

    if head == b.food{
        b.snake.Grow()
        b.spawnFood()
        b.score++
    }
    return nil
}

func (b *Board) spawnFood() {

    var p Position
    for {
        p = Position{rand.Intn(b.cols), rand.Intn(b.rows)}

        if !b.snake.HasPosition(&p) {
            break
        }
    }
    b.food = p
}

func (b *Board) isOutOfBounds(p *Position) bool {
    return p.x > b.cols-1 || p.y > b.rows-1 || p.x < 0 || p.y < 0
}
