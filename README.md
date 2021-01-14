# Task Manager

My attempt at the [CLI Task Manager Gophercise](https://github.com/gophercises/task)

## Instructions

To install: `go install github.com/leonseng/taskmanager`

Once installed, run `taskmanager -h` to see supported actions.

Some examples:

```
$ taskmanager add wash car
Added "wash car" to your task list.

$ taskmanager add feed dog
Added "feed dog" to your task list.

$ taskmanager list
You have the following tasks:
1. wash car
2. feed dog

$ taskmanager do 1
You have completed the "wash car" task.

$ taskmanager list
You have the following tasks:
1. some task description
```

## References
- [Using Cobra](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177)
- [boltDB repo](https://github.com/boltdb/bolt)
- [boltDB documentation](https://godoc.org/github.com/boltdb/bolt)
