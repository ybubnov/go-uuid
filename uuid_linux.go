package uuid

func init() {
	// For Linux operating system we will use the kernel UUID v4
	// generator in order to use it a default back-end for funcs
	// UUID and MustUUID.
	defaultSrc = &Kernel{MaxIdle: 128}
}
