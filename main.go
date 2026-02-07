package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const version = "0.1.0"

// AgentConfig is the unified configuration for all agent tools
type AgentConfig struct {
	Version     string            `json:"version"`
	Agent       AgentInfo         `json:"agent"`
	Wallet      WalletConfig      `json:"wallet"`
	Security    SecurityConfig    `json:"security"`
	APIKeys     APIKeysConfig     `json:"api_keys"`
	Monitoring  MonitoringConfig  `json:"monitoring"`
}

type AgentInfo struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	ERC8004ID   int    `json:"erc8004_id"`
	Website     string `json:"website"`
	GitHub      string `json:"github"`
}

type WalletConfig struct {
	Address       string   `json:"address"`
	Networks      []string `json:"networks"`
	DailyLimit    float64  `json:"daily_limit"`
	AlertThreshold float64 `json:"alert_threshold"`
}

type SecurityConfig struct {
	FirewallEnabled     bool     `json:"firewall_enabled"`
	HoneypotEnabled     bool     `json:"honeypot_enabled"`
	PromptGuardEnabled  bool     `json:"prompt_guard_enabled"`
	SimulatorEnabled    bool     `json:"simulator_enabled"`
	WhitelistedAddresses []string `json:"whitelisted_addresses"`
	BlacklistedAddresses []string `json:"blacklisted_addresses"`
}

type APIKeysConfig struct {
	Etherscan  string `json:"etherscan,omitempty"`
	Basescan   string `json:"basescan,omitempty"`
	OpenAI     string `json:"openai,omitempty"`
	Anthropic  string `json:"anthropic,omitempty"`
	Discord    string `json:"discord,omitempty"`
}

type MonitoringConfig struct {
	DashboardEnabled bool   `json:"dashboard_enabled"`
	DashboardPort    int    `json:"dashboard_port"`
	WebhookURL       string `json:"webhook_url,omitempty"`
	CheckInterval    int    `json:"check_interval_minutes"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "init":
		initConfig()
	case "show":
		showConfig()
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Usage: acm get <key>")
			os.Exit(1)
		}
		getValue(os.Args[2])
	case "set":
		if len(os.Args) < 4 {
			fmt.Println("Usage: acm set <key> <value>")
			os.Exit(1)
		}
		setValue(os.Args[2], os.Args[3])
	case "validate":
		validateConfig()
	case "export":
		exportConfig()
	case "version":
		fmt.Printf("agent-config-manager v%s\n", version)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("🔧 Agent Config Manager")
	fmt.Println("========================")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  acm init        - Create initial configuration")
	fmt.Println("  acm show        - Display current configuration")
	fmt.Println("  acm get <key>   - Get specific value (e.g., 'wallet.address')")
	fmt.Println("  acm set <key> <val> - Set specific value")
	fmt.Println("  acm validate    - Validate configuration")
	fmt.Println("  acm export      - Export config for all tools")
	fmt.Println("")
	fmt.Println("Config location: ~/.config/agent/config.json")
}

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "agent", "config.json")
}

func initConfig() {
	configPath := getConfigPath()
	configDir := filepath.Dir(configPath)
	
	// Create directory
	os.MkdirAll(configDir, 0755)
	
	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("⚠️  Config already exists at %s\n", configPath)
		fmt.Println("   Use 'acm show' to view or 'acm set' to modify")
		return
	}
	
	// Create default config
	config := AgentConfig{
		Version: version,
		Agent: AgentInfo{
			Name:      "Arithmos",
			ID:        "arithmos-quillsworth",
			ERC8004ID: 1941,
			Website:   "https://arithmos.dev",
			GitHub:    "https://github.com/arithmosquillsworth",
		},
		Wallet: WalletConfig{
			Address:        "0x120e011fB8a12bfcB61e5c1d751C26A5D33Aae91",
			Networks:       []string{"ethereum", "base"},
			DailyLimit:     0.5,
			AlertThreshold: 0.1,
		},
		Security: SecurityConfig{
			FirewallEnabled:     true,
			HoneypotEnabled:     true,
			PromptGuardEnabled:  true,
			SimulatorEnabled:    true,
			WhitelistedAddresses: []string{},
			BlacklistedAddresses: []string{},
		},
		APIKeys: APIKeysConfig{},
		Monitoring: MonitoringConfig{
			DashboardEnabled: true,
			DashboardPort:    8080,
			CheckInterval:    5,
		},
	}
	
	// Save config
	saveConfig(config)
	
	fmt.Printf("✅ Config created at %s\n", configPath)
	fmt.Println("")
	fmt.Println("Next steps:")
	fmt.Println("  1. Add API keys: acm set api_keys.etherscan YOUR_KEY")
	fmt.Println("  2. View config:  acm show")
	fmt.Println("  3. Validate:     acm validate")
}

func loadConfig() AgentConfig {
	configPath := getConfigPath()
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("❌ Config not found at %s\n", configPath)
		fmt.Println("   Run 'acm init' to create")
		os.Exit(1)
	}
	
	var config AgentConfig
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Printf("❌ Invalid config: %v\n", err)
		os.Exit(1)
	}
	
	return config
}

func saveConfig(config AgentConfig) {
	configPath := getConfigPath()
	
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("❌ Failed to marshal config: %v\n", err)
		os.Exit(1)
	}
	
	// Set restrictive permissions (no group/other read)
	if err := os.WriteFile(configPath, data, 0600); err != nil {
		fmt.Printf("❌ Failed to write config: %v\n", err)
		os.Exit(1)
	}
}

func showConfig() {
	config := loadConfig()
	
	fmt.Println("═".repeat(60))
	fmt.Println("  AGENT CONFIGURATION")
	fmt.Println("═".repeat(60))
	fmt.Println()
	
	fmt.Printf("Version: %s\n", config.Version)
	fmt.Println()
	
	fmt.Println("AGENT:")
	fmt.Printf("  Name:       %s\n", config.Agent.Name)
	fmt.Printf("  ID:         %s\n", config.Agent.ID)
	fmt.Printf("  ERC-8004:   #%d\n", config.Agent.ERC8004ID)
	fmt.Printf("  Website:    %s\n", config.Agent.Website)
	fmt.Printf("  GitHub:     %s\n", config.Agent.GitHub)
	fmt.Println()
	
	fmt.Println("WALLET:")
	fmt.Printf("  Address:    %s\n", config.Wallet.Address)
	fmt.Printf("  Networks:   %v\n", config.Wallet.Networks)
	fmt.Printf("  Daily Limit: %.2f ETH\n", config.Wallet.DailyLimit)
	fmt.Printf("  Alert Threshold: %.2f ETH\n", config.Wallet.AlertThreshold)
	fmt.Println()
	
	fmt.Println("SECURITY:")
	fmt.Printf("  Firewall:   %s\n", boolStatus(config.Security.FirewallEnabled))
	fmt.Printf("  Honeypot:   %s\n", boolStatus(config.Security.HoneypotEnabled))
	fmt.Printf("  Prompt Guard: %s\n", boolStatus(config.Security.PromptGuardEnabled))
	fmt.Printf("  Simulator:  %s\n", boolStatus(config.Security.SimulatorEnabled))
	fmt.Printf("  Whitelist:  %d addresses\n", len(config.Security.WhitelistedAddresses))
	fmt.Printf("  Blacklist:  %d addresses\n", len(config.Security.BlacklistedAddresses))
	fmt.Println()
	
	fmt.Println("API KEYS:")
	fmt.Printf("  Etherscan:  %s\n", keyStatus(config.APIKeys.Etherscan))
	fmt.Printf("  Basescan:   %s\n", keyStatus(config.APIKeys.Basescan))
	fmt.Printf("  OpenAI:     %s\n", keyStatus(config.APIKeys.OpenAI))
	fmt.Printf("  Anthropic:  %s\n", keyStatus(config.APIKeys.Anthropic))
	fmt.Printf("  Discord:    %s\n", keyStatus(config.APIKeys.Discord))
	fmt.Println()
	
	fmt.Println("MONITORING:")
	fmt.Printf("  Dashboard:  %s (port %d)\n", boolStatus(config.Monitoring.DashboardEnabled), config.Monitoring.DashboardPort)
	fmt.Printf("  Check Interval: %d minutes\n", config.Monitoring.CheckInterval)
	fmt.Printf("  Webhook:    %s\n", webhookStatus(config.Monitoring.WebhookURL))
	fmt.Println()
	fmt.Println("═".repeat(60))
}

func boolStatus(b bool) string {
	if b {
		return "✅ enabled"
	}
	return "❌ disabled"
}

func keyStatus(key string) string {
	if key == "" {
		return "❌ not set"
	}
	return "✅ set"
}

func webhookStatus(url string) string {
	if url == "" {
		return "not configured"
	}
	return "✅ configured"
}

func getValue(key string) {
	config := loadConfig()
	
	switch key {
	case "agent.name":
		fmt.Println(config.Agent.Name)
	case "agent.id":
		fmt.Println(config.Agent.ID)
	case "agent.erc8004_id":
		fmt.Println(config.Agent.ERC8004ID)
	case "wallet.address":
		fmt.Println(config.Wallet.Address)
	case "wallet.daily_limit":
		fmt.Println(config.Wallet.DailyLimit)
	case "wallet.alert_threshold":
		fmt.Println(config.Wallet.AlertThreshold)
	case "security.firewall_enabled":
		fmt.Println(config.Security.FirewallEnabled)
	case "security.honeypot_enabled":
		fmt.Println(config.Security.HoneypotEnabled)
	case "monitoring.dashboard_port":
		fmt.Println(config.Monitoring.DashboardPort)
	default:
		fmt.Printf("❌ Unknown key: %s\n", key)
		os.Exit(1)
	}
}

func setValue(key, value string) {
	config := loadConfig()
	
	switch key {
	case "api_keys.etherscan":
		config.APIKeys.Etherscan = value
	case "api_keys.basescan":
		config.APIKeys.Basescan = value
	case "api_keys.openai":
		config.APIKeys.OpenAI = value
	case "api_keys.anthropic":
		config.APIKeys.Anthropic = value
	case "api_keys.discord":
		config.APIKeys.Discord = value
	case "wallet.daily_limit":
		var limit float64
		fmt.Sscanf(value, "%f", &limit)
		config.Wallet.DailyLimit = limit
	case "wallet.alert_threshold":
		var threshold float64
		fmt.Sscanf(value, "%f", &threshold)
		config.Wallet.AlertThreshold = threshold
	case "monitoring.webhook_url":
		config.Monitoring.WebhookURL = value
	case "monitoring.check_interval":
		var interval int
		fmt.Sscanf(value, "%d", &interval)
		config.Monitoring.CheckInterval = interval
	default:
		fmt.Printf("❌ Unknown key: %s\n", key)
		os.Exit(1)
	}
	
	saveConfig(config)
	fmt.Printf("✅ Set %s\n", key)
}

func validateConfig() {
	config := loadConfig()
	
	fmt.Println("🔍 Validating configuration...")
	fmt.Println()
	
	issues := []string{}
	
	// Check required fields
	if config.Wallet.Address == "" {
		issues = append(issues, "❌ Wallet address not set")
	}
	
	if config.Wallet.DailyLimit <= 0 {
		issues = append(issues, "⚠️  Daily limit should be positive")
	}
	
	if config.APIKeys.Etherscan == "" {
		issues = append(issues, "⚠️  Etherscan API key not set (needed for monitoring)")
	}
	
	if config.APIKeys.Basescan == "" {
		issues = append(issues, "⚠️  Basescan API key not set (needed for monitoring)")
	}
	
	// Check security settings
	if !config.Security.FirewallEnabled && !config.Security.HoneypotEnabled {
		issues = append(issues, "⚠️  All security features disabled")
	}
	
	// Print results
	if len(issues) == 0 {
		fmt.Println("✅ Configuration is valid!")
	} else {
		for _, issue := range issues {
			fmt.Println(issue)
		}
		fmt.Println()
		fmt.Printf("Found %d issue(s)\n", len(issues))
	}
}

func exportConfig() {
	config := loadConfig()
	configPath := getConfigPath()
	
	// Export individual tool configs
	exportDir := filepath.Join(filepath.Dir(configPath), "exports")
	os.MkdirAll(exportDir, 0755)
	
	// Export for wallet-monitor
	walletConfig := map[string]interface{}{
		"address":         config.Wallet.Address,
		"etherscan_key":   config.APIKeys.Etherscan,
		"basescan_key":    config.APIKeys.Basescan,
		"check_interval":  config.Monitoring.CheckInterval,
		"alert_threshold": config.Wallet.AlertThreshold,
		"webhook_url":     config.Monitoring.WebhookURL,
	}
	exportToolConfig(exportDir, "wallet-monitor.json", walletConfig)
	
	// Export for reputation-scanner
	scannerConfig := map[string]interface{}{
		"address":      config.Wallet.Address,
		"etherscan_key": config.APIKeys.Etherscan,
		"basescan_key":  config.APIKeys.Basescan,
	}
	exportToolConfig(exportDir, "reputation-scanner.json", scannerConfig)
	
	// Export for security-dashboard
	dashboardConfig := map[string]interface{}{
		"port": config.Monitoring.DashboardPort,
	}
	exportToolConfig(exportDir, "security-dashboard.json", dashboardConfig)
	
	fmt.Printf("✅ Exported tool configs to %s/\n", exportDir)
	fmt.Println("   - wallet-monitor.json")
	fmt.Println("   - reputation-scanner.json")
	fmt.Println("   - security-dashboard.json")
}

func exportToolConfig(dir, filename string, config map[string]interface{}) {
	path := filepath.Join(dir, filename)
	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(path, data, 0600)
}

func (s string) repeat(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
