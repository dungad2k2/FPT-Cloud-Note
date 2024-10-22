# File cấu hình thực hiện cấu hình ferderation và remote write trong Prometheus

```
global:
    scrape_interval: 10s
scrape_configs:
    - job_name: 'ferderate'

      honor_labels: true
      metrics_path: '/federate'
      
      params: 
       'match[]':
          # Filter metrics that want to ferderate
          - '{__name__="api_server_readiness_status_value"}'
      static_configs:
       - targets:
         #IP and port of remote instance
         - '203.29.16.92:31547'
      basic_auth:
       # Ạuthentication 
       username: 'admin'
       password: 'cwET43NjWBWI'
remote_write:
      #URL of the remote_write endpoint on the instance
      - url: "http://192.168.101.31:9090/api/v1/write"
```