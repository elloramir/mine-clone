// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package gfx

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	position mgl32.Vec3
	uv       mgl32.Vec2
	normal   mgl32.Vec3
}

type Mesh struct {
	model    mgl32.Mat4
	vertices []Vertex
	bufferId uint32
	// @TODO: Material
}

// Consider all faces by looking top-down to the cube.
// Each face will be located on the edge of the cube (cube pivot is 0.5, 0.5).
const (
	FaceEast   = iota // ->
	FaceWest          // <-
	FaceNorth         // /\
	FaceSouth         // \/
	FaceTop           // °
	FaceBottom        // _
)
