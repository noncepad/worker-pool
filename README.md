# Worker Pool Manager

The Worker Pool Manager is a Go package designed to manage workers and distribute jobs efficiently in a concurrent environment.

## Overview

This package consists of several components:

- `manager`: Manages the creation and lifecycle of workers, as well as job distribution.
- `pool`: Provides interfaces for managing workers and submitting jobs.
- `meter`: Implements tracking the status of workers and streams worker capacity updates over gRPC.

## Components

### manager.go

The `manager` package includes functionalities for managing workers and distributing jobs. Here's an overview of key functionalities and implementations:

- **Worker Creation**: workers are created and managed by the `external` struct within the `Add` method. Workers are spawned as goroutines and communicate with the manager through channels.
- **Job Processing**: job processing occurs within the `loopWorkerListen` function. Jobs are received from the `jobC` channel, processed by workers, and results are sent back through channels.
- **Concurrency**: concurrency is achieved through goroutines and channels, allowing multiple workers to process jobs concurrently.
- **Worker Tracking**: the status of workers is tracked through the `internal` struct, where each worker is assigned a unique ID.

### pool.go

The `pool` package provides interfaces for managing workers and submitting jobs. Here's a summary of key functionalities:

- **Manager**: the `Manager` interface defines methods for adding workers to the pool and submitting jobs for processing.
- **Job Submission**: jobs can be submitted to the manager for processing using the `Submit` method.
- **Results**: results of job processing are in the `Result` struct, which contains payload data and error information.

### worker.go

The `worker` package defines the behavior of individual workers within the worker pool. Here's an overview of the key implementations:

- **Worker**: workers implement the `Worker` interface, which defines methods for running jobs and managing their lifecycle.
- **Jobs**: jobs are executed by workers within the `Run` method, where the actual job processing takes place.
- **Concurrency**: workers handle job processing concurrently, allowing multiple jobs to be executed simultaneously.

### meter.go

The `meter` package is for tracking the status of workers. Here's how worker tracking is implemented and integrates `solpipe.proto`:

- **Hooks**: the `Hook` struct tracks the number of busy and idle workers.
- **Concurrency Control**: concurrency is managed using a mutex to safely update worker status counters.
- **Hook Usage**: hooks are used within the manager to track worker activity, such as when a worker starts or finishes processing a job.
- **Integration with Protobuf**: the `meter` package implements the `WorkerStatus` service defined in `solpipe.proto`. This service allows clients to stream capacity updates from the worker pool.
- **Streaming Capacity Updates**: the `OnStatus` method in the `Hook` struct streams capacity responses to clients at 30 second intervals. Clients can connect to this service and receive updates on the current capacity of the worker pool.
- **Concurrency Control**: concurrency is managed using mutexes to safely update worker status counters and stream capacity updates.


## Usage

To use the Worker Pool Manager, import into your Go project and utilize the provided interfaces and functionalities to manage workers and distribute jobs efficiently.

```go
import (
    "github.com/noncepad/worker-pool"
)

