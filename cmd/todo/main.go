package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Denuwan-Kalubowila/todo-cli"
)

const (
	todofile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add new todo")
	completed := flag.Int("completed", 0, "Complete todo")
	delete := flag.Int("delete", 0, "Delete todo")
	list := flag.Bool("list", true, "list todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todofile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.AddTodo(task)
		err = todos.Store(todofile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *completed > 0:
		todos.CompleteTodo(*completed)
		err := todos.Store(todofile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *delete > 0:
		todos.DeleteTodo(*delete)
		err := todos.Store(todofile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.PrintTodo()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(scanner.Text()) == 0 {
		return "", errors.New("empty todo in not allowed.")
	}
	return scanner.Text(), nil
}
