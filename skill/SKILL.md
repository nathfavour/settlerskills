# SKILL.md - Central Token-Efficient AI Context Router

This document serves as the central context router for the `settlerskills` repository, mapping file extensions and operational domains to the strict rule sets defined in the codebase.

---

## 1. Skill Router Rule Ingestion

When modifying, inspecting, or executing operations within this workspace, AI agents must load specific rules dynamically based on the context of their task.

| Target Domain / Operation | Active Rule File | Primary Focus |
| :--- | :--- | :--- |
| **SVM RPC Interaction** | [01_functional_rpc.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/rules/01_functional_rpc.md) | Functional, modern Anza primitives, banned web3.js v1 API limits. |
| **Local Key Custody / State persistence** | [02_local_e2ee_vault.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/rules/02_local_e2ee_vault.md) | Local WebCrypto encryption (AES-GCM-256), PBKDF2 parameters, state isolation. |
| **Transaction Guarding / Simulation** | [03_drain_protection.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/rules/03_drain_protection.md) | Velocity limits, address whitelist assertions, simulated balance boundary checks. |

---

## 2. Dynamic Activation Rules

- **Context Activation Rule:** If the current working workspace buffer focuses on file operations within backend structures or raw data transport loops, the agent must load `rules/01_functional_rpc.md`.
- **Security Isolation Rule:** If code changes touch key storage handlers, configuration variables, or file-system reads, the agent must immediately append `rules/02_local_e2ee_vault.md` to its context window.
