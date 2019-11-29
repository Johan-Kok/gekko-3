package config

import "github.com/gottingen/gekko/gflag"

type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}

type FlagValue interface {
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}

type gflagValueSet struct {
	flags *gflag.FlagSet
}

func (g gflagValueSet) VisitAll(fn func(flag FlagValue)) {
	g.flags.VisitAll(func(flag *gflag.Flag) {
		fn(gflagValue{flag})
	})
}

// gflagValue is a wrapper aroung *gflag.flag
// that implements FlagValue
type gflagValue struct {
	flag *gflag.Flag
}

// HasChanged returns whether the flag has changes or not.
func (g gflagValue) HasChanged() bool {
	return g.flag.Changed
}

// Name returns the name of the flag.
func (g gflagValue) Name() string {
	return g.flag.Name
}

// ValueString returns the value of the flag as a string.
func (g gflagValue) ValueString() string {
	return g.flag.Value.String()
}

// ValueType returns the type of the flag as a string.
func (g gflagValue) ValueType() string {
	return g.flag.Value.Type()
}
