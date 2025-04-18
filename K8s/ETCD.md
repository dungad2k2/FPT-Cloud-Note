## ETCD là gì ?
- ETCD là một distributed reliable key-value databases. ETCD có một ưu điểm là đơn giản. bảo mật và nhanh 
- Thông thường các database truyền thống sử dụng dữ liệu quan hệ (SQL) thường được lưu trong nhiều cột và nhiều hàng. Nhưng key-value database chỉ lưu thông tin dưới format là key và value ![[Pasted image 20250411103537.png]]
- Trong kubernetes, etcd được dùng để lưu những cặp giá trị key-value cho cluster. Ngoài ra etcd còn được sử dụng như một backing store để có thể backup dữ liệu cho cluster. ETCD sẽ lưu những thông tin về cluster như: nodes, pods, configs, secrets, accounts, roles, bindings. 
- Tất cả thông tin khi ta dùng một câu lệnh quen thuộc là `kubectl get` đều được lấy từ `etcd server` . 
- Với một cluster được cài đặt bằng kubespray, etcd sẽ được cài đặt từ source và chạy như một service. Để biết thêm về etcd được advertise ở đâu, data được lưu ở đâu hãy xem config service của etcd.
- Trong một cluster có tính HA, etcd sẽ được đặt ở mỗi master node 