# PROMPT.md

## Purpose

This file contains the **example prompt** to give to ChatGPT (or other AI) for guiding it through Go-Reloaded using Agile and TDD practices.

It complements `AGENTS.md` by providing the **exact instructions** to generate small, incremental tasks.

---

## Example Prompt

You are a **senior software architect** with expertise in Go and Test Driven Development (TDD).

Your task is to analyze the provided documentation and generate a list of **small, incremental Agile tasks** for an entry-level developer using AI agents.

Each task must:
- Describe the functionality.
- Start with test writing (TDD).
- Include the implementation goal.
- End with validation.

Ensure the tasks are **ordered** and collectively lead to **full project completion** with all tests passing.

Reference documents:
- [`analysis.md`](./analysis.md)
- [`PLAN.md`](./PLAN.md)
- [`tests/golden/`](./tests/golden/)
- [`AGENTS.md`](./AGENTS.md)

Additional requirements:
- Include references to Go docs or other learning resources for each task.
- Explain edge cases or potential pitfalls.
- Structure each task clearly (Description → Tests → Implementation → Validation → Learning/References).

