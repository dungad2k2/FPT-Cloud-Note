# Mục đích 

Mục đích của bài này để benchmark VM Checkpoint firewall về các thông số CPU, Packet per second,...
# Outline

   [[#1. Mô hình hệ thống benchmark]]
   [[#2. Các bước dụng môi trường benchmark]]
   [[#3. Các bước thực hiện benchmark]]
   [[#4. Những lưu ý, kinh nghiệm rút ra khi benchmark]]
# 1. Mô hình hệ thống benchmark

![[Pasted image 20250207110628.png]]
# 2. Các bước dựng môi trường benchmark 
## Chuẩn bị tài nguyên và môi trường:

- Tiến hành tạo 2 network có subnet lần lượt là: *192.168.120.0/24* và *14.4.4.0/24*.
- Tạo 2 VM: CP-FW-GW và CP-FW-MNGT và attach 2 network đã tạo cho 2 VM.
- Tạo các VM client và server (tùy theo mục đích benchmark mà tạo số lượng các cặp client và server khác nhau). Các client và server sẽ thuộc về các subnet mà nó thuộc về
## Cấu hình checkpoint:

**Lưu ý**: Ở đây dù sử dụng checkpoint firewall có flavor 16 cores. Tuy nhiên trong phần thực hiện lab, benchmark này có sử dụng phiên bản trial nên firewall checkpoint chỉ sử dụng được tối đa 8 cores mà thôi.  

Mô hình checkpoint firewall có hai cách triển khai:
- Cách đầu tiên là một VM thực hiện cả hai tác vụ là gateway và management. Một cụm Checkpoint FW chỉ có một VM thực hiện tác vụ management mà thôi, các gateway khác sẽ được add và quản lý thôi qua node management này.
- Cách thứ hai là một VM thực hiện tác vụ gateway, một VM thực hiện tác vụ management. Bài hướng dẫn này sẽ thực hiện theo **cách này**.

Sau khi tạo xong 2 VM Management và VM Gateway, ta tiến hành cấu hình cho VM Gateway thông qua VM Management bằng SmartConsole. 

**B1**: Set up mật khẩu và activation key cho VM Gateway truy cập thông qua IP được attach cho VM Gateway:

![[Pasted image 20250210132540.png]]
***Lưu ý*:** Cần lưu lại activation key để quản lý gateway thông qua SmartConsole.

**B2:** Thực hiện cung cấp license key cho Firewall Checkpoint: 
Contact với bên Sec để cung cấp key cho Firewall Checkpoint. Mỗi key setup cho Firewall Checkpoint sẽ tương ứng với 8 cores được active (Vd: 16 cores thì cần 2 license key, 24 cores thì cần 3 license key). Sau khi đã có license key thực hiện vào console của checkpoint log in bằng account đã setup ở bước 1 và copy key đã nhận vào console để activate license
![[Pasted image 20250225174451.png]]
**B3**: Setup mật khẩu cho VM Management truy cập thông qua IP được attach cho VM Management và cài đặt SmartConsole trên JumpWindows:

![[Pasted image 20250210133511.png]]

**B4**: Cài đặt SmartConsole và đăng nhập thông qua thông tin tài khoản mật khẩu (của VM Management) được setup ở Bước 2:

![[Pasted image 20250210133645.png]]

**B5:** Sau khi đăng nhập thành công tiến hành add gateway: 
Chọn **Object** -> **More object types** -> **Network Object** -> **Gateways and Servers** -> **New Gateway** -> **Wizard Mode**.

![[Pasted image 20250210134726.png]]

Điền IP và tên của Gateway -> Next -> Điền activation key đã set up trước đó ở trên gateway portal.

**Lưu ý**: Để có thể add Gateway thành công cần phải add thêm gateway là IP của VM Management trên Gaia Portal của VM Gateway.

**B6:** Chỉnh lại policy allow any any hoặc bất cứ policies nào mong muốn. Ở trong bài lab này, chỉ sử dụng rule any any.

![[Pasted image 20250210135541.png]]

**B7:** Sau khi đã add xong gateway thực hiện Publish, sau đó Install Policy(Các rules cho gateway). 

![[Pasted image 20250210135326.png]]

Sau khi cấu hình xong các bước trên bây giờ để client và server có thể kết nối được với nhau ta cần phải setup tiếp các server và client

## Setup Server và Client

Trong bài lab benchmark này, ta sẽ sử dụng tool **K6** để thực hiện load testing đẩy traffic qua firewall checkpoint đến các server. Giống như Jmeter hay các load testing khác K6 cũng tạo ra các virtual users mỗi một virtual user sẽ được coi là một user riêng có trách nhiệm thực hiện đẩy traffic đến endpoint cần load testing. 
B1: Tiến hành cài đặt K6 cho các client và Nginx cho các server:
-   Để cài đặt K6 tham khảo link sau đây: https://grafana.com/docs/k6/latest/set-up/install-k6/

```
sudo gpg -k
sudo gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
sudo apt-get update
sudo apt-get install k6
```

-  Để cài đặt Nginx, thực hiện:  `apt update` -> `apt install nginx -y`
B2: Tiến hành tạo Script để chạy K6 

Tham khảo đoạn script dưới đây:

```
import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
        stages: [
           { duration: '5m', target: 20000 },

           { duration: '10m', target: 20000 },

           { duration: '5m', target: 0 },
        ],

        thresholds: {
                http_req_duration: ['p(100)<3000'],
        },
};

export default function () {
     const response = http.get('http://14.4.4.154'); //Endpoint's server
     check(response, {
                   'status is 200': (r) => r.status === 200,
     });

     sleep(1);
}
```

Để thực hiện tạo loadtesting thực hiện câu lệnh:
```
k6 run <tên file script vừa tạo>
```

Như vậy coi như ta đã tiến hành tạo xong môi trường benchmark.

**Lưu ý:** Đối với việc thực hiện đo băng thông giữa các clients và các servers, thực hiện câu lệnh sau:
- Trên server: `iperf3 -s -p <port_number>`
- Trên client: `iperf3 -c <ip_server> -p <port_number> --time <time to send iperf3 traffic>`
- port_number bắt buộc phải giống nhau ở 2 câu lệnh ở trên 
# 3. Các bước thực hiện benchmark

- Phát tải trên các node k6 client theo phương pháp tăng dần số client phát tải từ mỗi client có flavor (8C-8G) đến các web server có flavor (8C-8G).
- Thực hiện monitor biểu đồ packet/s của Checkpoint cho đến khi xuất hiện packetloss. Từ đó xác định được ngưỡng chịu tải của FW Checkpoint -> Check thông qua **mon.fptcloud.com**
- Ngoài ra để đo được throughput qua từ client đến server sử dụng công cụ iperf3 để đo.
# 4. Những lưu ý, kinh nghiệm rút ra khi benchmark:

- Việc phát traffic từ các node phải được phát lần lượt, nên để khi node k6 thực hiện khởi tạo xong hết virtual users và phát traffic được khoảng tầm 2-3ph. Sau đó mới được thực hiện phát tiếp traffic từ các node k6 khác (trong trường hợp muốn dùng nhiều cặp client và server để thực hiện benchmark).
- Đối với việc resize VM Checkpoint Firewall, khi thực hiện resize từ một flavor ít cores sang flavor nhiều core hơn (từ 16 cores -> 32 cores) thì việc resize có thể làm nhảy số core được chỉ định cho việc forward gói tin đi (VD: 16 cores - 4 cores fw -> 32 cores - 10 cores fw). Còn đối với việc resize 2 flavor cùng số cores nhưng khác các cờ properties thì việc resize sẽ không làm nhảy số cores forward.
- Đối với việc đo bandwidth thực hiện đo iperf3 trên nhiều cặp client và server đến khi có packet drop.
- Việc đo cần phải thời gian đợi cho các traffic được phát ổn định -> "Giục tốc bất đạt".