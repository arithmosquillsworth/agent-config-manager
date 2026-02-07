# Agent Config Manager

🔧 Unified configuration management for all Agent Security Stack tools.

## Overview

Centralizes configuration across all agent tools:
- Wallet settings
- API keys
- Security policies
- Monitoring endpoints
- Tool-specific exports

## Quick Start

```bash
# Initialize configuration
acm init

# View current config
acm show

# Set API keys
acm set api_keys.etherscan YOUR_KEY
acm set api_keys.basescan YOUR_KEY

# Validate configuration
acm validate

# Export tool-specific configs
acm export
```

## Configuration Structure

```json
{
  "version": "0.1.0",
  "agent": {
    "name": "Arithmos",
    "id": "arithmos-quillsworth",
    "erc8004_id": 1941,
    "website": "https://arithmos.dev",
    "github": "https://github.com/arithmosquillsworth"
  },
  "wallet": {
    "address": "0x120e011fB8a12bfcB61e5c1d751C26A5D33Aae91",
    "networks": ["ethereum", "base"],
    "daily_limit": 0.5,
    "alert_threshold": 0.1
  },
  "security": {
    "firewall_enabled": true,
    "honeypot_enabled": true,
    "prompt_guard_enabled": true,
    "simulator_enabled": true,
    "whitelisted_addresses": [],
    "blacklisted_addresses": []
  },
  "api_keys": {
    "etherscan": "...",
    "basescan": "...",
    "openai": "...",
    "anthropic": "...",
    "discord": "..."
  },
  "monitoring": {
    "dashboard_enabled": true,
    "dashboard_port": 8080,
    "webhook_url": "...",
    "check_interval_minutes": 5
  }
}
```

## Commands

| Command | Description |
|---------|-------------|
| `acm init` | Create initial configuration |
| `acm show` | Display current configuration |
| `acm get <key>` | Get specific value |
| `acm set <key> <value>` | Set specific value |
| `acm validate` | Validate configuration |
| `acm export` | Export tool-specific configs |

## Setting Values

```bash
# Set API keys
acm set api_keys.etherscan YOUR_ETHERSCAN_KEY
acm set api_keys.basescan YOUR_BASESCAN_KEY

# Set wallet limits
acm set wallet.daily_limit 1.0
acm set wallet.alert_threshold 0.5

# Set monitoring
acm set monitoring.webhook_url https://discord.com/api/webhooks/...
acm set monitoring.check_interval 10
```

## Getting Values

```bash
# Get specific values
acm get wallet.address
acm get wallet.daily_limit
acm get security.firewall_enabled
```

## Validation

```bash
$ acm validate
🔍 Validating configuration...

✅ Configuration is valid!

# Or with issues:
⚠️  Etherscan API key not set (needed for monitoring)
⚠️  Basescan API key not set (needed for monitoring)
Found 2 issue(s)
```

## Export

Export generates tool-specific config files:

```bash
$ acm export
✅ Exported tool configs to ~/.config/agent/exports/
   - wallet-monitor.json
   - reputation-scanner.json
   - security-dashboard.json
```

## Security

- Config stored at `~/.config/agent/config.json`
- File permissions: `0600` (owner read/write only)
- API keys are masked in `acm show` output
- Never commit config to version control

## Part of Agent Security Stack

- [agent-tx-firewall](https://github.com/arithmosquillsworth/agent-tx-firewall)
- [agent-honeypot](https://github.com/arithmosquillsworth/agent-honeypot)
- [prompt-guard](https://github.com/arithmosquillsworth/prompt-guard)
- [tx-simulator](https://github.com/arithmosquillsworth/tx-simulator)
- [agent-security-dashboard](https://github.com/arithmosquillsworth/agent-security-dashboard)
- [agent-wallet-monitor](https://github.com/arithmosquillsworth/agent-wallet-monitor)
- [agent-reputation-scanner](https://github.com/arithmosquillsworth/agent-reputation-scanner)
- **agent-config-manager** (this repo)

## License

MIT
