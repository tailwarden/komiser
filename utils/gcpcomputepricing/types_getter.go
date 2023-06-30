package gcpcomputepricing

type typeMachineGetter func(p *Pricing, opts Opts) (Subtype, Subtype, error)

type typeDiskGetter func(p *Pricing, opts Opts) (Subtype, error)
