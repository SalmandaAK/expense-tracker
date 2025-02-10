# Expense Tracker

## About
An Expense Tracker CLI made in Go Languange with Cobra CLI and Viper for [roadmap.sh backend project](https://roadmap.sh/projects/expense-tracker).

## Pre-requisite
Go 1.23.0

## Usage
### Clone Repository and Build app
1. Clone this Github repository
   ```
   git clone https://github.com/SalmandaAK/expense-tracker.git
   ```
2. Go to the project directory
   ```
   cd expense-tracker
   ```
3. Run go mod tidy
   ```
   go mod tidy
   ```
4. Build executable file
   ```
   go build
   ```

### Commands
```
# Display all commands
./expense-tracker --help

# Adding a new expense
./expense-tracker add --description "<Expense Description>" --amount <Expense Amount>
# Output: Expense added successfully (ID:1)

# Listing all expenses
./expense-tracker list

# Display summary of all expenses
./expense-tracker summary

# Display summary of all expenses for specific month (of current year)
./expense-tracker summary --month <Month Number>

# Deleting an expense
./expense-tracker delete --id <Expense ID>

# Set currency
./expense-tracker configure --currency <Currency>
```

### Usage Example
![expense-tracker](https://github.com/user-attachments/assets/320eff0d-019a-4e6e-ad42-bd7b5104e621)
