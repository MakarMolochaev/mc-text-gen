package fontloader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type FontData struct {
	Name            string   `yaml:"name"`
	RegisterSupport bool     `yaml:"register-support"`
	Symbols         []Symbol `yaml:"symbols"`
	LineHeight      int      `yaml:"line-height"`
	symbolCache     map[string]*Symbol
}

type Symbol struct {
	Char   string `yaml:"symbol"`
	SizeX  int    `yaml:"sizeX"`
	SizeY  int    `yaml:"sizeY"`
	Bias   int    `yaml:"bias"`
	Scheme string `yaml:"scheme"`
}

func LoadFont(filename string) (*FontData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while opening font file: %w", err)
	}

	var font FontData
	if err := yaml.Unmarshal(data, &font); err != nil {
		return nil, fmt.Errorf("Error in yaml parsing: %w", err)
	}

	font.InitCache()

	return &font, nil
}

func (f *FontData) InitCache() {
	f.symbolCache = make(map[string]*Symbol)
	for i := range f.Symbols {
		symbol := &f.Symbols[i]
		f.symbolCache[symbol.Char] = symbol
	}
}

func (f *FontData) GetSymbol(char string) (*Symbol, bool) {
	symbol, exists := f.symbolCache[char]
	return symbol, exists
}

func (f *FontData) IsSymbolSupported(char string) bool {
	_, exists := f.GetSymbol(char)
	return exists
}
