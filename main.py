from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('home.html')

@app.route('/secure')
def secure():
    return render_template('secure.html')

@app.route('/feedback')
def feedback():
    return render_template('feedback.html')

@app.route('/help')
def help():
    return render_template('help.html')

@app.route('/sign-in')
def sign_in():
    return render_template('sign-in.html')

@app.route('/sign-up')
def sign_up():
    return render_template('sign-up.html')

if __name__ == '__main__':
    app.run(debug=True)