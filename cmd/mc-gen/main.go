package main

import (
	"log"

	"github.com/makarmolochaev/mc-text-gen/internal/fontloader"
	"github.com/makarmolochaev/mc-text-gen/internal/utils"
)

func main() {
	fontData, err := fontloader.LoadFont("fonts/default-mm.yml")
	if err != nil {
		log.Fatalf("Error in font loading: %v", err)
	}

	lines, err := utils.ReadAllLines("examples/phrase.txt")
	if err != nil {
		log.Fatalf("Error in file reading: %v", err)
	}

	generator := fontloader.NewSchematicGenerator(fontData)
	//lines := []string{"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ!0123456789.,!?:;", "Пронос овощей на экзамен запрещён!!"}
	generator.Generate(lines, "mutanti", fontloader.Horizontal) //fontloader.Horizontal / Vertical
	generator.Save("mutanti.litematic")
}
