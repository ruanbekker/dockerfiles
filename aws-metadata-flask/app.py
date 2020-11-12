import socket
import requests
from flask import Flask, jsonify

app = Flask(__name__)

def get_instance_details():
    data = {}
    instance_id = requests.get('http://169.254.169.254/latest/meta-data/instance-id').content
    availability_zone = requests.get('http://169.254.169.254/latest/meta-data/placement/availability-zone').content
    instance_hostname = requests.get('http://169.254.169.254/latest/meta-data/hostname').content
    instance_lifecycle = requests.get('http://169.254.169.254/latest/meta-data/instance-life-cycle').content
    container_hostname = socket.gethostname()

    data['instance_id'] = instance_id.decode('utf-8')
    data['availability_zone'] = availability_zone.decode('utf-8')
    data['instance_hostname'] = instance_hostname.decode('utf-8').split('.')[0]
    data['instance_lifecycle'] = instance_lifecycle.decode('utf-8')
    data['container_hostname'] = container_hostname

    return data

@app.route('/')
def get():
    instance_details = get_instance_details()
    return jsonify(instance_details)

@app.route('/health')
def health():
    return jsonify({'message': 'ok'})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
