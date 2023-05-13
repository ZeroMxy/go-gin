package main

import "go-gin/bootstrap"

func main() {
	(&bootstrap.App{}).Run()
}