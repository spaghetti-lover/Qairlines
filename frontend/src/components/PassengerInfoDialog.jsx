import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { format, parse, isValid, isAfter, isBefore } from "date-fns";

export function PassengerInfoDialog({ isOpen, onClose, passengerCount, onInfoFilled }) {
  // State để lưu thông tin hành khách
  const [passengers, setPassengers] = useState(
    Array.from({ length: passengerCount }, () => ({
      idNumber: "",
      lastName: "",
      firstName: "",
      phoneNumber: "",
      birthDate: "",
      gender: "",
      address: "",
    }))
  );

  // State để lưu lỗi của mỗi hành khách
  const [errors, setErrors] = useState(
    Array.from({ length: passengerCount }, () => ({}))
  );

  // Hàm xử lý thay đổi dữ liệu input
  const handleInputChange = (index, field, value) => {
    const updatedPassengers = [...passengers];
    updatedPassengers[index][field] = value;
    setPassengers(updatedPassengers);

    // Xóa lỗi nếu người dùng bắt đầu nhập lại
    if (errors[index][field]) {
      const updatedErrors = [...errors];
      delete updatedErrors[index][field];
      setErrors(updatedErrors);
    }
  };

  // Hàm validate ngày sinh
  const validateBirthDate = (date) => {
    const parsedDate = parse(date, "dd/MM/yyyy", new Date());
    if (!isValid(parsedDate)) {
      return "Ngày sinh không hợp lệ.";
    }
    if (isAfter(parsedDate, new Date())) {
      return "Ngày sinh không thể là ngày trong tương lai.";
    }
    if (isBefore(parsedDate, new Date(1900, 0, 1))) {
      return "Ngày sinh không thể trước năm 1900.";
    }
    return null;
  };

  // Hàm xử lý khi nhấn submit
  const handleSubmit = async (e) => {
    e.preventDefault();

    const newErrors = passengers.map((passenger) => {
      const passengerErrors = {};
      if (!passenger.idNumber) passengerErrors.idNumber = "Vui lòng nhập số CCCD.";
      if (!passenger.lastName) passengerErrors.lastName = "Vui lòng nhập họ.";
      if (!passenger.firstName) passengerErrors.firstName = "Vui lòng nhập tên.";
      if (!passenger.phoneNumber) passengerErrors.phoneNumber = "Vui lòng nhập số điện thoại.";
      if (!passenger.birthDate) {
        passengerErrors.birthDate = "Vui lòng nhập ngày sinh.";
      } else {
        const birthDateError = validateBirthDate(passenger.birthDate);
        if (birthDateError) passengerErrors.birthDate = birthDateError;
      }
      if (!passenger.gender) passengerErrors.gender = "Vui lòng chọn giới tính.";
      if (!passenger.address) passengerErrors.address = "Vui lòng nhập địa chỉ.";
      return passengerErrors;
    });

    setErrors(newErrors);

    // Kiểm tra nếu không có lỗi
    if (newErrors.every((error) => Object.keys(error).length === 0)) {
      try {
        // Gửi dữ liệu hành khách hợp lệ về component cha
        onInfoFilled(passengers);
        onClose(); // Đóng dialog
      } catch (error) {
        console.error("Lỗi khi gửi thông tin hành khách:", error);
      }
    }
  };


  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[600px] max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Thông tin hành khách</DialogTitle>
        </DialogHeader>

        <form onSubmit={handleSubmit}>
          {passengers.map((passenger, index) => (
            <div key={index} className="mb-6 p-4 border border-gray-200 rounded-lg">
              <h3 className="text-lg font-semibold mb-4">Hành khách {index + 1}</h3>

              <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                {/* Số CCCD */}
                <div>
                  <Label htmlFor={`idNumber-${index}`}>Số CCCD</Label>
                  <Input
                    id={`idNumber-${index}`}
                    value={passenger.idNumber}
                    onChange={(e) => handleInputChange(index, "idNumber", e.target.value)}
                    placeholder="123456789012"
                  />
                  {errors[index]?.idNumber && <p className="text-red-500 text-sm">{errors[index].idNumber}</p>}
                </div>

                {/* Họ */}
                <div>
                  <Label htmlFor={`lastName-${index}`}>Họ</Label>
                  <Input
                    id={`lastName-${index}`}
                    value={passenger.lastName}
                    onChange={(e) => handleInputChange(index, "lastName", e.target.value)}
                    placeholder="Nguyen"
                  />
                  {errors[index]?.lastName && <p className="text-red-500 text-sm">{errors[index].lastName}</p>}
                </div>

                {/* Tên */}
                <div>
                  <Label htmlFor={`firstName-${index}`}>Tên</Label>
                  <Input
                    id={`firstName-${index}`}
                    value={passenger.firstName}
                    onChange={(e) => handleInputChange(index, "firstName", e.target.value)}
                    placeholder="Van A"
                  />
                  {errors[index]?.firstName && <p className="text-red-500 text-sm">{errors[index].firstName}</p>}
                </div>

                {/* Số điện thoại */}
                <div>
                  <Label htmlFor={`phoneNumber-${index}`}>Số điện thoại</Label>
                  <Input
                    id={`phoneNumber-${index}`}
                    value={passenger.phoneNumber}
                    onChange={(e) => handleInputChange(index, "phoneNumber", e.target.value)}
                    placeholder="0987654321"
                  />
                  {errors[index]?.phoneNumber && <p className="text-red-500 text-sm">{errors[index].phoneNumber}</p>}
                </div>

                {/* Ngày sinh */}
                <div>
                  <Label htmlFor={`birthDate-${index}`}>Ngày sinh (dd/MM/yyyy)</Label>
                  <Input
                    id={`birthDate-${index}`}
                    value={passenger.birthDate}
                    onChange={(e) => handleInputChange(index, "birthDate", e.target.value)}
                    placeholder="01/01/1990"
                  />
                  {errors[index]?.birthDate && <p className="text-red-500 text-sm">{errors[index].birthDate}</p>}
                </div>

                {/* Giới tính */}
                <div>
                  <Label htmlFor={`gender-${index}`}>Giới tính</Label>
                  <Select
                    value={passenger.gender}
                    onValueChange={(value) => handleInputChange(index, "gender", value)}
                  >
                    <SelectTrigger id={`gender-${index}`}>
                      <SelectValue placeholder="Chọn giới tính" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="Male">Nam</SelectItem>
                      <SelectItem value="Female">Nữ</SelectItem>
                      <SelectItem value="Other">Khác</SelectItem>
                    </SelectContent>
                  </Select>
                  {errors[index]?.gender && <p className="text-red-500 text-sm">{errors[index].gender}</p>}
                </div>
              </div>

              {/* Địa chỉ */}
              <div className="mt-4">
                <Label htmlFor={`address-${index}`}>Địa chỉ</Label>
                <Input
                  id={`address-${index}`}
                  value={passenger.address}
                  onChange={(e) => handleInputChange(index, "address", e.target.value)}
                  placeholder="123 Nguyen Trai, Ha Noi"
                />
                {errors[index]?.address && <p className="text-red-500 text-sm">{errors[index].address}</p>}
              </div>
            </div>
          ))}
          <Button type="submit" variant="orange" className="w-full mt-4">
            Lưu thông tin
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
