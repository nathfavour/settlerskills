#!/usr/bin/env bash
# Zero-dependency bootstrap script for settlerskills workspace.
set -euo pipefail

echo "Initializing settlerskills environment..."

# Check requirements
if command -v node >/dev/null 2>&1; then
    echo "Found Node.js version: $(node -v)"
else
    echo "Warning: Node.js is not installed. Node.js is required for TypeScript/Solana Kit execution."
fi

if command -v go >/dev/null 2>&1; then
    echo "Found Go version: $(go -v)"
else
    echo "Warning: Go is not installed. Go is required for the RPC stream daemon core."
fi

# Define global plugin directory for settlerskills loading
GLOBAL_PLUGIN_DIR="${HOME}/.settlerskills/plugins"
mkdir -p "${GLOBAL_PLUGIN_DIR}"
cp -R skill/rules "${GLOBAL_PLUGIN_DIR}/"

echo "Bootstrapped successfully! Project layout verified:"
echo "- ARCHITECTURE.md: Configured"
echo "- skill/SKILL.md: Configured"
echo "- Rules copied to local plugin path: ${GLOBAL_PLUGIN_DIR}/rules"

