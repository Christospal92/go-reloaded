# 🧩 Go-Reloaded

A text transformation and formatting tool written in **Go**.  
It reads an input file, applies multiple linguistic and formatting rules,  
and produces an **exactly formatted output** file.

---

## 🚀 Overview

`go-reloaded` performs automated text editing and correction based on  
directives and context rules. It reads an input file (e.g., `sample.txt`)  
and produces a corrected output file (`result.txt`).

### Example Usage

```bash
go run ./cmd/go-reloaded ./sample.txt ./result.txt
```

**Input (`sample.txt`):**
```
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '.
```

**Output (`result.txt`):**
```
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

---

## 🧠 Core Transformation Rules

| Rule | Description | Example |
|------|--------------|----------|
| `(hex)` | Converts preceding hexadecimal to decimal | `1E (hex)` → `30` |
| `(bin)` | Converts preceding binary to decimal | `10 (bin)` → `2` |
| `(up)` | Uppercases the previous word | `go (up)` → `GO` |
| `(low)` | Lowercases the previous word | `GO (low)` → `go` |
| `(cap)` | Capitalizes the previous word | `bridge (cap)` → `Bridge` |
| `(up, N)` / `(low, N)` / `(cap, N)` | Applies case rule to previous *N* words | `so exciting (up, 2)` → `SO EXCITING` |
| Punctuation | Ensures correct spacing around `. , ! ? : ;` | `Hello ,world !` → `Hello, world!` |
| Quotes `'` | Trims spaces inside quotes | `' awesome '` → `'awesome'` |
| Articles | Converts `a` → `an` before vowels or `h` | `a owl` → `an owl` |

---

## 🏗️ Architecture

### Pipeline Flow

```
Input → [Tokenize] → [Numbers] → [Casing] → [Articles] → [Format] → Output
```

Each stage is modular, testable, and independent:
- **tokenize.go** – splits and normalizes tokens  
- **numbers.go** – handles `(hex)` and `(bin)` conversions  
- **casing.go** – processes `(up)`, `(low)`, `(cap)` directives  
- **articles.go** – corrects “a” → “an”  
- **format.go** – cleans punctuation and quote spacing  
- **pipeline.go** – orchestrates all stages  
- **types.go** – defines token structures and enums  

---

## 📂 Project Structure

```
go-reloaded/
 ├─ cmd/
 │   └─ go-reloaded/
 │       └─ main.go          # CLI entry point
 ├─ pkg/
 │   └─ transform/           # All transformation logic
 │        ├─ tokenize.go
 │        ├─ numbers.go
 │        ├─ casing.go
 │        ├─ articles.go
 │        ├─ format.go
 │        ├─ pipeline.go
 │        └─ types.go
 ├─ tests/
 │   └─ golden/              # Input/output test pairs
 ├─ README.md
 ├─ analysis.md
 ├─ prompt.md
 └─ agents.md
```

---

## 🧪 Running Tests

You can manually test the tool using sample files:

```bash
echo "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure." > sample.txt
go run ./cmd/go-reloaded sample.txt result.txt
cat result.txt
```

Expected output:
```
I have to pack 5 outfits. Packed 26 just to be sure.
```

---

## 🧰 Requirements

- **Language:** Go ≥ 1.20  
- **Dependencies:** Only standard library (`fmt`, `os`, `strings`, `strconv`)  
- **Paradigm:** Pipeline architecture + Test-Driven Development (TDD)

---
