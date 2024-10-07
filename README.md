# Task Tracker CLI

A command-line interface (CLI) tool to track and manage tasks. The tool allows you to add, update, delete, mark tasks as in-progress or done, and list tasks based on their status. The sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Installation

To use the Task Tracker CLI, you need to have Go installed on your machine.

**Clone the Repository:**
```bash
git clone https://github.com/fhasnur/task-tracker.git
```

**Navigate to the project directory:**
```bash
cd task-tracker
```

**Build the CLI:**
```bash
go build -o task-cli
```

## Usage

Once built, you can run the CLI tool from your terminal. Below are the basic commands:

### Adding New Task
To add a new task, run:
```bash
./task-cli add [description]
```
Example:
```bash
./task-cli add "Buy groceries"
#Output: Task added successfully (ID: 1)
```

### Updating a Task
To update an existing task, use:
```bash
./task-cli update [id] [new description]
```
Example:
```bash
./task-cli update 1 "Buy groceries and cook dinner"
#Output: Task updated successfully
```

### Deleting a Task
To delete a task, run:
```bash
./task-cli delete [id]
```
Example:
```bash
./task-cli delete 1
#Output: Task deleted successfully
```

### Marking a Task as In-Progress
To mark a task as in-progress, use:
```bash
./task-cli mark-in-progress [id]
```
Example:
```bash
./task-cli mark-in-progress 1
#Output: Task marked as in-progress successfully
```

### Marking a Task as Done
To mark a task as done, run:
```bash
./task-cli mark-done [id]
```
Example:
```bash
./task-cli mark-done 1
#Output: Task marked as done successfully
```

### Listing All Tasks
To list all tasks, use:
```bash
./task-cli list
```
Example Output:
```bash
ID  | Description                    | Status       | Created At       | Updated At
------------------------------------------------------------------------------------------
1   | Buy groceries                  | in-progress  | 2024-09-10 22:01 | 2024-09-12 16:45
3   | Read a book                    | done         | 2024-09-11 23:06 | 2024-09-11 23:25
4   | Drink a juice                  | todo         | 2024-09-11 23:06 | 2024-09-11 23:06
5   | Learn coding                   | done         | 2024-09-12 16:56 | 2024-09-14 22:43
6   | Play snake game                | in-progress  | 2024-09-14 22:39 | 2024-09-14 22:42
```

### Listing Tasks by Status
To filter tasks by status (e.g., todo, in-progress or done), run:

**Todo**
```bash
./task-cli list todo
```
Example Output:
```bash
ID  | Description                    | Status       | Created At       | Updated At
------------------------------------------------------------------------------------------
4   | Drink a juice                  | todo         | 2024-09-11 23:06 | 2024-09-11 23:06
```

**In Progress**
```bash
./task-cli list in-progress
```
Example Output:
```bash
ID  | Description                    | Status       | Created At       | Updated At
------------------------------------------------------------------------------------------
1   | Buy groceries                  | in-progress  | 2024-09-10 22:01 | 2024-09-12 16:45
6   | Play snake game                | in-progress  | 2024-09-14 22:39 | 2024-09-14 22:42
```

**Done**
```bash
./task-cli list done
```
Example Output:
```bash
ID  | Description                    | Status       | Created At       | Updated At
------------------------------------------------------------------------------------------
3   | Read a book                    | done         | 2024-09-11 23:06 | 2024-09-11 23:25
5   | Learn coding                   | done         | 2024-09-12 16:56 | 2024-09-14 22:43
```

## Contributing

Feel free to submit pull requests or open issues for new features, improvements, or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.