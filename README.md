# <img src="https://github.com/user-attachments/assets/d604f059-ff5d-4b2d-a201-17cd3d211165" alt="logo-white" style="width: 200px; display: inline-block; vertical-align: middle;"> - <span style="padding-bottom: 50px;">Trang Web Đặt Vé Máy Bay</span>

![Screenshot 2024-12-22 024054](https://github.com/user-attachments/assets/3616554a-92ff-4023-b21e-979cf2a29660)

## Qairline là gì?

Trang web đặt vé máy bay QAirline là một nền tảng trực tuyến giúp người dùng dễ dàng tìm kiếm, so sánh và đặt vé máy bay. Với giao diện thân thiện và các tính năng thông minh, trang web cung cấp trải nghiệm đặt vé nhanh chóng, tiện lợi và an toàn

[Truy cập website chính thức tại đây](https://www.qairline.website/)

## Tính năng

- **Tìm kiếm chuyến bay**: Hỗ trợ tìm kiếm vé máy bay theo điểm đi, điểm đến, ngày bay và số lượng hành khách.

- **Đặt vé nhanh chóng**: Quy trình đặt vé đơn giản với hỗ trợ thanh toán trực tuyến.

- **Quản lý đặt vé**: Xem lại thông tin và trạng thái các đặt vé đã thực hiện.

## Công nghệ sử dụng

### Frontend

- NextJS
- TailwindCss
- Dùng của bạn này: https://github.com/oceantran27/QAirline (cảm ơn các bn nha :33)

### Backend

- Golang: gorilla, "net/http" package
- Kafka: Gửi mail
- Viper: Load config
- Testify: Chạy unit tests
- SQLC: Tương tác với database
- Github Action cho CI

### Cơ sở dữ liệu

- PostgreSQL

## Cách chạy

### Backend

- B1: Chạy PostgreSQL database và Kafka

```
cd backend
docker-compose up
```

- B2: Tạo database

```
make createdb
```

- B3: Tạo bảng và tạo dữ liệu

```
make migrateup
```

- B4: Test (optional)

```
make test
```

- B5: Chạy server

```
make server
```

### Frontend

```
npm install
npm run dev
```

## Ảnh Chụp Màn Hình

![Screenshot 2024-12-22 024234](https://github.com/user-attachments/assets/03ba9f8a-cef8-4a68-bf83-3544d0e5dd5a)
![image](https://github.com/user-attachments/assets/41e01cc0-613c-41b9-9287-8794c354bcf0)
![image](https://github.com/user-attachments/assets/3fe77d89-5bf3-47db-8f0f-9881c9145c15)
![image](https://github.com/user-attachments/assets/5a2119d8-0f9d-4005-9440-9b2dba689ca8)
![image](https://github.com/user-attachments/assets/f9a156c3-57fc-4c5a-bcfc-282fc5f84241)
![image](https://github.com/user-attachments/assets/47a73981-d7c9-464d-aeba-64d831ea348a)
![image](https://github.com/user-attachments/assets/e41a1f4b-e39f-4361-82c3-61abbd9f8ddc)
![image](https://github.com/user-attachments/assets/f73aa80e-6f95-4e40-9e2a-68f6436f62db)
![image](https://github.com/user-attachments/assets/b6724cf8-ec14-4c7c-a73f-1dd2bbf9139c)
![image](https://github.com/user-attachments/assets/146b978b-1b0d-4f0b-9e37-400716ff9a85)
![image](https://github.com/user-attachments/assets/d0800f80-1b12-4c59-942e-b5dea6d2a9c0)

### Take note:

- Tích hợp MinIO với backend, dùng cho việc upload, download file
- Hoan thien transaction booking voi concurrent request
- Test chuc nang sending mail va test performance sao lai cham vay
- Implement not phan BookingHistory cho cac /api/user
