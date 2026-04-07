
## Core Language & Basics

* Go is **statically typed, compiled**, designed for **simplicity + concurrency**
* Built mainly for **backend systems, cloud, DevOps tools**
* Program execution starts from **`package main` → `main()`**
* **Zero values** prevent uninitialized bugs (int → 0, string → "")
* `:=` → shorthand (inside functions), `var` → explicit declaration
* **Go has only one loop → `for`**
* **Slices > Arrays** in real-world usage (dynamic vs fixed)
* **Maps are NOT thread-safe**

---

## Functions & Memory

* Functions support **multiple return values** (common for `value, error`)
* **Defer executes in LIFO order** (important trap question)
* **Pointers modify original data**, values create copies
* Use **pointer receivers** when:

  * You want to modify struct
  * Avoid copying large structs
* Struct = **custom data model**, not OOP class

---

## Interfaces & Errors

* Interfaces are **implicitly implemented** (no `implements` keyword)
* Interface = **behavior, not data**
* `error` is just an **interface**
* **Custom error** = implement `Error() string`
* `panic` = crash, `recover` = handle panic (used rarely in production)

---

## Concurrency Fundamentals

* **Goroutine = lightweight thread (~2KB stack)**
* Start using `go func()`
* Go uses **M:N scheduler (G-M-P model)** → goroutines mapped to OS threads
* **Concurrency ≠ Parallelism**

  * Concurrency = structure
  * Parallelism = execution

---

## Channels (Very Important)

* Channels = **communication, not shared memory**
* **Unbuffered channel → blocking (sync)**
* **Buffered channel → non-blocking (async until full)**
* Sending blocks until receiver is ready (unbuffered)
* Receiving blocks until value is available
* **Closing channel → returns zero value**
* Use `range` to read until channel closes
* `select` = handle multiple channels (like switch)

---

## Advanced Channel Concepts

* **Directional channels** improve compile-time safety:

  * `chan<-` (send only)
  * `<-chan` (receive only)
* Buffered channel works like **FIFO queue**
* Full buffered channel → sender blocks
* Empty channel → receiver blocks
* Producer-consumer pattern → **most common interview pattern**

---

## Synchronization (Critical)

* **WaitGroup → wait for goroutines to finish**

  * Always match `Add()` and `Done()`
* **Mutex → protects shared data**

  * Lock → critical section → Unlock
* **Maps require mutex** for concurrent access
* Misusing WaitGroup → can cause deadlock

---

## Race Conditions & Deadlocks (Must Know)

* Race condition = **multiple goroutines access shared data unsafely**
* Happens due to **non-atomic operations**
* Fix using:

  * Mutex
  * Channels
  * Atomic operations
* Deadlock = **all goroutines blocked**
* Common causes:

  * No receiver for channel
  * Circular dependency
* Go runtime detects → `fatal error: all goroutines are asleep`

---

## Context & Control

* `context` is used for:

  * Cancellation
  * Timeout
  * Deadlines
* Prevents **goroutine leaks**
* Always call `cancel()` to release resources

---

## Runtime & Performance

* **Atomic operations** → faster than mutex (lock-free)
* Goroutine leaks happen when:

  * Channel not closed
  * Context not cancelled
* Scheduler issues:

  * Starvation
  * Long-running goroutines block others

---

## System Design Patterns (Go Specific)

* **Worker Pool** → fixed workers process jobs from channel
* **Fan-out** → multiple workers consume tasks
* **Fan-in** → merge results into single channel

---

## Golden Interview Lines (Use These)

* “Don’t communicate by sharing memory; share memory by communicating”
* “Channels are for coordination, mutex is for protection”
* “Concurrency is about structure, parallelism is about execution”
* “Use context to avoid goroutine leaks”
