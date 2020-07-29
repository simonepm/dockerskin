FROM python:3.8.5-slim-buster

WORKDIR /app

ADD . /app

RUN python -m pip install --upgrade pip

RUN pip install -r requirements.txt

RUN python test.py

CMD python -u main.py
