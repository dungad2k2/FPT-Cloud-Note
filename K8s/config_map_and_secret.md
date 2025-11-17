## ConfigMaps

Configmap trong k8s là một kubernetes API object chỉ đơn giản bao gồm một list các cặp key/value. Pod có thể được truy cập hoặc đọc được các
cặp key/value này trong config map. Một pod có thể được sử dụng một hoặc nhiều configmaps. Các key/value này trong config map có thể được đưa 
vào containers như các biến môi trường hoặc có thể được mount như một file trong filesystem của containers thông qua configmap volume. 

Việc sử dụng configmap để lưu các các cấu hình giúp ta tách biệt phần cấu hình khỏi code từ đó có thể hoạt động một cách linh hoạt dễ bảo trì
và triển khai dễ dàng đối với từng môi trường dev/test, staging, production. Hơn nữa, ta có thẻ quản lý các cấu hình một cách tập trung và không 
cần build lại image hoặc sửa code. 

## Tạo một configmap object 

Tạo một configmap với kubectl create và một cặp key/value thông qua `--from-literal`

```
kubectl create configmap kiada-config --from-literal status-message="This status message is set in the kiada-config config map"
```

Ngoài ra ta có thể tạo một configmap từ một file bằng cách thay thế `--from-literal` bằng `--from-file`

Tạo configmap bằng YAML manifest

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: kiada-config
data:
  status-message: This status message is set in the kiada-config config map
```

Ta có thể truyền giá trị cho biến env thông qua configmap entry

```
kind: Pod
...
spec:
  containers:
  - name: kiada
    env:
    - name: INITIAL_STATUS_MESSAGE
      valueFrom:
        configMapKeyRef:
          name: kiada-config
          key: status-message
          optional: true
    volumeMounts:
    - ...
```

Để ngăn chặn user có thể thay đổi data trong configmap ta cần mark configmap đó là `immutable` với `immutable=true`. 

Ngoài ra ta có thể mount configmap như là một volume. 
## Secret

Secret có chức năng tương tự như configmap tuy nhiên nó thường được dùng để lưu các thông tin nhạy cảm như là mật khẩu, token, ssh key,..... 