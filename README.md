# settlerskills

An open-source, local-first, zero-dependency configuration and progressive rule engine for AI runtime agents building non-custodial payment rails, streaming RPC daemons, and secure local cryptographic storage wrappers on the Solana virtual machine (SVM).

## Sovereign Payment Infrastructure Philosophy

In autonomous agent environments, relying on third-party cloud key management or external SaaS indexers introduces critical security risks and operational dependencies. `settlerskills` enforces:
1. **Absolute Autonomy:** All keys, transactions, and state operations remain strictly within the local host environment.
2. **Minimal Resource Overhead:** Uses low-allocation streaming architectures (e.g., zero-GC Go ring buffers) to eliminate compute spikes.
3. **Hardened Security boundaries:** Strict local transaction simulations, encrypted whitelists, and ephemeral memory sanitization.

## Architectural Layers

1. **Local E2EE Vault**: TypeScript/WebCrypto wrapper utilizing AES-GCM-256 for local persistence and PBKDF2 for key derivation.
2. **Transaction Serializer**: Tree-shakeable functional `@solana/kit` pipelines compiling and packing transactions.
3. **Functional RPC Rail**: Modern stateless Solana RPC interaction wrappers.
4. **RPC Event Stream**: Ultra-low-overhead Go-based event listeners and confirmation loops.

## Repository Topography

See [ARCHITECTURE.md](file:///home/nathfavour/code/nathfavour/settlerskills/ARCHITECTURE.md) for full engineering specifications.

- `install.sh`: Zero-dependency environment bootstrap script.
- `skill/SKILL.md`: Central Token-Efficient AI Context Router.
- `skill/rules/`: Code style laws and local security boundaries.
- `skill/templates/`: Reference implementations for transactions and streaming RPC.

## Manual Usage & Setup

### Bootstrap Setup
Initialize and verify the project directory on your local machine:
```bash
chmod +x install.sh
./install.sh
```

### Context Registration
For AI runtime agents (such as Claude Code or Cursor), feed `/skill/SKILL.md` as the entry-point context router. The agent will dynamically lazy-load rules depending on your operational scope.

