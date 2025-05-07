import Image from "next/image";
import { User } from "lucide-react";

export function FlightHeader({ departureCode, arrivalCode, departureCity, arrivalCity, departureDate, returnDate, passengers }) {
  return (
    <>
      <div className="relative h-[300px] w-full">
        <Image
          src="/tours_background.jpg"
          alt=""
          fill
          className="object-cover"
          priority
        />
        <div className="absolute inset-0 bg-black/30" />
        <div className="absolute inset-0 flex flex-col items-center justify-center text-white">
          <h1 className="text-4xl font-bold mb-2">Explore The Worlds</h1>
          <p className="text-lg">People Dont Take Trips Take People</p>
        </div>
      </div>

      {/* Flight Info Bar */}
      <div className="sticky top-24 border-t-2 border-orange bg-white shadow-md z-40">
        <div className="max-w-7xl mx-auto px-4 py-3">
          <div className="flex items-center justify-between flex-wrap gap-4 max-w-5xl m-auto">
            <div className="flex items-center gap-12">
              <div className="flex items-center gap-2">
                <span className="font-semibold">{departureCode}</span>
                <span className="text-gray-400">⋯⋯⋯⋯⋯⋯⋯⋯➜</span>
                <span className="font-semibold">{arrivalCode}</span>
              </div>
              <div className="text-sm text-gray-600">
                <div>{departureCity}</div>
                <div>{arrivalCity}</div>
              </div>
            </div>
            <div className="flex items-center gap-12">
              <div>
                <div className="text-sm text-gray-600">Chuyến đi</div>
                <div className="font-semibold">{departureDate}</div>
              </div>
              <div>
                <div className="text-sm text-gray-600">Chuyến về</div>
                <div className="font-semibold">{returnDate}</div>
              </div>
              <div>
                <div className="text-sm text-gray-600">Hành khách</div>
                <div className="flex items-center gap-1">
                  <User className="h-4 w-4" />
                  <span className="font-semibold">{passengers}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
