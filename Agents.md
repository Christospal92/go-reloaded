# AGENTS.md

## 🧭 Purpose

This file provides **AI coding agents** with precise, structured context for working on the **Go-Reloaded** project.

While `README.md` focuses on humans (goals, setup, contributions),  
`AGENTS.md` focuses on **machine collaborators** — defining:
- how to reason about the project,
- how to break down tasks (Agile & TDD style),
- and how to generate consistent, valid contributions.

---

## 🧩 Project Overview

**Project Name:** go-reloaded  
**Language:** Go (Golang)  
**Paradigm:** Functional modular design via a Pipeline Architecture  
**Testing Philosophy:** Test-Driven Development (TDD)  
**Goal:** Transform input text according to predefined rules (numbers, casing, punctuation, articles) and produce exact output verified against golden test sets.

Reference documents:
- [`analysis.md`](./analysis.md) — detailed rule specifications
- [`tests/golden/`](./tests/golden/) — golden input/output test pairs
- [`PLAN.md`](./PLAN.md) — design skeleton
- [`README.md`](./README.md) — human overview

---

## 🧠 Meta-Prompt for Agile AI Agents

You are a **Senior Software Architect** guiding an **entry-level AI developer**.

### 🔁 Workflow Philosophy
- Work incrementally (small Agile sprints).
- Begin each task with **test creation** (TDD first).
- Implement functionality only after tests are defined.
- Validate outputs using the **golden test set**.
- Document any learning resources or reasoning references used per task.

---

## 🧱 Architecture Summary

**Pipeline (“Car Wash”) Approach**

```
Input → [Tokenize] → [Numbers] → [Casing] → [Articles] → [Format] → Output
```

Each stage is implemented as a small, pure, testable function:
- `tokenize.go` — normalization and token creation  
- `numbers.go` — numeric base detection and conversion  
- `casing.go` — capitalization handling  
- `articles.go` — “a” → “an” correction  
- `format.go` — punctuation and spacing normalization  
- `pipeline.go` — orchestration logic connecting all stages  
- `types.go` — core data types, enums, and helper utilities  

---

## 🧩 Folder Structure (for reference)

```
go-reloaded/
 ├─ cmd/                # main entrypoint for the executable
 ├─ pkg/transform/      # transformation logic (pipeline + stages)
 ├─ tests/golden/       # input/output golden files for verification
 ├─ README.md           # human-readable guide
 ├─ analysis.md         # detailed rules and examples
 ├─ PLAN.md             # system design & architecture plan
 └─ AGENTS.md           # (this file)
```

---

## 🪄 Agile Task Template (Meta-Prompt Pattern)

Each new task should follow this structure:

### **Task [n]: [Short Title]**

#### 🧠 Description
Explain what functionality to add or improve.

#### 🧪 Step 1: Write Tests
- Define test cases (use golden files or new inputs).
- Use Go’s built-in testing package.
- Ensure edge cases are included.

#### 🛠️ Step 2: Implement
- Write minimal code to make the new test pass.
- Keep code modular and pure.
- Follow project conventions from `pkg/transform/`.

#### ✅ Step 3: Validate
- Run all tests with `go test ./...`
- Confirm that outputs match golden references.

#### 📘 Step 4: Learn & Document
- Note any new Go features, regex patterns, or parsing logic.
- Record useful links or concepts (for collective learning).

---

## 🚀 Suggested Task Roadmap

| # | Module | Task | Output |
|---|---------|------|--------|
| 1 | Tokenize | Write unit tests for token splitting and normalization | `tokenize_test.go` |
| 2 | Tokenize | Implement token parsing logic (spaces, punctuation) | `tokenize.go` |
| 3 | Numbers | Write tests for numeric base detection `(hex/bin)` | `numbers_test.go` |
| 4 | Numbers | Implement base conversion using `strconv` | `numbers.go` |
| 5 | Casing | Add tests for directives `(up,n)` `(low,n)` | `casing_test.go` |
| 6 | Casing | Implement casing transformation logic | `casing.go` |
| 7 | Articles | Write tests for `a` → `an` rule | `articles_test.go` |
| 8 | Articles | Implement vowel detection & article correction | `articles.go` |
| 9 | Format | Test punctuation & spacing normalization | `format_test.go` |
| 10 | Format | Implement cleanup logic for quotes, commas, spacing | `format.go` |
| 11 | Pipeline | Integrate all stages and verify flow | `pipeline.go` |
| 12 | Validation | Run integration tests with all golden files | `tests/golden/` |

---

## 🧩 Agent Behavior Guidelines

1. **No code generation before tests.**  
   Always start from tests (TDD discipline).

2. **Cite references** when using Go libraries or regex patterns.

3. **Keep commits atomic:**  
   One task → one feature → one test suite.

4. **Use golden files for validation:**  
   Compare generated output with `*.expected.txt`.

5. **Maintain developer-friendly tone in prompts:**  
   Encourage understanding, not just automation.

---

## 📚 Recommended Learning References

- **Go Documentation:** [https://go.dev/doc/](https://go.dev/doc/)
- **TDD in Go (Tutorial):** [https://quii.gitbook.io/learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests)
- **Clean Architecture (Uncle Bob)**: foundational principles for modular design
- **Meta-Prompting in AI:** structured reasoning for code agents
- **Regex in Go:** [https://pkg.go.dev/regexp](https://pkg.go.dev/regexp)

---

## 🧩 Expected Deliverables from Agents

After following this file, AI agents should produce:
1. `*_test.go` files for each stage.
2. Clean, modular implementation for every transformation.
3. Updated golden test verification results.
4. Summary markdown (`DEVLOG.md`) of what was learned per iteration.

---

## 🧠 Closing Note

This project blends **learning** and **delivery**.  
Each Agile cycle builds both **working code** and **understanding**.  
AI agents are expected not only to code — but to **explain, justify, and evolve** the system design as they proceed.

