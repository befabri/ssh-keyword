import json
from ssh_keyword_tools import checkQuit
from ssh_keyword_json import ManageJson
from ssh_keyword_newconnection import CreateConnection

class Connection(object):
    'Manage connections'
    def __init__(self, ip='0'):
        self.ip = ip
        self.fJson = ManageJson()

    def setConnection(self):
        if not self.isExist(self.ip):
            newConnection = CreateConnection()
            print('Entry for new connection')
            print("Enter 'Quit' for exit")
            print()
            self.user = newConnection.qUser()
            self.port = newConnection.qPort()
            self.keywords = newConnection.qKeywords()
            self.default = newConnection.qDefault()
            datajson = self.toJson()
            self.reloadDatas()
            self.fJson.data.append(json.loads(datajson))
            self.toFile()
        else:
            print('Connection already exist, you can edit it')

    def getConnection(self, value, dictKey):
        'Get connection from json'
        if self.fJson.search(value, dictKey):
            return self.fJson.search(value, dictKey)
        return None

    def isExist(self, ip):
        if self.fJson.search(ip, 'ip'):
            return True
        return False

    def remove(self):
        'Remove a connection from json'
        if self.isExist(self.ip):
            connection = self.getConnection(self.ip, 'ip')
            self.fJson.data.remove(connection)
            self.toFile()
        else:
            print("connection not found")

    def update(self, connection, value, dictKey):
        'Update a connection from json'
        indexConnection = self.fJson.data.index(connection)
        self.fJson.data[indexConnection].update({value:dictKey})
        self.toFile()

    def edit(self):
        'Edit a connection from json'
        if self.isExist(self.ip):
            connection = self.getConnection(self.ip, 'ip')
            while True:
                ip, user = connection.get('ip'), connection.get('user')
                port, keyword = connection.get('port'), connection.get('keywords')
                default = connection.get('default')
                print(f'Edit connection {self.ip}')
                print(f"ip:{ip}  user:{user}  port:{port}  keywords:{keyword}  default:{default}")
                print('')
                enter = input("What do you want to edit ([Q]uit for exit): ")
                checkQuit(enter)
                if enter.lower() in ['ip', 'user', 'port', 'key', 'keys', 'keyword', 'keywords', 'default', 'def']:
                    editCo = CreateConnection()
                    if enter.lower() == 'ip':
                        ip = editCo.qIp()
                        if not self.isExist(ip):
                            self.update(connection, 'ip', ip)
                            self.ip = ip
                            print('Ip edited')
                        else:
                            print('Ip already exist')
                    elif enter.lower() == 'user':
                        self.update(connection, 'user', editCo.qUser())
                        print('User edited')
                    elif enter == 'port':
                        self.update(connection, 'port', editCo.qPort())
                        print('Port edited')
                    elif enter.lower() in ['key', 'keys', 'keyword', 'keywords']:
                        self.update(connection, 'keywords', editCo.qKeywords())
                        print('Keywords edited')
                    elif enter.lower() in ['def', 'default']:
                        self.update(connection, 'default', editCo.qDefault())
                        print('Server default edited')
        else:
            print("Connection not found")

    def addDefault(self):
        'Add a default connection'
        if self.isExist(self.ip):
            connection = self.getConnection(self.ip, 'ip')
            if self.getConnection(True, 'default'):
                previousConnection = self.getConnection(True, 'default')
                print(f"Previous connection '{previousConnection.get('ip')}' default is now set to False")
                self.update(previousConnection, 'default', False)
            self.update(connection, 'default', True)
            self.toFile()
        else:
            print("Connection not found")

    def toJson(self):
        del self.fJson
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True, indent=4)

    def toFile(self):
        self.fJson.saveToJson()
        self.reloadDatas()

    def reloadDatas(self):
        'Reload json file'
        self.fJson = ManageJson()