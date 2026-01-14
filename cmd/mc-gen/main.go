package main

import (
	"log"

	"github.com/makarmolochaev/mc-text-gen/internal/fontloader"
)

func main() {
	fontData, err := fontloader.LoadFont("fonts/default-mm.yml")
	if err != nil {
		log.Fatalf("Error in font loading: %v", err)
	}

	/*
		lines, err := utils.ReadAllLines("examples/phrase.txt")
		if err != nil {
			log.Fatalf("Error in file reading: %v", err)
		}
	*/

	generator := fontloader.NewSchematicGenerator(fontData)
	lines := []string{"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЭЮЯ!0123456789.,!?:;", "Пронос овощей на экзамен запрещён!!"}
	generator.Generate(lines)
	generator.Save("test.litematic")
}
