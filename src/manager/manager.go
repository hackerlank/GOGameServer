// manager project manager.go
package manager

type Manager interface {
	Manage(item interface{}) bool
}
