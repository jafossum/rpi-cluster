import platform


class SayHello:

    def __init__(self):
        self.un = platform.uname()[1]

    def get(self, event):
        st = ''
        
        user = event.query.get('user', '')
        time = event.query.get('time', 0)

        if user != '':
            st += ', user: ' + user
        if time != 0:
            st += ', time: ' + time

        return {
            "statusCode": 201,
            "body": 'Hello from ' + self.un + st + '\n'
        }


def handle(event, context):

    print(event.method + " Received")
    hello = SayHello()
    return hello.get(event)
