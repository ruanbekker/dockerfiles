from flask import Flask, jsonify, request
import os, socket, json

app = Flask(__name__)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
    HOSTNAME  = socket.gethostname()
    MESSAGE   = os.environ['APP_MESSAGE']
    NAMESPACE = os.environ['NAMESPACE']
    IP        = request.environ.get('HTTP_X_REAL_IP', request.remote_addr)
    
    return jsonify({
        'message': MESSAGE, 
        'hostname': HOSTNAME,
        'namespace': NAMESPACE,
        'ip-address': IP,
        'request-path': path,
        'headers': json.loads(json.dumps({k:v for k, v in request.headers}))
    })

if __name__ == '__main__':
    app.run(host='0.0.0.0')
