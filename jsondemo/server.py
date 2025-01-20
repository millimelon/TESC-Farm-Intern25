from flask import Flask, render_template, request
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

@app.route("/")
def json_frontend():
    url = request.base_url[:-6]
    return render_template("json.html", url=url)
