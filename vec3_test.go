package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPSubtract(t *testing.T) {
	testV := Vec3{3, 2, 3}
	testV.SubInPlace(Vec3{2, 0, 0})
	assert.Equal(t, Vec3{1, 2, 3}, testV)
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, Vec3{-2, 3, -10}, Vec3{0, 0, 0}.Sub(Vec3{2, -3, 10}))
}

func TestPAdd(t *testing.T) {
	x := Vec3{1, 2, 3}
	x.AddInPlace(Vec3{1.1, 2.2, 3.3})
	assert.Equal(t, Vec3{2.1, 4.2, 6.3}, x)
}

func TestAdd(t *testing.T) {
	assert.Equal(t, Vec3{2.1, 4.2, 6.3}, Vec3{1, 2, 3}.Add(Vec3{1.1, 2.2, 3.3}))
}

func TestPMultiply(t *testing.T) {
	x := Vec3{1, 2, 3}
	x.MulInPlace(1.5)
	assert.Equal(t, Vec3{1.5, 3, 4.5}, x)
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, Vec3{1.5, 3, 4.5}, Vec3{1, 2, 3}.Mul(1.5))
}

func TestPDivide(t *testing.T) {
	x := Vec3{10, 5, 4}
	x.DivInPlace(0.5)
	assert.Equal(t, Vec3{20, 10, 8}, x)
}

func TestDivide(t *testing.T) {
	assert.Equal(t, Vec3{20, 10, 8}, Vec3{10, 5, 4}.Divide(0.5))
}

func TestLengthSquared(t *testing.T) {
	assert.Equal(t, 14.0, Vec3{1, 2, 3}.LenSquared())
}

func TestLength(t *testing.T) {
	assert.Equal(t, math.Sqrt(14), Vec3{1, 2, 3}.Len())
}

func TestUnitVector(t *testing.T) {
	assert.Equal(t, Vec3{1, 2, 3}.Divide(Vec3{1, 2, 3}.Len()), Vec3{1, 2, 3}.UnitVector())
}
