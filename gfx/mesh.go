// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package gfx

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	pos    mgl32.Vec3
	uv     mgl32.Vec2
	normal mgl32.Vec3
}

const vertexStride = 4 * (3 + 2 + 3)

type Mesh struct {
	model    mgl32.Mat4
	vertices []Vertex
	indices  []uint32
	vao      uint32
	vbo      uint32
	ebo      uint32
	// @TODO: Material
}

func NewMesh() *Mesh {
	m := &Mesh{}

	gl.GenVertexArrays(1, &m.vao)
	gl.GenBuffers(1, &m.vbo)
	gl.GenBuffers(1, &m.ebo)

	return m
}

func (m *Mesh) AddQuad(v1, v2, v3, v4 mgl32.Vec3) {
	i := uint32(len(m.indices))
	m.vertices = append(m.vertices, Vertex{pos: v1})
	m.vertices = append(m.vertices, Vertex{pos: v2})
	m.vertices = append(m.vertices, Vertex{pos: v3})
	m.vertices = append(m.vertices, Vertex{pos: v4})
	m.indices = append(m.indices, i+0, i+1, i+3, i+1, i+2, i+3)
}

func (m *Mesh) Upload() {
	// @TODO: Warning about that?
	if len(m.vertices) == 0 {
		return
	}

	gl.BindVertexArray(m.vao)

	// Load vertex buffer data.
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.vertices)*4, gl.Ptr(m.vertices), gl.STATIC_DRAW)

	// Load index buffer data.
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.indices)*4, gl.Ptr(m.indices), gl.STATIC_DRAW)

	// Unbind handles.
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (m *Mesh) Unload() {
	gl.DeleteBuffers(1, &m.vbo)
	gl.DeleteBuffers(1, &m.ebo)
	gl.DeleteVertexArrays(1, &m.vao)
}

func (m *Mesh) Clear() {
	// Clear the length/index but still the same capacity.
	// m.verties = m.verties[:0]
	// m.indices = m.indices[:0]

	// Just clear the memory sounds better for now, but
	// we need to add some priority factor to just clear
	// or reset the index.
	m.vertices = nil
	m.indices = nil
}

func (m *Mesh) Render() {
	// @TODO: Warning about that?
	if len(m.vertices) == 0 {
		return
	}

	gl.BindVertexArray(m.vao)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)

	// @TODO: Move to render pipeline?
	// Position vector 3.
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, vertexStride, 3*4)

	// UV vector 2.
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, vertexStride, 2*4)

	// Normals vector 3.
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointerWithOffset(2, 3, gl.FLOAT, false, vertexStride, 3*4)

	// Draw command
	gl.DrawElements(gl.TRIANGLES, int32(len(m.indices)), gl.UNSIGNED_INT, nil)

	// Unbind handles.
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}
