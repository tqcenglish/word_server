package web

import "embed"

//go:embed www/*
var WWWStatic embed.FS

//go:embed apidoc/*
var ApiDocStatic embed.FS
