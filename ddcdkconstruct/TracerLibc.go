package ddcdkconstruct


// C standard library of the application image.
type TracerLibc string

const (
	// Used by most Linux images, such as Debian and Ubuntu.
	TracerLibc_GLIBC TracerLibc = "GLIBC"
	// Used by Alpine Linux images.
	TracerLibc_MUSL TracerLibc = "MUSL"
)

