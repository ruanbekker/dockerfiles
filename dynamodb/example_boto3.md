```
import boto3
dynamodb = boto3.resource('dynamodb', endpoint_url='http://localhost:8000')
table = dynamodb.create_table(TableName='staffs',KeySchema=[{'AttributeName': 'username', 'KeyType': 'HASH'}, {'AttributeName': 'last_name', 'KeyType': 'RANGE'}],AttributeDefinitions=[{'AttributeName': 'username','AttributeType': 'S'},{'AttributeName': 'last_name', 'AttributeType': 'S'}], ProvisionedThroughput={'ReadCapacityUnits': 1,'WriteCapacityUnits': 1})
table.item_count
table.put_item(Item={'username': 'ruanb','first_name': 'ruan','last_name': 'bekker','age': 30,'account_type': 'administrator'})
table.scan()
```

More examples: https://sysadmins.co.za/interfacing-amazon-dynamodb-with-python-using-boto3/
