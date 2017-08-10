package reuse

// loading modules system in App
type Module interface {
	LoadModule(app Application) bool
}
