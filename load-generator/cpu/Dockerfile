FROM python:3.7-alpine
RUN apk add autoconf automake g++ make --no-cache
RUN pip install py-bcrypt
ADD app.py /app.py
CMD ["python", "app.py"]
