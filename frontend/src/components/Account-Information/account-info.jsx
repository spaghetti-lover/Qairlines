import Image from "next/image";
import { Card } from "@/components/ui/card";

export default function AccountInfo({
  personalInfo,
  setPersonalInfo,
  isEditing,
  setIsEditing,
  handleUpdate,
}) {
  if (!personalInfo) return <p>Không tìm thấy thông tin cá nhân.</p>;

  const uid = personalInfo.uid || "";
  const firstName = personalInfo.firstName || "";
  const lastName = personalInfo.lastName || "";
  const fullName = `${firstName} ${lastName}`.trim();
  const email = personalInfo.email || "";
  const address = personalInfo.address || "";
  const phoneNumber = personalInfo.phoneNumber || "";
  const passportNumber = personalInfo.passportNumber || "";

  let genderLabel = "";
  switch (personalInfo.gender) {
    case "male":
      genderLabel = "Nam";
      break;
    case "female":
      genderLabel = "Nữ";
      break;
    default:
      genderLabel = "Khác";
      break;
  }

  let birthDateDisplay = "";
  let birthDateValue = "";
  if (personalInfo.dateOfBirth) {
    const dateObj = new Date(personalInfo.dateOfBirth);

    if (dateObj) {
      const day = dateObj.getDate().toString().padStart(2, "0");
      const month = (dateObj.getMonth() + 1).toString().padStart(2, "0");
      const year = dateObj.getFullYear();
      birthDateDisplay = `${day} tháng ${month} năm ${year}`;
      birthDateValue = `${year}-${month}-${day}`;
    }
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    const parts = fullName.split(" ");
    const lastNameUpdate = parts.pop() || "";
    const firstNameUpdate = parts.join(" ");

    const isoDate = birthDateValue ? `${birthDateValue}T00:00:00.000Z` : null;

    const updatedData = {
      ...personalInfo,
      firstName: firstNameUpdate,
      lastName: lastNameUpdate,
      address: address,
      phoneNumber: phoneNumber,
      passportNumber: passportNumber,
      gender: personalInfo.gender || "other",
      dateOfBirth: isoDate,
    };

    handleUpdate(updatedData);
  };

  return (
    <div className="p-4 sm:p-6">
      <h2 className="text-xl font-medium mb-4 sm:mb-6">Thông tin cá nhân</h2>

      <div className="flex flex-col lg:flex-row lg:space-x-8 space-y-6 lg:space-y-0">
        <div className="w-full lg:w-1/3">
          <Card className="p-0 bg-gradient-to-r from-sky-100 to-sky-200 shadow-lg rounded-lg relative h-48 sm:h-64">
            <Image
              src="/QAirline-card.png"
              alt="Lotus Miles"
              fill
              style={{ objectFit: "cover" }}
              className="rounded-lg"
            />
          </Card>
        </div>

        <div className="w-full lg:w-2/3">
          {isEditing ? (
            <form
              className="grid grid-cols-1 md:grid-cols-2 gap-4"
              onSubmit={handleSubmit}
            >
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Mã thẻ hội viên</label>
                <input
                  type="text"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={uid}
                  disabled
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Họ và tên</label>
                <input
                  type="text"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={fullName}
                  onChange={(e) => {
                    const value = e.target.value.trim();
                    const nameParts = value.split(" ");
                    const last = nameParts.pop() || "";
                    const first = nameParts.join(" ");
                    setPersonalInfo({
                      ...personalInfo,
                      firstName: first,
                      lastName: last,
                    });
                  }}
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Email</label>
                <input
                  type="email"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={email}
                  disabled
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Địa chỉ</label>
                <input
                  type="text"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={address}
                  onChange={(e) =>
                    setPersonalInfo({ ...personalInfo, address: e.target.value })
                  }
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Số điện thoại</label>
                <input
                  type="text"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={phoneNumber}
                  onChange={(e) =>
                    setPersonalInfo({
                      ...personalInfo,
                      phoneNumber: e.target.value,
                    })
                  }
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Giới tính</label>
                <select
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={personalInfo.gender || "other"}
                  onChange={(e) =>
                    setPersonalInfo({
                      ...personalInfo,
                      gender: e.target.value,
                    })
                  }
                >
                  <option value="male">Nam</option>
                  <option value="female">Nữ</option>
                  <option value="other">Khác</option>
                </select>
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Ngày sinh</label>
                <input
                  type="date"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={birthDateValue}
                  onChange={(e) => {
                    const val = e.target.value;
                    const newDate = val ? `${val}T00:00:00.000Z` : null;
                    setPersonalInfo({ ...personalInfo, dateOfBirth: newDate });
                  }}
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm text-gray-600">Số hộ chiếu</label>
                <input
                  type="text"
                  className="block w-full px-3 py-2 border border-gray-300 rounded-lg"
                  value={passportNumber}
                  onChange={(e) =>
                    setPersonalInfo({
                      ...personalInfo,
                      passportNumber: e.target.value,
                    })
                  }
                />
              </div>
              <div className="md:col-span-2 pt-2 flex space-x-4">
                <button
                  type="submit"
                  className="px-6 py-2 bg-orange text-white rounded-lg hover:bg-orangeLight transition-all"
                >
                  Lưu thông tin
                </button>
                <button
                  type="button"
                  onClick={() => setIsEditing(false)}
                  className="px-6 py-2 bg-gray-300 text-black rounded-lg hover:bg-gray-400 transition-all"
                >
                  Hủy
                </button>
              </div>
            </form>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-x-4 gap-y-6">
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Mã thẻ hội viên</div>
                <div className="font-semibold">{uid}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Họ và tên</div>
                <div className="font-semibold">{fullName}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Email</div>
                <div className="font-semibold">{email}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Địa chỉ</div>
                <div className="font-semibold">{address}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Số điện thoại</div>
                <div className="font-semibold">{phoneNumber}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Giới tính</div>
                <div className="font-semibold">{genderLabel}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Ngày sinh</div>
                <div className="font-semibold">{birthDateDisplay}</div>
              </div>
              <div className="space-y-1">
                <div className="text-sm text-gray-600">Số hộ chiếu</div>
                <div className="font-semibold">{passportNumber}</div>
              </div>
              <div className="md:col-span-2 pt-2">
                <button
                  onClick={() => setIsEditing(true)}
                  className="w-full sm:w-auto px-6 py-2 bg-orange text-white rounded-lg hover:bg-orangeLight transition-all"
                >
                  Cập nhật thông tin
                </button>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
