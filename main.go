package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Userid    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Queue struct {
	UUID    string `json:"uuid"`
	Maxsize uint   `json:"maxsize"`
	Users   []User `json:"users"`
}

func handler(c *gin.Context) {
	var queue Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	processQueue := process(&queue)
	c.JSON(http.StatusOK, gin.H{
		"message": "Queue processed successfully",
		"queue":   processQueue,
	})
}

func process(queue *Queue) *Queue {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(queue.Users))
	for i := range queue.Users {
		go func(i int) {
			defer wg.Done()
			time.Sleep(1000 * time.Millisecond)
			mu.Lock()
			queue.Users[i].FirstName = "Processed-" + queue.Users[i].FirstName
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return queue
}

// func router01() http.Handler {
// 	e := gin.Default()
// 	e.POST("/queue", handler)
// 	return e
// }

// func router02() http.Handler {
// 	e := gin.Default()
// 	e.POST("/queue", handler)
// 	return e
// }

// func router03() http.Handler {
// 	e := gin.Default()
// 	e.POST("/queue", handler)
// 	return e
// }

// var g errgroup.Group

func main() {
	r := gin.Default()
	r.POST("/queue", handler)
	r.Run(":8199")

	// server01 := &http.Server{
	// 	Addr:    ":8199",
	// 	Handler: router01(),
	// 	// ReadTimeout:  5 * time.Second,
	// 	// WriteTimeout: 10 * time.Second,
	// }

	// server02 := &http.Server{
	// 	Addr:    ":8200",
	// 	Handler: router02(),
	// 	// ReadTimeout:  5 * time.Second,
	// 	// WriteTimeout: 10 * time.Second,
	// }

	// server03 := &http.Server{
	// 	Addr:    ":8201",
	// 	Handler: router03(),
	// 	// ReadTimeout:  5 * time.Second,
	// 	// WriteTimeout: 10 * time.Second,
	// }

	// g.Go(func() error {
	// 	return server01.ListenAndServe()
	// })
	// g.Go(func() error {
	// 	return server02.ListenAndServe()
	// })
	// g.Go(func() error {
	// 	return server03.ListenAndServe()
	// })

	// if err := g.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
}
