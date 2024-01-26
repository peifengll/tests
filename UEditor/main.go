package main

import "github.com/peifengll/tests/UEditor/router"

func main() {
	s := router.Server()
	s.Run("192.168.0.194:8080")
}
