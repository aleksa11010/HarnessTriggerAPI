package harness

import (
	"log"
	"os"
	"regexp"
)

func ReadTriggerYaml(f, pipelineIdentifier string, c *Config) string {
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	log.Println(string(data))
	vars := map[string]string{
		"ORG":      c.OrgIdentifier,
		"PROJECT":  c.ProjectIdentifier,
		"PIPELINE": pipelineIdentifier,
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

	log.Println(result)

	return result
}
