import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import Image from 'next/image';
import Head from 'next/head';
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { ArrowLeftIcon, CalendarIcon } from "lucide-react";
import { toast } from '@/hooks/use-toast';
import mockNewsDataService from '@/pages/mockNewsData';

const NewsDetail = () => {
  const router = useRouter();
  const { slug } = router.query;
  const [article, setArticle] = useState(null);
  const [relatedArticles, setRelatedArticles] = useState([]);

  useEffect(() => {
    if (slug) {
      getNewsDetail();
      getRelatedNews();
    }
  }, [slug]);

  const getNewsDetail = async () => {
    try {
      const response = await mockNewsDataService.getNewsById(slug);
      const articleData = response.data;
      setArticle({
        slug: articleData.id,
        image: articleData.image,
        title: articleData.title,
        description: articleData.description,
        author: articleData.authorName,
        content: articleData.content,
        date: new Date(articleData.createdAt).toLocaleDateString('vi-VN'),
        buttonText: "Đọc thêm",
        authorTitle: "Nhà báo",
        authorImage: "/AvatarUser/no_avatar.jpg",
      });
    } catch (error) {
      toast({
        title: "Lỗi",
        description: "Không thể tải bài viết. Vui lòng thử lại sau.",
        variant: "destructive"
      });
      router.push('/news');
    }
  };

  const getRelatedNews = async () => {
    try {
      const response = await mockNewsDataService.getAllNews();
      const allArticles = response.data
        .filter(item => item.id !== slug)
        .slice(0, 3)
        .map(item => ({
          slug: item.id,
          image: item.image,
          title: item.title,
          description: item.description,
          author: item.authorName,
          content: item.content,
          date: new Date(item.createdAt).toLocaleDateString('vi-VN'),
        }));
      setRelatedArticles(allArticles);
    } catch (error) {
      console.error("Lỗi khi tải bài viết liên quan:", error);
    }
  };

  if (!article) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 dark:bg-gray-900">
      <Head>
        <title>{`${article.title} | QAirlines News`}</title>
        <meta name="description" content={article.description} />
      </Head>

      <main className="container mx-auto px-4 py-8">
        <article className="max-w-4xl mx-auto">
          <Button variant="ghost" className="mb-4 -ml-2" onClick={() => router.back()}>
            <ArrowLeftIcon className="mr-2 h-4 w-4" />
            Quay lại trang trước
          </Button>

          <header className="mb-8">
            <h1 className="text-3xl md:text-5xl font-bold mb-4 text-gray-900 dark:text-white">
              {article.title}
            </h1>
            <div className="flex items-center space-x-4 text-gray-600 dark:text-gray-300 mb-4">
              <div className="flex items-center">
                <CalendarIcon className="h-5 w-5 mr-2" />
                <span>{article.date}</span>
              </div>
            </div>
            <div className="flex items-center space-x-4">
              <Avatar>
                <AvatarImage src={article.authorImage} alt={article.author} />
                <AvatarFallback>{article.author[0]}</AvatarFallback>
              </Avatar>
              <div>
                <p className="font-semibold text-gray-900 dark:text-white">{article.author}</p>
                <p className="text-sm text-gray-600 dark:text-gray-300">{article.authorTitle}</p>
              </div>
            </div>
          </header>

          <Image
            src={article.image}
            alt={article.title}
            width={800}
            height={400}
            className="w-full h-[300px] md:h-[400px] object-cover rounded-lg mb-8"
            priority
          />

          <div className="prose dark:prose-invert max-w-none mb-8">
            <p className="whitespace-pre-wrap text-gray-800 dark:text-gray-200">
              {article.content}
            </p>
          </div>

          {relatedArticles.length > 0 && (
            <div className="bg-white dark:bg-gray-800 rounded-lg p-6 mb-8">
              <h3 className="text-xl font-semibold mb-4 text-gray-900 dark:text-white">
                Bài viết liên quan
              </h3>
              <ul className="space-y-4">
                {relatedArticles.map((item) => (
                  <li key={item.slug} className="flex space-x-4">
                    <Link href={`/news/${item.slug}`} className="flex space-x-4 hover:opacity-80">
                      <Image
                        src={item.image}
                        alt={item.title}
                        width={70}
                        height={50}
                        className="object-cover rounded"
                      />
                      <div>
                        <h4 className="font-semibold text-gray-900 dark:text-white">{item.title}</h4>
                        <p className="text-sm text-gray-600 dark:text-gray-300">{item.description}</p>
                      </div>
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          )}
        </article>
      </main>
    </div>
  );
};

export default NewsDetail;