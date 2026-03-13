package main

import "strings"

type Book struct {
	Num     int
	Ko      string // Korean abbreviation
	En      string // English display name
	Aliases []string
}

var books = []Book{
	{1, "창", "Genesis", []string{"창", "gen", "genesis"}},
	{2, "출", "Exodus", []string{"출", "exo", "ex", "exodus"}},
	{3, "레", "Leviticus", []string{"레", "lev", "leviticus"}},
	{4, "민", "Numbers", []string{"민", "num", "numbers"}},
	{5, "신", "Deuteronomy", []string{"신", "deut", "dt", "deuteronomy"}},
	{6, "수", "Joshua", []string{"수", "josh", "jos", "joshua"}},
	{7, "삿", "Judges", []string{"삿", "judg", "jdg", "judges"}},
	{8, "룻", "Ruth", []string{"룻", "ruth", "ru"}},
	{9, "삼상", "1 Samuel", []string{"삼상", "1sam", "1sa", "1samuel"}},
	{10, "삼하", "2 Samuel", []string{"삼하", "2sam", "2sa", "2samuel"}},
	{11, "왕상", "1 Kings", []string{"왕상", "1ki", "1kgs", "1kings"}},
	{12, "왕하", "2 Kings", []string{"왕하", "2ki", "2kgs", "2kings"}},
	{13, "대상", "1 Chronicles", []string{"대상", "1chr", "1ch", "1chronicles"}},
	{14, "대하", "2 Chronicles", []string{"대하", "2chr", "2ch", "2chronicles"}},
	{15, "스", "Ezra", []string{"스", "ezr", "ezra"}},
	{16, "느", "Nehemiah", []string{"느", "neh", "nehemiah"}},
	{17, "에", "Esther", []string{"에", "est", "esth", "esther"}},
	{18, "욥", "Job", []string{"욥", "job"}},
	{19, "시", "Psalms", []string{"시", "ps", "psa", "psalm", "psalms"}},
	{20, "잠", "Proverbs", []string{"잠", "prov", "pro", "proverbs"}},
	{21, "전", "Ecclesiastes", []string{"전", "eccl", "ecc", "ecclesiastes"}},
	{22, "아", "Song of Solomon", []string{"아", "song", "sos", "songofsolomon"}},
	{23, "사", "Isaiah", []string{"사", "isa", "isaiah"}},
	{24, "렘", "Jeremiah", []string{"렘", "jer", "jeremiah"}},
	{25, "애", "Lamentations", []string{"애", "lam", "lamentations"}},
	{26, "겔", "Ezekiel", []string{"겔", "ezek", "eze", "ezekiel"}},
	{27, "단", "Daniel", []string{"단", "dan", "daniel"}},
	{28, "호", "Hosea", []string{"호", "hos", "hosea"}},
	{29, "욜", "Joel", []string{"욜", "joel"}},
	{30, "암", "Amos", []string{"암", "amos"}},
	{31, "옵", "Obadiah", []string{"옵", "obad", "ob", "obadiah"}},
	{32, "욘", "Jonah", []string{"욘", "jonah"}},
	{33, "미", "Micah", []string{"미", "mic", "micah"}},
	{34, "나", "Nahum", []string{"나", "nah", "nahum"}},
	{35, "합", "Habakkuk", []string{"합", "hab", "habakkuk"}},
	{36, "습", "Zephaniah", []string{"습", "zeph", "zep", "zephaniah"}},
	{37, "학", "Haggai", []string{"학", "hag", "haggai"}},
	{38, "슥", "Zechariah", []string{"슥", "zech", "zec", "zechariah"}},
	{39, "말", "Malachi", []string{"말", "mal", "malachi"}},
	{40, "마", "Matthew", []string{"마", "matt", "mt", "mat", "matthew"}},
	{41, "막", "Mark", []string{"막", "mark", "mk"}},
	{42, "눅", "Luke", []string{"눅", "luke", "lk"}},
	{43, "요", "John", []string{"요", "john", "jn"}},
	{44, "행", "Acts", []string{"행", "acts", "act"}},
	{45, "롬", "Romans", []string{"롬", "rom", "romans"}},
	{46, "고전", "1 Corinthians", []string{"고전", "1co", "1cor", "1corinthians"}},
	{47, "고후", "2 Corinthians", []string{"고후", "2co", "2cor", "2corinthians"}},
	{48, "갈", "Galatians", []string{"갈", "gal", "galatians"}},
	{49, "엡", "Ephesians", []string{"엡", "eph", "ephesians"}},
	{50, "빌", "Philippians", []string{"빌", "phil", "php", "philippians"}},
	{51, "골", "Colossians", []string{"골", "col", "colossians"}},
	{52, "살전", "1 Thessalonians", []string{"살전", "1th", "1thess", "1thessalonians"}},
	{53, "살후", "2 Thessalonians", []string{"살후", "2th", "2thess", "2thessalonians"}},
	{54, "딤전", "1 Timothy", []string{"딤전", "1ti", "1tim", "1timothy"}},
	{55, "딤후", "2 Timothy", []string{"딤후", "2ti", "2tim", "2timothy"}},
	{56, "딛", "Titus", []string{"딛", "tit", "titus"}},
	{57, "몬", "Philemon", []string{"몬", "phm", "phlm", "philemon"}},
	{58, "히", "Hebrews", []string{"히", "heb", "hebrews"}},
	{59, "약", "James", []string{"약", "jas", "james"}},
	{60, "벧전", "1 Peter", []string{"벧전", "1pe", "1pet", "1peter"}},
	{61, "벧후", "2 Peter", []string{"벧후", "2pe", "2pet", "2peter"}},
	{62, "요일", "1 John", []string{"요일", "1jn", "1john"}},
	{63, "요이", "2 John", []string{"요이", "2jn", "2john"}},
	{64, "요삼", "3 John", []string{"요삼", "3jn", "3john"}},
	{65, "유", "Jude", []string{"유", "jude"}},
	{66, "계", "Revelation", []string{"계", "rev", "revelation"}},
}

var bookByAlias map[string]*Book
var bookByNumber map[int]*Book

func init() {
	bookByAlias = make(map[string]*Book)
	bookByNumber = make(map[int]*Book)
	for i := range books {
		b := &books[i]
		bookByNumber[b.Num] = b
		for _, alias := range b.Aliases {
			bookByAlias[strings.ToLower(alias)] = b
		}
	}
}

func lookupBook(name string) *Book {
	return bookByAlias[strings.ToLower(name)]
}
