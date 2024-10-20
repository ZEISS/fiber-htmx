package imports

// Registry is the enum that represents the registry of a package.
type Registry string

const (
	// NPM is the NPM registry.
	NPM Registry = "npm"
	// JSDelivr is the JSDelivr registry.
	JSDelivr Registry = "jsdelivr"
)

// String returns the string representation of the registry.
func (r Registry) String() string {
	return string(r)
}

// ExactPackage is a provider that provides an exact package.
type ExactPackage interface {
	// Name returns the name of the package.
	Name() string
	// Version returns the version of the package.
	Version() string
	// Registry returns the registry of the package.
	Registry() Registry
}

// Provider is the interface that must be implemented by all providers.
type Provider interface {
	// ToURL returns the URL of the provider.
	ToURL(pkg ExactPackage) string
}

var _ Provider = (*UnimplementedProvider)(nil)

// UnimplementedProvider is a provider that is not implemented.
type UnimplementedProvider struct{}

// ToURL returns the URL of the provider.
func (p *UnimplementedProvider) ToURL(pkg ExactPackage) string {
	return ""
}
