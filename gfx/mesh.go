// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package gfx

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	pos    mgl32.Vec3
	uv     mgl32.Vec2
	normal mgl32.Vec3
}

type Mesh struct {
	model    mgl32.Mat4
	vertices []Vertex
	bufferId uint32
	// @TODO: Material
}

// It automatically sets the correct UV and normal based
// only in vertex order.
func (m *Mesh) AddQuad(v1, v2, v3, v4 mgl32.Vec3) {

} 