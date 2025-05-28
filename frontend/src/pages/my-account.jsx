import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import AccountInfo from "@/components/Account-Information/account-info";
import ActivityHistory from "@/components/Account-Information/activity-history";
import PasswordChange from "@/components/Account-Information/password-change";
import { Skeleton } from "@/components/ui/skeleton";
import { Card } from "@/components/ui/card";
import { useAccountInfo } from "@/hooks/useAccountInfo";
import { useEffect, useState } from "react";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export default function AccountPage() {
  const {
    personalInfo,
    setPersonalInfo,
    isEditing,
    setIsEditing,
    loading,
    errorMessage,
    handleUpdate,
  } = useAccountInfo();

  const [activityData, setActivityData] = useState([]);

  useEffect(() => {
    const fetchBookingHistory = async () => {
      if (!personalInfo || !personalInfo.bookingHistory) return;

      const token = localStorage.getItem("token");
      if (!token) {
        console.error("Token is missing. Please log in.");
        return;
      }

      const bookings = await Promise.all(
        personalInfo.bookingHistory.map(async (bookingId) => {
          try {
            const response = await fetch(
              `${API_BASE_URL}/api/booking?id=${bookingId}`,
              {
                headers: {
                  Authorization: `Bearer ${token}`,
                },
              }
            );

            if (!response.ok) {
              throw new Error("Failed to fetch booking.");
            }

            const { data } = await response.json();
            return {
              date: new Date(data.createdAt.seconds * 1000),
              type: "Đặt vé máy bay",
              bookingId: data.bookingId,
              details: `${data.departureCity} → ${data.arrivalCity}`,
            };
          } catch (error) {
            console.error(`Error fetching booking ${bookingId}:`, error);
            return null;
          }
        })
      );

      setActivityData(bookings.filter(Boolean));
    };

    fetchBookingHistory();
  }, [personalInfo]);

  if (loading) {
    return (
      <div className="p-4 sm:p-6">
        <h2 className="text-xl font-medium mb-4 sm:mb-6">Thông tin cá nhân</h2>
        <div className="flex flex-col lg:flex-row lg:space-x-8 space-y-6 lg:space-y-0">
          <div className="w-full lg:w-1/3">
            <Card className="p-0 bg-gradient-to-r from-sky-100 to-sky-200 shadow-lg rounded-lg relative h-48 sm:h-64">
              <Skeleton className="w-full h-full rounded-lg" />
            </Card>
          </div>

          <div className="w-full lg:w-2/3">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {/* Mã thẻ hội viên */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Họ và tên */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Email */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Địa chỉ */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Số điện thoại */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Giới tính */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Ngày sinh */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>
              {/* Số hộ chiếu */}
              <div className="space-y-1">
                <Skeleton className="h-5 w-1/2" />
                <Skeleton className="h-6 w-full" />
              </div>

              {/* Nút chỉnh sửa */}
              <div className="md:col-span-2 pt-2 flex space-x-4">
                <Skeleton className="h-10 w-32 rounded-lg" />
                <Skeleton className="h-10 w-20 rounded-lg" />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  if (errorMessage) {
    return <p className="text-red-600">{errorMessage}</p>;
  }

  return (
    <div className="min-h-screen bg-[#e6f2f7] bg-opacity-50 py-8">
      <div className="container mx-auto px-4">
        <h1 className="text-center text-2xl font-semibold mb-8 text-gray-800">
          THÔNG TIN TÀI KHOẢN
        </h1>

        <Tabs defaultValue="account" className="w-full">
          <TabsList className="w-full bg-orangeLight h-auto flex flex-wrap shadow-xl">
            <TabsTrigger
              value="account"
              className="flex-1 py-2 transition-all duration-300 ease-in-out"
            >
              Thông tin tài khoản
            </TabsTrigger>
            <TabsTrigger
              value="history"
              className="flex-1 py-2 transition-all duration-300 ease-in-out"
            >
              Lịch sử hoạt động
            </TabsTrigger>
            <TabsTrigger
              value="password"
              className="flex-1 py-2 transition-all duration-300 ease-in-out"
            >
              Thay đổi mật khẩu
            </TabsTrigger>
          </TabsList>

          <Card className="mt-4 border-0 shadow-xl">
            {/* Tab Thông tin tài khoản */}
            <TabsContent value="account">
              <AccountInfo
                personalInfo={personalInfo}
                setPersonalInfo={setPersonalInfo}
                isEditing={isEditing}
                setIsEditing={setIsEditing}
                handleUpdate={handleUpdate}
              />
            </TabsContent>

            {/* Tab Lịch sử hoạt động */}
            <TabsContent value="history">
              <ActivityHistory activityData={activityData} />
            </TabsContent>

            {/* Tab Thay đổi mật khẩu */}
            <TabsContent value="password">
              <PasswordChange personalInfo={personalInfo} />
            </TabsContent>
          </Card>
        </Tabs>
      </div>
    </div>
  );
}
