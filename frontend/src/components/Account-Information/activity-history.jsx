import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useState, useEffect } from "react";
import { format, isWithinInterval } from "date-fns";

export default function ActivityHistory({ activityData }) {
  const [fromDate, setFromDate] = useState(null);
  const [toDate, setToDate] = useState(null);
  const [filteredData, setFilteredData] = useState(activityData);

  useEffect(() => {
    setFilteredData(activityData);
  }, [activityData]);

  const handleSearch = () => {
    if (!fromDate || !toDate) {
      setFilteredData(activityData);
      return;
    }

    const filtered = activityData.filter((activity) =>
      isWithinInterval(activity.date, { start: fromDate, end: toDate })
    );

    setFilteredData(filtered);
  };

  const handleReset = () => {
    setFromDate(null);
    setToDate(null);
    setFilteredData(activityData);
  };

  return (
    <div className="p-6">
      <h2 className="text-xl font-medium mb-6">Lịch sử hoạt động</h2>

      <div className="flex gap-4 mb-6 items-end">
        <div className="flex-1">
          <label className="block text-sm mb-1">Từ ngày:</label>
          <Popover>
            <PopoverTrigger asChild>
              <button className="text-sm font-bold w-full justify-start text-left pl-0 border-b border-gray-400 py-1">
                {fromDate ? format(fromDate, "dd/MM/yyyy") : "Chọn ngày"}
              </button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0 border shadow-lg">
              <Calendar
                mode="single"
                selected={fromDate}
                onSelect={setFromDate}
                className="border-none shadow-none"
              />
            </PopoverContent>
          </Popover>
        </div>

        <div className="flex-1">
          <label className="block text-sm mb-1">Đến ngày:</label>
          <Popover>
            <PopoverTrigger asChild>
              <button className="text-sm font-bold w-full justify-start text-left pl-0 border-b border-gray-400 py-1">
                {toDate ? format(toDate, "dd/MM/yyyy") : "Chọn ngày"}
              </button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0 border shadow-lg">
              <Calendar
                mode="single"
                selected={toDate}
                onSelect={setToDate}
                className="border-none shadow-none"
              />
            </PopoverContent>
          </Popover>
        </div>

        <Button
          onClick={handleSearch}
          className="h-10 px-6 bg-orange text-white hover:bg-orangeLight transition-all"
        >
          TÌM KIẾM
        </Button>

        <Button
          onClick={handleReset}
          className="h-10 px-6 bg-gray-300 text-black hover:bg-gray-400 transition-all"
        >
          RESET
        </Button>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Ngày</TableHead>
            <TableHead>Loại giao dịch</TableHead>
            <TableHead>Mã đặt chỗ</TableHead>
            <TableHead>Chi tiết</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {filteredData.length > 0 ? (
            filteredData.map((activity, index) => (
              <TableRow key={index}>
                <TableCell>{format(activity.date, "dd/MM/yyyy")}</TableCell>
                <TableCell>{activity.type}</TableCell>
                <TableCell>{activity.bookingId}</TableCell>
                <TableCell>{activity.details}</TableCell>
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={4} className="text-center text-gray-500">
                Không có dữ liệu phù hợp.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}