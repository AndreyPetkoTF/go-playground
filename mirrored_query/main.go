package main

func main() {

}

func mirroredQuery() string {
	responses := make(chan string, 3)

	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()

	return <-responses
}

func request(hostname string) string {
	return ""
}
