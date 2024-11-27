# Tìm hiểu về k8s
# Pod là gì ?
Trong K8s, pod là một nhóm gồm một hoặc nhiều những containers liên quan trực tiếp đến nhau và sẽ thường chạy cùng nhau ở trên cùng một worker node and cùng một Linux namespaces. Mỗi pod được coi như là một 'logical machine' với một địa chỉ IP riêng, hostname cũng như process riêng, có thể chạy một application riêng biệt. Application có thể là một process đơn lẻ, chạy trên một container hoặc nó có thể là một process chính của một application,....Tất cả các containers trong một pod sẽ xuất hiện để chạy trên cùng một 'logical machine', nếu một container khác pods hoặc chạy trên cùng một worker node thì cũng không thể chạy cùng container ở pods khác. 
![alt text](images/image.png)

Bởi vì ta không thể list được những containers riêng lẻ, vì chúng không phải là một objects có thể đứng được riêng lẻ trong K8s. Để liệt kê tất cả các pods ta sử dụng câu lệnh
```
kubectl get pods
```

![alt text](images/image1.png)

Trong câu lệnh này ta có thể thấy các trường tên của pods, trong pod này có bao nhiêu container được thể hiện qua trường ready, trạng thái của pod thông qua trường STATUS. Để xem thêm thông tin chi tiết của một pod ta có thể dùng câu lệnh `kubectl describe pod`. Nếu pod bị stuck ở trạng thái Pending có thể là do K8s không thể pull được image từ registry. 

Với một pod đang chạy, làm sao để ta có thể kết nối với nó ?. Như ta đã biết từ trước mỗi pod có một IP riêng, IP này là internal trong cluster và không thể kết nối từ bên ngoài. Để khiến một pod có thể kết nối được từ bên ngoài, ta sẽ thực hiện expose nó thông qua service object. Ta sẽ thực hiện tạo service mang tên là 'Loadbalancer'. Bằng cách tạo một loadbalancer service, một external LB có thể được tạo và ta có thể kết nối tới pod thông qua public IP của LB. 
**Lưu ý:** Minikube không hỗ trợ LB services, vì thể service sẽ không bao giờ có được một external IP. Nhưng ta có thể kết nối service thông qua external port của nó.