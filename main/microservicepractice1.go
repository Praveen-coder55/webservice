package main

// func main() {

// 	l := log.New(os.Stdout, "product-api", log.LstdFlags)

// 	// hh := handlers.NewHello(l)
// 	// gg := handlers.NewGoodbye(l)

// 	sl := handlers.NewStudents(l)
// 	sm := http.NewServeMux()
// 	sm.Handle("/", sl)
// 	//sm.Handle("/goodbye", gg)

// 	s := &http.Server{
// 		Addr:         ":9090",
// 		Handler:      sm,
// 		IdleTimeout:  120 * time.Second,
// 		ReadTimeout:  1 * time.Second,
// 		WriteTimeout: 1 * time.Second,
// 	}

// 	go func() {
// 		err := s.ListenAndServe()
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 	}()

// 	sigChan := make(chan os.Signal)
// 	signal.Notify(sigChan, os.Interrupt)
// 	signal.Notify(sigChan, os.Kill)

// 	sig := <-sigChan
// 	l.Println("Received end request , succesfully terminate", sig)

// 	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	s.Shutdown(tc)
// }
