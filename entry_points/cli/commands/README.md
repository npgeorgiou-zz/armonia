# Here are the CLI Commands

They are conveniently broken each in its own file, to resemble the familiar "One file per command" convention of frameworks.

Notice that each command satisfies the DiAwareCommand Type. We like consistency :)

Some commands wrap a business case, while others can just do their own thing.

The commands that wrap a business case can also just embed the case directly, but to maintain symmetry 
with the http adapters, cli adapters have been created to wrap Business Cases.
 