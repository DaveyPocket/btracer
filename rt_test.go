package main

import (
	"image/color"
	"testing"
)

func TestMain(t *testing.T) {
	sphereCenter := point{10, 0, 0}
	testSphere := sphere{sphereCenter, 2, color.RGBA{255, 0, 0, 255}}
	rayDir := vector{point{0, -2, 0}, point{25.5, -2, 0}}
	testRay := ray{point{0, -2, 0}, rayDir}
	t.Log("begin")
	c := testSphere.collision(testRay)
	if c != 10 {
		t.Error(c)
	}
	t.Log("Scale factor:", c)
	t.Log("Ray direction vector:", testRay.direction)
	t.Log("Ray origin point:", testRay.origin)
	normal := vector{testSphere.center, testRay.direction.unit().scale(c).actualPoint(testRay.origin)}
	if normal.p1.x != 10 {
		t.Error(normal)
	}
	if normal.p1.y != -2 {
		t.Error(normal)
	}
	if normal.p1.z != 0 {
		t.Error(normal)
	}
	resultDot := normal.dot(testRay.direction.unit())
	t.Log("Dot product with ray:", resultDot)
	if resultDot != 0 {
		t.Error(resultDot)
	}
	lightRay := vector{point{0, 0, 0}, point{0, 5, 0}}
	result2Dot := normal.unit().dot(lightRay.unit())
	if result2Dot != -1 {
		t.Error(result2Dot)
	}
}
