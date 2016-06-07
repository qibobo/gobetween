/**
 * roundrobin.go - roundrobin balance impl
 *
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 */

package balance

import (
	"../core"
	"errors"
)

/**
 * Roundrobin balancer
 */
type RoundrobinBalancer struct {

	/* Current backend position */
	current int
}

/**
 * Elect backend using roundrobin strategy
 */
func (b *RoundrobinBalancer) ElectBackend(context *core.Context, backends []core.Backend) (*core.Backend, error) {

	if len(backends) == 0 {
		return nil, errors.New("Can't elect backend, Backends empty")
	}

	if b.current >= len(backends) {
		b.current = 0
	}

	backend := &backends[b.current]
	b.current += 1

	return backend, nil
}
