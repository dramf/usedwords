package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/dramf/usedwords/movie"
	"github.com/dramf/usedwords/stat"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	defaultBaseName = "usedwords"
)

type Config struct {
	movie         string
	basename      string
	ignoreNumeric bool
	vocabulary    string
}

func parseFlags(progname string, args []string) (config *Config, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	conf := &Config{}
	flags.StringVar(&conf.movie, "movie", "", "a link to a movie for processing")
	flags.StringVar(&conf.basename, "base", defaultBaseName, "a base name for output files")
	flags.StringVar(&conf.vocabulary, "vocabulary", "", "a vocablary of known words")
	flags.BoolVar(&conf.ignoreNumeric, "numignore", true, "ignore numerics")
	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	return conf, buf.String(), nil
}

func processingMovie(conf *Config) {
	checkError := func(err error, title string) {
		if err != nil {
			log.Fatalf("%s: %v", title, err)
		}
	}

	log.Printf("Movie link processing: %q\n", conf.movie)
	mv, err := movie.ParseLink(conf.movie)
	checkError(err, "download error")

	data, err := mv.DownloadCaptions()
	checkError(err, "parse link error")

	filename := fmt.Sprintf("%s.sub", conf.basename)
	log.Printf("Writing a received data to %q\n", filename)
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	checkError(err, "Writing data error")

	st, err := stat.InitWordStat(data, conf.ignoreNumeric)
	checkError(err, "InitWordStat error")

	filewords := fmt.Sprintf("%s.words", conf.basename)
	log.Printf("Writeting unique words to %q\n", filewords)
	words := strings.Join(st.Words(), "\n")
	err = ioutil.WriteFile(filewords, []byte(words), 0644)
	checkError(err, "writing words to .words file error")

	if conf.vocabulary != "" {
		log.Printf("Usedword uses a vocabulary %q\n", conf.vocabulary)

		v, err := ioutil.ReadFile(conf.vocabulary)
		checkError(err, "read vocabulary error")

		oldWords := strings.Split(string(v), "\n")
		newWords := st.NewWords(oldWords)

		newwordsfile := fmt.Sprintf("%s.new", conf.basename)
		log.Printf("Writeting new words to %q\n", newwordsfile)

		words := strings.Join(newWords, "\n")
		err = ioutil.WriteFile(newwordsfile, []byte(words), 0644)
		checkError(err, "writing new words to .new.words file error")
	}

	filestat := fmt.Sprintf("%s.stat", conf.basename)
	log.Printf("Writeting a statistic to %q\n", filestat)
	err = ioutil.WriteFile(filestat, []byte(st.String()), 0644)
	checkError(err, "writing stats to .stat file error")
}

func main() {
	log.Println("The Used Words App")
	conf, output, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(output)
		os.Exit(2)
	} else if err != nil {
		log.Print("Error:", err, output)
		os.Exit(1)
	}
	if conf.movie != "" {
		processingMovie(conf)
	} else {
		log.Print("The link to a movie for processing is empty")
	}
}
