// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package world

import (
	"github.com/elloramir/mine-clone/gfx"
	"github.com/go-gl/mathgl/mgl32"
	simplex "github.com/ojrac/opensimplex-go"
)

// 256 block types looks enough.
type Block uint8

const (
	BlockAir Block = iota
	BlockDirty
	BlockWater
)

const ChunkSize = 16
const ChunkHeight = 16
const ChunkVolume = (ChunkSize * ChunkSize * ChunkHeight)

type Chunk struct {
	X, Z   float32
	Blocks [ChunkSize][ChunkHeight][ChunkSize]Block
	Mesh   *gfx.Mesh
}

var noise32 simplex.Noise32 = simplex.New32(0x8739018fe1)

// Smooth means how much detail on noise.
const noiseSmooth = 20

func NewChunk(x, z float32) *Chunk {
	c := &Chunk{
		Mesh: gfx.NewMesh(),
		X:    x,
		Z:    z}

	c.generateTerrain()
	c.generateMesh()

	return c
}

func (c *Chunk) generateTerrain() {
	// Absolute chunck position.
	offsetX := int(c.X * ChunkSize)
	offsetZ := int(c.Z * ChunkSize)

	// Generate world blocks from noise.
	for x := 0; x < ChunkSize; x++ {
		for z := 0; z < ChunkSize; z++ {
			noiseX := float32(offsetX+x) / noiseSmooth
			noiseY := float32(offsetZ+z) / noiseSmooth

			// Normalize from [-1, 1] to [0, 1].
			value := (noise32.Eval2(noiseX, noiseY) + 1) / 2
			height := int32(value * ChunkHeight)

			// Grass Block.
			for height >= 0 {
				c.Blocks[x][height][z] = BlockDirty
				height -= 1
			}
		}
	}
}

func (c *Chunk) GetBlock(x, y, z int) Block {
	if y < 0 || y >= ChunkHeight {
		return BlockAir
	}

	// @TODO: Neighbor blocks
	if x < 0 || x >= ChunkSize || z < 0 || z >= ChunkSize {
		return BlockAir
	}

	return c.Blocks[x][y][z]
}

func (c *Chunk) generateMesh() {
	for i := 0; i < ChunkSize; i++ {
		for j := 0; j < ChunkHeight; j++ {
			for k := 0; k < ChunkSize; k++ {
				// Skip empty blocks.
				if c.GetBlock(i, j, k) == BlockAir {
					continue
				}

				// Usefull precast values.
				x := float32(i)
				y := float32(j)
				z := float32(k)

				// Precomputed vertices.
				// We may not use some of these vertices, but still
				// reasonable to have most of them already done.
				//   0 ------ 1
				//  /        /|
				// 3 ------ 2 |
				// |  4     | 5
				// |        |/
				// 7 ------ 6
				v0 := mgl32.Vec3{-0.5 + x, -0.5 + y, -0.5 + z}
				v1 := mgl32.Vec3{+0.5 + x, -0.5 + y, -0.5 + z}
				v2 := mgl32.Vec3{+0.5 + x, -0.5 + y, +0.5 + z}
				v3 := mgl32.Vec3{-0.5 + x, -0.5 + y, +0.5 + z}
				v4 := mgl32.Vec3{-0.5 + x, +0.5 + y, -0.5 + z}
				v5 := mgl32.Vec3{+0.5 + x, +0.5 + y, -0.5 + z}
				v6 := mgl32.Vec3{+0.5 + x, +0.5 + y, +0.5 + z}
				v7 := mgl32.Vec3{-0.5 + x, +0.5 + y, +0.5 + z}

				// Creating block faces
				if c.GetBlock(i, j, k-1) == BlockAir {
					c.Mesh.AddQuad(v1, v0, v4, v5)
				}
				if c.GetBlock(i, j, k+1) == BlockAir {
					c.Mesh.AddQuad(v3, v2, v6, v7)
				}
				if c.GetBlock(i-1, j, k) == BlockAir {
					c.Mesh.AddQuad(v0, v3, v7, v4)
				}
				if c.GetBlock(i+1, j, k) == BlockAir {
					c.Mesh.AddQuad(v2, v1, v5, v6)
				}
				if c.GetBlock(i, j-1, k) == BlockAir {
					c.Mesh.AddQuad(v0, v3, v2, v1)
				}
				if c.GetBlock(i, j+1, k) == BlockAir {
					c.Mesh.AddQuad(v4, v5, v6, v7)
				}
			}
		}
	}

	c.Mesh.Upload()
	c.Mesh.Clear()
}
