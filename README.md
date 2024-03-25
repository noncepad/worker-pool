# Worker Pool Manager

The Worker Pool Manager is a Go package designed to manage workers and distribute jobs efficiently in a concurrent environment.

## Overview

This package consists of several components:

- `manager`: Manages the creation and lifecycle of workers, as well as job distribution.
- `pool`: Provides interfaces for managing workers and submitting jobs.
- `meter`: Implements tracking the status of workers.

## Components

### manager.go

The `manager` package includes functionalities for managing workers and distributing jobs. Here's an overview of key functionalities and implementations:

- **Worker Creation**: Workers are created and managed by the `external` struct within the `Add` method. Workers are spawned as goroutines and communicate with the manager through channels.
- **Job Processing**: Job processing occurs within the `loopWorkerListen` function. Jobs are received from the `jobC` channel, processed by workers, and results are sent back through channels.
- **Concurrency**: Concurrency is achieved through goroutines and channels, allowing multiple workers to process jobs concurrently.
- **Worker Tracking**: The status of workers is tracked through the `internal` struct, where each worker is assigned a unique ID.

### pool.go

The `pool` package provides interfaces for managing workers and submitting jobs. Here's a summary of key functionalities:

- **Manager**: The `Manager` interface defines methods for adding workers to the pool and submitting jobs for processing.
- **Job Submission**: Jobs can be submitted to the manager for processing using the `Submit` method.
- **Results**: Results of job processing are in the `Result` struct, which contains payload data and error information.

### worker.go

The `worker` package defines the behavior of individual workers within the worker pool. Here's an overview of the key implementations:

- **Worker**: Workers implement the `Worker` interface, which defines methods for running jobs and managing their lifecycle.
- **Jobs**: Jobs are executed by workers within the `Run` method, where the actual job processing takes place.
- **Concurrency**: Workers handle job processing concurrently, allowing multiple jobs to be executed simultaneously.

### meter.go

The `meter` package is for tracking the status of workers. Here's how worker tracking is implemented:

- **Hooks**: The `Hook` struct tracks the number of busy and idle workers.
- **Concurrency Control**: Concurrency is managed using a mutex to safely update worker status counters.
- **Hook Usage**: Hooks are used within the manager to track worker activity, such as when a worker starts or finishes processing a job.

## Usage

To use the Worker Pool Manager, import into your Go project and utilize the provided interfaces and functionalities to manage workers and distribute jobs efficiently.

```go
import (
    "github.com/noncepad/worker-pool"
)
