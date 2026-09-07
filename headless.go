package fieldcraft

// NonInteractive disables all interactive prompts in the framework. Set it
// before gathering to force a headless/flag-driven run. Gather reads it to
// decide whether an unresolved required field is a hard error (headless) or an
// interactive prompt.
//
// It is a package-level flag for behavioral parity with legacy consumers; new
// composition roots should prefer request-scoped state instead.
var NonInteractive bool
