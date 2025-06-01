// components/Benefit.js
import dynamic from "next/dynamic";
import {
  MdArrowCircleRight,
  MdPerson,
} from "react-icons/md";
import { RiFlightTakeoffFill } from "react-icons/ri";
import { GiRuleBook } from "react-icons/gi";
import "react-multi-carousel/lib/styles.css";
import Link from "next/link";
import { HomeNewsCard } from "./NewsCards";
import Image from "next/image";
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
// Import `Carousel` với dynamic import để tắt SSR
const Carousel = dynamic(() => import("react-multi-carousel"), { ssr: false });

const responsive = {
  superLargeDesktop: {
    breakpoint: { max: 4000, min: 3000 },
    items: 3,
  },
  desktop: {
    breakpoint: { max: 3000, min: 1024 },
    items: 3,
  },
  tablet: {
    breakpoint: { max: 1024, min: 464 },
    items: 2,
  },
  mobile: {
    breakpoint: { max: 464, min: 0 },
    items: 1,
  },
};

export default function Benefit() {
  const router = useRouter()
  const [featuredArticles, setFeaturedArticles] = useState([])

  useEffect(() => {
    getAllNews()
  }, [router]);

  const getAllNews = async () => {
    const getAllNewsApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}GET /api/news/all`

    try {
      const response = await fetch(getAllNewsApi, {
          method: "GET",
      })
      if (!response.ok) {
          throw new Error("Send request failed")
      }

      const res = await response.json()

      setFeaturedArticles(res.map(a => {return {
        "slug": a.newsId,
        "image": a.image,
        "title": a.title,
        "description": a.description,
        "author": a.authorId,
        "content": a.content,
        "date": a.createAt.seconds ? new Date(a.arrival_time).toISOString().replace("T", " ").slice(0, -5) : a.createAt.split('T')[0],
        "buttonText": "Đọc thêm",
        "authorTitle": "Nhà báo",
        "authorImage": "/AvatarUser/no_avatar.jpg",
      }}))

    } catch (error) {
        // alert("Đã xảy ra lối, vui lòng thử lại(fixed)")
    }
  }
  return (
    <div>
      {/* Phần giới thiệu */}
      <section className="lg:flex">
        <div className="lg:w-[45%]" data-aos="fade-down">
        <Image
          src="/bg-2.jpg"
          alt="Background"
          width={1000}
          height={900}
          className="h-full"
        />
        </div>
        <div className="relative lg:w-[55%] bg-[url(/bg-map-2.png)] bg-[#313041] bg-right-bottom bg-contain flex flex-col justify-center py-16">
          <div className="absolute lg:w-3 w-0 bg-orange left-0 rounded-r-[10px] top-[120px] bottom-[120px]"></div>
          <div className="lg:w-[55%] lg:pl-24 px-6">
            <p className="text-xl text-orange" data-aos="fade-up">Danh sách lợi ích của chúng tôi</p>
            <h4 className="text-white lg:text-[50px] text-[30px] py-4" data-aos="fade-up">Tại sao chọn QAirline</h4>
            <p className="text-gray leading-8" data-aos="fade-down">
              Chúng tôi tự hào mang đến trải nghiệm đặt vé máy bay trực tuyến tiện lợi và an toàn, với những tính năng vượt trội và dịch vụ khách hàng tận tâm, giúp bạn dễ dàng tìm kiếm chuyến bay tốt nhất với giá cả hợp lý.
            </p>

            {/* Các lợi ích */}
            <div className="flex lg:flex-row flex-col items-center gap-8 pt-16" data-aos="fade-up">
              <span>
                <RiFlightTakeoffFill className="text-orange text-5xl" />
              </span>
              <span>
                <h6 className="text-white text-xl">Chuyên nghiệp và được chứng nhận</h6>
                <p className="text-gray leading-8 py-4">
                Chúng tôi hợp tác với các hãng hàng không uy tín và đã được chứng nhận, đảm bảo chuyến bay của bạn luôn an toàn và chất lượng. Đội ngũ chuyên nghiệp luôn hỗ trợ bạn 24/7.
                </p>
              </span>
            </div>
            <div className="flex lg:flex-row flex-col items-center gap-8 pt-6" data-aos="fade-up">
              <span>
                <GiRuleBook className="text-orange text-5xl" />
              </span>
              <span>
                <h6 className="text-white text-xl">Đặt vé ngay lập tức</h6>
                <p className="text-gray leading-8 py-4">
                  Dễ dàng lên lịch trình và đặt vé chỉ trong vài bước. Nhận xác nhận tức thì và tận hưởng những ưu đãi hấp dẫn khi đặt vé trên nền tảng của chúng tôi.
                </p>
              </span>
            </div>
          </div>
        </div>
      </section>

      {/* Phần carousel bài viết */}
      <section className="lg:pt-28 pt-14 pb-14" data-aos="fade-up">
        <div className="max-w-[1200px] xl:px-0 px-6 mx-auto">
          <div className="lg:flex items-center justify-between">
            <span>
              <p className="text-xl text-orange pb-4">Từ trang tin tức</p>
              <p className="lg:text-[50px] text-3xl text-textColor font-semibold">Tin tức và Bài viết</p>
            </span>
            <Link href="/news">
              <button className="bg-orange text-white text-xs font-bold rounded-md px-8 h-12 hoverBtn">
                XEM TOÀN BỘ
              </button>
            </Link>
          </div>
          <div className="pt-16">
            <Carousel
              partialVisible={false}
              swipeable
              draggable={false}
              responsive={responsive}
              ssr={false}
              infinite
              autoPlay
              arrows
              keyBoardControl
              itemClass="carouselItem"
            >
              {featuredArticles.map((article, index) => (
                <HomeNewsCard key={index} {...article} />
              ))}
              {/* <HomeNewsCard
                slug="nhung-dieu-can-xem-va-lam-khi-den-tham-nhat-ban"
                image="/News/home_post-1.jpg"
                title="Những điều cần xem và làm khi đến thăm Nhật Bản"
                date={{ day: "08", month: "Tháng 12" }}
                author="Phạm Nhật Dương"
                comments={0}
                description="Khám phá văn hóa Nhật Bản với những địa điểm đẹp và những hoạt động thú vị."
              />

              <HomeNewsCard
                slug="mot-noi-de-bat-dau-cuoc-song-moi-voi-su-binh-yen"
                image="/News/home_post-3.jpg"
                title="Một nơi để bắt đầu cuộc sống mới với sự bình yên"
                date={{ day: "15", month: "Tháng 12" }}
                author="Nguyễn Văn An"
                comments={3}
                description="Địa điểm lý tưởng để nghỉ ngơi và tận hưởng cuộc sống chậm lại."
              />

              <HomeNewsCard
                slug="nhung-hanh-trinh-duoc-do-luong-tot-nhat-voi-ban-be"
                image="/News/home_post-4.jpg"
                title="Những hành trình được đo lường tốt nhất với bạn bè"
                date={{ day: "20", month: "Tháng 12" }}
                author="Trần Thị Bảo"
                comments={5}
                description="Cùng bạn bè khám phá những hành trình đầy ý nghĩa và kỷ niệm đáng nhớ."
              />

              <HomeNewsCard
                slug="du-lich-den-nhung-noi-dep-nhat-tren-the-gioi"
                image="/News/home_post-2.jpg"
                title="Du lịch đến những nơi đẹp nhất trên thế giới"
                date={{ day: "25", month: "Tháng 12" }}
                author="Lê Minh Cường"
                comments={2}
                description="Những địa danh nổi tiếng mang vẻ đẹp mê hoặc trên thế giới."
              /> */}

            </Carousel>
          </div>
        </div>
      </section>

      {/* Phần tham gia thành viên */}
      <div className="bg-[#faf5ee]" data-aos="fade-down">
        <div className="flex flex-wrap max-w-[1200px] xl:px-0 px-6 mx-auto lg:pt-28 pt-14">
          <div className="lg:w-1/2">
            <p className="text-xl text-orange pb-4">Tham gia với chúng tôi</p>
            <p className="lg:text-[50px] text-3xl text-textColor font-semibold">
              Bạn đã là thành viên chưa?
            </p>
            <p className="text-[#757783] py-4">
              Tham gia với chúng tôi! Các thành viên của chúng tôi có thể tiết kiệm lên đến 50%.
            </p>
            <div className="flex gap-4 pt-6">
              <Link href="/login">
                <button className="bg-orange text-white text-xs font-bold rounded flex gap-2 px-8 h-12 items-center hoverBtn hover:bg-orange hover:text-textColor">
                  <MdArrowCircleRight size={20} /> &nbsp;ĐĂNG NHẬP&nbsp;
                </button>
              </Link>
              <Link href="/signup">
                <button className="bg-white text-textColor text-xs font-bold rounded flex gap-2 px-8 h-12 items-center hoverBtn hover:bg-orange">
                  <MdPerson size={20} /> ĐĂNG KÝ
                </button>
              </Link>
            </div>
          </div>
          <div className="flex justify-center lg:w-1/2 w-full lg:mt-0 mt-14">
          <Image
            src="/image-app.png"
            alt="Image description"
            width={500}
            height={500}
            priority
          />
          </div>
        </div>
      </div>

    </div>
  );
}
