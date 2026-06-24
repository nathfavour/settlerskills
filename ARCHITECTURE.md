# ARCHITECTURE.md - System Architecture and Engineering Specification

> **SYSTEM INGESTION DIRECTIVE FOR AI AGENTS:** This document defines the unalterable system architecture, data flow state machines, cryptographic boundaries, and engineering invariants for the `settlerskills` repository. Every code generation cycle, pull request, or architectural modification must comply with the mathematical and structural constraints detailed herein. Zero deviation is permitted.

---

## 1. System Vision & Sovereign Foundations

The `settlerskills` ecosystem is an open-source, local-first, zero-dependency configuration and progressive rule engine. It transforms AI runtime agents (e.g., Claude Code, Cursor, Codex) into hyper-precise, deterministic operators capable of building non-custodial payment rails, streaming RPC daemons, and secure local cryptographic storage wrappers on the Solana virtual machine (SVM).

### Core Architectural Axioms
1. **Absolute Autonomy:** Zero reliance on third-party cloud infrastructure, SaaS key managers, or hosted indexers. All compute, indexing, and state mutations occur within the local host environment.
2. **Minimal Resource Allocation:** High-velocity loops must maintain deterministic execution characteristics, memory boundaries, and clear lifecycle definitions.
3. **Immutability of Key Assets:** Cryptographic seed generation and transaction serialization paths are physically isolated from general runtime memory scopes.

---

## 2. Component Pipeline & Data Flow

The system architecture is structured into four distinct execution layers. Each layer operates as an independent state machine with strict boundary isolation:


```

[Local E2EE Vault] ────(Ephemeral Key)────► [Transaction Serializer]
│
(Signed Transaction)
▼
[RPC Event Stream] ◄──(Subscription Loop)─── [Functional RPC Rail]

```

### 2.1 Component Matrix

| Layer | Primary Technology | Responsibility | Memory Boundary Constraints |
| :--- | :--- | :--- | :--- |
| **Local E2EE Vault** | TypeScript / WebCrypto API | Secure storage, key derivation, disk-level state persistence. | Zero persistent plaintext keys in heap space. Ephemeral execution only. |
| **Transaction Serializer** | Functional `@solana/kit` | Direct binary instruction packing, transaction pipeline construction via `pipe()`. | No intermediate buffer copying; zero-allocation memory mutations. |
| **Functional RPC Rail** | Modern Anza Primitive | Stateless transaction broadcast, retry orchestration, blockhash acquisition. | Purely functional, immutable input-to-output transitions. |
| **RPC Event Stream** | Go (Daemon Core) / WebSockets | Low-overhead confirmation loops, block parsing, webhook trigger routing. | Fixed ring-buffer allocations to completely eliminate garbage collection spikes. |

---

## 3. Technical Specifications & Invariants

### 3.1 The Cryptographic Vault Layer (`rules/02_local_e2ee_vault.md`)
The vault layer handles file-system state management without cloud dependency. 
* **Encryption Primitive:** Authenticated encryption via AES-GCM-256 with an explicit 12-byte cryptographically secure pseudorandom initialization vector (IV) per mutation.
* **Key Derivation:** Master keying derived locally through PBKDF2 with SHA-512, utilizing a minimum of 600,000 iterations and a local system-level unique hardware salt.
* **State Isolation:** Plaintext private key material must never cross into logs, shell outputs, or unmanaged object heaps. Keys are materialized as `Uint8Array` allocations and explicitly overwritten with zeroes immediately after transaction signature execution.

### 3.2 The Functional Transaction Pipeline (`rules/01_functional_rpc.md`)
All interaction with the Solana blockchain must drop the legacy monolithic programming model in favor of the modern, tree-shakeable 2026 Anza architecture.
* **Instantiation Primitive:** Utilize stateless endpoints derived solely via `createSolanaRpc()` and `createSolanaRpcSubscriptions()`.
* **Execution Pipelining:** Code generation must force the use of functional pipelines to assemble messages:
  ```typescript
  // Canonical Target Construction Pattern
  const transaction = pipe(
      createTransaction({ version: 0 }),
      tx => appendTransactionMessageInstruction(instruction, tx),
      tx => setTransactionMessageLifetimeUsingBlockhash(latestBlockhash, tx)
  );
  ```

* **Banned Runtimes:** Any invocation of the old `@solana/web3.js` v1 framework (e.g., `new Connection()`, `new Keypair()`, `Transaction.add()`) will instantly trigger a system compilation failure.

### 3.3 Agentic Drain Protection Engine (`rules/03_drain_protection.md`)

To permit automated execution by autonomous cron or event loops, a strict, local risk-evaluation state engine is interposed before the broadcast layer.

* **Invariant Invariants:**
1. **Velocity Cap:** Maximum volume allowances capped at a strictly defined local constant per 24-hour UTC window.
2. **Destination Whitelisting:** Transactions targeting addresses outside an encrypted local configuration lookup manifest are rejected automatically at the compilation level.
3. **Simulated State Assertion:** Prior to raw broadcast, every transaction payload must undergo local RPC simulation (`simulateTransaction`). If the simulation reports modified account balances deviating from predefined boundary differentials, execution is halted immediately.

---

## 4. Repository Layout & Lazy-Loading Topography

The workspace structure is optimized for progressive context injection. It ensures LLM agents load context dynamically, preventing token saturation during iterative development loops.

```
settlerskills/
├── README.md                      # Vision, installation, quick-start matrix, manual usage
├── LICENSE                        # MIT License for unencumbered upstream ingestion
├── ARCHITECTURE.md                # Global master engineering specification & boundary layer
├── install.sh                     # Zero-dependency bootstrap script copying rules to ~/.settlerskills/plugins
└── skill/
    ├── SKILL.md                   # Central context router and entry point for LLM execution
    ├── rules/
    │   ├── 01_functional_rpc.md   # Functional code style laws for SVM interaction
    │   ├── 02_local_e2ee_vault.md # Technical guardrails for local key encapsulation & memory safety
    │   └── 03_drain_protection.md # Risk verification logic and execution constraints
    └── templates/
        ├── loop.go                # Reference Go implementation for low-allocation RPC stream parsing
        └── sign.ts                # Reference modern TS instruction-packing pipeline
```

### 4.1 Detailed Codebase Directory Specifications

1. `/README.md`
   - **Type**: Markdown Documentation File
   - **Path**: [README.md](file:///home/nathfavour/code/nathfavour/settlerskills/README.md)
   - **Description**: Details sovereign payment infrastructure philosophy, architectural layers, and manual setup guides.
2. `/LICENSE`
   - **Type**: Plain Text License File
   - **Path**: [LICENSE](file:///home/nathfavour/code/nathfavour/settlerskills/LICENSE)
   - **Description**: MIT License permitting unencumbered upstream integration.
3. `/ARCHITECTURE.md`
   - **Type**: Markdown System Specification
   - **Path**: [ARCHITECTURE.md](file:///home/nathfavour/code/nathfavour/settlerskills/ARCHITECTURE.md)
   - **Description**: Defines unalterable system boundaries, components, memory constraints, and quality gates.
4. `/install.sh`
   - **Type**: Bash Executable Script
   - **Path**: [install.sh](file:///home/nathfavour/code/nathfavour/settlerskills/install.sh)
   - **Description**: Setups dependencies and copies local skill rule files into the target plugin loading path (`~/.settlerskills/plugins/rules`).
5. `/skill/SKILL.md`
   - **Type**: Markdown Skill Router
   - **Path**: [SKILL.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/SKILL.md)
   - **Description**: Entry point router utilizing mapping tables to direct agent focus to subset rules without bloated files.
6. `/skill/rules/01_functional_rpc.md`
   - **Type**: Markdown Rule Document
   - **Path**: [01_functional_rpc.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/rules/01_functional_rpc.md)
   - **Description**: Enforces functional `@solana/kit` usage and prohibits legacy `@solana/web3.js` API wrappers.
7. `/skill/rules/02_local_e2ee_vault.md`
   - **Type**: Markdown Rule Document
   - **Path**: [02_local_e2ee_vault.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/rules/02_local_e2ee_vault.md)
   - **Description**: Enforces WebCrypto AES-GCM-256 specifications, PBKDF2 parameters, and memory zeroing invariants.
8. `/skill/rules/03_drain_protection.md`
   - **Type**: Markdown Rule Document
   - **Path**: [03_drain_protection.md](file:///home/nathfavour/code/nathfavour/settlerskills/skill/rules/03_drain_protection.md)
   - **Description**: Outlines velocity, whitelists, and local transaction simulation balance differential checks.
9. `/skill/templates/loop.go`
   - **Type**: Go Code Template
   - **Path**: [loop.go](file:///home/nathfavour/code/nathfavour/settlerskills/skill/templates/loop.go)
   - **Description**: Pristine reference loop mimicking a low-allocation websocket connection framework.
10. `/skill/templates/sign.ts`
    - **Type**: TypeScript Code Template
    - **Path**: [sign.ts](file:///home/nathfavour/code/nathfavour/settlerskills/skill/templates/sign.ts)
    - **Description**: Reference modern transactional compilation and memory sanitization implementation.

### 4.2 SKILL.md Routing Matrix

The central routing mechanism maps file extensions and incoming automated agent prompts directly to their respective rule bases.

* **Context Activation Rule:** If the current working workspace buffer focuses on file operations within backend structures or raw data transport loops, the agent must load `rules/01_functional_rpc.md`.
* **Security Isolation Rule:** If code changes touch key storage handlers, configuration variables, or file-system reads, the agent must immediately append `rules/02_local_e2ee_vault.md` to its context window.


---

## 5. Absolute Execution Bans (Anti-Slop Directives)

To guarantee software quality matches high-end systems engineering criteria, the following implementation behaviors are banned from the codebase:

1. **NO Implicit Typing:** General `any` expressions or implicit runtime structural inferences are strictly forbidden. All state modifications must handle errors explicitly.
2. **NO Cloud Leakage:** Code blocks must never use external telemetry services, tracking payloads, remote logger endpoints, or non-local secrets management APIs.
3. **NO Block-Stalling Sync Calls:** All local file-system or network operations must use non-blocking, asynchronous, event-driven interfaces to prevent daemon pipeline lockup.
4. **NO Monolithic Bloat:** Avoid importing heavy wrapper frameworks or secondary kitchen-sink utilities. Use native platform primitives and clean, localized utility blocks.

---

## 6. Meta-Architecture Integrity & Quality Control Enforcement

### 6.1 Meta-Boundary Declaration
`ARCHITECTURE.md` functions exclusively as the top-level technical specification and invariant map for the entire `settlerskills` workspace. 

* **The Distinction Axiom:** This file is NOT an agentic skill, tool, configuration profile, or executable plugin within the project topology. It is the structural boundary layer that governs how the container itself is arranged. 
* **The Containment Axiom:** The skills inside `settlerskills` are isolated modules living exclusively inside the `/skill` directory. Agents must never interpret this structural specification as an actionable codebase behavior or a functional router node.

### 6.2 Quality Control & Anti-Slop Compilation Gates
Before any cycle commit, automated pull request generation, or workspace write operation, the AI agent must validate its target files against this meta-specification. A file is classified as non-compliant "AI slop" and will fail the architectural compilation gate if it contains any of the following metadata characteristics:

1. **Conversational Scaffolding:** No raw LLM meta-instructions, prompt design drafts, step-by-step assistant setups, chat history transcripts, or conversational placeholders (e.g., `// TODO: implement later`, `// Ask user for parameter`).
2. **Context Bleed:** The internal rule files inside `/skill/rules/` must maintain absolute functional isolation. No payment rail execution guidelines (`01_functional_rpc.md`) may bleed into cryptographic key storage instructions (`02_local_e2ee_vault.md`).
3. **Instruction Integrity Check:** Run an explicit structural linting validation pass to confirm that `SKILL.md` contains only clean context routing maps and workflow triggers, devoid of inline implementation templates or block commentary.

