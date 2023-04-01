from flask import Flask, render_template
from os import environ as envar
from socket import gethostname

app = Flask(__name__)

@app.route('/')
def root():
    hostname = gethostname()
    return render_template('index.html', app_title=envar['APP_TITLE'], hostname=hostname)

@app.route('/health')
def health():
    return ('', 204)

@app.route('/health/liveness')
def liveness():
    return ('ok', 200)

@app.route('/health/readiness')
def readiness():
    return ('ok', 200)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
