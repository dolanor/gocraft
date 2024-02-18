package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dolanor/rip"
	"github.com/go-gl/mathgl/mgl32"
)

type Block struct {
	X, Y, Z float32
	Kind    int
}

func (b *Block) IDFromString(s string) error {
	x, y, z, err := coordsStringToFloat32(s)
	if err != nil {
		return err
	}

	b.X = x
	b.Y = y
	b.Z = z

	return nil
}

func (b *Block) IDString() string {
	return fmt.Sprintf("%fx%fx%f", b.X, b.Y, b.Z)
}

type BlockProvider struct {
	game *Game
}

func (p *BlockProvider) Get(ctx context.Context, id rip.Entity) (*Block, error) {
	return nil, nil
}

func (p *BlockProvider) Create(ctx context.Context, b *Block) (*Block, error) {
	id := b.IDString()
	x, y, z, err := coordsStringToFloat32(id)
	if err != nil {
		return nil, err
	}

	pos := mgl32.Vec3{x, y, z}
	block := NearBlock(pos)

	p.game.world.UpdateBlock(block, b.Kind)
	p.game.dirtyBlock(block)
	go ClientUpdateBlock(block, b.Kind)

	return b, nil
}

func (p *BlockProvider) Delete(ctx context.Context, b rip.Entity) error {
	id := b.IDString()
	x, y, z, err := coordsStringToFloat32(id)
	if err != nil {
		return err
	}

	pos := mgl32.Vec3{x, y, z}
	block := NearBlock(pos)

	p.game.world.UpdateBlock(block, 0)
	p.game.dirtyBlock(block)
	go ClientUpdateBlock(block, 0)
	return nil
}

func (p *BlockProvider) Update(ctx context.Context, block *Block) error {
	return nil
}

func (p *BlockProvider) ListAll(ctx context.Context) ([]*Block, error) {
	return nil, nil
}

func coordsStringToFloat32(id string) (x, y, z float32, err error) {
	coords := strings.Split(id, "x")
	if len(coords) < 3 {
		return 0, 0, 0, errors.New("bad coords format")
	}

	xx, err := strconv.ParseFloat(coords[0], 32)
	if err != nil {
		return 0, 0, 0, err
	}

	yy, err := strconv.ParseFloat(coords[1], 32)
	if err != nil {
		return 0, 0, 0, err
	}

	zz, err := strconv.ParseFloat(coords[2], 32)
	if err != nil {
		return 0, 0, 0, err
	}

	x = float32(xx)
	y = float32(yy)
	z = float32(zz)

	return x, y, z, nil
}
