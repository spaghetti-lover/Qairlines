// components/NewsCards.js
import React from 'react';
import Link from 'next/link';
import { MdArrowRight, MdPerson } from "react-icons/md";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { cn } from "@/lib/utils";
import { CalendarIcon } from 'lucide-react';
import Image from 'next/image';
// HomeNewsCard Component
const HomeNewsCard = React.forwardRef(
  (
    { slug, image, title, date, author, comments, description, className, ...props },
    ref
  ) => (
    <Card ref={ref} className={cn("overflow-hidden", className)} {...props}>
      <div className="relative overflow-hidden rounded-t-lg">
        <Link href={`/news/${slug}`}>
          {/* <img src={image} alt={title} className="rounded-t-lg hoverImg" /> */}
          <Image 
            src={image} 
            alt={title} 
            // layout="intrinsic" 
            width={500} 
            height={100}
            className="w-full h-48 object-cover rounded-t-lg hoverImg"  
          />
          <p className="font-bold text-white bg-orange flex flex-col w-fit pt-3 pb-7 px-6 absolute right-0 bottom-0 rounded-tl-lg">
            {date.day} <span className="font-normal text-xs">{date.month}</span>
          </p>
        </Link>
      </div>
      <div className="border border-[#ebe6de] rounded-b-lg relative">
        <div className="absolute w-full h-5 -top-5 bg-white rounded-t-[20px]"></div>
        <div className="p-6">
          <div className="flex items-center gap-4">
            <span className="flex items-center gap-1 text-xs">
              <MdPerson className="text-orange text-base" /> {author}
            </span>
          </div>
          <Link href={`/news/${slug}`}>
            <CardTitle className="text-2xl font-semibold py-2 hover:text-orange">
              {title}
            </CardTitle>
          </Link>
          <CardContent className="p-0">
            <p className="pt-2 leading-8">{description}</p>
          </CardContent>
          <CardFooter className="p-0">
            <Link href={`/news/${slug}`} className="flex items-center gap-2 text-orange text-sm font-bold mt-2">
              ĐỌC THÊM <MdArrowRight />
            </Link>
          </CardFooter>
        </div>
      </div>
    </Card>
  )
);
HomeNewsCard.displayName = "HomeNewsCard";

// NewsCard Component
const NewsCard = React.forwardRef(
  (
    { slug, image, title, description, content, buttonText, className, ...props },
    ref
  ) => (
    <Card ref={ref} className={cn(className)} {...props}>
      <Link href={`/news/${slug}`}>
        {image && (
          // <img src={image} alt={title} className="w-full h-48 object-cover rounded-t-lg" />
          <Image 
            src={image} 
            alt={title} 
            width={500}
            height={192}
            className="w-full h-48 object-cover rounded-t-lg" 
          />
        )}
        <CardHeader>
          <CardTitle className="text-xl font-semibold">{title}</CardTitle>
          {description && <CardDescription>{description}</CardDescription>}
        </CardHeader>
      </Link>
      {content && (
        <CardContent>
          <p className="text-gray-600 dark:text-gray-300 line-clamp-3">{content}</p>
        </CardContent>
      )}
      <CardFooter>
        <Button asChild variant="orange">
          <Link href={`/news/${slug}`}>
            {buttonText}
          </Link>
        </Button>
      </CardFooter>
    </Card>
  )
);
NewsCard.displayName = "NewsCard";


// FeaturedNewsCard Component
const FeaturedNewsCard = React.forwardRef(
  (
    { slug, image, title, description, date, buttonText, className, ...props },
    ref
  ) => (
    <div
      ref={ref}
      className={cn(
        "relative h-[400px] md:h-[500px] overflow-hidden rounded-lg shadow-lg",
        className
      )}
      {...props}
    >
      <Link href={`/news/${slug}`}>
      <Image
        src={image}
        alt={title}
        layout="fill"
        objectFit="cover"
        className="object-cover"
      />
      </Link>
      <div className="absolute inset-0 bg-gradient-to-t from-black to-transparent"></div>
      <div className="absolute bottom-0 left-0 right-0 p-6 text-white">
        <Link href={`/news/${slug}`}>
          <h3 className="text-2xl md:text-4xl font-bold mb-2 leading-tight">
            {title}
          </h3>
        </Link>
        <p className="text-lg mb-4 line-clamp-2 md:line-clamp-3">
          {description}
        </p>
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <CalendarIcon className="h-5 w-5" />
            <span className="text-sm">{date}</span>
          </div>
          <Button asChild variant="orange">
            <Link href={`/news/${slug}`}>
              {buttonText}
            </Link>
          </Button>
        </div>
      </div>
    </div>
  )
);
FeaturedNewsCard.displayName = "FeaturedNewsCard";
export { HomeNewsCard, FeaturedNewsCard, NewsCard };
