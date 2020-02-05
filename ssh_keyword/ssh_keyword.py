import sys
from ssh_keyword_connection import Connection
from ssh_keyword_tools import isIp
from ssh_keyword_json import ManageJson
from subprocess import call

def main(args, value=None):
    if args in ['-a', '--add'] and value:
        if isIp(value):
            connection = Connection(value)
            connection.setConnection()
        else:
            print('Invalid')
    elif args in ['-d', '--default'] and value:
        if isIp(value):
            connection = Connection(value)
            connection.addDefault()
        else:
            print('Invalid')
    elif args in ['-rm', '--remove'] and value:
        if isIp(value):
            connection = Connection(value)
            connection.remove()
        else:
            print('Invalid')
    elif args in ['-ls', '--list']:
        if not value:
            search = ManageJson()
            print(search)
        elif isIp(value):
            connection = Connection()
            if connection.getConnection(value, 'ip'):
                connection = connection.getConnection(value, 'ip')
                ip, user = connection.get('ip'), connection.get('user')
                port, keyword = connection.get('port'), connection.get('keywords')
                default = connection.get('default')
                print(f"ip:{ip}  user:{user}  port:{port}  keywords:{keyword}  default:{default}")
            else:
                print('Not found')
        else:
            print('Invalid')
    elif args in ['-e', '--edit'] and value:
        if isIp(value):
            connection = Connection(value)
            connection.edit()
        else:
            print('Invalid')
    elif args != '-default':
        search = ManageJson()
        if search.searchList(args):
            connection = search.searchList(args)
            connectToSSH(connection.get('ip'), connection.get('user'), connection.get('port'))
        else:
            print('Not Found')
    else:
        connection = Connection()
        if connection.getConnection(True, 'default'):
            connection = connection.getConnection(True, 'default')
            connectToSSH(connection.get('ip'), connection.get('user'), connection.get('port'))
        else:
            print('No default server')

def connectToSSH(ip, user, port):
    if isIp(ip):
        cmd = f'ssh -p {port} {user}@{ip}'
        call(cmd, shell = True)

def help():
	print('Usage: ssh_keyword [OPTIONS option]  or ssh_keyword [KEYWORD]')
	print('Keyword recognition in ssh command')
	print('Specify a keyword in list of keywords of the connection')
	print('')
	print('Optional arguments:')
	print('-a',' '*4, '--add', ' '*6, 'add a new connection (ssh_keyword -a [IP])')
	print('-d',' '*4, '--default', ' '*2, 'add/change default connection (ssh_keyword -d [IP])')
	print('-rm',' '*3, '--remove', ' '*3, 'remove connection (ssh_keyword -rm [IP])')
	print('-ls',' '*3, '--list', ' '*5, 'list connection (ssh_keyword -ls or ssh_keyword -ls [IP])')
	print('-e', ' '*4, '--edit', ' '*5, 'edit connection (ssh_keyword -e [IP])')
	print('-h', ' '*4, '--help', ' '*5, 'show this help message and exit')

if __name__ == '__main__':

	arguments = ['-a', '--add',
				 '-d', '--default',
				 '-rm', '--remove',
				 '-ls', '--list',
				 '-e', '--edit'
				]

	if len(sys.argv) > 1:
		if sys.argv[1] in arguments and len(sys.argv) > 2:
			main(sys.argv[1], sys.argv[2])
		elif sys.argv[1] in ['-ls', '--list']:
			main(sys.argv[1])
		elif sys.argv[1] in ['-h', '--help']:
			help()
		elif sys.argv[1] in arguments:
			print(f'ssh_keyword {sys.argv[1]} [IP]')
		elif sys.argv[1][:1] != '-' and sys.argv[1] !='':
			main(sys.argv[1])
		elif sys.argv[1] =='':
			main('-default')
		else:
			print("Bad arguments passed.")
	else:
		main('-default')
