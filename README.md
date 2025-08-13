(Under development)

#  High-Performance Distributed API Rate Limiter (With redis )

A production-grade, extensible rate limiting service written in Go, featuring multiple algorithms, distributed state management via Redis, dynamic configuration, and real-time observability via Prometheus and Grafana.

>  Built for developers, APIs, and backend systems requiring scalable, reliable request throttling.

---

##  Features
`
-  **Multiple Algorithms**:
  -  Fixed Window Counter
  -  Sliding Window Log
  -  Token Bucket (request-based, byte-based optional)
  
-  **Distributed Rate Limiting**:
  - Shared state stored in **Redis**
  - Uses **Lua scripts** for atomic operations

- ⚙ **Dynamic Configuration**:
  - Per-endpoint + per-user-tier rate limits
  - Hot-reload support (via file or Redis)

-  **Monitoring & Alerts**:
  - **Prometheus** for metrics
  - **Grafana dashboards** for real-time visualization
  - Custom alerts for high block rates

-  **Dockerized Stack**:
  - Redis, Prometheus, Grafana, and the API server via `docker-compose`

---

##  Use Cases

- Enforce rate limits on public/private APIs
- Apply different limits for Free, Premium, and Admin users
- Build internal throttling middleware for microservices
- Simulate CDN-style bandwidth constraints (optional)

---

## Project Structure

