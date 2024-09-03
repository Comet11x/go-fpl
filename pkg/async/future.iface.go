package async

type Future[T any] interface {
	// Extends Awaiter
	Awaiter[T]

	// Returns Promise[T]
	Promise() Promise[T]
}
