# biblecli

Bible verse lookup from your terminal. Multiple translations, Korean input support, and local translation files.

```
$ biblecli jn3:16

[kjv] John 3:16
For God so loved the world, that he gave his only begotten Son, that whosoever
believeth in him should not perish, but have everlasting life.
```

## Install

```bash
go install github.com/twpark/biblecli@latest
```

Or build from source:

```bash
git clone https://github.com/twpark/biblecli.git
cd biblecli
go build -o biblecli .
cp biblecli ~/.local/bin/  # or anywhere in your PATH
```

## Usage

```bash
# Single verse (default translation: kjv)
biblecli jn3:16
biblecli john3:16

# Verse range
biblecli 1co13:4-7

# Whole chapter
biblecli ps23

# Multiple translations at once
biblecli jn3:16 -t kjv,asv

# Set default translations
biblecli -d kjv,web

# Then just:
biblecli rom8:28
# outputs both kjv and web
```

Korean book abbreviations also work:

```bash
biblecli 요3:16           # John 3:16
biblecli 고전13:4-7       # 1 Corinthians 13:4-7
biblecli 시23             # Psalm 23
```

## Translations

### Remote (GetBible API)

Fetched live from [api.getbible.net](https://api.getbible.net). No API key needed.

| ID | Translation |
|---|---|
| `kjv` | King James Version |
| `asv` | American Standard Version |
| `web` | World English Bible |
| `ylt` | Young's Literal Translation |
| `korean` | Korean Revised Version (1952/1961) |
| `koreankjv` | Korean King James Version |

[Full list of available translations](https://api.getbible.net/v2/translations.json)

### Local translations

Place JSON files in `~/.bible/` for translations not available via API (e.g., copyrighted ones).

```
~/.bible/nkrv.json    # any local translation
~/.bible/custom.json
```

**File format** — flat key-value with Korean book abbreviations as keys:

```json
{
  "gen1:1": "In the beginning...",
  "gen1:2": "And the earth was...",
  "jn3:16": "For God so loved..."
}
```

The CLI tries the GetBible API first. If the translation isn't found there, it falls back to `~/.bible/<translation>.json`.

## Book names

Both English and Korean abbreviations are supported:

| English | Korean | Book |
|---|---|---|
| gen | 창 | Genesis |
| ps | 시 | Psalms |
| isa | 사 | Isaiah |
| matt, mt | 마 | Matthew |
| john, jn | 요 | John |
| rom | 롬 | Romans |
| 1co | 고전 | 1 Corinthians |
| rev | 계 | Revelation |
| ... | ... | all 66 books supported |

## Config

Stored at `~/.bible/config.json`:

```json
{
  "default": ["kjv"]
}
```

Set via CLI: `biblecli -d kjv,web`

If no config exists, defaults to `kjv`.

## License

MIT
