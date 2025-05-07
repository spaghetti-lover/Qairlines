import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Calendar, Clock, CreditCard, User, Plane } from 'lucide-react'

export function FlightDetailCard({
  type,
  flightNumber,
  departureTime,
  arrivalTime,
  departureCity,
  arrivalCity,
  date,
  duration,
  passengers,
  paymentMethod,
}) {
  return (
    <Card className="w-full border-primary/20 shadow-md">
      <CardHeader className="flex flex-row items-center justify-between bg-primary/5 border-b border-primary/10">
        <CardTitle className="text-lg text-primary">Chi tiết {type === "departure" ? "chuyến đi" : "chuyến về"}</CardTitle>
        <span className="text-sm font-medium text-primary/80">{flightNumber}</span>
      </CardHeader>
      <CardContent className="pt-6">
        <div className="flex justify-between items-center mb-6">
          <div className="text-center">
            <div className="text-2xl font-bold text-primary">{departureTime}</div>
            <div className="text-sm font-medium text-gray-600">{departureCity}</div>
          </div>
          <div className="flex-1 mx-4 border-t-2 border-primary/30 relative">
            <Plane className="text-primary absolute top-1/2 left-1/2 transform -translate-y-1/2 -translate-x-1/2" />
          </div>
          <div className="text-center">
            <div className="text-2xl font-bold text-primary">{arrivalTime}</div>
            <div className="text-sm font-medium text-gray-600">{arrivalCity}</div>
          </div>
        </div>
        
        <div className="grid grid-cols-2 gap-4">
          <div className="flex items-center gap-2">
            <Calendar className="h-4 w-4 text-primary" />
            <div>
              <div className="text-xs text-gray-500">Ngày khởi hành</div>
              <div className="font-medium">{date}</div>
            </div>
          </div>
          
          <div className="flex items-center gap-2">
            <Clock className="h-4 w-4 text-primary" />
            <div>
              <div className="text-xs text-gray-500">Thời gian bay</div>
              <div className="font-medium">{duration}</div>
            </div>
          </div>
          
          <div className="flex items-center gap-2">
            <User className="h-4 w-4 text-primary" />
            <div>
              <div className="text-xs text-gray-500">Hành khách</div>
              <div className="font-medium">{passengers}</div>
            </div>
          </div>
          
          <div className="flex items-center gap-2">
            <CreditCard className="h-4 w-4 text-primary" />
            <div>
              <div className="text-xs text-gray-500">Phương thức thanh toán</div>
              <div className="font-medium">{paymentMethod}</div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}

