from random import choice, randint
from string import ascii_letters, punctuation, digits
from flask import Flask, jsonify
import bcrypt
import time
import os
import logging
import socket
import requests

app = Flask(__name__)

LOG_ROUNDS = os.environ.get('LOG_ROUNDS', 10)

def generate_hash(log_rounds):
    string_format = ascii_letters + punctuation + digits
    generated_string = "".join(choice(string_format) for x in range(randint(5, 20)))
    bcrypt_hash = bcrypt.hashpw(generated_string, bcrypt.gensalt(log_rounds))
    return bcrypt_hash

def get_instance_details(instruction):
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

    # generate cpu load
    if instruction == 'cpuload':
        generate_hash(int(LOG_ROUNDS))

    return data

@app.route('/')
def get():
    instance_details = get_instance_details('noload')
    return jsonify(instance_details)

@app.route('/cpu')
def get_cpu():
    instance_details = get_instance_details('cpuload')
    return jsonify(instance_details)

@app.route('/health')
def health():
    return jsonify({'message': 'ok'})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
