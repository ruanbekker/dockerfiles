from flask import Flask, request

app = Flask(__name__)

@app.route("/", methods=["POST", "GET"])
def main():
    response = {}

    if request.headers["Content-Type"] == 'application/json':
        response["event"] = request.get_json()
    else:
        response["event"] = request.get_data()

    print("Received Event:")
    print(response)
    return "event: {} \n".format(response)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
