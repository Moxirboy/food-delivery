package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	// Initialize the enforcer with a model configuration file and policy file.
	e, err := casbin.NewEnforcer("internal/configs/rbac.conf", "internal/configs/main_policy.csv")
	if err != nil {
		fmt.Println("Error creating enforcer:", err)
		return
	}

	// Add a policy to the database
	_, err = e.AddPolicy("alice", "data1", "read")
	if err != nil {
		log.Fatalf("Failed to add policy: %v", err)
	}
	// Check permissions.
	allowed, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		fmt.Println("Error checking permissions:", err)
		return
	}

	if allowed {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}
}
