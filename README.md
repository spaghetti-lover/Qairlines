<a id="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/spaghetti-lover/Qairlines">
    <img src="https://github.com/user-attachments/assets/d604f059-ff5d-4b2d-a201-17cd3d211165" alt="Logo" width="auto" height="auto">
  </a>

<h3 align="center">QAirline</h3>

  <p align="center">
    High Performance Flight Booking Management System
    <br />
    <a href="https://github.com/spaghetti-lover/Qairlines/tree/main/backend/docs"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://www.youtube.com/watch?v=EIDpxah3Ugw&t=53s">View Demo</a>
    &middot;
    <a href="https://github.com/spaghetti-lover/Qairlines/issues">Report Bug</a>
    &middot;
    <a href="https://github.com/spaghetti-lover/Qairlines/pulls">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
### <img src="https://github.com/user-attachments/assets/d604f059-ff5d-4b2d-a201-17cd3d211165" alt="logo-white" style="width: 200px; display: inline-block; vertical-align: middle;">

![Screenshot 2024-12-22 024054](https://github.com/user-attachments/assets/3616554a-92ff-4023-b21e-979cf2a29660)

### What is QAirline?

QAirline is an online flight booking platform that helps users easily search, compare, and book airline tickets. With a user-friendly interface and smart features, the website offers a fast, convenient, and secure booking experience. Demo link: https://www.youtube.com/watch?v=EIDpxah3Ugw&t=53s


[Visit the official website here](https://www.qairline.website/)

## Features

- **Flight Search**: Search for flights by departure, destination, travel date, and number of passengers.

- **Quick Booking**: Simple booking process.

- **Booking Management**: Review your booking details and statuses.

- **Online Payment**: Integrated with the Stripe payment platform.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* [![Next][Next.js]][Next-url]
* [![Golang][Go]][Go-url]
* [![Postgres][PostgreSQL]][PostgreSQL-url]
* [![Stripe][Stripe]][Stripe-url]
* [![Redis][Redis]][Redis-url]
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started
### Prerequisites
- Go: Install [here](https://go.dev/doc/install)
- Docker: Install [here](https://docs.docker.com/engine/install/)

### Installation
Clone the repo
```sh
git clone https://github.com/spaghetti-lover/Qairlines.git
```
#### Backend
1. Go to backend folder
```
cd backend
```
2. Add environment variable in app.env
```
APP_EVN=production //can use development for debugging with logs

MAIL_SENDER_NAME="Qairlines Support"
MAIL_SENDER_ADDRESS = "your email"
MAIL_SENDER_PASSWORD = "your email password"

REDIS_ADDRESS=0.0.0.0:6379
DB_DRIVER = "postgres"
DB_SOURCE = "postgresql://root:secret@localhost:5432/qairline?sslmode=disable"
SERVER_ADDRESS_PORT = :8080
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=7h

RATE_LIMITER_REQUEST_SEC=5
RATE_LIMITER_REQUEST_BURST=10

STRIPE_SECRET_KEY=<Stripe secret key>
STRIPE_WEBHOOK_SECRET=<Stripe webhook secret>
```

3. Start PostgreSQL service
```
make postgres
```

4. Start Redis service
```
make redis
```

5. Setup database
```
make createdb
make migrateup
```

6. Test (optional)
```
make test
```

7. Run server
```
make server
```

#### Frontend
```
npm install
npm run dev
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

_For more examples, you can the read [Documentation](https://github.com/spaghetti-lover/Qairlines/tree/main/backend/docs/documentation)_

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


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap
- [ ] Add MINIO for image management
- [ ] Change booking feature to work with concurrent requests
- [ ] Add BookingHistory

See the [open issues](https://github.com/spaghetti-lover/Qairlines/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Top contributors:

<a href="https://github.com/spaghetti-lover/QAirlines/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=spaghetti-lover/QAirlines" alt="contrib.rocks image" />
</a>


<!-- LICENSE -->
## License

Distributed under the project_license. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Phung Duc Anh - [@facebook](https://www.facebook.com/duc.anh.phung.2511/) - [phungducanh2511@gmail.com](mailto:phungducanh2511@gmail.com)

Project Link: [https://github.com/spaghetti-lover/Qairlines](https://github.com/spaghetti-lover/Qairlines)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[Go]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge
[Go-url]: https://go.dev/
[PostgreSQL]: https://img.shields.io/badge/postgresql-4169e1?style=for-the-badge&logo=postgresql&logoColor=white
[PostgreSQL-url]: https://www.postgresql.org/
[Stripe]: https://img.shields.io/badge/-Stripe-008CDD?style=flat&logo=stripe&logoColor=white
[Stripe-url]: https://stripe.com/
[Redis]: https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white
[Redis-url]: https://redis.io/