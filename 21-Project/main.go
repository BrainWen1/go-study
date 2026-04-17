package main

func main() {
	ser := NewServer("127.0.0.1", 8080)
	ser.Start()
}
