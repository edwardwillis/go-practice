# go-practice

Scratch repo for practicing Go interview questions. One package per problem
category under `problems/`, each solution paired with a table-driven test.

## Layout

```
problems/
  arrays/       TwoSum
  strings/      IsAnagram
  linkedlist/   Node + Reverse
  stack/        IsValid (balanced brackets)
  trees/        TreeNode + MaxDepth
  dp/           ClimbStairs
  graphs/       NumIslands
template/       copy this to start a new problem
scratch/        `go run ./scratch` for quick one-off experiments
```

## Workflow

Add a new problem:

```
make new CATEGORY=arrays NAME=three_sum
```

This copies `template/TEMPLATE.go` and `TEMPLATE_test.go` into
`problems/arrays/three_sum.go` / `three_sum_test.go`. Fill in the signature
and test cases, then implement.

Run everything:

```
make test            # go test ./...
make test-verbose    # -v output
make cover           # coverage summary
make lint            # gofmt + go vet
```

Run a single package while iterating:

```
go test ./problems/arrays/... -run TestTwoSum -v
```

## Conventions

- One exported function (or small related set) per file, named after the
  problem.
- Every solution file's doc comment states time/space complexity.
- Tests are table-driven (`[]struct{ name, input..., want }`) with
  `t.Run(tt.name, ...)` subtests.
- Shared types (`linkedlist.Node`, `trees.TreeNode`) live in a
  non-test file in their package and are reused across problems in that
  category.
