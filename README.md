# Encode and Decode Strings

> NeetCode 150 · Arrays & Hashing · Blind 75 — length-prefix framing: write `len#bytes`, read exactly that many bytes back; payload content can never confuse the parser. Go, stdlib only, tested.

## Problem

Design a reversible encoding packing a list of strings into one string — where the strings may contain *anything*, including whatever delimiter you were planning to use. The toy version of a real problem: message framing.

## Approach

Delimiters fail because data can contain them; escaping works but complicates both sides. **Length prefixing** sidesteps the issue entirely: each item is `<decimal length>#<raw bytes>`, and the decoder reads the number, skips the `#`, then consumes exactly that many bytes without inspecting them — `co#de` and empty strings travel untouched. The `#` after the digits isn't a data delimiter, only a terminator for the number itself. This is precisely how real protocols frame messages (HTTP Content-Length, bencode, protobuf's length-delimited fields), which is why interviewers love it. Round-trip tests over adversarial payloads are the natural verification.

**Complexity:** O(total bytes) both ways.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
