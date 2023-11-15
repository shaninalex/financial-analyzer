import time
import os
import sys
from pika import BlockingConnection, ConnectionParameters
from pika.exchange_type import ExchangeType
from pika.exceptions import AMQPConnectionError

connection = None

while connection == None:
    try:
        connection = BlockingConnection(
            ConnectionParameters(
                host=os.getenv("RABBITMQ_HOST"),
                port=os.getenv("RABBITMQ_PORT"),
            )
        )
    except AMQPConnectionError as e:
        print(e)
        print("reconnect")
        time.sleep(1)


channel = connection.channel()
print("Connected\nDeclare exchanges.")

# EXCHANGES
# direct exchanges
# deliver messages to client queue by routing key
# every client should subscribe to this exchange (Bind all this exchanges to user queue)
channel.exchange_declare(
    exchange="ex.datasource", exchange_type=ExchangeType.fanout, durable=True
)
channel.exchange_declare(
    exchange="ex.email", exchange_type=ExchangeType.direct, durable=True
)
channel.exchange_declare(
    exchange="ex.notifications", exchange_type=ExchangeType.direct, durable=True
)
channel.exchange_declare(
    exchange="ex.report", exchange_type=ExchangeType.direct, durable=True
)


# lobal exchange to broadcast notifications to all users
channel.exchange_declare(
    exchange="ex.global.notifications", exchange_type=ExchangeType.fanout, durable=True
)

print("Declare queues.")
# QUEUES
# this queues require for services to get messages and trigger actions
channel.queue_declare(queue="q.datasource", durable=True)
channel.queue_declare(queue="q.email", durable=True)
channel.queue_declare(queue="q.report", durable=True)


# BINDNING
# receive all messages from datasource exchange
channel.queue_bind("q.datasource", "ex.datasource")

print("Successfully declared queues and exchanges. Exiting...")
sys.exit()
