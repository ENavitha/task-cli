Task Tracker CLI
A simple command-line tool to manage tasks efficiently. Allows users to add, update, delete, and track tasks using a JSON file.

Features
Add new tasks
Update and delete tasks
Mark tasks as to-do, in-progress, or done
List all tasks or filter by status
Stores tasks persistently in a JSON file


Installation
Clone the repository:
git clone https://github.com/ENavitha/task-cli.git


Build the project:
  go build -o task-cli

Add a Task

task-cli add "Buy groceries"

Update a Task
  task-cli update 1 "Buy groceries and cook dinner"

Delete a Task
    task-cli delete 1

Mark Task Status

    task-cli mark-in-progress 1
    task-cli mark-done 1

List Tasks

task-cli list          # All tasks
task-cli list done     # Completed tasks
task-cli list todo     # Pending tasks
task-cli list in-progress  # Tasks in progress
