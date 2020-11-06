
from flask import Flask, request
from flask_restful import Resource, Api
import platform
import os

app = Flask(__name__)
api = Api(app)

class sayHello(Resource):

    def __init__(self):
        self.un = platform.uname()[1]

    def get(self):
        st = ''
        
        user = request.args.get('user', '')
        time = request.args.get('time', 0)

        if user != '':
            st += ', user: ' + user
        if time != 0:
            st += ', time: ' + time

        return 'Hello from ' + self.un + st

api.add_resource(sayHello, '/')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=int(os.environ.get('PORT', 8080)))