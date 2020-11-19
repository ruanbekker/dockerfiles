
### Non CPU Intensive

On a EC2 Instance:

```
$ docker run -it -p 5000:5000 ruanbekker/aws-metadata-container:latest
```

Test:

```
$ curl -s http://ec2-ip-address:5000 | jq .
{
  "availability_zone": "eu-west-1c",
  "container_hostname": "83cad2586360",
  "instance_hostname": "ip-172-31-84-235",
  "instance_lifecycle": "on-demand",
  "instance_id": "i-00000000000000000"
}
```

### CPU Intensive

On a EC2 Instance:

```
$ docker run -it -p 5000:5000 ruanbekker/aws-metadata-container:cpu
```

Test:

```
$ curl -s http://ec2-ip-address:5000/cpu | jq .
{
  "availability_zone": "eu-west-1c",
  "container_hostname": "83cad2586360",
  "instance_hostname": "ip-172-31-84-235",
  "instance_lifecycle": "on-demand",
  "instance_id": "i-00000000000000000"
}
```
