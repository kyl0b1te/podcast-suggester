FROM python:3.11-slim

RUN apt-get update && apt-get install -y ffmpeg
RUN pip install poetry

WORKDIR /app

COPY poetry.lock pyproject.toml ./

RUN poetry config virtualenvs.create false && poetry install --no-interaction --no-ansi

COPY . .
