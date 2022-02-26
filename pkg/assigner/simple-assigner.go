package assigner

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type SimplerAssigner struct {
}

func NewSimplerAssigner() SimplerAssigner {
	return SimplerAssigner{}
}

func (sa SimplerAssigner) Assign(items []string) string {
	position := rand.Intn(len(items))
	return items[position]
}
