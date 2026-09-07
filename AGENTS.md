# AGENTS.md

This file provides development guidelines and architectural documentation for
the fieldcraft project.

## Common Commands

### Building
```bash
# Build all packages
go build -v ./...
```

### Testing
```bash
# Run all tests with race detection and coverage
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# Run tests for the root package only
go test -v -race .

# View coverage report
go tool cover -func=coverage.out
```

### Mock Generation

```bash
# Generate mocks for interfaces (uses .mockery.yaml; mockery is
# pre-installed at $HOME/go/bin/mockery; never reinstall it)
mockery
```

### Dependency Management
```bash
# Download dependencies
go mod download

# Verify dependencies
go mod verify

# Tidy dependencies
go mod tidy
```

## Project Overview

fieldcraft is a declarative field framework: hosts declare the typed values
they need, and the framework resolves each one in precedence order
(derivation, switch/flag, operator decision, persisted env file, fallback
default) instead of hand-rolling per-field conditionals. It serves both
interactive workflows (via a pluggable `Prompter`) and headless ones, and
projects the same field declarations to JSON Schema.

## Architecture

### Package Structure
- **Module path**: `go.lumeweb.com/fieldcraft`
- **Root package** (`fieldcraft`): all functionality as a flat package;
  `doc.go` holds the package documentation, `errors.go` the sentinel errors

### Design
- **Two-channel provenance**: `Decided` (operator switch/prompt, set only by
  an explicit decision) vs `Operational` (derived, env-folded, or defaulted
  working value). A fold must never look like a decision.
- **Precedence**: `Field.Derived` (0) > `Flag` (1) > surviving `Decided` (2) >
  `EnvFileKey` fold (3) > `DefaultVal` (4); unresolved required fields
  hard-error headless and prompt interactively.
- **Type erasure**: `AnyField[S]` lets a heterogeneous field set resolve in one
  `GatherAny` pass; the `Str`/`Bool`/`Int`/`Enum` constructors return
  already-erased fields.
- **Headless**: the `ValueSource` interface abstracts switches and the
  persisted env file; `NonInteractive` forces a prompt-free run where an
  unresolved required field is a `FieldError` (wrapping `ErrUnresolvedField`).
- **Prompter interfaces**: `Prompter` is the only input abstraction; terminal
  rendering (pterm or otherwise) explicitly does NOT belong in this module —
  hosts implement the adapter.
- **Schema projection**: `Field.Schema` / `FormSchema` emit invopop/jsonschema
  from the same declarations that drive gathering.

### Testing Conventions
- Tests are colocated next to source (`*_test.go`); scripted `ValueSource`
  and `Prompter` fakes live in the test files
- Tests pin current behavior (characterization) — the precedence and
  provenance rules are regression-guarded; do not change semantics silently
- Do not add a README/board copyright header to source files; attribution
  lives only in the LICENSE file
- Generated mocks live in `mocks/` (mockery, testify templates)
