FROM arm32v7/python:3-buster

LABEL maintainer="John A. Fossum"

RUN pip3 install --no-cache-dir blinkt

WORKDIR /app

COPY app .

CMD ["python", "pixel-controller.py"]