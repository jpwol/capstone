# Project Overview

This project is composed of two parts:

- A backend generic API consisting of a DAG dependency manager and a decentralized greedy worker-pool
- A frontend cli (with config file) build-runner layered over the backend

## Backend API

The API will consist of two main parts

1. A DAG for dependency tracking
1. A decentralized greedy worker-pool that communicates with the DAG to complete tasks based on satisfied dependencies

#### Implementation

Worker-pool will take jobs (functions) from a global queue, published by the DAG manager.

Consider making the DAG dynamic, so jobs being completed by workers can create more jobs.

1. DAG representation

   - Nodes have tasks and directed edges
   - Use adjacency lists (acquired from frontend) to resolve dependencies and dependents
   - Detect cycles
   - Validate the graph
   - Track task state

   Here we should be able to build an inspect graphs. Build tests to validate them.

   Dependency Tracking

   - Initialize a *remaining-dependency* counter.
   - Track task completion (to allow tasks with dependencies to run)
   - "Unlock" dependent tasks
   - Handle failed dependencies (Backend should propagate error and let frontend decide)
   - Publish newly ready tasks

   Basic Execution Engine (Optional)

   Work-stealing scheduler (greedy worker-pool)

   - Worker local deques
   - Work acquisition (global queue)
   - Work stealing (thieve from other worker's queues)
   - Queue synchronization
   - Idle worker behavior (sleep and wake signals?)
   - Acquire scheduler metrics

   Tie Backend API Together

   - Public task and graph interfaces
   - Engine configuration
   - Execution lifecycle
   - Cancellation (cancelling the frontend should terminate child processes)
   - Error reporting/propagation
   - Scheduler selection (if centralized is implemented as well)

   Adjacency list example

   ```go
   // pseudo-code
   A: [B, C] // A depends on B and C
   B: [D, E] // B depends on D and E
   ```

   From this we can build a DAG, where:

   - `Node A` has outgoing edges to `B` and `C`
   - `Node B` has outgoing edges to `D` and `E`

   `A` has no dependencies, so it can be run immediately. Once it finishes, `B` and `C` can be run concurrently as they do not depend on eachother.

   ```go
   type NodeId u32

   type Graph struct {
        adjacency map[NodeId][]NodeId
   }

   adjacency := [][]uint32{
     {1,2},
     {3,4},
     {4},
     {},
     {},
   }
   ```

## Frontend Build-Runner
