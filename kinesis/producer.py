import boto3
import random
import json
import time

names = ['james', 'stefan', 'pete', 'tom', 'frank', 'peter', 'ruan']

client = boto3.Session(region_name='eu-west-1').client('kinesis', aws_access_key_id='', aws_secret_access_key='', endpoint_url='http://localhost:4567')

list_streams = client.list_streams()

if 'mystream' not in list_streams['StreamNames']:
    client.create_stream(StreamName='mystream', ShardCount=1)
    time.sleep(1)

count = 0
print("Starting at {}".format(time.strftime("%H:%m:%S")))

while count != 25:
    count += 1
    response = client.put_record(StreamName='mystream', Data=json.dumps({"number": count, "name": random.choice(names), "age": random.randint(20,50)}), PartitionKey='a01')
    time.sleep(1)

print("Finished at {}".format(time.strftime("%H:%m:%S")))
