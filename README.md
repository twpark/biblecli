# biblecli

Bible verse lookup from your terminal. Supports Korean and English, multiple translations, and local translation files.

```
$ biblecli 요3:16 -t nkrv,kjv

[nkrv] 요 3:16
하나님이 세상을 이처럼 사랑하사 독생자를 주셨으니 이는 그를 믿는 자마다
멸망하지 않고 영생을 얻게 하려 하심이라

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
biblecli 요3:16
biblecli jn3:16
biblecli john3:16

# Verse range
biblecli 고전13:4-7
biblecli 1co13:4-7

# Whole chapter
biblecli 시23
biblecli ps23

# Multiple translations
biblecli 요3:16 -t nkrv,kjv
biblecli 롬8:28 -t korean,kjv

# Set default translations
biblecli -d nkrv,kjv

# Then just:
biblecli 요3:16
# outputs both nkrv and kjv
```

## Translations

### Remote (GetBible API)

Fetched live from [api.getbible.net](https://api.getbible.net). No API key needed.

| ID | Translation |
|---|---|
| `kjv` | King James Version |
| `korean` | 개역한글 (1952/1961) |
| `koreankjv` | 한글 킹제임스 |
| `asv` | American Standard Version |
| `web` | World English Bible |
| `ylt` | Young's Literal Translation |

[Full list of available translations](https://api.getbible.net/v2/translations.json)

### Local translations

Place JSON files in `~/.bible/` for translations not available via API (e.g., copyrighted ones).

```
~/.bible/nkrv.json    # 개역개정
~/.bible/custom.json  # any custom translation
```

**File format** — flat key-value with Korean book abbreviations:

```json
{
  "창1:1": "태초에 하나님이 천지를 창조하시니라",
  "창1:2": "땅이 혼돈하고 공허하며...",
  "요3:16": "하나님이 세상을 이처럼 사랑하사..."
}
```

The CLI tries the GetBible API first. If the translation isn't found there, it falls back to `~/.bible/<translation>.json`.

## Book names

Both Korean and English abbreviations are supported:

| Korean | English | Book |
|---|---|---|
| 창 | gen | Genesis |
| 출 | exo | Exodus |
| 시 | ps | Psalms |
| 사 | isa | Isaiah |
| 마 | matt, mt | Matthew |
| 요 | john, jn | John |
| 롬 | rom | Romans |
| 고전 | 1co | 1 Corinthians |
| 계 | rev | Revelation |
| ... | ... | all 66 books supported |

## Config

Stored at `~/.bible/config.json`:

```json
{
  "default": ["nkrv", "kjv"]
}
```

Set via CLI: `biblecli -d nkrv,kjv`

If no config exists, defaults to `kjv`.

## License

MIT
