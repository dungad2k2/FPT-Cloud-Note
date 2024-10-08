# Tìm hiểu một cách cẩn thận về prometheus:

## Prothemeus metric types 
1. Gauges
Gauges là một metric mà giá trị có thể đi lên và đi xuống một cách tự nhiên bởi vì nó đại diện cho tình trạng hiện tại của hệ thống như là disk space, memory usage,....
2. Counter
Counter là đại điện cho một giá trị được đếm trong một khoảng thời gian ví dụ như là tổng số lượng HTTP request mà application có thể handle được. Không giống gauges, counter chỉ có thể đi lên và không bao giờ đi xuống. Nếu hệ thống bị crashed hoặc restart bởi một số lý do nào đó thì counter sẽ trở về giá trị 0 -> điểm này được gọi là *counter reset*

Counter metric có 2 method để update các giá trị: Inc() tăng một số với giá trị nguyên và Add() tăng một số với giá trị thập phân. 

Do đặc thù phụ thuộc vào điểm *counter reset* đo đó mà giá trị thô của những metric counter thường không thể hiện được gì nhiều. Vì vậy để biết được các metric counter đi lên nhanh như nào. Để làm được điều này PromQL cho ta một số hàm như là **rate()**, **irate()**, **increase()** bằng cách lấy những metric counter làm tham số cho các hàm trên.