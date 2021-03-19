from flask import Flask, render_template
from os import environ as envar

app = Flask(__name__)

@app.route('/')
def root():
    return render_template('index.html', app_title=envar['APP_TITLE'], app_url=envar['APP_URL'], app_text=envar['APP_TEXT'])

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
