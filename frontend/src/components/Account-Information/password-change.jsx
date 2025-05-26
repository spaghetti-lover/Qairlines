import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Skeleton } from "@/components/ui/skeleton";
import { toast } from "@/hooks/use-toast";
const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export default function PasswordChange({ personalInfo }) {
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!personalInfo) {
        toast({
            title: "Lỗi",
            description: "Không thể tải thông tin người dùng. Vui lòng thử lại.",
            variant: "destructive",
        });
        return;
    }

    if (newPassword !== confirmPassword) {
        toast({
            title: "Lỗi",
            description: "Mật khẩu mới và nhắc lại mật khẩu không khớp.",
            variant: "destructive",
        });
        return;
    }

    try {
        const token = localStorage.getItem("token");
        if (!token) {
            toast({
                title: "Lỗi xác thực",
                description: "Bạn cần đăng nhập để thực hiện thao tác này.",
                variant: "destructive",
            });
            return;
        }

        const response = await fetch(
            `${API_BASE_URL}/api/auth/id=${personalInfo.uid}/password`,
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify({
                    email: personalInfo.email,
                    oldPassword: currentPassword,
                    newPassword: newPassword,
                }),
            }
        );

        if (!response.ok) {
            const data = await response.json();
            throw new Error(data.message || "Có lỗi xảy ra, vui lòng thử lại.");
        }

        toast({
            title: "Thành công",
            description: "Mật khẩu đã được thay đổi thành công.",
        });

        // Reset các input về trạng thái rỗng
        setCurrentPassword("");
        setNewPassword("");
        setConfirmPassword("");
    } catch (err) {
        toast({
            title: "Lỗi",
            description: err.message,
            variant: "destructive",
        });
    }
};


  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h2 className="text-xl font-medium mb-6">Thông tin mật khẩu</h2>

      <p className="text-red-500 mb-6 text-sm">
        Mật khẩu tối thiểu phải có 8 ký tự, không giới hạn độ dài tối đa. Mật khẩu phải bao gồm ít nhất 1 ký tự số, 1 chữ cái hoa, 1 chữ cái thường và 1 ký tự đặc biệt (@ $ ! % * ? &). Ví dụ: Matkhau@123
      </p>

      <form className="space-y-4" onSubmit={handleSubmit}>
        <div>
          <label className="block mb-1">
            Mật khẩu hiện tại <span className="text-red-500">*</span>
          </label>
          <Input
            type="password"
            required
            value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)}
          />
        </div>

        <div>
          <label className="block mb-1">
            Mật khẩu mới <span className="text-red-500">*</span>
          </label>
          <Input
            type="password"
            required
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
          />
        </div>

        <div>
          <label className="block mb-1">
            Nhắc lại mật khẩu <span className="text-red-500">*</span>
          </label>
          <Input
            type="password"
            required
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
        </div>

        <div className="text-right">
          <Button type="submit" className="bg-orange text-white hover:bg-orangeLight">
            LƯU
          </Button>
        </div>
      </form>
    </div>
  );
}
