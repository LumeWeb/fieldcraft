// Package fieldcraft provides declarative fields, provenance, derivation,
// validation, and gathering shared by interactive and headless workflows.
//
// A host describes the values it needs as typed Fields and the framework
// resolves each one in precedence order — derivation (precedence 0), an
// explicit switch/flag (precedence 1), a surviving operator decision
// (precedence 2), a persisted env file (precedence 3), and a declared
// fallback default (precedence 4) — instead of hand-rolling per-field
// conditionals on state and interactivity.
//
// Two channels of provenance are kept per value:
//
//	Decided      — set only by an operator switch or interactive prompt
//	Operational  — the current working value (derived, folded, or defaulted)
//
// Gathering is headless-capable: with prompts disabled an unresolved required
// field is a hard error; otherwise it loads a Prompter from the context and
// asks the operator. Terminal rendering is deliberately out of scope — hosts
// implement Prompter with their UI toolkit of choice (e.g. pterm) and bind it
// via WithPrompter.
//
// The same declarative fields also project to JSON Schema (FormSchema) so one
// declaration can drive both an interactive prompt and a structured form.
//
// A minimal heterogeneous form:
//
//	dec := fieldcraft.Decided[*Config, string]{
//	    Read:  func(s *Config, n string) *string { return s.decisions[n] },
//	    Write: func(s *Config, n, v string) { s.decisions[n] = &v },
//	}
//	fields := []fieldcraft.AnyField[*Config]{
//	    fieldcraft.Str(dec, "Domain", getDomain, setDomain,
//	        fieldcraft.Meta{Flag: "domain", EnvFileKey: "APP_DOMAIN"}),
//	}
//	seeded, fullyDecided, err := fieldcraft.GatherAny(ctx, src, cfg, fields)
package fieldcraft
