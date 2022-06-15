package main

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
func exampleSplitRatio(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return SplitRatio{Ratio: -0.3}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", red)
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", blue)
	})
}

// END EXAMPLE OMIT

// START WIDGET OMIT
type SplitRatio struct {
	// Ratio keeps the current layout.
	// 0 is center, -1 completely to the left, 1 completely to the right.
	Ratio float32
}

func (s SplitRatio) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	proportion := (s.Ratio + 1) / 2
	leftsize := int(proportion * float32(gtx.Constraints.Max.X))

	rightoffset := leftsize
	rightsize := gtx.Constraints.Max.X - rightoffset

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
		left(gtx)
	}

	{
		trans := op.Offset(image.Pt(rightoffset, 0)).Push(gtx.Ops)
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
		right(gtx)
		trans.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

// END WIDGET OMIT
