package fontloader

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tnze/go-mc/level/block"
	"github.com/elvis972602/go-litematica-tools/schematic"
)

type SchematicGenerator struct {
	LineSpacing   int
	LetterSpacing int
	font          *FontData
	project       *schematic.Project
	width         int
	height        int
}

func NewSchematicGenerator(font *FontData) *SchematicGenerator {
	return &SchematicGenerator{
		LineSpacing:   1,
		LetterSpacing: 1,
		font:          font,
	}
}

func (g *SchematicGenerator) getWidth(textLines []string) int {
	w := 0
	for _, text := range textLines {
		currentW := 0
		for _, char := range strings.ToUpper(text) {
			currentW += g.font.symbolCache[string(char)].SizeX + g.LetterSpacing
		}
		w = max(w, currentW)
	}

	return w
}

func (g *SchematicGenerator) getHeight(textLines []string) int {
	return len(textLines) * (g.font.LineHeight + g.LineSpacing)
}

func (g *SchematicGenerator) Generate(textLines []string, orientation Orientation) {
	if orientation.index != 0 && orientation.index != 1 {
		orientation.index = 1
	}
	w := g.getWidth(textLines)
	h := g.getHeight(textLines)
	fmt.Println(w)
	fmt.Println(h)
	if orientation.index == 1 {
		g.project = schematic.NewProject("Test", w+g.font.LineHeight, h+g.font.LineHeight, 1)
	} else {
		g.project = schematic.NewProject("Test", w+g.font.LineHeight, 1, h+g.font.LineHeight)
	}

	posX := 0
	posY := 0
	for _, line := range textLines {
		for _, letter := range strings.ToUpper(line) {
			symbol := g.font.symbolCache[string(letter)]
			i := 0
			j := 0
			for _, ch := range symbol.Scheme {
				if ch == '1' {
					if orientation.index == 1 {
						g.project.SetBlock(posX+i, h-(posY+j+symbol.Bias), 0, block.WhiteConcrete{})
					} else {
						g.project.SetBlock(w-(posX+i), 0, h-(posY+j+symbol.Bias), block.WhiteConcrete{})
					}
				}
				i++
				if i == symbol.SizeX {
					i = 0
					j++
					continue
				}
				if j == symbol.SizeY {
					break
				}
			}
			posX += symbol.SizeX + g.LetterSpacing
		}
		posY += g.font.LineHeight + g.LineSpacing
		posX = 0
	}
}

func (g *SchematicGenerator) Save(filename string) {
	os.Mkdir("export", 0755)
	file, err := os.Create("export/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	g.project.Encode(file)
}
