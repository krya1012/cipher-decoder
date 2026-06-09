# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Hyperskill](https://hyperskill.org/projects/236) educational project implementing the **Diffie-Hellman key exchange protocol** with **Caesar cipher** encryption in Go. You play the role of "Bob" exchanging encrypted messages with "Alice" (the test harness).

## Commands

**Run the program:**
```bash
cd "Cipher Decoder/task" && go run main.go
```

**Run tests (requires Python + hstest):**
```bash
pip install -r requirements.txt
cd "Cipher Decoder/task" && python tests.py
```

**Build:**
```bash
cd "Cipher Decoder/task" && go build main.go
```

## Structure

The project has three sequential stages, each building on the previous. The active code file is always `Cipher Decoder/task/main.go`. Each stage directory (`The protocol/`, `A secret that we share/`, `Encrypt and decrypt/`) contains `task.html` (requirements) and `tests.py` (stage-specific test harness).

## Stage Progression

| Stage | Task | Input → Output |
|-------|------|----------------|
| 1 — The protocol | Parse g and p from Alice | `g is 245 and p is 999` → `g=245 and p=999` |
| 2 — A secret that we share | Diffie-Hellman key exchange using **modular exponentiation** (iterative `c = (c*g) % p`, repeated b times) | Receive g/p, print `OK`, receive A, print `B is [B]`, `A is [A]`, `S is [S]` |
| 3 — Encrypt and decrypt | Caesar cipher using shared secret S as shift | Full conversation: encrypt `Will you marry me?`, receive/decrypt Alice's reply, respond with encrypted `Great!` or `What a pity!` |

## Key Implementation Notes

- **Modular exponentiation**: Do NOT use `math.Pow` — use the iterative algorithm to avoid integer overflow. Start with `c = 1`, loop b times doing `c = (c * g) % p`.
- **Caesar cipher**: Shift only letters cyclically (preserve case); non-letter characters pass through unchanged.
- The shared secret `S = A^b mod p` uses the same iterative modular exponentiation (b iterations of `c = (c * A) % p`).
- Bob's private key `b` must be a random integer where `1 < b < p`.
- Input format for g/p: `"g is [g] and p is [p]"` — parse with `fmt.Sscanf` or string splitting.
- Input format for A: `"A is [A]"`.
