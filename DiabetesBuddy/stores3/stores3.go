package stores3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func stores3() {
	// Create an S3 client

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		panic(err)
	}

	client := s3.New(sess)

	// Prompt the user for the glucose reading, meal, and exercise information
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter glucose reading: ")
	glucoseStr, _ := reader.ReadString('\n')
	glucose, err := strconv.ParseFloat(strings.TrimSpace(glucoseStr), 64)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter meal: ")
	meal, _ := reader.ReadString('\n')
	meal = strings.TrimSpace(meal)

	fmt.Print("Did you exercise 2 hours prior to eating? (Y/N): ")
	exercise, _ := reader.ReadString('\n')
	exercise = strings.TrimSpace(exercise)

	exerciseBool := false
	if exercise == "Y" || exercise == "y" {
		exerciseBool = true
	}

	// Store the data in S3
	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("diabetes-buddy-cold-store"),
		Key:    aws.String(fmt.Sprintf("glucose-%d.txt", time.Now().Unix())),
		Body:   strings.NewReader(fmt.Sprintf("%f,%s,%t", glucose, meal, exerciseBool)),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Data stored in S3 successfully!")
}
