// components/Hero.js
import {
  MdUpgrade,
  MdEventSeat,
  MdShoppingBag,
  MdHotel,
  MdMiscellaneousServices,
} from "react-icons/md";
import dynamic from "next/dynamic";
import "react-multi-carousel/lib/styles.css";
import FlightBookingTabs from '@/components/FlightsSearch/FlightBookingTabs';


const Carousel = dynamic(() => import("react-multi-carousel"), { ssr: false });

const responsive = {
  superLargeDesktop: {
    breakpoint: { max: 4000, min: 3000 },
    items: 5,
  },
  desktop: {
    breakpoint: { max: 3000, min: 1024 },
    items: 5,
  },
  tablet: {
    breakpoint: { max: 1024, min: 464 },
    items: 3,
  },
  mobile: {
    breakpoint: { max: 464, min: 0 },
    items: 1,
  },
};

export default function Hero() {
  return (
    <div className="relative bg-black lg:h-[80vh]" data-aos="fade-down" data-aos-delay="300" data-aos-duration="3000">
      <video
        autoPlay
        muted
        loop
        loading="lazy"
        className="absolute z-10 w-full h-full lg:top-0 -top-[12vh] object-cover opacity-55"
      >
        <source src="/video.mp4" type="video/mp4" />
      </video>
    

      <div className="relative z-20 lg:h-full h-screen max-w-[1200px] px-2 lg:pt-0 pt-16 mx-auto flex flex-col items-center justify-center">
        <div className="text-center mb-8">
          <p className="text-3xl text-orange">Hãy cùng khám phá</p>
          <h4 className="lg:text-[52px] text-3xl text-white mt-5">
            Bạn muốn bay đến những đâu?
          </h4>
          <p className="text-gray text-2xl my-8">
            Khám phá những điểm đến tuyệt đẹp.
          </p>
        </div>

        <div className="w-full flex items-center justify-center">
          {/* <SearchForm /> */}
          <FlightBookingTabs />
        </div>
      </div>

      <div className="py-16 lg:-mt-24 relative z-10 max-w-[1200px] px-6 mx-auto">
      <Carousel
            partialVisible={false}
            swipeable={true}
            draggable={false} 
            responsive={responsive}
            ssr={true}
            infinite
            autoPlay={true}
            arrows={true}
            keyBoardControl={true}
            itemClass="carouselItem"
          >
            <div className="shadow-xl bg-white p-8 mb-10 rounded-lg flex flex-col justify-center items-center gap-4">
              <MdUpgrade className="rounded-full w-20 h-20 p-4 bg-[#e5faf5] text-[#3fd2a8] hover:bg-orange hover:text-white" />
              <p className="font-bold">Nâng cấp vé</p>
            </div>

            <div className="shadow-xl bg-white p-8 mb-10 rounded-lg flex flex-col justify-center items-center gap-4">
              <MdEventSeat className="rounded-full w-20 h-20 p-4 bg-[#26B2EC24] text-[#06aff6] hover:bg-orange hover:text-white" />
              <p className="font-bold">Chọn chỗ ngồi</p>
            </div>

            <div className="shadow-xl bg-white p-8 mb-10 rounded-lg flex flex-col justify-center items-center gap-4">
              <MdShoppingBag className="rounded-full w-20 h-20 p-4 bg-[#f5ecfd] text-[#9e60e5] hover:bg-orange hover:text-white" />
              <p className="font-bold">Mua sắm</p>
            </div>

            <div className="shadow-xl bg-white p-8 mb-10 rounded-lg flex flex-col justify-center items-center gap-4">
              <MdHotel className="rounded-full w-20 h-20 p-4 bg-[#fff4de] text-[#f6b23b] hover:bg-orange hover:text-white" />
              <p className="font-bold">Khách sạn</p>
            </div>

            <div className="shadow-xl bg-white p-8 mb-10 rounded-lg flex flex-col justify-center items-center gap-4">
              <MdMiscellaneousServices className="rounded-full w-20 h-20 p-4 bg-[#D036321C] text-[#d03632] hover:bg-orange hover:text-white" />
              <p className="font-bold">Những dịch vụ khác</p>
            </div>
          </Carousel>
      </div>
    
    </div>
  );
}
