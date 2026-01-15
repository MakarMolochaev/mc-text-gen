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
			sym, err := g.font.GetSymbol(string(char))
			if !err {
				//fmt.Printf("Font %s doesn't contains symbol '%s'. It will be not displayed in schematic!\n", g.font.Name, string(char))
				currentW += g.font.LineHeight + g.LetterSpacing
				continue //not panic but notifying about it
			}
			currentW += sym.SizeX + g.LetterSpacing
		}
		w = max(w, currentW)
	}

	return w
}

func (g *SchematicGenerator) getHeight(textLines []string) int {
	return len(textLines) * (g.font.LineHeight + g.LineSpacing)
}

func (g *SchematicGenerator) Generate(textLines []string, projectName string, orientation Orientation) {
	if orientation.index != 0 && orientation.index != 1 {
		orientation.index = 1
	}
	w := g.getWidth(textLines)
	h := g.getHeight(textLines)
	if orientation.index == 1 {
		g.project = schematic.NewProject(projectName, w+g.font.LineHeight, h+g.font.LineHeight, 1)
	} else {
		g.project = schematic.NewProject(projectName, w+g.font.LineHeight, 1, h+g.font.LineHeight)
	}

	posX := 0
	posY := 0
	for _, line := range textLines {
		for _, letter := range strings.ToUpper(line) {
			symbol, err := g.font.GetSymbol(string(letter))
			if !err {
				//fmt.Printf("Font %s doesn't contains symbol '%s'. It will be not displayed in schematic!\n", g.font.Name, string(letter))
				continue
			}
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
	fmt.Printf("Schematic '%s' was exported into 'export/%s'. size: (%d x %d x %d)\n", g.project.RegionName, filename, g.project.Size().X, g.project.Size().Y, g.project.Size().Z)
	defer file.Close()
	g.project.Encode(file)
}
