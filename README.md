A simple command-line diary written in Go. Allows you to add, read, and clear entries saved in the diary.txt file.
🔧 Commands
Command	Description
write	Add a new diary entry
read	Display all diary entries
clear	Clear all entries from diary
exit	Exit the program
📦 Example Usage

go run main.go write
go run main.go read
go run main.go clear

💾 Where are entries stored?

All entries are saved in a local file called diary.txt.
⏰ Entry Format

Each entry includes a timestamp:

2025-07-25 14:33:12 — Learned how to write to a file in Go.
