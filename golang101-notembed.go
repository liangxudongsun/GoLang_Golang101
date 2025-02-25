//go:build !go1.16
// +build !go1.16

package main

import "html/template"

var staticFilesHandler = staticFilesHandler_NonEmbedding
var resFilesHandler = resFilesHandler_NonEmbedding

func loadArticleFile(file string) ([]byte, error) {
	return loadArticleFile_NonEmbedding(file)
}

func parseTemplate(commonPaths []string, files ...string) *template.Template {
	return parseTemplate_NonEmbedding(commonPaths, files...)
}

func updateGolang101() {
	updateGolang101_NonEmbedding()
}
