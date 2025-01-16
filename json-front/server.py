from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def json_frontend():
    return render_template("json.html")
