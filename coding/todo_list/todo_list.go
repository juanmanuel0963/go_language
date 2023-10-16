package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Task represents a to-do task with a name and a completion status.
type Task struct {
	Name string
	Done bool
}

var taskList []Task
var taskFile = "tasks.txt"

func main() {
	// Create a root command for the CLI application.
	var rootCmd = &cobra.Command{Use: "todolist"}

	// Define the "add" command for adding tasks.
	var cmdAdd = &cobra.Command{
		Use:   "add [Task name]",
		Short: "Add task to the list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskName := strings.Join(args, " ")
			addTask(taskName)
		},
	}

	// Define the "list" command for listing tasks.
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List tasks that have not been done",
		Run: func(cmd *cobra.Command, args []string) {
			listTasks()
		},
	}

	// Define the "done" command for marking a task as done.
	var cmdDone = &cobra.Command{
		Use:   "done [Task ID]",
		Short: "Mark a task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Invalid task ID: %v\n", args[0])
			}
			markTaskDone(taskID)
		},
	}

	// Define the "undone" command for marking a task as not done.
	var cmdUndone = &cobra.Command{
		Use:   "undone [Task ID]",
		Short: "Mark a task as not done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Invalid task ID: %v\n", args[0])
			}
			markTaskUndone(taskID)
		},
	}

	// Define the "cleanup" command for removing tasks marked as done.
	var cmdCleanup = &cobra.Command{
		Use:   "cleanup",
		Short: "Remove tasks marked as done",
		Run: func(cmd *cobra.Command, args []string) {
			cleanupDoneTasks()
		},
	}

	// Add all the subcommands to the root command.
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDone)
	rootCmd.AddCommand(cmdUndone)
	rootCmd.AddCommand(cmdCleanup)

	// Load tasks from the file.
	loadTasks()

	// Execute the CLI application.
	rootCmd.Execute()
}

// addTask adds a new task to the task list.
func addTask(name string) {
	task := Task{Name: name, Done: false}
	taskList = append(taskList, task)
	saveTasks()
	fmt.Printf("Added task: %s\n", name)
}

// listTasks lists tasks that have not been marked as done.
func listTasks() {
	//fmt.Println("To-Do List:")
	for i, task := range taskList {
		if !task.Done {
			status := " "
			fmt.Printf("%d: [%s] %s\n", i+1, status, task.Name)
		}
	}
}

// markTaskDone marks a task as done by setting its completion status to true.
func markTaskDone(taskID int) {
	if taskID <= 0 || taskID > len(taskList) {
		log.Fatalf("Invalid task ID: %d\n", taskID)
	}
	taskList[taskID-1].Done = true
	saveTasks()
	fmt.Printf("Marked task %d as done: %s\n", taskID, taskList[taskID-1].Name)
}

// markTaskUndone marks a task as not done by setting its completion status to false.
func markTaskUndone(taskID int) {
	if taskID <= 0 || taskID > len(taskList) {
		log.Fatalf("Invalid task ID: %d\n", taskID)
	}
	taskList[taskID-1].Done = false
	saveTasks()
	fmt.Printf("Marked task %d as not done: %s\n", taskID, taskList[taskID-1].Name)
}

// cleanupDoneTasks removes tasks marked as done from the task list.
func cleanupDoneTasks() {
	undoneTasks := []Task{}
	for _, task := range taskList {
		if !task.Done {
			undoneTasks = append(undoneTasks, task)
		}
	}
	removedCount := len(taskList) - len(undoneTasks)
	taskList = undoneTasks
	saveTasks()
	fmt.Printf("Removed %d done task(s).\n", removedCount)
}

// saveTasks saves the task data to the flat file.
func saveTasks() {
	taskData := []string{}
	for _, task := range taskList {
		taskData = append(taskData, fmt.Sprintf("%t:%s", task.Done, task.Name))
	}
	data := []byte(strings.Join(taskData, "\n"))
	err := ioutil.WriteFile(taskFile, data, 0644)
	if err != nil {
		log.Fatalf("Error saving tasks: %v\n", err)
	}
}

// loadTasks loads task data from the flat file.
func loadTasks() {
	data, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return
	}
	taskData := strings.Split(string(data), "\n")
	for _, entry := range taskData {
		if entry != "" {
			parts := strings.SplitN(entry, ":", 2)
			if len(parts) == 2 {
				done, name := parts[0], parts[1]
				task := Task{Name: name, Done: done == "true"}
				taskList = append(taskList, task)
			}
		}
	}
}

/*
type Task struct {
	Name string
	Done bool
}

var taskList []Task
var taskFile = "tasks.txt"

func main() {
	var rootCmd = &cobra.Command{Use: "todolist"}
	var cmdAdd = &cobra.Command{
		Use:   "add [Task name]",
		Short: "Add task to the list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskName := strings.Join(args, " ")
			addTask(taskName)
		},
	}
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List tasks that have not been done",
		Run: func(cmd *cobra.Command, args []string) {
			listTasks()
		},
	}
	var cmdDone = &cobra.Command{
		Use:   "done [Task ID]",
		Short: "Mark a task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Invalid task ID: %v\n", args[0])
			}
			markTaskDone(taskID)
		},
	}
	var cmdUndone = &cobra.Command{
		Use:   "undone [Task ID]",
		Short: "Mark a task as not done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Invalid task ID: %v\n", args[0])
			}
			markTaskUndone(taskID)
		},
	}
	var cmdCleanup = &cobra.Command{
		Use:   "cleanup",
		Short: "Remove tasks marked as done",
		Run: func(cmd *cobra.Command, args []string) {
			cleanupDoneTasks()
		},
	}

	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDone)
	rootCmd.AddCommand(cmdUndone)
	rootCmd.AddCommand(cmdCleanup)

	// Load tasks from the file
	loadTasks()

	rootCmd.Execute() // Don't change this
}

func addTask(name string) {
	task := Task{Name: name, Done: false}
	taskList = append(taskList, task)
	saveTasks()
	fmt.Printf("Added task: %s\n", name)
}

func listTasks() {
	//fmt.Println("To-Do List:")
	for i, task := range taskList {
		if !task.Done {
			status := " "
			fmt.Printf("%d: [%s] %s\n", i+1, status, task.Name)
		}
	}
}

func markTaskDone(taskID int) {
	if taskID <= 0 || taskID > len(taskList) {
		log.Fatalf("Invalid task ID: %d\n", taskID)
	}
	taskList[taskID-1].Done = true
	saveTasks()
	fmt.Printf("Marked task %d as done: %s\n", taskID, taskList[taskID-1].Name)
}

func markTaskUndone(taskID int) {
	if taskID <= 0 || taskID > len(taskList) {
		log.Fatalf("Invalid task ID: %d\n", taskID)
	}
	taskList[taskID-1].Done = false
	saveTasks()
	fmt.Printf("Marked task %d as not done: %s\n", taskID, taskList[taskID-1].Name)
}

func cleanupDoneTasks() {
	undoneTasks := []Task{}
	for _, task := range taskList {
		if !task.Done {
			undoneTasks = append(undoneTasks, task)
		}
	}
	removedCount := len(taskList) - len(undoneTasks)
	taskList = undoneTasks
	saveTasks()
	fmt.Printf("Removed %d done task(s).\n", removedCount)
}

func saveTasks() {
	taskData := []string{}
	for _, task := range taskList {
		taskData = append(taskData, fmt.Sprintf("%t:%s", task.Done, task.Name))
	}
	data := []byte(strings.Join(taskData, "\n"))
	err := ioutil.WriteFile(taskFile, data, 0644)
	if err != nil {
		log.Fatalf("Error saving tasks: %v\n", err)
	}
}

func loadTasks() {
	data, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return
	}
	taskData := strings.Split(string(data), "\n")
	for _, entry := range taskData {
		if entry != "" {
			parts := strings.SplitN(entry, ":", 2)
			if len(parts) == 2 {
				done, name := parts[0], parts[1]
				task := Task{Name: name, Done: done == "true"}
				taskList = append(taskList, task)
			}
		}
	}
}
*/
