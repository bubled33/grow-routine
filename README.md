# 🌱 Grow-Routine (IoT Control Plane Simulator)

**High-throughput IoT simulation platform written in Go.**

> *"Where concurrency meets cultivation."*

Grow-Routine emulates a massive network of concurrent IoT sensors for a smart greenhouse complex. It demonstrates advanced Go concurrency patterns, lock-free data structures, and high-load system architecture without external dependencies.

## 🚀 Key Features

- **Massive Concurrency:** Simulates 10,000+ independent sensor goroutines.
- **Fan-In / Fan-Out:** Efficient data aggregation using channel multiplexing.
- **Non-Blocking I/O:** Implements drop patterns to handle backpressure.
- **Thread-Safe State:** Uses `sync.RWMutex` and `sync/atomic` for safe data access.
- **Dynamic Control:** Broadcasts commands (e.g., "Emergency Stop") via channel closure.
- **Graceful Shutdown:** Robust orchestration using `context` and `sync.WaitGroup`.

## 🛠 Tech Stack

- **Language:** Go 1.25+
- **Core Concepts:** Goroutines, Buffered Channels, Select, Mutexes, Atomics.
- **Patterns:** Worker Pool, Pipeline, Generator, Semaphore.

## 🏗 Architecture Modules

1.  **Sensor Mesh (Generator):** Thousands of long-lived goroutines generating telemetry.
2.  **Data Bus (Fan-In):** Non-blocking aggregation pipeline with backpressure handling.
3.  **Control Plane (State):** In-memory storage of sensor states protected by RWMutex.
4.  **Action Workers (Pool):** Worker pool processing complex events (e.g., "Cooling System").
5.  **Orchestrator:** Manages lifecycle, configuration updates, and system shutdown.

## 🚦 Getting Started

### Prerequisites

- Go 1.25 or higher
- Make (optional)

### Installation

