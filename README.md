Pull requests are welcome

 basically I realized another simple sequence of doing it:

func main() {
	password := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	rand.Read(password)

	fmt.Printf("password: %x", password)
}

This sequence uses the math/rand package which generates random numbers
The read function provided by the rand package is used to fill the slice with values

