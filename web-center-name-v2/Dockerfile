FROM python:3-alpine

ENV APP_TITLE Welcome
ENV APP_URL https://google.com
ENV APP_TEXT Visit our Website

RUN pip install flask
WORKDIR /app
ADD . .

CMD ["python", "app.py"]
