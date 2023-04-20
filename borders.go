package tview

// Borders defines various borders used when primitives are drawn.
// These may be changed to accommodate a different look and feel.
type BorderStyle struct {
	Horizontal  rune
	Vertical    rune
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune

	TopHorizontal  rune
	BottomHorizontal  rune
	LeftVertical    rune
	RightVertical    rune

	LeftT   rune
	RightT  rune
	TopT    rune
	BottomT rune
	Cross   rune

	HorizontalFocus  rune
	VerticalFocus    rune
	TopLeftFocus     rune
	TopRightFocus    rune
	BottomLeftFocus  rune
	BottomRightFocus rune
}
var DefaultBorders = &BorderStyle{
	Horizontal:  BoxDrawingsLightHorizontal,
	Vertical:    BoxDrawingsLightVertical,
	TopLeft:     BoxDrawingsLightDownAndRight,
	TopRight:    BoxDrawingsLightDownAndLeft,
	BottomLeft:  BoxDrawingsLightUpAndRight,
	BottomRight: BoxDrawingsLightUpAndLeft,

	LeftT:   BoxDrawingsLightVerticalAndRight,
	RightT:  BoxDrawingsLightVerticalAndLeft,
	TopT:    BoxDrawingsLightDownAndHorizontal,
	BottomT: BoxDrawingsLightUpAndHorizontal,
	Cross:   BoxDrawingsLightVerticalAndHorizontal,

	HorizontalFocus:  BoxDrawingsDoubleHorizontal,
	VerticalFocus:    BoxDrawingsDoubleVertical,
	TopLeftFocus:     BoxDrawingsDoubleDownAndRight,
	TopRightFocus:    BoxDrawingsDoubleDownAndLeft,
	BottomLeftFocus:  BoxDrawingsDoubleUpAndRight,
	BottomRightFocus: BoxDrawingsDoubleUpAndLeft,
}
var Borders = *DefaultBorders

func ResetBorderStyle() {
  Borders = *DefaultBorders
}
func SetActiveBorderStyle(b *BorderStyle) {
  Borders = *b
}
