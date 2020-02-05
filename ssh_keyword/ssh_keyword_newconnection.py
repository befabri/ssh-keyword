from ssh_keyword_tools import isIp, checkQuit

class CreateConnection(object):
    'Initialization of variables to record a new connection'

    def __init__(self):
        pass

    def qPort(self):
        while True:
            enter = input("Port: ")
            checkQuit(enter)
            if enter.isdigit():
                return enter
            else:
                print("Invalid port")

    def qUser(self):
        while True:
            enter = input("Enter a user: ")
            checkQuit(enter)
            if bool(enter.strip()):
                return enter.lower()
            else:
                print('Invalid user')

    def qIp(self):
        while True:
            enter = input(f'Enter a ip adress: ')
            checkQuit(enter)
            if isIp(enter):
                return enter
            else:
                print('Invalid ip')

    def qKeywords(self):
        while True:
            enter = input("Enter a list of names separate by ',': ")
            checkQuit(enter)
            try:
                enter = enter.replace(' ', '').split(",")
                return enter
            except:
                print('Invalid keywords')

    def qDefault(self):
        while True:
            enter = input(f"Default server ([Y]es | [N]o): ")
            checkQuit(enter)
            if enter.lower() in ['y', 'yes']:
                return True
            elif enter.lower() in ['n', 'no']:
                return False
            else:
                print('Invalid')
