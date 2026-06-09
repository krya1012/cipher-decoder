# Cipher Decoder

A Go implementation of the [Diffie-Hellman key exchange protocol](https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange) with Caesar cipher encryption. Built as [Hyperskill project #236](https://hyperskill.org/projects/236).

The program plays the role of **Bob**, exchanging encrypted messages with **Alice** (the test harness) over a simulated unsecured channel.

## Stages

| # | Name | Description |
|---|------|-------------|
| 1 | The Protocol | Parse public parameters `g` and `p` from Alice's input |
| 2 | A Secret That We Share | Perform Diffie-Hellman key exchange using modular exponentiation to establish shared secret `S` |
| 3 | Encrypt and Decrypt | Exchange messages encrypted with Caesar cipher using `S` as the shift |

## Running

```bash
go run "Cipher Decoder/task/main.go"
```

**Quick test (stage 1):**
```bash
echo "g is 245 and p is 999" | go run "Cipher Decoder/task/main.go"
# g=245 and p=999
```

**Run stage tests** (requires Python + [hstest](https://github.com/hyperskill/hs-test-python)):
```bash
pip install -r requirements.txt
cd "Cipher Decoder/task" && python tests.py
```
