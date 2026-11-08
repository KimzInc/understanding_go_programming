# Go Escape Analysis: Benchmark Results and Takeaways

Notes from benchmarking `MakePerson` (returns by value) and `MakePersonPtr` (returns a pointer) with inlining on and off.

## Results

The `allocs/op` and `B/op` columns are exact. Treat the `ns/op` figures as approximate.

| Benchmark | Inlining on | Inlining off (`-gcflags=-l`) |
| --- | --- | --- |
| `MakePerson` | 5.4 ns, 0 allocs | 7.8 ns, 0 allocs |
| `MakePersonPtr_Escapes` | 37.5 ns, 48 B, 1 alloc | 53.3 ns, 48 B, 1 alloc |
| `MakePersonPtr_Local` | 1.3 ns, 0 allocs | 58.7 ns, 48 B, 1 alloc |

Commands used:

```bash
go test -bench . -benchmem
go test -bench . -benchmem -gcflags=-l
```

## What the numbers show

### 1. `Local` is the clearest result

The code is identical in both runs, yet turning off inlining took it from 0 to 1 allocation and from 1.3 ns to 58.7 ns, roughly 44 times slower. With inlining, the compiler could see the pointer never left the loop body, so the `Person` stayed on the stack. Without it, `MakePersonPtr` is a standalone function that must be safe for any caller, so it heap-allocates. This is the same behavior as lines 28 and 158 in the `-gcflags="-m"` output.

### 2. `Escapes` costs the same 1 alloc either way

The pointer is stored in a global, so it has to be on the heap no matter what. Inlining makes no difference to the allocation count, and the small time difference is mostly function call overhead and noise.

### 3. A heap allocation is expensive compared to the work itself

`Escapes` takes about 7 times longer than `MakePerson` (37.5 ns vs 5.4 ns) for the same logical job of building a 40-byte struct. Most of that gap is the allocator, and it doesn't include the later cost of the garbage collector tracing and freeing that memory. In a hot path called millions of times, that adds up.

### 4. `MakePerson` is stable

It goes from 5.4 to 7.8 ns when inlining is disabled, which is just the cost of a real function call. It never allocates, so it doesn't depend on escape analysis at all.

## Why `Local` (1.3 ns) beats `MakePerson` (5.4 ns)

This is not a pointer-vs-value result. `MakePerson` writes a full `Person` (two strings plus an int) into a global each iteration. `Local` only reads `Age` and adds it to an int, so the compiler can skip building the unused fields. The two benchmarks do different amounts of work, so compare each one against its own inlining-off number rather than against each other.

The 48 B figure matches expectations: the struct is 40 bytes (two 16-byte strings plus an 8-byte int), rounded up to Go's 48-byte size class.

## Caveats

- The `allocs/op` and `B/op` columns are exact and reliable. These are the real lesson.
- For trustworthy timings, run `go test -bench . -benchmem -count=10` and compare runs with `benchstat`.

## What the evidence supports

- Using a pointer does not automatically mean heap, and using a value does not automatically mean stack. Escape analysis decides, based on whether the data can outlive the frame.
- Inlining changes what the compiler can prove, so it can change where data lives.
- Heap allocation has a real cost, so returning small structs by value is a sensible default in Go.

## 
