package configuration

import (
	"fmt"
	"net/http"
)

func StartServer(port string, handler http.Handler) error {
	// address := fmt.Sprintf(":%s", port)
	fmt.Printf("🚀 Server starting on port %s...\n", port)

	if err := http.ListenAndServe(port, handler); err != nil {
		return fmt.Errorf("❌ error starting server on port %s: %w", port, err)
	}

	return nil
}
