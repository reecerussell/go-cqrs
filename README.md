# CQRS

This is a basic commanding package with interfaces used to command and query data, which aid the separation of concerns. Commands and queries are executed using Executors such as `QueryExecutor` or `CommandExecutor` which enforce every command to be executed the same.
