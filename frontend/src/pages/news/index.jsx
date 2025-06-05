import React, { useEffect, useState } from 'react';
import { FeaturedNewsCard, NewsCard } from '@/components/NewsCards';
import Carousel from 'react-multi-carousel';
import 'react-multi-carousel/lib/styles.css';
import { useRouter } from 'next/router';
import { toast } from '@/hooks/use-toast';
import mockNewsDataService from '@/pages/mockNewsData';

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
  const [featuredArticles, setFeaturedArticles] = useState([]);

  useEffect(() => {
    getAllNews();
  }, []);

  const getAllNews = async () => {
    try {
      const response = await mockNewsDataService.getAllNews();
      setFeaturedArticles(response.data.map(article => ({
        slug: article.id,
        image: article.image,
        title: article.title,
        description: article.description,
        author: article.authorName,
        content: article.content,
        date: new Date(article.createdAt).toLocaleDateString('vi-VN'),
        buttonText: "Đọc thêm",
        authorTitle: "Nhà báo",
        authorImage: "/AvatarUser/no_avatar.jpg",
      })));
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Không thể tải tin tức. Vui lòng thử lại sau.",
        variant: "destructive"
      });
    }
  };

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
          {featuredArticles.slice(0, 3).map((article, index) => (
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
