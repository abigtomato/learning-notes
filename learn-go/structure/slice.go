package main

/*
#include <stdlib.h>
 */
import "C"
import (
	"fmt"
	"unsafe"
)

const TAG = 8

type Slice struct {
	Data	unsafe.Pointer
	Len		int
	Cap		int
}

func (this *Slice) Create(l, c int, data ...int) {
	if len(data) == 0 {
		return
	}

	if l < 0 || c < 0 || l > c || len(data) > c {
		return
	}

	this.Len = l
	this.Cap = c
	this.Data = C.malloc(C.ulonglong(c) * 8)

	point := uintptr(this.Data)
	for _, val := range data {
		*(*int)(unsafe.Pointer(point)) = val
		point += TAG
	}
}

func (this *Slice) Print()  {
	point := uintptr(this.Data)
	for i := 0; i < this.Len; i++ {
		fmt.Println(*(*int)(unsafe.Pointer(point)))
		point += TAG
	}
}

func (this *Slice) Append(data ...int) error {
	if this.Data == nil {
		return fmt.Errorf("slice未被创建")
	}

	if this.Len + len(data) > this.Cap {
		this.Data = C.realloc(this.Data, C.ulonglong(this.Cap) * 2 * 8)
		this.Cap = this.Cap * 2
	}

	point := uintptr(this.Data)
	for i := 0; i < this.Len; i++ {
		point += TAG
	}

	for _, val := range data {
		*(*int)(unsafe.Pointer(point)) = val
		point += TAG
	}

	this.Len = this.Len + len(data)
	return nil
}

func (this *Slice) Get(index int) (data int, err error) {
	if this.Data == nil {
		err = fmt.Errorf("slice未被创建")
		return
	}

	if index < 0 || index > this.Len - 1 {
		err = fmt.Errorf("下标越界")
		return
	}

	point := uintptr(this.Data)
	for i := 0; i < index; i++ {
		point += TAG
	}

	data = *(*int)(unsafe.Pointer(point))
	return
}

func (this *Slice) Search(data int) (index int, err error) {
	if this.Data == nil {
		err = fmt.Errorf("slice未被创建")
		return
	}

	point := uintptr(this.Data)
	for i := 0; i < this.Len; i++ {
		if *(*int)(unsafe.Pointer(point)) == data {
			data = i
			return
		}
		point += TAG
	}

	err = fmt.Errorf("未找到符合条件的数据")
	return
}

func (this *Slice) Delete(index int) error {
	if this.Data == nil {
		return fmt.Errorf("slice未被创建")
	}

	if index < 0 || index > this.Len - 1 {
		return fmt.Errorf("下标越界")
	}

	point := uintptr(this.Data)
	for i := 0; i < index; i++ {
		point += TAG
	}

	temp := point
	for i := index; i < this.Len; i++ {
		temp += TAG
		*(*int)(unsafe.Pointer(point)) = *(*int)(unsafe.Pointer(temp))
		point += TAG
	}

	this.Len--
	return nil
}

func (this *Slice) Insert(index int, data int) error {
	if this.Data == nil {
		return fmt.Errorf("slice未被创建")
	}

	if index < 0 || index > this.Len - 1 {
		return fmt.Errorf("下标越界")
	}

	if index == this.Len {
		if err := this.Append(data); err != nil {
			fmt.Println(err.Error())
		}
		return nil
	}

	point := uintptr(this.Data)
	for i := 0; i < this.Len; i++ {
		point += TAG
	}

	first := uintptr(this.Data) + TAG * uintptr(this.Len)
	for i := this.Len; i > index; i-- {
		*(*int)(unsafe.Pointer(first)) = *(*int)(unsafe.Pointer(first - TAG))
		first -= TAG
	}

	*(*int)(unsafe.Pointer(point)) = data
	this.Len++
	return nil
}

func (this *Slice) Destroy() error {
	if this.Data == nil {
		return fmt.Errorf("slice未被创建")
	}

	C.free(this.Data)
	this.Data = nil
	this.Len = 0
	this.Cap = 0
	this = nil

	return nil
}
