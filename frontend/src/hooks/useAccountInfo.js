// hooks/useAccountInfo.js
import { useState, useEffect } from "react";
import { fetchCustomerInfo, updateCustomerInfo } from "@/services/customerService";

export function useAccountInfo() {
  const [personalInfo, setPersonalInfo] = useState(null);
  const [isEditing, setIsEditing] = useState(false);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    const loadCustomerInfo = async () => {
      try {
        const data = await fetchCustomerInfo();
        setPersonalInfo(data);
      } catch (error) {
        setErrorMessage("Không thể tải thông tin cá nhân.");
      } finally {
        setLoading(false);
      }
    };
    loadCustomerInfo();
  }, []);

  const handleUpdate = async (updatedInfo) => {
    try {
      await updateCustomerInfo(updatedInfo);
      setPersonalInfo(updatedInfo);
      setIsEditing(false);
      alert("Thông tin đã được cập nhật thành công!");
    } catch (error) {
      setErrorMessage("Cập nhật thông tin thất bại. Vui lòng thử lại.");
    }
  };

  return {
    personalInfo,
    isEditing,
    loading,
    errorMessage,
    setIsEditing,
    handleUpdate,
    setPersonalInfo,
  };
}
