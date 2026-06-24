import {
    pipe,
    createTransaction,
    appendTransactionMessageInstruction,
    setTransactionMessageLifetimeUsingBlockhash,
    createSolanaRpc,
    type Blockhash,
    type Instruction,
} from '@solana/kit';

/**
 * Builds a modern, stateless Solana transaction using tree-shakeable functional pipelines.
 * Complies with skill/rules/01_functional_rpc.md rules.
 * 
 * @param rpcEndpoint The Solana RPC endpoint.
 * @param instruction The instruction to append to the transaction.
 * @param latestBlockhash The blockhash information.
 */
export async function buildAndSerializeTransaction(
    rpcEndpoint: string,
    instruction: Instruction,
    latestBlockhash: { blockhash: Blockhash; lastValidBlockHeight: bigint }
) {
    const rpc = createSolanaRpc(rpcEndpoint);

    // Modern 2026 Anza pipeline composition
    const transaction = pipe(
        createTransaction({ version: 0 }),
        tx => appendTransactionMessageInstruction(instruction, tx),
        tx => setTransactionMessageLifetimeUsingBlockhash(latestBlockhash, tx)
    );

    return transaction;
}

/**
 * Safely clears secret buffers in memory after use.
 * Complies with skill/rules/02_local_e2ee_vault.md requirements.
 */
export function secureZeroMemory(buffer: Uint8Array): void {
    buffer.fill(0);
}
