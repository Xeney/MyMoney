from flask_wtf import FlaskForm
from wtforms import StringField, TextAreaField
from wtforms.validators import DataRequired, Email

class FeedbackForm(FlaskForm):
    name = StringField("ФИО: ", validators=[DataRequired()])
    email = StringField("Email: ", validators=[Email(), DataRequired()])
    title_message = StringField("Тема сообщения", validators=[DataRequired()])
    message = TextAreaField("Сообщение", validators=[DataRequired()])
