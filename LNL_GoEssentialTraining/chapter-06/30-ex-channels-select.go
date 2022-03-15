package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func parseLogFile(filePath string) (map[string]string, error){

 	file, err := os.Open(filePath)
 	if err != nil {
 		return nil, err
	}
	defer file.Close()

 	checkSums := make(map[string]string)
 	scanner := bufio.NewScanner(file)
 	for lineNum := 1; scanner.Scan(); lineNum++ {
 		parts := strings.Fields(scanner.Text())

 		if len(parts) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", filePath, lineNum)
		}
		checkSums[parts[1]] = parts[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return checkSums, nil
}

func getChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

type Md5MatchResult struct {
	filePath string
	match bool
	err error
}

func md5Worker(filePath string, providedCheckSum string, out chan *Md5MatchResult) {
	matchResult := &Md5MatchResult{filePath: filePath}
	calculatedChecksum, err := getChecksum(filePath)

	if err != nil {
		matchResult.err = err
		out <- matchResult
		return
	}

	matchResult.match = calculatedChecksum == providedCheckSum
	out <- matchResult
}

func main() {

	checkSums, err := parseLogFile("nasa-logs/md5sum.txt")
	if err != nil {
		log.Fatalf("err: can't read a md5sum file - %s", err)
	}

 	outMd5checkSums := make(chan *Md5MatchResult)
 	for localFileName, localMd5Sum := range checkSums {
 		localFilePath := path.Join("nasa-logs", localFileName)
 		go md5Worker(localFilePath, localMd5Sum, outMd5checkSums)
	}

	allValid := false

	for range checkSums {
		validationResult := <-outMd5checkSums

		switch {
		case validationResult.err != nil:
			allValid = false
			log.Fatalf("%s err: %s\n", validationResult.filePath, validationResult.err)
		case !validationResult.match:
			allValid = false
			log.Fatalf("%s signature mismatch\n", validationResult.filePath)
		default:
			allValid = true
		}
	}

	if !allValid {
		fmt.Println("Mismatch founded")
		os.Exit(1)
	}

	fmt.Println("All checksums valid")
}