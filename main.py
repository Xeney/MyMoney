from flask import Flask, render_template, request, redirect
from forms import forms as Fm

app = Flask(__name__)
app.config['SECRET_KEY'] = 'a really really really really long secret key'
feedbacks = []

@app.route('/')
def home():
    return render_template('home.html')

@app.route('/secure')
def secure():
    return render_template('secure.html')

@app.route('/feedback', methods=["POST", "GET"])
def feedback():
    form = Fm.FeedbackForm()
    if request.method == "POST":
        if form.validate_on_submit():
            name = form.name.data
            email = form.email.data
            title_message = form.title_message.data
            message = form.message.data
            feedbacks.append([name, email, title_message, message])
            print(feedbacks)
            return redirect('/feedback')
    return render_template('feedback.html', form=form)

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
    app.run(debug=True, port=5000, host="0.0.0.0")