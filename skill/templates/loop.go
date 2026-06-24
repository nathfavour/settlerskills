package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

// BlockNotification represents a simple parsed websocket block notification structure.
type BlockNotification struct {
	Slot uint64 `json:"slot"`
}

func main() {
	// Setup shutdown handler
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rpcURL := os.Getenv("SOLANA_RPC_WS_URL")
	if rpcURL == "" {
		rpcURL = "ws://localhost:8900" // default local validator ws
	}

	u, err := url.Parse(rpcURL)
	if err != nil {
		log.Fatalf("invalid websocket URL: %v", err)
	}

	log.Printf("Connecting to RPC Subscription stream at %s", u.String())
	
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	defer conn.Close()

	// Example subscription payload for slotNotifications
	subPayload := `{"jsonrpc":"2.0","id":1,"method":"slotSubscribe"}`
	if err := conn.WriteMessage(websocket.TextMessage, []byte(subPayload)); err != nil {
		log.Fatalf("subscribe request failed: %v", err)
	}

	// Ring-buffer simulation for low-allocation parsing
	const bufferSize = 1024
	ringBuffer := make([][]byte, bufferSize)
	writeIdx := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Printf("read error: %v", err)
					return
				}
				
				// Zero-allocation ring buffer placement
				ringBuffer[writeIdx] = message
				writeIdx = (writeIdx + 1) % bufferSize
				
				// Process the message asynchronously or inline
				// In production, parse block transactions and execute webhooks here
				fmt.Printf("Received slot/block update chunk, size: %d bytes\n", len(message))
			}
		}
	}()

	<-ctx.Done()
	log.Println("Gracefully shutting down RPC Event Stream daemon...")
}
