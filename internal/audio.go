package internal

import (
	"log"
	"regexp"
	"strings"

	"github.com/bogem/id3v2"
)

func ReadTagsFromFile(audiofilePath string) (string, string) {
	tag, err := id3v2.Open(audiofilePath, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	defer tag.Close()

	
	return tag.Title(), tag.Artist()
}

var artistSeparator = regexp.MustCompile(`(?i)\s*(?:;|/|,|&|\band\b|\bfeat\.?\b|\bfeaturing\b|\bft\.?\b|\bwith\b|\bx\b)\s*`)

func ParseArtists(artistTag string) []string{
	artistTag = strings.TrimSpace(artistTag)
	if artistTag == "" {
		return nil
	}

	exceptions := map[string]struct{}{
		"AC/DC": {},
		"K/DA": {},
		"Earth, Wind & Fire": {},
		"Simon & Garfunkel": {},
	}

	if _, ok := exceptions[artistTag]; ok {
		return []string{artistTag}
	}

	parts := artistSeparator.Split(artistTag, -1)

	seen := make(map[string]struct{})
	result := make([]string, 0, len(parts))

	for _, p := range parts {
		p = strings.Join(strings.Fields(strings.TrimSpace(p)), " ")
		if p == "" {
			continue
		}

		key := strings.ToLower(p)
		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, p)
	}

	return result
}