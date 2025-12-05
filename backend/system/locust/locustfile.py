from locust import HttpUser, task, between

class HelloUser(HttpUser):
  wait_time = between(1, 3)

  # APIs Liên quan đến Booking và Payment


  # APIs Search và List
  @task(3)
  def search_flights(self):
    self.client.get("/api/flights?from=Hanoi&to=Ho-Chi-Minh-City")
  @task(3)
  def hello(self):
    self.client.get("/metrics")
    self.client.get("/health")

  # APIs Authentication