# fieldcraft

[![Go Version](https://img.shields.io/badge/Go-1.26.0-blue)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build Status](https://github.com/LumeWeb/fieldcraft/actions/workflows/go.yml/badge.svg)](https://github.com/LumeWeb/fieldcraft/actions/workflows/go.yml)

A declarative field framework for Go: typed fields, value-source provenance,
derived values with precedence, headless gathering, and JSON-schema projection,
shared by interactive and headless workflows.

A host declares the values it needs; `fieldcraft` resolves each one in
precedence order (derivation → flag → operator decision → persisted env →
default) and prompts only when nothing else resolves the field. Terminal
rendering is out of scope: implement `Prompter` with your UI toolkit and bind
it with `WithPrompter`.

## Example usage

```go
dec := fieldcraft.Decided[*Config, string]{
    Read:  func(s *Config, n string) *string { return s.decisions[n] },
    Write: func(s *Config, n, v string) { s.decisions[n] = &v },
}
fields := []fieldcraft.AnyField[*Config]{
    fieldcraft.Str(dec, "Domain",
        func(s *Config) string { return s.Domain },
        func(s *Config, v string) { s.Domain = v },
        fieldcraft.Meta{Flag: "domain", EnvFileKey: "APP_DOMAIN"}),
}

seeded, fullyDecided, err := fieldcraft.GatherAny(ctx, src, cfg, fields)
```

## License

MIT
