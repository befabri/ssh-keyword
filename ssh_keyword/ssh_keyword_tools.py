import ipaddress, sys

def isIp(ip):
    try:
        ipaddress.ip_address(ip)
        return True
    except:
        return False

def checkQuit(enter):
    'Check if quit'
    if enter.lower() == 'quit' or enter.lower() == 'q':
        print('Exit')
        sys.exit()