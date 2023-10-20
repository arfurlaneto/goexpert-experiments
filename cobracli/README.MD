More info about cobra at https://github.com/spf13/cobra.


```bash
# install cobra-cli
go install github.com/spf13/cobra-cli@latest

# init project
cobra-cli init

# add a new command
cobra-cli add ping

# add a new nested command
cobra-cli add pong -p "pingCmd"
```

