package harbor

//go:generate mockery -all -case=underscore

// ClientAPI describes a client API.
type ClientAPI interface {
	// ProfileGet returns a profile.
	ProfileGet() (*Profile, error)

	// ProfilePatch updates a profile.
	ProfilePatch(*Profile) (*Profile, error)
}
