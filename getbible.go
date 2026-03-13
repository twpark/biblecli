package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type gbVerse struct {
	Verse int    `json:"verse"`
	Text  string `json:"text"`
}

type gbResponse struct {
	BookName string          `json:"book_name"`
	Chapter  int             `json:"chapter"`
	Verses   json.RawMessage `json:"verses"`
}

func fetchGetBible(translation string, ref *Reference) ([]VerseResult, error) {
	url := fmt.Sprintf("https://api.getbible.net/v2/%s/%d/%d.json",
		translation, ref.Book.Num, ref.Chapter)

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("not found in API")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data gbResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	verses, err := parseGBVerses(data.Verses)
	if err != nil {
		return nil, err
	}

	sort.Slice(verses, func(i, j int) bool {
		return verses[i].Verse < verses[j].Verse
	})

	if ref.VerseStart == 0 {
		var results []VerseResult
		for _, v := range verses {
			results = append(results, VerseResult{
				Verse: v.Verse,
				Text:  strings.TrimSpace(v.Text),
			})
		}
		return results, nil
	}

	verseMap := make(map[int]string)
	for _, v := range verses {
		verseMap[v.Verse] = strings.TrimSpace(v.Text)
	}

	var results []VerseResult
	for i := ref.VerseStart; i <= ref.VerseEnd; i++ {
		if text, ok := verseMap[i]; ok {
			results = append(results, VerseResult{Verse: i, Text: text})
		} else {
			results = append(results, VerseResult{Verse: i, Missing: true})
		}
	}
	return results, nil
}

func parseGBVerses(raw json.RawMessage) ([]gbVerse, error) {
	// Try array
	var arr []gbVerse
	if err := json.Unmarshal(raw, &arr); err == nil {
		return arr, nil
	}
	// Try object with string keys
	var obj map[string]gbVerse
	if err := json.Unmarshal(raw, &obj); err == nil {
		var out []gbVerse
		for _, v := range obj {
			out = append(out, v)
		}
		return out, nil
	}
	return nil, fmt.Errorf("unexpected verses format")
}
