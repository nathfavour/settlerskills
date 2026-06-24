# Rule 02: Local E2EE Vault Guardrails

This rule defines the security, key custody, and cryptographic invariants for local state persistence.

---

## 1. Cryptographic Specifications
- **Encryption Primitive:** AES-GCM-256 with an explicit, secure, 12-byte initialization vector (IV) generated uniquely for each write mutation.
- **Key Derivation (KDF):** PBKDF2 with SHA-512.
- **Iteration Count:** Strictly $\ge 600,000$ iterations.
- **Salt Source:** System-level local unique hardware salt (combined with user passcode if necessary).

## 2. Memory Isolation & Plaintext Boundaries
- **No Plaintext Key Persistence:** Do not persist plaintext keys or secrets in logs, terminal outputs, or unmanaged object heaps.
- **Explicit Zeroing (Memory Sanitization):** Keys must be instantiated as `Uint8Array` views or buffers. Immediately after transaction signing or key utilization, the buffer memory must be explicitly overwritten with zeroes.
- **Non-blocking Operations:** All vault reads and writes must be fully asynchronous. Use `fs/promises` or other non-blocking asynchronous file APIs.
