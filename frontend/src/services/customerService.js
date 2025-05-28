import axios from "axios";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export const fetchCustomerInfo = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("Token không tồn tại. Vui lòng đăng nhập lại.");
    }

    const response = await axios.get(`${API_BASE_URL}/api/customer`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data.data;
  } catch (error) {
    console.error("Lỗi khi lấy thông tin khách hàng:", error);
    throw error.response?.data?.message || "Đã xảy ra lỗi. Vui lòng thử lại.";
  }
};

export const updateCustomerInfo = async (updateData) => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("Token không tồn tại. Vui lòng đăng nhập lại.");
    }

    const response = await axios.put(`${API_BASE_URL}/api/user`, updateData, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data.message;
  } catch (error) {
    console.error("Lỗi khi cập nhật thông tin khách hàng:", error);
    throw error.response?.data?.message || "Đã xảy ra lỗi. Vui lòng thử lại.";
  }
};
