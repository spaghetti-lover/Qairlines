import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import SeatSelector from "@/components/checkin/Seat-selector";
import { ArrowLeftRight } from "lucide-react";
import { useToast } from "@/hooks/use-toast"; // Import useToast

export function SeatSelectionStep({
  seats,
  passengers,
  onSeatSelect,
  onContinue,
  onBack,
  onSwitchTrip,
  currentTrip,
}) {
  const { toast } = useToast(); // Sử dụng useToast

  const handleContinue = () => {
    const allSelected = passengers.every((customer) => customer.seat);
    if (!allSelected) {
      // Thay alert bằng toast
      toast({
        title: "Chưa chọn đủ ghế",
        description:
          "Vui lòng chọn ghế cho tất cả hành khách trước khi tiếp tục.",
        variant: "destructive",
      });
    } else {
      onContinue();
    }
  };

  return (
    <Card className="w-full max-w-4xl mx-auto">
      <CardContent className="p-6">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-semibold">
            Chọn Ghế Ngồi (
            {currentTrip === "departure" ? "Chuyến đi" : "Chuyến về"})
          </h2>
          {onSwitchTrip && (
            <Button
              variant="outline"
              onClick={onSwitchTrip}
              className="flex items-center gap-2 hover:bg-gray-100"
            >
              <ArrowLeftRight className="w-4 h-4" />
              <span className="hidden sm:inline">
                {currentTrip === "departure"
                  ? "Chuyển sang Chuyến về"
                  : "Chuyển sang Chuyến đi"}
              </span>
              <span className="sm:hidden">Đổi chuyến</span>
            </Button>
          )}
        </div>

        <SeatSelector
          seats={seats}
          passengers={passengers}
          onSeatSelect={onSeatSelect}
        />

        <div className="flex justify-between gap-4 mt-6">
          <Button variant="outline" onClick={onBack}>
            Quay Lại
          </Button>
          <Button variant="orange" onClick={handleContinue}>
            Tiếp Tục
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
