import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Eye, Download } from "lucide-react";

export function FlightInfoCard({ passengerName, seatNumber, onView, onDownload }) {
  return (
    <Card className="w-full border-primary/20 shadow-sm hover:shadow-md transition-shadow">
      <CardContent className="p-4">
        <div className="flex justify-between items-center mb-4">
          <div>
            <div className="text-xs text-gray-500">Hành khách</div>
            <div className="text-sm font-semibold">{passengerName}</div>
          </div>
          <div>
            <div className="text-xs text-gray-500">Ghế</div>
            <div className="text-sm font-semibold">{seatNumber}</div>
          </div>
        </div>
        <div className="flex gap-2">
          <Button
            variant="outline"
            size="sm"
            className="flex items-center gap-2 text-primary hover:text-primary-foreground hover:bg-orange transition-colors"
            onClick={onView} // Kích hoạt sự kiện Xem vé
          >
            <Eye className="h-4 w-4" />
            <span>Xem vé</span>
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
