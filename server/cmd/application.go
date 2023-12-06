package main

import (
	_ "go_archetype/core/config"
	_ "go_archetype/core/storage"
	"go_archetype/core/web"
)

func main() {
	web.Listen()
}
