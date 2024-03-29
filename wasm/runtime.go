package wasm

import (
	"fmt"
)

// SetString set the string type arg for wasm
func (w *Wasm) SetString(str string) (int32, error) {
	alloc := w.Instance.GetFunc(w.Store, "allocate")
	if alloc == nil {
		return 0, fmt.Errorf("not found allocate method")
	}
	lengthOfStr := len(str)

	allocResult, err := alloc.Call(w.Store, lengthOfStr+1)
	if err != nil {
		return 0, err
	}
	inputPointer := allocResult.(int32)

	store := w.Instance.GetExport(w.Store, "memory").Memory()
	buf := store.UnsafeData(w.Store)
	memory := buf[inputPointer:]

	var i int
	for i = 0; i < lengthOfStr; i++ {
		memory[i] = str[i]
	}

	memory[i] = 0
	w.context[CONTEXT_ARGMAP].(map[int32]int32)[inputPointer] = int32(len(str))

	return inputPointer, nil
}

// SetBytes set bytes type arg for wasm
func (w *Wasm) SetBytes(b []byte) (int32, error) {
	alloc := w.Instance.GetFunc(w.Store, "allocate")
	if alloc == nil {
		return 0, fmt.Errorf("not found allocate method")
	}
	lengthOfBytes := len(b)

	allocResult, err := alloc.Call(w.Store, lengthOfBytes+1)
	if err != nil {
		return 0, err
	}
	inputPointer := allocResult.(int32)

	store := w.Instance.GetExport(w.Store, "memory").Memory()
	buf := store.UnsafeData(w.Store)
	memory := buf[inputPointer:]

	var i int
	for i = 0; i < lengthOfBytes; i++ {
		memory[i] = b[i]
	}

	memory[i] = 0
	w.context[CONTEXT_ARGMAP].(map[int32]int32)[inputPointer] = int32(len(b))

	return inputPointer, nil
}

// FreeString free the string type arg for wasm
func (w *Wasm) FreeString(inputPointer interface{}, str string) error {
	dealloc := w.Instance.GetFunc(w.Store, "deallocate")
	if dealloc == nil {
		return fmt.Errorf("not found allocate method")
	}
	lengthOfStr := len(str)

	_, err := dealloc.Call(w.Store, inputPointer, lengthOfStr+1)
	if err != nil {
		return err
	}
	delete(w.context[CONTEXT_ARGMAP].(map[int32]int32), inputPointer.(int32))

	return nil
}

// FreeBytes free the bytes type arg for wasm
func (w *Wasm) FreeBytes(inputPointer interface{}, b []byte) error {
	dealloc := w.Instance.GetFunc(w.Store, "deallocate")
	if dealloc == nil {
		return fmt.Errorf("not found allocate method")
	}
	lengthOfBytes := len(b)

	_, err := dealloc.Call(w.Store, inputPointer, lengthOfBytes+1)
	if err != nil {
		return err
	}
	delete(w.context[CONTEXT_ARGMAP].(map[int32]int32), inputPointer.(int32))

	return nil
}
