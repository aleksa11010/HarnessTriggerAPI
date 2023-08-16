package harness

import (
	"log"
	"os"
	"regexp"
	"unicode"
)

func ReadTriggerYaml(f, pipelineIdentifier, name, id string, c *Config) string {
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var vars = make(map[string]string)
	log.Println(string(data))
	if name == "" && id == "" {
		vars = map[string]string{
			"ORG":      c.OrgIdentifier,
			"PROJECT":  c.ProjectIdentifier,
			"PIPELINE": pipelineIdentifier,
		}
	} else {
		vars = map[string]string{
			"ORG":        c.OrgIdentifier,
			"PROJECT":    c.ProjectIdentifier,
			"PIPELINE":   pipelineIdentifier,
			"NAME":       name,
			"IDENTIFIER": id,
		}
	}
	re := regexp.MustCompile(`<\+\w+>`)
	result := re.ReplaceAllStringFunc(string(data), func(match string) string {
		varName := match[2 : len(match)-1]
		replace, ok := vars[varName]
		if ok {
			return replace
		}
		return match
	})

	return result
}

func ConvertToCamelCase(input string) string {
	result := ""
	capitalizeNext := true
	for _, r := range input {
		if unicode.IsSpace(r) || r == '_' || r == '-' {
			capitalizeNext = true
			continue
		}
		if capitalizeNext {
			result += string(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result += string(unicode.ToLower(r))
		}
	}
	return result
}
