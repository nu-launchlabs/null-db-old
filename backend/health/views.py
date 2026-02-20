from django.http import HttpRequest, JsonResponse


def health_check(request: HttpRequest) -> JsonResponse:
    """Return a simple health check response."""
    return JsonResponse({"status": "ok"})
