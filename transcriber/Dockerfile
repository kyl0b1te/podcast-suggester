FROM python:3.13-alpine

RUN apk add --no-cache curl
RUN curl -sSL https://install.python-poetry.org | python3 -

# Add `poetry` to PATH
ENV PATH="${PATH}:$HOME/.local/bin"
ENV POETRY_VIRTUALENVS_IN_PROJECT=true

WORKDIR /app

# Install dependencies
COPY pyproject.toml ./
RUN poetry install

# Run your app
COPY . /app
CMD [ "poetry", "--version" ]