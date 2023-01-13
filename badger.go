package jpostcode

import (
	"embed"
)

//go:embed data/*
var staticFS embed.FS
