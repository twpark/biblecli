package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func fetchLocal(translation string, ref *Reference) ([]VerseResult, error) {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".bible", translation+".json")

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("local file not found: ~/.bible/%s.json", translation)
	}

	var data map[string]string
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("invalid JSON in %s: %v", path, err)
	}

	ko := ref.Book.Ko

	if ref.VerseStart == 0 {
		// Whole chapter
		prefix := fmt.Sprintf("%s%d:", ko, ref.Chapter)
		var results []VerseResult
		for key, text := range data {
			if strings.HasPrefix(key, prefix) {
				verseStr := strings.TrimPrefix(key, prefix)
				verse, err := strconv.Atoi(verseStr)
				if err != nil {
					continue
				}
				results = append(results, VerseResult{
					Verse: verse,
					Text:  strings.TrimSpace(text),
				})
			}
		}
		sort.Slice(results, func(i, j int) bool {
			return results[i].Verse < results[j].Verse
		})
		if len(results) == 0 {
			return nil, fmt.Errorf("chapter not found: %s %d", ko, ref.Chapter)
		}
		return results, nil
	}

	// Specific verses
	var results []VerseResult
	for i := ref.VerseStart; i <= ref.VerseEnd; i++ {
		key := fmt.Sprintf("%s%d:%d", ko, ref.Chapter, i)
		if text, ok := data[key]; ok {
			results = append(results, VerseResult{
				Verse: i,
				Text:  strings.TrimSpace(text),
			})
		} else {
			results = append(results, VerseResult{Verse: i, Missing: true})
		}
	}
	return results, nil
}
