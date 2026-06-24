# Rule 01: Functional RPC Style Laws

This rule governs all interactions with the Solana Virtual Machine (SVM) within the `settlerskills` workspace.

---

## 1. Allowed Primitives & Frameworks
- You must exclusively use `@solana/kit` and tree-shakeable modern Anza APIs.
- Instantiation of RPC interfaces must use:
  - `createSolanaRpc(endpoint)`
  - `createSolanaRpcSubscriptions(endpoint)`

## 2. Functional Pipelining Invariant
- Transaction and message assembly must use the functional pipeline pattern via `pipe()`.
- Example of correct functional composition:
  ```typescript
  import { pipe, createTransaction, appendTransactionMessageInstruction, setTransactionMessageLifetimeUsingBlockhash } from '@solana/kit';
  
  const transaction = pipe(
      createTransaction({ version: 0 }),
      tx => appendTransactionMessageInstruction(instruction, tx),
      tx => setTransactionMessageLifetimeUsingBlockhash(latestBlockhash, tx)
  );
  ```

## 3. Absolute Execution Bans (Anti-Slop)
- **BANNED RUNTIMES:** Any reference, import, or invocation of the legacy `@solana/web3.js` (v1) framework is strictly banned:
  - NO `new Connection(...)`
  - NO `new Keypair(...)`
  - NO `Transaction.add(...)`
  - NO `SystemProgram.transfer(...)` from the legacy package.
- Violations will trigger system compilation failure.
