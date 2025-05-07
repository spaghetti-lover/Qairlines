// components/About.js
import { FaPlay } from "react-icons/fa";
import { GiDeer, GiHangGlider } from "react-icons/gi";
import { CiFlag1 } from "react-icons/ci";
import { MdOutlineParagliding } from "react-icons/md";
import { FaPlaneDeparture } from 'react-icons/fa';
import { MdFlight } from 'react-icons/md';
import { GiPriceTag } from 'react-icons/gi';
import { AiOutlineCheckCircle } from 'react-icons/ai'
import Image from 'next/image';
export default function About() {
  return (
    <div>
      <div className="bg-[url(/bg-1.jpg)] bg-cover bg-center bg-no-repeat py-24 w-full">
        <div className="lg:flex gap-12 max-w-[1200px] px-6 mx-auto">
          <div className="lg:w-1/2 w-full">
            <span data-aos="fade-up-right">
              <button className="animate-pulse rounded-lg bg-orange w-20 h-20 flex justify-center items-center mb-8">
                <FaPlay className="text-white" />
              </button>
              <p className="text-orange text-xl pb-2">
                Bạn đã sẵn sàng để bay chưa?
              </p>
            </span>
            <p
              className="lg:text-[50px] leading-normal text-2xl font-bold text-white lg:pb-0 pb-4"
              data-aos="fade-up-right"
            >
              QAirline là nền tảng đặt vé máy bay hàng đầu thế giới
            </p>
          </div>


          <div
            className="grid md:grid-cols-2 grid-cols-1 gap-2"
            data-aos="fade-up-right"
          >
            <div
              className="border border-gray rounded-lg p-10 flex flex-col items-center gap-4 text-orange hover:bg-orange hover:text-white transition-bg hover:border-transparent"
              data-aos="fade-down-right"
            >
              <FaPlaneDeparture className="text-2xl w-20 h-20" />
              <p className="text-white text-xl font-semibold">Flight Booking</p>
            </div>

            <div
              className="border border-gray rounded-lg p-10 flex flex-col items-center gap-4 text-orange hover:bg-orange hover:text-white transition-bg hover:border-transparent"
              data-aos="fade-down-right"
            >
              <MdFlight className="text-2xl w-20 h-20" />
              <p className="text-white text-xl font-semibold">Flight Status</p>
            </div>

            <div
              className="border border-gray rounded-lg p-10 flex flex-col items-center gap-4 text-orange hover:bg-orange hover:text-white transition-bg hover:border-transparent"
              data-aos="fade-down-right"
            >
              <GiPriceTag className="text-2xl w-20 h-20" />
              <p className="text-white text-xl font-semibold">Travel Deals</p>
            </div>

            <div
              className="border border-gray rounded-lg p-10 flex flex-col items-center gap-4 text-orange hover:bg-orange hover:text-white transition-bg hover:border-transparent"
              data-aos="fade-down-right"
            >
              <AiOutlineCheckCircle className="text-2xl w-20 h-20" />
              <p className="text-white text-xl font-semibold">Check-in Online</p>
            </div>
          </div>

        </div>
      </div>

      <div
        className="bg-[url(/bg-line.png)] bg-contain bg-bottom bg-orange py-[100px]"
        data-aos="fade-down"
      >
        <div className="flex lg:flex-nowrap flex-wrap justify-between gap-16 max-w-[1200px] xl:px-0 px-6 mx-auto">
          <p className="text-[40px] font-bold text-white whitespace-pre">
            Đối tác của chúng tôi
          </p>
          <div className="flex flex-wrap gap-8 justify-between w-full">
            <Image src="/vietnam_airline.png" alt="vietnam airline logo" width={300} height={300} loading="lazy" />
            <Image src="/vietjet.png" alt="vietjet logo" width={300} height={300} loading="lazy" />
            <Image src="/china_airline.png" alt="china airline logo" width={300} height={300} loading="lazy" />
            <Image src="/hongkong_airline.png" alt="hongkong airline logo" width={300} height={300} loading="lazy" />
          </div>
        </div>
      </div>
    </div>
  );
}
