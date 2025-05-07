import { useState } from "react"; 
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Slider } from "@/components/ui/slider";

/**
 * Component Filter: Dùng để chọn khoảng ngân sách và giờ khởi hành.
 * Nhận vào `filters` (state từ useFlightData) và `setFilters` (để cập nhật filters).
 */
export function FlightSideFilter({ filters, setFilters }) {
  // Sử dụng state cục bộ để người dùng điều chỉnh trước khi áp dụng chính thức
  const [localBudget, setLocalBudget] = useState(filters.budget); // Giá vé
  const [localDepartureTime, setLocalDepartureTime] = useState(filters.departureTime); // Giờ khởi hành

  // Hàm áp dụng bộ lọc (khi click "Áp dụng")
  const handleApplyFilters = () => {
    setFilters((prev) => ({
      ...prev,
      budget: localBudget,
      departureTime: localDepartureTime,
    }));
  };

  // Hàm thiết lập lại filter về giá trị mặc định
  const handleResetFilters = () => {
    const defaultBudget = [1929000, 6400000];
    const defaultDepartureTime = "all";
    setLocalBudget(defaultBudget);
    setLocalDepartureTime(defaultDepartureTime);
    setFilters({
      budget: defaultBudget,
      departureTime: defaultDepartureTime,
    });
  };

  return (
    <Card className="w-full md:w-64 h-fit">
      <CardContent className="p-4">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-lg font-semibold">Bộ lọc</h2>
        </div>

        {/* Bộ lọc ngân sách */}
        <div className="mb-4">
          <Label>Ngân sách</Label>
          <div className="flex justify-between text-sm mt-2">
            <span>{localBudget[0].toLocaleString()} VND</span>
            <span>{localBudget[1].toLocaleString()} VND</span>
          </div>
          <Slider
            min={100000}
            max={4000000}
            step={1000}
            value={localBudget}
            onValueChange={setLocalBudget}
            className="mt-2"
          />
        </div>

        {/* Bộ lọc giờ khởi hành */}
        <div className="mb-4">
          <Label>Giờ khởi hành</Label>
          <RadioGroup
            value={localDepartureTime}
            onValueChange={setLocalDepartureTime}
          >
            <div className="flex items-center space-x-2 mt-2">
              <RadioGroupItem value="all" id="all" />
              <Label htmlFor="all">Tất cả</Label>
            </div>
            <div className="flex items-center space-x-2 mt-2">
              <RadioGroupItem value="morning" id="morning" />
              <Label htmlFor="morning">00:00 - 11:59 Sáng</Label>
            </div>
            <div className="flex items-center space-x-2 mt-2">
              <RadioGroupItem value="afternoon" id="afternoon" />
              <Label htmlFor="afternoon">12:00 - 17:59 Chiều</Label>
            </div>
            <div className="flex items-center space-x-2 mt-2">
              <RadioGroupItem value="evening" id="evening" />
              <Label htmlFor="evening">18:00 - 23:59 Tối</Label>
            </div>
          </RadioGroup>
        </div>

        {/* Nút thao tác */}
        <div className="flex justify-between mt-4">
          <Button variant="outline" onClick={handleResetFilters}>
            Thiết lập lại
          </Button>
          <Button className="bg-orange text-white" onClick={handleApplyFilters}>
            Áp dụng
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
