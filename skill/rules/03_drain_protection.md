# Rule 03: Agentic Drain Protection Engine

This rule outlines risk evaluation checks that must occur locally before broadcasting any transaction.

---

## 1. Safety Checklist & Simulation Invariants

Before any signed payload is sent to the network, the executing runtime must:

1. **Velocity Check:** Verify total transaction volume/value within the last 24 hours UTC does not exceed the local safe-boundary velocity limit constant.
2. **Whitelist Validation:** Verify that target destination addresses reside in the local encrypted address whitelist manifest. Reject automatically if not.
3. **RPC Simulation:**
   - Always run `simulateTransaction` on the built payload.
   - Parse simulation responses. If account balance differentials show a decrease/drain exceeding the expected transaction fee plus the calculated output payload size, halt and abort execution immediately.
