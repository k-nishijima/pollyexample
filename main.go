package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

func main() {
	cfg := aws.NewConfig().WithRegion("us-east-1")
	sess := session.Must(session.NewSession(cfg))

	svc := polly.New(sess)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "quit" {
			break
		}

		params := &polly.SynthesizeSpeechInput{
			OutputFormat: aws.String("mp3"),
			Text:         aws.String(input),
			VoiceId:      aws.String("Mizuki"),
		}
		resp, err := svc.SynthesizeSpeech(params)
		if err != nil {
			panic(err.Error)
		}

		content, err := ioutil.ReadAll(resp.AudioStream)
		ioutil.WriteFile("/tmp/gopolly.mp3", content, os.ModePerm)

		exerr := exec.Command("afplay", "/tmp/gopolly.mp3").Run()
		if exerr != nil {
			panic(exerr.Error)
		}
	}

}
