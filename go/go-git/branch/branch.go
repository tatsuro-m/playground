package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"log"
)

func main() {
	repo := git.Repository{Storer: memory.NewStorage()}
	branches, err := repo.Branches()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(branches)

	tree, err := repo.Worktree()
	log.Print(tree)
}
