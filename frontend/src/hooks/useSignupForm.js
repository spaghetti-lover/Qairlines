import { toast } from "@/hooks/use-toast";
import { useState } from "react";
const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;
export const useSignup = (onSuccess) => {
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    confirmPassword: '',
  });
  const [loading, setLoading] = useState(false);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      toast({
        title: "Lỗi đăng ký",
        description: "Mật khẩu và xác nhận mật khẩu không khớp.",
        variant: "destructive",
      });
      return;
    }

    setLoading(true);

    try {
      const response = await fetch(`${API_BASE_URL}/api/user`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          first_name: formData.firstName,
          last_name: formData.lastName,
          email: formData.email,
          password: formData.password,
        }),
      });

      if (response.ok) {
        const data = await response.json();
        toast({
          title: "Đăng ký thành công!",
          description: "Tài khoản đã được tạo thành công.",
        });
        if (onSuccess) onSuccess();
      } else {
        const errorData = await response.json();
        toast({
          title: "Lỗi đăng ký",
          description: errorData.message || "Không thể tạo tài khoản.",
          variant: "destructive",
        });
      }
    } catch (error) {
      console.error('Lỗi khi đăng ký:', error);
      toast({
        title: "Lỗi hệ thống",
        description: "Đã xảy ra lỗi khi kết nối đến máy chủ.",
        variant: "destructive",
      });
    } finally {
      setLoading(false);
    }
  };

  return {
    formData,
    loading,
    handleInputChange,
    handleSubmit,
  };
};
