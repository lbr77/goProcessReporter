## ProcessReporter with go

used for [Shiro](https://github.com/Innei/Shiro)

## feature

1.Active window detect.
2.Check now playing netease music by reading logs. (provided by https://egg.moe/2020/07/get-netease-cloudmusic-playing/)

## usage

```bash
goProcessreporter
A Process reporter used by Shiro

Usage:
  goProcessReporter [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  start        Start service
  start-daemon Start in background
  stop-daemon  Stop in background
  version      Show version

Flags:
  -h, --help   help for goProcessReporter

Use "goProcessReporter [command] --help" for more information about a command.
```

set the config file and run `goProcessReporter start`

## todo

[]support all media platform by ("SMTC protocal")
