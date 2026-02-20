# Makefile for NullDB backend

install:
	@echo "Installing dependencies..."
	@pip install -r requirements.txt

run:
	@python manage.py runserver 0.0.0.0:8080

test:
	@echo "Testing..."
	@python manage.py test

check:
	@python manage.py check

fmt:
	@echo "Formatting..."
	@ruff format .
	@ruff check --fix .

lint:
	@echo "Linting..."
	@ruff check .
	@ruff format --check .

typecheck:
	@echo "Type checking..."
	@mypy .

clean:
	@echo "Cleaning..."
	@find . -type d -name "__pycache__" -exec rm -rf {} + 2>/dev/null || true
	@find . -type f -name "*.pyc" -delete 2>/dev/null || true
	@rm -rf .mypy_cache .ruff_cache

.PHONY: install run test check fmt lint typecheck clean
