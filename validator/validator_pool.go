package validator

import (
	"fmt"
	"sync"
)

type ValidatorPools struct {
	pools map[string]*ValidatorPool

	sync.RWMutex
}

func NewValidationPools() *ValidatorPools {
	return &ValidatorPools{
		pools: make(map[string]*ValidatorPool),
	}
}

func (p *ValidatorPools) SetPool(address string, pool *ValidatorPool) {
	p.Lock()
	defer p.Unlock()

	p.pools[address] = pool
}

func (p *ValidatorPools) GetPool(address string) (*ValidatorPool, bool) {
	p.Lock()
	defer p.Unlock()

	pool, ok := p.pools[address]
	return pool, ok
}

type ValidatorPool struct {
	validators chan Validator
	size       int
	length     int

	sync.RWMutex
}

func NewValidationPool(size int) *ValidatorPool {
	return &ValidatorPool{
		validators: make(chan Validator, size),
		size:       size,
		length:     0,
	}
}

func (pool *ValidatorPool) Add(validator Validator) error {
	pool.Lock()
	defer pool.Unlock()

	pool.length += 1
	if pool.length > pool.size {
		return fmt.Errorf("exceed max chan size")
	}
	pool.validators <- validator
	return nil
}

func (pool *ValidatorPool) GetValidator() Validator {
	return <-pool.validators
}

func (pool *ValidatorPool) SetValidator(vlt Validator) {
	pool.validators <- vlt
}
