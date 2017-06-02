from flask import Flask
import json

app = Flask(__name__)

@app.route('/')
def index():
    return json.dumps({
    	"fib10": fib(10),
    	"fib20": fib(20)
    })


def fib (n):
  if n <= 1: return n
  return fib(n - 1) + fib(n - 2)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
