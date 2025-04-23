# Architecture Overview

VictoriaMetrics có 2 kiểu deploy: single node và multi node. Trong phần này, ta sẽ tìm hiểu về cách triển khai multi node (Cluster version) luôn. Phần single node sẽ được tìm hiểu sau.

VictoriaMetrics cluster bao gồm những services như sau:
- `vmstorage`: lưu trữ dữ liệu thô và trả về dữ liệu đã được query trong khoảng thời gian đối với một metrics hay label nào đó.
- `vminsert`: là một điểm mà dữ liệu được scrape đi vào và thực hiện việc chia dữ liệu được scrape giữa các node **vmstorage**.
- `vmselect`: thực hiện những queries tới bằng cách truy vấn những metrics cần tới tất cả các `vmstorage` node.
# Cài dặt victoria-metrics thông qua Docker Compose 
```
services:
  vmstorage-1:
    container_name: vmstorage-1
    image: victoriametrics/vmstorage:v1.112.0-cluster
    ports:
      - 8482
      - 8400
      - 8401
    volumes:
      - ./vmstorage1:/storage
    command:
      - "--storageDataPath=/storage"
    restart: always
  vmstorage-2:
    container_name: vmstorage-2
    image: victoriametrics/vmstorage:v1.112.0-cluster
    ports:
      - 8482
      - 8400
      - 8401
    volumes:
      - ./vmstorage2:/storage
    command:
      - "--storageDataPath=/storage"
    restart: always
  vminsert:
    container_name: vminsert
    image: victoriametrics/vminsert:v1.112.0-cluster
    depends_on:
      - "vmstorage-1"
      - "vmstorage-2"
    command:
      - "--storageNode=vmstorage-1:8400"
      - "--storageNode=vmstorage-2:8400"
    ports:
      - 8480:8480
    restart: always
  vmselect-1:
    container_name: vmselect-1
    image: victoriametrics/vmselect:v1.112.0-cluster
    depends_on:
    depends_on:
      - "vmstorage-1"
      - "vmstorage-2"
    command:
      - "--storageNode=vmstorage-1:8401"
      - "--storageNode=vmstorage-2:8401"
    ports:
      - 8481:8481
    restart: always
  vmselect-2:
    container_name: vmselect-2
    image: victoriametrics/vmselect:v1.112.0-cluster
    depends_on:
      - "vmstorage-1"
      - "vmstorage-2"
    command:
      - "--storageNode=vmstorage-1:8401"
      - "--storageNode=vmstorage-2:8401"
    ports:
      - 8481
    restart: always
```