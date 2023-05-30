Pull requests are welcome

 basically I realized another simple sequence of doing it:

func main() {
	password := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	rand.Read(password)

	fmt.Printf("password: %x", password)
}
