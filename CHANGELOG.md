# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [1.2.0] - 2026-06-09

### Added
- Stage 3 (Encrypt and Decrypt): Caesar cipher using shared secret S as shift. Bob sends encrypted "Will you marry me?", decrypts Alice's reply, and responds with encrypted "Great!" or "What a pity!" accordingly.

## [1.1.0] - 2026-06-09

### Added
- Stage 2 (A Secret That We Share): Diffie-Hellman key exchange using iterative modular exponentiation. Bob picks random private key `b`, computes and sends `B`, receives Alice's `A`, and derives shared secret `S`.

## [1.0.0] - 2026-06-09

### Added
- Stage 1 (The Protocol): parse public Diffie-Hellman parameters `g` and `p` from input `"g is X and p is Y"` and echo them as `g=X and p=Y`.
