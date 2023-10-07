// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package main

import (
	_ "embed"
	"fmt"
	"github.com/elloramir/mine-clone/gfx"
	"github.com/elloramir/mine-clone/world"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

const windowWidth = 800
const windowHeight = 600

//go:embed shaders/vert.glsl
var vertShaderSource string

//go:embed shaders/frag.glsl
var fragShaderSource string

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Create window
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "MineClone - v1.0", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	program, err := gfx.CompileProgram(vertShaderSource, fragShaderSource)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteProgram(program)

	// Create world stuff.
	chunk := world.NewChunk(0, 0)
	camera := world.NewCamera()

	// Handle events.
	for !window.ShouldClose() {
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(program)
		camera.Uniforms(program)
		chunk.Mesh.Render()

		// Next frame
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
