# ssh-keyword

A CLI tool for managing SSH connections via keywords.

## Example

Your ssh host is 192.168.1.250.

```
$ ssh-keyword -a 192.168.1.250
Entry for new connection
Enter 'quit' for exit

Enter a user: john
Port: 22
Enter a list of names separate by ',': server, john
Default server ([Y]es | [N]o): no

$ ssh-keyword -ls
ip:192.168.1.250  user:john  port:22  keywords:['server', 'john']  default:False

$ ssh-keyword server
```

You are now connected to 192.168.1.250 !

## Installation

### Windows

To install on Windows, download the latest `.exe` file from the GitHub releases and add the directory containing `ssh-keyword.exe` to your environment path to access it from anywhere.

### Linux

To install `ssh-keyword` on Linux, follow these steps:

1. Download the latest release from the GitHub repository:
```
wget https://github.com/befabri/ssh-keyword/releases/download/v1.2.2/ssh-keyword-linux-amd64
```
2. Make the executable accessible:
```
chmod +x ssh-keyword-linux-amd64
```
3. Move the executable to a bin directory:
```
sudo mv ssh-keyword-linux-amd64 /usr/local/bin/ssh-keyword
```
4. Verify the installation by checking the version or help:
```
ssh-keyword -h
```

## Documentation

First add your connection with `ssh-keyword -a [IP]`.\
Now you can connect to your ssh host by typing `ssh-keyword [YOUR_KEYWORD]` or just `ssh-keyword` if you have set a default connection.

## Help

```
$ ssh-keyword -h

Usage: ssh-keyword [keyword]
       ssh-keyword [options] [command]

Options:
  -a,  --add [IP]                   Add a connection using the specified IP address.
  -d,  --default [IP]               Set the specified IP as the default connection.
  -rm, --remove [IP]                Remove the connection with the specified IP.
  -ls, --list [IP]                  List all connections or a specific connection by IP.
  -e,  --edit [IP]                  Edit the connection with the specified IP.
  -h,  --help                       Show this help message and exit.

Examples:
  ssh-keyword server                 Connects directly to the connection labeled 'server'.
  ssh-keyword --add 192.168.1.1      Add a connection for 192.168.1.1.
  ssh-keyword --remove 192.168.1.1   Remove the connection for 192.168.1.1.
  ssh-keyword --list                 List all connections.
  ssh-keyword --edit 192.168.1.1     Edit the connection for 192.168.1.1.
  ssh-keyword --help                 Show the help message.
```
