FROM python:3.11-slim

RUN apt-get update && apt-get install -y ffmpeg \
  && python3 -m pip install pipenv

WORKDIR /app

CMD ["bash"]
