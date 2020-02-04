# ssh-keyword
A keywords ssh connection.


Example
-------------

Your ssh host is 192.168.1.250.

```
ssh_keyword -a 192.168.1.250
Entry for new connection
Enter 'Quit' for exit

Enter a user: john
Port: 22
Enter a list of name separate by ',': server, john
Default server ([Y]es | [N]o): no

python ssh_keyword.py -ls
ip:192.168.1.250  user:jhon  port:22  keywords:['server', 'john']  default:False

python ssh_keyword.py server
```
You are now connected to 192.168.1.250 !


Installation
------------

For installation you can clone this repo.


If you are on **Linux**, you can add **ssh_keyword** to `$HOME/.bashrc` to access it anywhere.
You can also specify a shortcut name there.

Change SHORTCUTNAME with the desired shortcut name and the path of ssh_keyword.
You need to restart your terminal to take effect.

Run in terminal :
```
printf 'SHORTCUTNAME() {\n    python PATH/ssh-keyword/ssh_keyword.py "$1" "$2"\n}' >> $HOME/.bashrc
```

If you are on **Windows**, you can add the directory of **ssh_keyword** to your env path to access it anywhere. 


Documentation
-------------

First add your connection with `ssh_keyword -a [IP]`. 
Now you can connect to your ssh host by typing `ssh_keyword [YOURKEYWORD]`
or just `ssh_keyword` if you have set a default connection.


Help
----

```
#ssh_keyword -h

Usage: ssh_keyword [OPTIONS option] or ssh_keyword [KEYWORD]
Keyword recognition in ssh command
Specify a keyword in list of keywords of the connection

Optional arguments:
-a      --add        add a new connection (ssh_keyword -a [IP])
-d      --default    add/change default connection (ssh_keyword -d [IP])
-rm     --remove     remove connection (ssh_keyword -rm [IP])
-ls     --list       list connection (ssh_keyword -ls or ssh_keyword -ls [IP])
-e      --edit       edit connection (ssh_keyword -e [IP])
-h      --help       show this help message and exit
```
