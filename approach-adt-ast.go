package main

// ADT approach --- vs. TTF approach

import (
	"fmt"
)

type iAdtExpr interface {
	fmt.Stringer
	parseJoinPrev(iAdtExpr) (iAdtExpr, error)
}

type adtExprLit struct {
	Num num
}

type adtExprOp1 struct {
	Op    string
	Right iAdtExpr
}

type adtExprOp2 struct {
	Left  iAdtExpr
	Op    string
	Right iAdtExpr
}

func adtNum(n float64) iAdtExpr {
	return &adtExprLit{Num: num(n)}
}

func (me *adtExprLit) String() string { return me.Num.String() }

func adtOp1(op string, right iAdtExpr) iAdtExpr {
	return &adtExprOp1{Op: op, Right: right}
}

func (me *adtExprOp1) String() string { return fmt.Sprintf("(%s%s)", me.Op, stringer(me.Right)) }

func adtOp2(left iAdtExpr, op string, right iAdtExpr) iAdtExpr {
	return &adtExprOp2{Left: left, Op: op, Right: right}
}

func (me *adtExprOp2) String() string {
	return fmt.Sprintf("(%s %s %s)", stringer(me.Left), me.Op, stringer(me.Right))
}
