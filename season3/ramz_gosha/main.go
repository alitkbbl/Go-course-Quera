package main

func StartDecipher(senderChan chan string, decipherer func(encrypted string) string) chan string {
	c := make(chan string, 5)

	go func() {
		for encrypted := range senderChan {
			decoded := decipherer(encrypted)
			c <- decoded
		}
		close(c)
	}()

	return c
}
