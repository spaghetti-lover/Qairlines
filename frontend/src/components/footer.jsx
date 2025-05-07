// components/Footer.js
import {
  MdPhone,
  MdOutlineMailOutline,
  MdOutlineLocationOn,
} from "react-icons/md";
import { Checkbox } from "@/components/ui/checkbox"
import Image from "next/image";

export default function Footer() {
  return (
    <footer className="bg-[url(/bg-map-2.png)] bg-[#313041] py-16 bg-contain">
      <div className="text-[#A9A8B6] flex lg:flex-nowrap flex-wrap justify-between max-w-[1200px] xl:px-0 px-6 mx-auto gap-10">
        
        {/* Phần thông tin liên hệ */}
        <div className="lg:w-1/3">
        <Image 
          src="/logo-white.png" 
          alt="Logo" 
          width={130} 
          height={130} 
        />
          <p className="my-5">
            Chào mừng đến với QAirline. Chúng tôi cung cấp dịch vụ đặt vé máy bay trực tuyến hàng đầu thế giới.
          </p>
          <div className="border-t border-[#A9A8B6] pt-[15px]">
            <ul className="flex flex-col gap-4">
              <a href="#" className="flex items-center gap-2 hover:text-orange">
                <MdPhone className="text-orange" /> +84 123 456 789
              </a>
              <a href="#" className="flex items-center gap-2 hover:text-orange">
                <MdOutlineMailOutline className="text-orange" /> qairline-support@qairline.website
              </a>
              <a href="#" className="flex items-center gap-2 hover:text-orange">
                <MdOutlineLocationOn className="text-orange" /> Nhà E3, 144 Xuân Thủy, quận Cầu Giấy, Hà Nội , Việt Nam
              </a>
            </ul>
          </div>
        </div>

          <div className="lg:w-1/5">
            <h6 className="text-white text-xl font-bold my-5 lg:mt-0">Công ty</h6>
            <ul className="flex flex-col gap-4">
              <a href="#" className="hover:text-white">Về chúng tôi</a>
              <a href="#" className="hover:text-white">Blog cộng đồng</a>
              <a href="#" className="hover:text-white">Phần thưởng</a>
              <a href="#" className="hover:text-white">Làm việc với chúng tôi</a>
              <a href="#" className="hover:text-white">Liên hệ</a>
            </ul>
          </div>

          {/* Phần khám phá */}
          <div className="lg:w-1/5">
            <h6 className="text-white text-xl font-bold my-5 lg:mt-0">Khám phá</h6>
            <ul className="flex flex-col gap-4">
              <a href="#" className="hover:text-white">Tài khoản</a>
              <a href="#" className="hover:text-white">Chính sách bảo mật</a>
              <a href="#" className="hover:text-white">Chương trình liên kết</a>
              <a href="#" className="hover:text-white">Đối tác của chúng tôi</a>
              <a href="#" className="hover:text-white">Sự kiện</a>
            </ul>
          </div>

          {/* Phần đăng ký nhận tin */}
        <div className="flex flex-col lg:w-1/3">
          <h6 className="text-white text-xl font-bold my-5 lg:mt-0">Nhận tin tức</h6>
          <input
            type="email"
            placeholder="Hòm thư điện tử"
            className="bg-[#00000033] rounded-lg text-center hover:bg-[#0000004d] h-14 outline-none"
          />
          <button className="mt-2 mb-4 bg-orange rounded-lg text-center text-white h-14 outline-none hover:opacity-90">
            ĐĂNG KÝ
          </button>
          <label className="text-[#A9A8B6]">
            <input type="checkbox" className="mr-2" />Tôi đồng ý với tất cả các điều khoản và chính sách
          </label>
        </div>
      </div>
    </footer>
  );
}
