import React, { useEffect, useState } from 'react';
import { FeaturedNewsCard, NewsCard } from '@/components/NewsCards';
// import latestNews from '../../data/latestNews.json';
// import featuredArticles from '../../data/featuredArticles.json';
import Carousel from 'react-multi-carousel';
import 'react-multi-carousel/lib/styles.css';
import { useRouter } from 'next/router';
import { toast } from '@/hooks/use-toast';

// Định nghĩa responsive
const responsive = {
  superLargeDesktop: {
    breakpoint: { max: 4000, min: 1280 },
    items: 1,
  },
  desktop: {
    breakpoint: { max: 1280, min: 1024 },
    items: 1,
  },
  tablet: {
    breakpoint: { max: 1024, min: 640 },
    items: 1,
  },
  mobile: {
    breakpoint: { max: 640, min: 0 },
    items: 1,
  },
};

const NewsPage = () => {
  const router = useRouter();
  const [featuredArticles, setFeaturedArticles] = useState([])

  useEffect(() => {
    getAllNews()
  }, [router]);

  const getAllNews = async () => {
    const getAllNewsApi = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/news/all`

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
        "date": a.createdAt.seconds ? new Date(a.arrival_time).toISOString().replace("T", " ").slice(0, -5) : a.createdAt.split("T")[0],
        "buttonText": "Đọc thêm",
        "authorTitle": "Nhà báo",
        "authorImage": "/AvatarUser/no_avatar.jpg",
      }}))

    } catch (error) {

      toast({
        title: "Lỗi",
        description: "Đã có lỗi xảy ra khi kết nối với máy chủ, vui lòng tải lại trang hoặc đăng nhập lại",
        variant: "destructive"
      })
    }
  }
  return (
    <main className="container mx-auto px-4 py-8">
      {/* Phần Bài Viết Nổi Bật */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6 text-gray-900 dark:text-white">
          Bài viết nổi bật
        </h2>
        <Carousel
          swipeable
          draggable={false}
          showDots
          responsive={responsive}
          ssr={true}
          infinite
          autoPlay
          autoPlaySpeed={5000}
          keyBoardControl
          customTransition="all .5s"
          transitionDuration={500}
          containerClass="carousel-container"
          removeArrowOnDeviceType={['tablet', 'mobile']}
          dotListClass="custom-dot-list-style"
          itemClass="carousel-item-padding-40-px"
        >
          {featuredArticles.map((article, index) => (
            <FeaturedNewsCard key={index} {...article} />
          ))}
        </Carousel>
      </section>

      {/* Phần Tin Tức Mới Nhất */}
      <section>
        <h2 className="text-3xl font-bold mb-6 text-gray-900 dark:text-white">
          Tin tức mới nhất
        </h2>
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {featuredArticles.map((article, index) => (
            <NewsCard key={index} {...article} />
          ))}
        </div>
      </section>
    </main>
  );
};

export default NewsPage;
