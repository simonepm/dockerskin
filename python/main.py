from flask import Flask
from waitress import serve
from hello import hello

app = Flask(__name__)
port = 8080


@app.route("/")
def index():
    return hello.SayHello() + '\n'


if __name__ == '__main__':

    print(f"Listening on port {port}...")

    serve(app, host='0.0.0.0', port=port)
