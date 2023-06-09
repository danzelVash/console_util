package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type InputParams struct {
	boolFlags map[string]bool
	intFlags  map[string]uint
	data      []string
	input     string
	output    string
}

func (i InputParams) defineProgramBehavior() func([]PreparedForOutput) string {
	f := unique
	if ok, val := i.boolFlags["c"]; ok && val {
		f = strWithNumOfRepeat
	} else if ok, val = i.boolFlags["d"]; ok && val {
		f = repeated
	} else if ok, val = i.boolFlags["u"]; ok && val {
		f = unrepeated
	}

	return f
}

type PreparedForOutput struct {
	str string
	num int
}

func getInput() (InputParams, bool) {
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic happend in getInput:", err)
			}
		}()
	*/
	c := flag.Bool("c", false, "counting the number of lines in the input")
	d := flag.Bool("d", false, "output only those lines that are repeated in the input")
	u := flag.Bool("u", false, "output only those lines that are not repeated in the input")
	i := flag.Bool("i", false, "ignore case")

	f := flag.Uint("f", 0, "ignore the first <num_fields> fields in a row")
	s := flag.Uint("s", 0, "ignore the first <num_chars> characters in a string")

	flag.Parse()

	boolFlags := map[string]bool{
		"c": *c,
		"d": *d,
		"u": *u,
		"i": *i,
	}

	intFlags := map[string]uint{
		"f": *f,
		"s": *s,
	}

	var stdin = InputParams{
		boolFlags: boolFlags,
		intFlags:  intFlags,
	}

	var input io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		stdin.input = filename
		if err != nil {
			panic(err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatalf("Can`t close file.")
			}
		}()

		input = f
	} else {
		input = os.Stdin
		stdin.input = "STDIN"
	}

	if filename := flag.Arg(1); filename != "" {
		stdin.output = filename
	} else {
		stdin.output = "STDOUT"
	}

	fileScanner := bufio.NewScanner(input)
	data := make([]string, 0, 4)

	for i := 0; fileScanner.Scan(); i++ {
		data = append(data, fileScanner.Text())
	}

	stdin.data = data

	return stdin, len(data) == 0
}

func validate(s string, intFlags map[string]uint, boolFlags map[string]bool) (newStr string) {
	arr := strings.SplitN(s, " ", -1)
	if val, ok := intFlags["f"]; ok && val >= uint(len(arr)) {
		newStr = ""
	} else {
		newStr = strings.Join(arr[intFlags["f"]:], " ")
	}
	if val, ok := intFlags["s"]; ok && val > uint(len(newStr)) {
		newStr = ""
	} else {
		newStr = newStr[intFlags["s"]:]
	}

	if val, ok := boolFlags["i"]; ok && val {
		newStr = strings.ToLower(newStr)
	}
	return newStr
}

func preparing(s InputParams) []PreparedForOutput {
	arr := make([]PreparedForOutput, 0)
	el := PreparedForOutput{s.data[0], 1}
	prevValid := validate(s.data[0], s.intFlags, s.boolFlags)
	for i := 0; i < len(s.data)-1; i++ {
		if prevValid == validate(s.data[i+1], s.intFlags, s.boolFlags) {
			el.num++
		} else {
			arr = append(arr, el)
			el = PreparedForOutput{s.data[i+1], 1}
			prevValid = validate(s.data[i+1], s.intFlags, s.boolFlags)
		}
	}

	return append(arr, el)
}

func strWithNumOfRepeat(arr []PreparedForOutput) (result string) {
	for i := 0; i < len(arr); i++ {
		result += strconv.Itoa(arr[i].num) + " " + arr[i].str + "\n"
	}
	return result
}

func repeated(arr []PreparedForOutput) (result string) {
	for i := 0; i < len(arr); i++ {
		if arr[i].num > 1 {
			result += arr[i].str + "\n"
		}
	}
	return result
}

func unrepeated(arr []PreparedForOutput) (result string) {
	for i := 0; i < len(arr); i++ {
		if arr[i].num == 1 {
			result += arr[i].str + "\n"
		}
	}
	return result
}

func unique(arr []PreparedForOutput) (result string) {
	for i := 0; i < len(arr); i++ {
		result += arr[i].str + "\n"
	}
	return result
}

func output(s InputParams, arr []PreparedForOutput) (err error) {
	f := s.defineProgramBehavior()
	str := f(arr)
	if s.output == "STDOUT" {
		fmt.Printf(str)
		err = nil
	} else {
		err = writeInFile(str, s.output)
	}

	return err
}

func writeInFile(str string, file string) (err error) {
	f, err := os.Create(file)

	if err != nil {
		return err
	}

	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Printf("error with closing %s file in writing func: err:\n%s\n", file, err)
		}
	}()

	_, err = f.WriteString(str)

	return err
}

func main() {
	stdParams, ifDataEmpty := getInput()
	if ifDataEmpty {
		fmt.Println("file is empty")
		return
	}

	preparedToDisplay := preparing(stdParams)

	err := output(stdParams, preparedToDisplay)
	if err != nil {
		fmt.Printf("error with writing result in file: err:\n%s", err)
	}
}
