
## Golang Interview Summary Table (Questions → Concepts)

| Topic                          | Why                             | What is                                | Example                  |
| ------------------------------ | ------------------------------- | -------------------------------------- | ------------------------ |
| **Go Purpose**                 | To build fast, scalable systems | Designed for concurrency + performance | Backend, cloud systems   |
| **Go Usage**                   | To solve real-world systems     | Used in backend, DevOps, cloud         | Docker, Kubernetes       |
| **Language Type**              | To ensure safety + speed        | Statically typed, compiled             | `go build`               |
| **Run Program**                | To execute code                 | CLI execution                          | `go run main.go`         |
| **Modules**                    | To manage dependencies          | Go dependency system                   | `go mod init app`        |
| **package main**               | To define entry package         | Required for executable program        | `package main`           |
| **main()**                     | To start execution              | Entry function                         | `func main(){}`          |
| **Built-in Types**             | To store data                   | `int`, `string`, `bool`, `float`       | `var x int`              |
| **Zero Value**                 | To avoid uninitialized bugs     | Default values                         | `int → 0`                |
| **var vs :=**                  | To declare variables            | `var` explicit, `:=` shorthand         | `x := 10`                |
| **Constants**                  | To define fixed values          | Immutable variables                    | `const Pi = 3.14`        |
| **Control Flow**               | To control logic                | `if`, `for`, `switch`                  | `if x > 0 {}`            |
| **Function**                   | To reuse logic                  | Block of code                          | `func add()`             |
| **Defer**                      | To delay execution              | Runs at end (LIFO)                     | `defer close()`          |
| **Multiple Returns**           | To return more info             | Multiple values                        | `return x, err`          |
| **Pointers**                   | To modify original data         | Store memory address                   | `*int`                   |
| **Struct**                     | To model data                   | Custom type                            | `type User struct{}`     |
| **Array**                      | Fixed storage                   | Fixed size collection                  | `[3]int{}`               |
| **Slice**                      | Flexible storage                | Dynamic array                          | `[]int{}`                |
| **Slice Capacity**             | To optimize memory              | Max before reallocation                | `cap(slice)`             |
| **Map**                        | Fast lookup                     | Key-value store                        | `map[string]int{}`       |
| **Range**                      | To iterate                      | Loop helper                            | `for k,v := range`       |
| **Interface**                  | To achieve abstraction          | Defines behavior                       | `type Shape interface{}` |
| **Interface Implementation**   | For flexibility                 | Implicit implementation                | No `implements` keyword  |
| **Error**                      | To handle failures              | Built-in interface                     | `error`                  |
| **Custom Error**               | To extend errors                | Implement `Error()`                    | `type MyErr struct{}`    |
| **Panic**                      | To handle critical failure      | Runtime crash                          | `panic("err")`           |
| **Recover**                    | To handle panic                 | Catch panic safely                     | `recover()`              |
| **Goroutine**                  | For concurrency                 | Lightweight thread                     | `go func()`              |
| **Start Goroutine**            | To run async tasks              | `go` keyword                           | `go task()`              |
| **Scheduler**                  | To manage execution             | M:N scheduling                         | Goroutines ↔ threads     |
| **Concurrency vs Parallelism** | To understand execution         | Structure vs execution                 | Concurrent ≠ Parallel    |
| **Channel**                    | To communicate                  | Data pipe between goroutines           | `make(chan int)`         |
| **Buffered Channel**           | To avoid blocking               | Channel with capacity                  | `make(chan int,2)`       |
| **Blocking Channel**           | To sync tasks                   | Waits sender/receiver                  | `<-ch`                   |
| **Select**                     | To handle multiple channels     | Channel switch                         | `select {}`              |
| **Range over Channel**         | To consume data                 | Reads until closed                     | `for v := range ch`      |
| **Channel Close Behavior**     | To avoid bugs                   | Returns zero value                     | `close(ch)`              |
| **Directional Channels**       | For safety                      | Send-only / receive-only               | `chan<- int`             |
| **Channel Conversion**         | For flexibility                 | Bi → directional                       | Func param               |
| **Channel as Queue**           | To buffer tasks                 | FIFO behavior                          | Buffered channel         |
| **Producer-Consumer**          | To manage workload              | One produces, one consumes             | `go producer()`          |
| **WaitGroup**                  | To sync goroutines              | Wait for completion                    | `wg.Wait()`              |
| **WaitGroup Misuse**           | To avoid bugs                   | Missing Add/Done                       | Deadlock risk            |
| **Mutex**                      | To protect data                 | Lock mechanism                         | `mu.Lock()`              |
| **Maps Thread Safety**         | To avoid race                   | Not safe by default                    | Use mutex                |
| **Race Condition**             | To ensure correctness           | Concurrent unsafe access               | Shared variable          |
| **Fix Race Condition**         | To ensure safety                | Mutex / channels                       | `mu.Lock()`              |
| **Atomic Operation**           | To optimize performance         | Lock-free ops                          | `atomic.AddInt32`        |
| **Mutex vs Channel**           | To choose approach              | Lock vs communication                  | Depends use case         |
| **Deadlock**                   | To avoid freeze                 | All goroutines blocked                 | Channel misuse           |
| **Deadlock Cause**             | To debug issues                 | Circular wait                          | No receiver              |
| **Deadlock Detection**         | To identify issues              | Runtime panic                          | `fatal error`            |
| **Goroutine Leak**             | To avoid memory issue           | Goroutine never exits                  | Missing cancel           |
| **Context Package**            | To control lifecycle            | Timeout/cancel                         | `context.WithTimeout()`  |
| **Time Package**               | To manage time                  | Sleep, timers                          | `time.Sleep()`           |
| **Pointer vs Value Receiver**  | To control mutation             | Pointer modifies original              | `func (p *T)`            |
| **Defer Internals**            | To understand execution         | LIFO stack                             | Multiple defers          |
| **M-P-G Model**                | To understand runtime           | Goroutine-Thread mapping               | G-M-P                    |
| **Scheduling Issues**          | To optimize performance         | Starvation/preemption                  | Long-running goroutine   |
| **Worker Pool**                | To scale tasks                  | Fixed workers process jobs             | Jobs channel             |
| **Fan-In Fan-Out**             | To handle pipelines             | Multiple producers/consumers           | Merge channels           |


---

## Golang Golden Lines (Top 20)

1. **“Go is designed for simplicity, concurrency, and performance.”**

2. **“Don’t communicate by sharing memory; share memory by communicating.”**

3. **“Goroutines are lightweight threads managed by the Go runtime.”**

4. **“Channels are used for communication, mutex is used for protection.”**

5. **“Concurrency is about structure, parallelism is about execution.”**

6. **“Go uses an M:N scheduler (G-M-P model) to manage goroutines efficiently.”**

7. **“An unbuffered channel is synchronous; a buffered channel is asynchronous.”**

8. **“Sending or receiving on a channel blocks until the other side is ready.”**

9. **“Closing a channel signals no more values; reads return zero value.”**

10. **“Interfaces in Go are implemented implicitly, not explicitly.”**

11. **“Interface defines behavior, not data.”**

12. **“Error in Go is just an interface, not an exception system.”**

13. **“Defer statements execute in LIFO order after the function returns.”**

14. **“Use pointer receivers when you want to modify the original value.”**

15. **“Maps are not thread-safe; use mutex or sync mechanisms.”**

16. **“Race conditions occur due to unsynchronized access to shared memory.”**

17. **“Deadlock happens when all goroutines are blocked indefinitely.”**

18. **“Context is used to control cancellation, timeout, and request lifecycle.”**

19. **“WaitGroup is used to wait for multiple goroutines to complete.”**

20. **“Atomic operations provide lock-free concurrency for simple cases.”**

---

## How to Use These

* Use **2–3 lines naturally in answers**, not all at once
* Combine with **code examples** for strong impact
