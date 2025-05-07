import { useState } from "react";
import { useAuth } from "@/context/AuthContext";
import { toast } from "@/hooks/use-toast"; 
const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;
export const useLogin = (onSuccess) => {
  const { login } = useAuth(); // Lấy hàm login từ AuthContext
  const [formData, setFormData] = useState({ email: "", password: "" });
  const [loading, setLoading] = useState(false);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
      const response = await fetch(`${API_BASE_URL}/api/login/customer/`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      const data = await response.json();

      if (response.ok) {
        const { token } = data;
        login(token);
        toast({
          title: "Đăng nhập thành công!",
          description: "Chào mừng bạn trở lại.",
        });
        if (onSuccess) onSuccess();
      } else {
        toast({
          title: "Lỗi đăng nhập",
          description: data.message || "Đã xảy ra lỗi, vui lòng thử lại.",
          variant: "destructive",
        });
      }
    } catch (error) {
      console.error("Lỗi khi đăng nhập:", error);
      toast({
        title: "Lỗi hệ thống",
        description: "Không thể kết nối đến server. Vui lòng thử lại sau.",
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
