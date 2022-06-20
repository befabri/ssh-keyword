import sys, os, json

class ManageJson(object):
    def __init__(self):
        self.dir = os.path.dirname(__file__)+'/data/'
        if self.readJson() == "Not exist":
            self.createJson()
        self.data = self.readJson()

    def search(self, valeur, keyDict):
        for i in self.data:
            if valeur == i.get(keyDict):
                return i

    def searchList(self, keyword):
        for i in self.data:
            if keyword in i.get('keywords'):
                return i

    def searchIndex(self, valeur):
        return self.data[valeur]

    def getLen(self):
        return len(self.data)

    def readJson(self):
        try:
            with open(self.dir+"data.json", "r") as f:
                data = json.load(f)
            return data
        except:
            return "Not exist"

    def createJson(self):
        if not os.path.exists(self.dir):
            os.makedirs(self.dir)
        try:
            with open(self.dir+"data.json", "a") as f:
                    json.dump([], f)
        except:
            print('Cant create data/data.json')
            sys.exit()

    def saveToJson(self):
        with open(self.dir+"data.json", "w+") as f:
            json.dump(self.data, f, indent=4)

    def __str__(self):
        string = ''
        x = 1
        for connection in self.data:
            ip, user = connection.get('ip'), connection.get('user')
            port, keyword = connection.get('port'), connection.get('keywords')
            default = connection.get('default')
            string += f"{x}. ip:{ip}  user:{user}  port:{port}  keywords:{keyword}  default:{default}\n"
            x+=1
        return string
