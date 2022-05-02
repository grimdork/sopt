package sopt

// Options base definition.
type Options struct {
	short      map[string]*Option
	long       map[string]*Option
	positional map[string]*Option
	groups     map[string]*Group
	commands   map[string]*Command
	// Order of groups.
	order []string
	// Remainder contains args not parsed as options, commands or positional args.
	Remainder []string
	// hashelp is true if default help is defined.
	hashelp bool
}

// New options instance.
func New() *Options {
	opt := &Options{
		short:      make(map[string]*Option),
		long:       make(map[string]*Option),
		positional: make(map[string]*Option),
		groups:     make(map[string]*Group),
		commands:   make(map[string]*Command),
	}

	opt.AddGroup("default")
	return opt
}

// GroupCount returns the number of groups.
func (opt *Options) GroupCount() int {
	return len(opt.order)
}

// AddGroup adds a new group. This ensures the order for help listing.
func (opt *Options) AddGroup(group string) *Group {
	opt.groups[group] = &Group{Name: group}
	opt.order = append(opt.order, group)
	return opt.groups[group]
}

// GetGroup returns a pointer to a group.
func (opt *Options) GetGroup(name string) *Group {
	if name == "" {
		return opt.groups["default"]
	}

	g := opt.groups[name]
	return g
}

// GetGroups returns a slice of groups.
func (opt *Options) GetGroups() []*Group {
	list := [](*Group){}
	for _, g := range opt.order {
		list = append(list, opt.groups[g])
	}
	return list
}

// RemoveGroup from map and order.
func (opt *Options) RemoveGroup(name string) {
	delete(opt.groups, name)
	for i, v := range opt.order {
		if v == name {
			opt.order = append(opt.order[:i], opt.order[i+1:]...)
			return
		}
	}
}

// GetOption returns a pointer to an option.
func (opt *Options) GetOption(name string) (*Option, bool) {
	var o *Option
	var ok bool
	if len(name) > 1 {
		o, ok = opt.long[name]
	} else {
		o, ok = opt.short[name]
	}

	if !ok {
		return nil, false
	}

	return o, true
}

// GetBool returns a bool option's value.
func (opt *Options) GetBool(name string) bool {
	o, ok := opt.GetOption(name)
	if !ok {
		return false
	}

	if o.Value == nil {
		if o.Default != nil {
			return o.Default.(bool)
		}

		return false
	}

	return o.Value.(bool)
}

// GetString returns a string option's value.
func (opt *Options) GetString(name string) string {
	o, ok := opt.GetOption(name)
	if !ok {
		return ""
	}

	if o.Value == nil {
		return o.Default.(string)
	}

	return o.Value.(string)
}

// GetInt returns an int option's value.
func (opt *Options) GetInt(name string) int {
	o, ok := opt.GetOption(name)
	if !ok {
		return 0
	}

	if o.Value == nil {
		return o.Default.(int)
	}

	return o.Value.(int)
}

// GetFloat returns a float option's value.
func (opt *Options) GetFloat(name string) float64 {
	o, ok := opt.GetOption(name)
	if !ok {
		return 0
	}

	if o.Value == nil {
		return o.Default.(float64)
	}

	return o.Value.(float64)
}
