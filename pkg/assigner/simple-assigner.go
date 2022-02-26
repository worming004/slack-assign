package assigner

import (
	"math/rand"
	"time"
)

type SimplerAssigner struct {
	usersToRemove map[string]struct{}
}

func NewSimplerAssigner(usersToRemove []string) SimplerAssigner {
	excluded := make(map[string]struct{})
	for _, u := range usersToRemove {
		excluded[u] = struct{}{}
	}
	return SimplerAssigner{excluded}
}

func (sa SimplerAssigner) Assign(items []string) string {
	withoutExcluded := withoutExcluded(items, sa.usersToRemove)

	rand.Seed(time.Now().Unix())
	position := rand.Int() % len(withoutExcluded)

	return withoutExcluded[position]
}

func withoutExcluded(items []string, usersToRemove map[string]struct{}) []string {
	var withoutExcluded []string
	for _, i := range items {
		if _, ok := usersToRemove[i]; !ok {
			withoutExcluded = append(withoutExcluded, i)
		}
	}
	return withoutExcluded
}
