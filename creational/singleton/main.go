package main

import (
	"fmt"
	"sync"
)

// JoshBankConfigManager is our singleton that manages JoshBank application configuration.
// In a real-world scenario, this would load config from files, environment
// variables, or remote config servers.
type JoshBankConfigManager struct {
	config map[string]string
	mu     sync.RWMutex // Protects concurrent access to config map
}

var (
	// instance holds the single instance of JoshBankConfigManager
	instance *JoshBankConfigManager
	// once ensures the instance is created only once, even with concurrent calls
	once sync.Once
)

// GetInstance returns the singleton instance of JoshBankConfigManager.
// This is the only way to access the ConfigManager.
// Thread-safe: multiple goroutines can call this simultaneously.
func GetInstance() *JoshBankConfigManager {
	once.Do(func() {
		fmt.Println("Creating JoshBankConfigManager instance (this should appear only once)")
		instance = &JoshBankConfigManager{
			config: make(map[string]string),
		}
		// Simulate loading configuration from a file
		instance.loadDefaultConfig()
	})
	return instance
}

// loadDefaultConfig simulates loading configuration from external source
func (c *JoshBankConfigManager) loadDefaultConfig() {
	c.config["bank_name"] = "JoshBank"
	c.config["version"] = "1.0.0"
	c.config["database_url"] = "postgres://localhost:5432/joshbank"
	c.config["max_connections"] = "100"
	c.config["api_timeout"] = "30s"
	c.config["transaction_limit"] = "10000"
	c.config["kyc_provider"] = "internal"
}

// GetConfig retrieves a configuration value by key
func (c *JoshBankConfigManager) GetConfig(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.config[key]
}

// SetConfig updates a configuration value
func (c *JoshBankConfigManager) SetConfig(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.config[key] = value
}

// GetAllConfig returns all configuration (for demonstration)
func (c *JoshBankConfigManager) GetAllConfig() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Return a copy to prevent external modification
	configCopy := make(map[string]string)
	for k, v := range c.config {
		configCopy[k] = v
	}
	return configCopy
}

func main() {
	fmt.Println("=== Singleton Pattern: JoshBank Configuration Manager ===\n")

	// Simulate multiple parts of the application accessing config
	fmt.Println("1. Main application initializing...")
	config1 := GetInstance()
	fmt.Printf("   Bank Name: %s\n", config1.GetConfig("bank_name"))
	fmt.Printf("   Version: %s\n\n", config1.GetConfig("version"))

	fmt.Println("2. Payment processing module accessing config...")
	config2 := GetInstance()
	fmt.Printf("   Database URL: %s\n", config2.GetConfig("database_url"))
	fmt.Printf("   Max Connections: %s\n", config2.GetConfig("max_connections"))
	fmt.Printf("   Transaction Limit: %s\n\n", config2.GetConfig("transaction_limit"))

	fmt.Println("3. Updating configuration...")
	config2.SetConfig("transaction_limit", "20000")
	fmt.Printf("   Updated Transaction Limit: %s\n\n", config2.GetConfig("transaction_limit"))

	fmt.Println("4. Verifying singleton behavior...")
	config3 := GetInstance()
	fmt.Printf("   config1 == config2: %v\n", config1 == config2)
	fmt.Printf("   config2 == config3: %v\n", config2 == config3)
	fmt.Printf("   All point to same instance: %v\n\n", config1 == config2 && config2 == config3)

	fmt.Println("5. Testing concurrent access...")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := GetInstance()
			fmt.Printf("   Goroutine %d got instance: %s\n", id, cfg.GetConfig("bank_name"))
		}(i)
	}
	wg.Wait()

	fmt.Println("\n✓ All goroutines accessed the same singleton instance")
	fmt.Println("✓ JoshBank configuration is managed consistently across all modules")
}
