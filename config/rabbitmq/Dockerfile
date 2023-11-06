FROM python:3.11.2

ENV PYTHONUNBUFFERED=1

COPY ./requirements.txt /app/requirements.txt

WORKDIR /app

RUN pip install -r requirements.txt

COPY . /app

ENTRYPOINT [ "python", "-u" ]

CMD ["main.py"]