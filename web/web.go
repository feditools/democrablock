package web

import "embed"

// Files contains static files required by the application
//go:embed template/*
var Files embed.FS
