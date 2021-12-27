package defer

func checkDefers() error {
	var a *string = ptr.String("sda")
	defer func() {
		log.Print(a)
	}()
	defer log.Print(a)
	
	a = ptr.String("1111111")
	log.Print("создана")
	
	for i := 0; i < 100; i++ {
			defer log.Printf("asd %d", i)
	}

	for i := 0; i < 100; i++ {
		func() {
			defer log.Printf("asd %d", i)
		}()
	}

	for i := 0; i < 100; i++ {
		wrapdefer(i)
	}

	return nil
}

func wrapdefer(i int) {
	defer log.Printf("asd %d", i)
}