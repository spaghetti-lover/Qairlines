import { useState } from 'react';
import { useRouter } from 'next/router';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Plane, TicketIcon, UserCircle } from 'lucide-react';
import SearchForm from "./SearchFlightsForm";

const FlightBookingTabs = () => {
  const router = useRouter();

  // State cho các tab khác nhau
  const [bookingID, setBookingID] = useState(''); // Quản lý mã đặt chỗ
  const [email, setEmail] = useState('');

  const handleSearch = (data) => {
    router.push({
      pathname: '/flights',
      query: {
        fromAirport: data.fromAirport,
        toAirport: data.toAirport,
        departureDate: data.departureDate,
        returnDate: data.returnDate || null,
        tripType: data.tripType,
        passengerCount: data.passengerCount,
      },
    });
  };

  const handleBookingManagement = () => {
    if (!bookingID || !email) {
      alert('Vui lòng nhập mã đặt chỗ và email!');
      return;
    }

    router.push({
      pathname: '/booking-management',
      query: {
        bookingID: bookingID,
        email: email,
      },
    });
  };

  const handleCheckIn = () => {
    if (!bookingID || !email) {
      alert('Vui lòng nhập mã đặt chỗ và email!');
      return;
    }

    router.push({
      pathname: '/check-in',
      query: {
        bookingID: bookingID,
        email: email,
      },
    });
  };

  return (
    <div className="w-full flex items-center justify-center">
      <div className="w-full mx-auto p-3 bg-white shadow-lg rounded-lg">
        <Tabs defaultValue="buy" className="w-full">
          {/* Tab Navigation */}
          <TabsList className="grid w-full grid-cols-3 bg-gray text-textColor h-auto rounded-t-lg">
            <TabsTrigger
              value="buy"
              className="flex items-center gap-2 py-4 text-textColor data-[state=active]:bg-orange data-[state=active]:text-white"
            >
              <Plane className="h-5 w-5" />
              MUA VÉ
            </TabsTrigger>
            <TabsTrigger
              value="manage"
              className="flex items-center gap-2 py-4 text-textColor data-[state=active]:bg-orange data-[state=active]:text-white"
            >
              <TicketIcon className="h-5 w-5" />
              QUẢN LÝ ĐẶT VÉ
            </TabsTrigger>
            <TabsTrigger
              value="checkin"
              className="flex items-center gap-2 py-4 text-textColor data-[state=active]:bg-orange data-[state=active]:text-white"
            >
              <UserCircle className="h-5 w-5" />
              LÀM THỦ TỤC
            </TabsTrigger>
          </TabsList>

          {/* Mua Vé Tab */}
          <TabsContent value="buy" className="mt-2">
            <SearchForm onSearch={handleSearch} />
          </TabsContent>

          {/* Quản Lý Đặt Chỗ Tab */}
          <TabsContent value="manage" className="mt-2">
            <div className="flex flex-col gap-2">
              <Input
                placeholder="Mã đặt chỗ/Số vé điện tử"
                value={bookingID}
                onChange={(e) => setBookingID(e.target.value)}
              />
              <Input
                placeholder="Hòm thư điện tử"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              <Button
                className="w-full bg-orange"
                onClick={handleBookingManagement}
              >
                TÌM KIẾM
              </Button>
            </div>
          </TabsContent>

          {/* Làm Thủ Tục Tab */}
          <TabsContent value="checkin" className="mt-2">
            <div className="flex flex-col gap-2">
              <Input
                placeholder="Mã đặt chỗ"
                value={bookingID}
                onChange={(e) => setBookingID(e.target.value)}
              />
              <Input
                placeholder="Hòm thư điện tử"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              <Button
                className="w-full bg-orange"
                onClick={handleCheckIn}
              >
                LÀM THỦ TỤC
              </Button>
            </div>
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
};

export default FlightBookingTabs;
