package parsers

type Parser interface {
	GenerateConfig(... any) (map[string]any, error)
}