package models

import (
	"time"
	"unsafe"
)

type Attachment struct {
	Size int64
	Path string
}

type Attachments []Attachment

type Thread struct {
	Posts
}

type Post struct {
	Time        time.Time
	Body        string
	Hash        string
	Attachments Attachments
}

func (a Attachment) Sizeof() (size uintptr) {
	size += unsafe.Sizeof(a)
	size += uintptr(len(a.Path))
	return
}

func (a Attachments) Size() (size uintptr) {
	for _, attachment := range a {
		size += attachment.Sizeof()
	}
	return
}

func (p Post) Size() (size uintptr) {
	size += unsafe.Sizeof(p)
	size += uintptr(len(p.Body))
	size += uintptr(len(p.Hash))
	size += p.Attachments.Size()
	return
}

type Posts []Post

func (p *Posts) Add(post Post) {
	*p = append(*p, post)
}

func (p Posts) Len() int {
	return len(p)
}
