# Architecture Overview

VictoriaMetrics có 2 kiểu deploy: single node và multi node. Trong phần này, ta sẽ tìm hiểu về cách triển khai multi node (Cluster version) luôn. Phần single node sẽ được tìm hiểu sau.

VictoriaMetrics cluster bao gồm những services như sau:
- `vmstorage`: lưu trữ dữ liệu thô và trả về dữ liệu đã được query trong khoảng thời gian đối với một metrics hay label nào đó.
- `vminsert`: là một điểm mà dữ liệu được scrape đi vào và thực hiện việc chia dữ liệu được scrape giữa các node **vmstorage**.
- `vmselect`: thực hiện những queries tới bằng cách truy vấn những metrics cần tới tất cả các `vmstorage` node.
