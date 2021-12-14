package wasm

import (
	"fmt"
)

// SetString set the string type arg for wasm
func (w *Wasm) SetString(str string) (int32, error) {
	alloc := w.Instance.Exports["allocate"]
	if alloc == nil {
		return 0, fmt.Errorf("not found allocate method")
	}
	lengthOfStr := len(str)

	allocResult, err := alloc(lengthOfStr + 1)
	if err != nil {
		return 0, err
	}
	inputPointer := allocResult.ToI32()

	memory := w.Instance.Memory.Data()[inputPointer:]

	var i int
	for i = 0; i < lengthOfStr; i++ {
		memory[i] = str[i]
	}

	memory[i] = 0
	w.argMap[int(inputPointer)] = len(str)

	return inputPointer, nil
}

// SetBytes set bytes type arg for wasm
func (w *Wasm) SetBytes(b []byte) (int32, error) {
	alloc := w.Instance.Exports["allocate"]
	if alloc == nil {
		return 0, fmt.Errorf("not found allocate method")
	}
	lengthOfBytes := len(b)

	allocResult, err := alloc(lengthOfBytes + 1)
	if err != nil {
		return 0, err
	}
	inputPointer := allocResult.ToI32()

	memory := w.Instance.Memory.Data()[inputPointer:]

	var i int
	for i = 0; i < lengthOfBytes; i++ {
		memory[i] = b[i]
	}

	memory[i] = 0
	w.argMap[int(inputPointer)] = len(b)

	return inputPointer, nil
}

// FreeString free the string type arg for wasm
func (w *Wasm) FreeString(inputPointer interface{}, str string) error {
	dealloc := w.Instance.Exports["deallocate"]
	if dealloc == nil {
		return fmt.Errorf("not found allocate method")
	}
	lengthOfStr := len(str)

	_, err := dealloc(inputPointer, lengthOfStr+1)
	if err != nil {
		return err
	}
	delete(w.argMap, int(inputPointer.(int32)))

	return nil
}

// FreeBytes free the bytes type arg for wasm
func (w *Wasm) FreeBytes(inputPointer interface{}, b []byte) error {
	dealloc := w.Instance.Exports["deallocate"]
	if dealloc == nil {
		return fmt.Errorf("not found allocate method")
	}
	lengthOfBytes := len(b)

	_, err := dealloc(inputPointer, lengthOfBytes+1)
	if err != nil {
		return err
	}
	delete(w.argMap, int(inputPointer.(int32)))

	return nil
}
