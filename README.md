# Expense-Tracker

A command-line tool build in Go to track your expenses.

## Features

- Add expenses with description and amount

- Update expenses

- Delete expenses

- View all expenses

- View a summary of the expenses

- View a summary of expenses for a specific month

## Getting Started

1. Clone the repository
   
   ```bash
   git clone https://github.com/vic-eco/CLI-expense-tracker.git
   cd CLI-expense-tracker
   ```

2. Build the executable
   
   ```bash
   go build -o expense-tracker.exe .\main.go
   ```

3. Run commands
   
   ```bash
   #See all available commands
   ./expense-tracker.exe --help
   
   #Add an expense
   ./expense-tracker.exe add --description "my_expense" --amount 7.50
   
   #List all expenses
   ./expense-tracker.exe list
   
   #Update an expense
   ./expense-tracker.exe update --id 1 --amount 35
   
   #Delete an expense
   ./expense-tracker.exe delete --id 1
   
   #View summary
   ./expense-tracker.exe summary
   
   ./expense-tracker.exe summary --month 2
   
   
   ```


