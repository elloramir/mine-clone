// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package world

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	view mgl32.Mat4
	proj mgl32.Mat4
}

var upVector = mgl32.Vec3{0, 1, 0}

func NewCamera() *Camera {
	c := &Camera{}

	// @TODO: Get window dimensions from elsewhere
	c.proj = mgl32.Perspective(mgl32.DegToRad(45.0), float32(800)/600, 0.1, 1000.0)
	// @TODO: Camera position and controller.
	c.view = mgl32.LookAtV(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 0, 0}, upVector)

	return c
}

func (c *Camera) Uniforms(program uint32) {
	viewLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
	projLoc := gl.GetUniformLocation(program, gl.Str("proj\x00"))

	gl.UniformMatrix4fv(viewLoc, 1, false, &c.view[0])
	gl.UniformMatrix4fv(projLoc, 1, false, &c.proj[0])
}