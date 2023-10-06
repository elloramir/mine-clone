// Copyright 2023 Elloramir. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package world

import (
	//"github.com/go-gl/mathgl/mgl32"
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
}

var noise32 simplex.Noise32 = simplex.New32(0x8739018fe1)

// Smooth means how much detail on noise.
const noiseSmooth = 20

func NewChunk(x, z float32) *Chunk {
	c := &Chunk{
		X: x,
		Z: z}

	c.generateTerrain()

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
