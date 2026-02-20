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

clean:
	@echo "Cleaning..."
	@find . -type d -name "__pycache__" -exec rm -rf {} + 2>/dev/null || true
	@find . -type f -name "*.pyc" -delete 2>/dev/null || true

.PHONY: install run test check clean
