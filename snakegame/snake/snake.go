package snake

import "time"

type Snake struct {
    body []Position
    facing Direction
    growing bool
    lastMoved time.Time
    speed float64
}

func NewSnake(body []Position, facing Direction) *Snake {
    return &Snake{
        body: body,
        facing: facing,
        growing: false,
        lastMoved: time.Now(),
        speed: 5.0,
    }
}

func (s *Snake) Head() Position{
    return s.body[len(s.body)-1]
}

func (s *Snake) Move() {
    if time.Since(s.lastMoved).Seconds() < float64(1 / s.speed) {
        return
    }
    s.lastMoved = time.Now()

    head := s.Head()
    newHead := Position{x: head.x, y: head.y}
    switch s.facing {
    case Up:
        newHead.x--
    case Down:
        newHead.x++
    case Left:
        newHead.y--
    case Right:
        newHead.y++
    }

    if s.growing{
            s.body = append(s.body, newHead)
            s.growing = false
        } else {
            s.body = append(s.body[1:], newHead)
    }
}

func (s *Snake) Grow() {
    s.growing = true
    s.speed = s.speed + 0.5

}

func (s *Snake) Turn(d Direction) {
    if d == Up && s.facing != Down {
        s.facing = Up
    } else if d == Down && s.facing != Up {
        s.facing = Down
    } else if d == Left && s.facing != Right {
        s.facing = Left
    } else if d == Right && s.facing != Left {
        s.facing = Right
    }
}

func (s *Snake) HasPosition(p *Position) bool {
    for _, b := range s.body {
        if *p == b {
            return true
        }
    }
    return false
}

func (s *Snake) HitsSelf() bool {
    h := s.Head()
    bodyWithoutHead := s.body[:len(s.body)-1]
    for _, b := range bodyWithoutHead {
        if b == h {
            return true
        }
    }
    return false
}
