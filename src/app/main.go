package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"movie"
	"os"
)

type Config struct {
	movie string
}

func parseFlags(progname string, args []string) (config *Config, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	conf := &Config{}
	flags.StringVar(&conf.movie, "movie", "", "a link to a movie for processing")
	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	return conf, buf.String(), nil
}

func processingMovie(link string) {
	fmt.Printf("Movie link processing: %q\n", link)
	mv, err := movie.ParseLink(link)
	if err != nil {
		fmt.Printf("parse link error: %v", err)
		return
	}
	data, err := mv.DownloadCaptions()
	if err != nil {
		fmt.Printf("download error: %v", err)
		return
	}
	filename := "captions.sub"

	fmt.Printf("Writing a received data to %q\n", filename)
	if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		fmt.Printf("Writing file error: %v", err)
	}
}

func main() {
	fmt.Println("The Used Words App")
	conf, output, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(output)
		os.Exit(1)
	}
	if conf.movie != "" {
		processingMovie(conf.movie)
	}
}