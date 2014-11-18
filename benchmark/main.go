package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	. "github.com/ernado/gib/models"
)

/*
	Advanced imageboard in golang.
	Goals:
		REST API
		speed & efficency

	/root
		/:board
			/:thread
				/:post
*/

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func randString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func main() {
	var (
		PostSize           = 256
		HashSize           = 12
		AttachmentPathSize = 16
		BumpLimit          = 500
		Boards             = 70
		Pages              = 30
	)
	post := Post{time.Now(), randString(PostSize), randString(HashSize),
		Attachments{Attachment{0, randString(AttachmentPathSize)}},
	}
	threadSize := int(post.Size()) * BumpLimit
	boardSize := threadSize * Pages
	chanSize := Boards * boardSize
	count := BumpLimit * Boards * Pages
	fmt.Printf("With average post %dx%d, in %d pages on %d boards:\n", PostSize, BumpLimit, Pages, Boards)
	fmt.Printf("Posts: %d\n", count)
	fmt.Printf("Size of post: %s\n", ByteSize(post.Size()))
	fmt.Printf("Size of thread: %s\n", ByteSize(threadSize))
	fmt.Printf("Size of board: %s\n", ByteSize(boardSize))
	fmt.Printf("Size of chan: %s\n", ByteSize(chanSize))
	fmt.Print("Generating")
	var (
		posts Posts
		last  float64
		rate  = 0.1
	)
	start := time.Now()
	for i := 0; i < count; i++ {
		p := Post{time.Now(), randString(PostSize), randString(HashSize),
			Attachments{Attachment{0, randString(AttachmentPathSize)}},
		}
		posts.Add(p)
		current := float64(i) / float64(count)
		if (current - last) > rate {
			last = current
			fmt.Print(".")
		}
	}
	fmt.Printf("ok\n")
	end := time.Now()
	duration := end.Sub(start)
	perPost := duration / time.Duration(count)
	fmt.Printf("Completed for %v, %v per post\n", duration, perPost)
	fmt.Scanln()
}
