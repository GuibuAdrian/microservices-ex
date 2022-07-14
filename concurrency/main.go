package main

import (
	"bufio"
	"fmt"
	"microservices-ex-app/src/api/domain/repositories"
	"microservices-ex-app/src/api/services"
	"microservices-ex-app/src/api/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("C:/Progs/go/microservices-ex-app/concurrency/requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{Name: line}
		result = append(result, request)

	}
	return result
}

func main() {
	requests := getRequests()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10) // We create a channel and limiting the amounts of go routines
	var wg sync.WaitGroup

	go handleResults(&wg, input)

	for _, request := range requests {
		buffer <- true // We make sure that we don't create more than 10 go routines at the same time
		wg.Add(1)
		go createRepo(buffer, input, request)
	}

	wg.Wait()
	close(input)

	//Now we can write success and failed maps to disk or notify them via email.
}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
			continue
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done() //This is done here to process the las output
	}
}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)

	//After sending data to the channel
	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}

	// write: chan <- dataToSend
	// read: data <- readFromBuffer

	// We take an element from the buffer and realising one of the buffer's slots, so we can create another go routine
	<-buffer
}
