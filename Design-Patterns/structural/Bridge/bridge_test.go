package Bridge

import "testing"

func TestCircle_Draw(t *testing.T) {
	redCircle := Circle{}
	redCircle.Constructor(100, 100, 10, &RedCircle{})
	redCircle.Draw()

	yellowCircle := Circle{}
	yellowCircle.Constructor(200, 200, 20, &YellowCircle{})
	yellowCircle.Draw()
}
