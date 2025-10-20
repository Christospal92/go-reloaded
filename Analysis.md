# Go-Reloaded: Executive Summary

## Goal: Go-reloaded is a tool that reads text, applies predefined transformations, and produces exact output.

**Main Deliverables:**
1. Analysis Document: Describes all transformation rules (number conversions, casing directives, punctuation & quotes, articles) and edge cases.
2. Golden Test Set: Input/output files serving as the source of truth for testing and verification.
3. Skeleton Plan: Blueprint for project architecture (Pipeline stages, folders, testing strategy).

**Key Rules:**
- Numbers in hex/bin are converted to decimal.
- Directives like (up,2), (low,1) apply to previous words.
- Punctuation and quotes are cleaned to ensure correct spacing.
- Articles (a → an) are corrected before words starting with a vowel.

**Architecture:**
- Pipeline recommended for modularity.
- Stages: Tokenize → Numbers → Casing → Articles → Format → Output.
- Folder structure organized for easy maintenance and testing.

**Testing:**
- Golden Test Set covers basic and complex cases.
- Unit tests per stage + Integration tests with golden files.

**Presentation Purpose:**
- To show the logic, edge cases, and strategy for developing the tool before coding.

## Analysis Document — Full Specification

Go-reloaded applies text transformations. It reads text, applies rules, and produces exact expected output.

### 1.1 Number Conversion
Detects numeric literals with a base and converts them to decimal.
Examples:
- 1E (hex) → 30
- 10 (bin) → 2

### 1.2 Casing Directives
Applies capitalization changes to previous words according to directives.
Example: This is so exciting (up, 2) → This is SO EXCITING

### 1.3 Punctuation & Quotes
Cleans spacing around commas, periods, and quotes.
Example: ' I am a optimist , ' → 'I am an optimist,'

### 1.4 Articles (a → an)
Corrects article 'a' to 'an' before words starting with a vowel.
Example: a apple → an apple

## 2) Golden Test Set — Small example
Includes basic and combined input/output examples. Used as the source of truth.

**Example 1 — Numbers**
```
Input: 1E (hex) and 10 (bin)
Expected: 30 and 2
```

**Example 2 — Casing**
```
Input: This is so exciting (up, 2)
Expected: This is SO EXCITING
```

**Example 3 — Punctuation & Quotes**
```
Input: ' I am a optimist , '
Expected: 'I am an optimist,'
```

## 3) Skeleton Plan
Pipeline architecture ("Car Wash") recommended for easy testing and clear organization.

### Stages:
1. Tokenize / Normalize
2. Numbers Stage
3. Casing Stage
4. Articles Stage
5. Format Stage
6. Detokenize / Output

**Explanation:** Each stage has a unique role and the next stage works on the cleaned output from the previous one.

### Recommended Folder Structure:
```
go-reloaded/
 ├─ cmd/
 ├─ pkg/transform/
 ├─ tests/golden/
 └─ README.md
```

- `cmd/` contains another folder with the executable: `go-reloaded/main.go`
- `pkg/transform/` contains the full transformation logic:
  - `tokenize.go` (implements Tokenize / Normalize)
  - `numbers.go` (implements Numbers Stage)
  - `casing.go` (implements Casing Stage)
  - `articles.go` (implements Articles Stage)
  - `format.go` (implements Format Stage)
  - `pipeline.go` (manages execution of all stages)
  - `types.go` (defines Token struct, enums, helpers)

**types.go meanings:**
- Token struct: holds data for each token
- Enums: define fixed token types (Word, Number, Directive) to prevent errors
- Helpers: small functions used across multiple stages

- `tests/golden/` contains all golden test sets: `Input.txt` and `Output.txt`, used as tests to verify correctness.

## Final Deliverables:
1. `analysis.md` — full analysis document.
2. Golden tests — input/output test files.
3. `PLAN.md` — skeleton plan.

