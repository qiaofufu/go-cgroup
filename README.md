## support 

support quick difine cgroup subsytems 

- cpu 
  - limit cpu use ratio
- memory
  - limit memory use size

## Usage

### Limit CPU ratio

```go
    cg := cgroup.New([]rescource.IResource{&rescource.CPU{
        Strategy: "test_cpu",
        Quota:    10 * rescource.MS,
        Period:   1 * rescource.MS,
    }})

    cg.AddPid(1001)

```

### Limit Memory Size

```go
    cg := cgroup.New([]rescource.IResource{&rescource.Memory{
        Strategy:      "test_memory",
        LimitInMemory: 512 * rescource.MB,
        Swappiness:    0,
    }})
    
    cg.AddPid(1001)
```


