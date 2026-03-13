package main

import (
	"fmt"
	"os"
	"strings"
)

type VerseResult struct {
	Verse   int
	Text    string
	Missing bool // verse number exists in range but not in this translation
}

// Known GetBible API translations (try API first for these)
var knownAPI = map[string]bool{
	"kjv": true, "kjva": true, "akjv": true,
	"korean": true, "koreankjv": true,
	"asv": true, "web": true, "ylt": true,
	"basicenglish": true, "douayrheims": true,
	"weymouth": true, "wb": true, "tyndale": true, "wycliffe": true,
}

func isKoreanTranslation(name string) bool {
	n := strings.ToLower(name)
	return strings.Contains(n, "korean") || strings.Contains(n, "krv") || strings.Contains(n, "nkrv")
}

func fetchVerses(translation string, ref *Reference) ([]VerseResult, error) {
	// Try API first for known translations
	if knownAPI[translation] {
		results, err := fetchGetBible(translation, ref)
		if err == nil {
			return results, nil
		}
	}

	// Try local file
	results, err := fetchLocal(translation, ref)
	if err == nil {
		return results, nil
	}

	// If not known API, try API as fallback
	if !knownAPI[translation] {
		results, err := fetchGetBible(translation, ref)
		if err == nil {
			return results, nil
		}
	}

	return nil, fmt.Errorf("translation '%s' not found (not in API, no local file at ~/.bible/%s.json)", translation, translation)
}

func formatHeader(ref *Reference, translation string) string {
	var bookName string
	if isKoreanTranslation(translation) {
		bookName = ref.Book.Ko
	} else {
		bookName = ref.Book.En
	}

	if ref.VerseStart == 0 {
		return fmt.Sprintf("%s %d", bookName, ref.Chapter)
	}
	if ref.VerseStart == ref.VerseEnd {
		return fmt.Sprintf("%s %d:%d", bookName, ref.Chapter, ref.VerseStart)
	}
	return fmt.Sprintf("%s %d:%d-%d", bookName, ref.Chapter, ref.VerseStart, ref.VerseEnd)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	var refStr string
	var transFlag string
	var defaultFlag string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-t":
			i++
			if i < len(args) {
				transFlag = args[i]
			}
		case "-d":
			i++
			if i < len(args) {
				defaultFlag = args[i]
			}
		case "-h", "--help":
			printUsage()
			return
		default:
			refStr = args[i]
		}
	}

	// Set default translations
	if defaultFlag != "" {
		parts := splitTrim(defaultFlag)
		cfg := &Config{Default: parts}
		if err := saveConfig(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Default set to: %s\n", strings.Join(parts, ", "))
		return
	}

	if refStr == "" {
		printUsage()
		os.Exit(1)
	}

	ref, err := parseReference(refStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Determine translations
	var transList []string
	if transFlag != "" {
		transList = splitTrim(transFlag)
	} else {
		cfg := loadConfig()
		transList = cfg.Default
	}

	multiVerse := ref.VerseStart == 0 || ref.VerseStart != ref.VerseEnd

	for i, trans := range transList {
		if i > 0 {
			fmt.Println()
		}

		results, err := fetchVerses(trans, ref)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[%s] %v\n", trans, err)
			continue
		}

		fmt.Printf("[%s] %s\n", trans, formatHeader(ref, trans))
		for _, v := range results {
			if v.Missing {
				fmt.Printf("%d. (verse not available in this translation)\n", v.Verse)
			} else if multiVerse {
				fmt.Printf("%d. %s\n", v.Verse, v.Text)
			} else {
				fmt.Println(v.Text)
			}
		}
	}
}

func splitTrim(s string) []string {
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `biblecli — Bible verse lookup from terminal

Usage:
  biblecli <reference> [-t translations]
  biblecli -d <default-translations>

Examples:
  biblecli 요3:16              Look up John 3:16 (default translations)
  biblecli jn3:16 -t nkrv,kjv  John 3:16 in NKRV and KJV
  biblecli 1co13:4-7           1 Corinthians 13:4-7
  biblecli 요3                  John chapter 3 (full)
  biblecli -d nkrv,kjv         Set default translations

Translations:
  API: kjv, korean, koreankjv, asv, web, ylt, ...
  Local: ~/.bible/<name>.json (e.g., ~/.bible/nkrv.json)
`)
}
