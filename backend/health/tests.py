import json

from django.test import Client, SimpleTestCase


class HealthCheckTest(SimpleTestCase):
    def test_health_check_returns_ok(self):
        """GET /health should return {"status": "ok"} with a 200 status."""
        client = Client()
        response = client.get("/health")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.content), {"status": "ok"})
