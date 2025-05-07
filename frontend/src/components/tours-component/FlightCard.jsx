import { useState } from "react";
import { ChevronsRight, Info, Plane, RefreshCw, Luggage } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";

export function FlightCard({ flights, passengerCount, onSelectFlight }) {
  const [expandedFlight, setExpandedFlight] = useState(null);
  const [expandedClass, setExpandedClass] = useState(null);

  // Toggle expanded state
  const handleExpand = (flightIndex, classType) => {
    if (expandedFlight === flightIndex && expandedClass === classType) {
      setExpandedFlight(null);
      setExpandedClass(null);
    } else {
      setExpandedFlight(flightIndex);
      setExpandedClass(classType);
    }
  };

  return (
    <div className="flex-1 space-y-4">
      {flights.map((flight, index) => (
        <Card key={index}>
          <CardContent className="p-4">
            <div className="flex flex-col">
              {/* Header */}
              <div className="flex justify-between items-center">
                {/* Flight Info */}
                <div className="flex-1">
                  <div className="flex items-center gap-4 mb-4">
                    <div className="flex flex-col items-center">
                      <span className="text-2xl font-bold">{flight.departureTime}</span>
                      <span className="text-sm text-gray-500">{flight.departureCode}</span>
                    </div>
                    <div className="flex-1 relative">
                      <div className="border-t border-gray-300 absolute w-full top-4"></div>
                      <ChevronsRight className="text-orange-500 absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2" />
                    </div>
                    <div className="flex flex-col items-center mr-10">
                      <span className="text-2xl font-bold">{flight.arrivalTime}</span>
                      <span className="text-sm text-gray-500">{flight.arrivalCode}</span>
                    </div>
                  </div>
                  <div className="items-center gap-4 text-sm text-gray-600 mb-2">
                    <span className="inline-block mr-8">• Thời gian bay dự kiến: {flight.duration}</span>
                    <span className="inline-block">• Số hiệu: {flight.flightNumber}</span>
                  </div>
                  <Dialog>
                    <DialogTrigger asChild>
                      <button className="text-sm text-blue-600 hover:underline flex items-center gap-1">
                        Chi tiết hành trình
                        <Info className="h-4 w-4" />
                      </button>
                    </DialogTrigger>
                    <DialogContent className="sm:max-w-[550px]">
                      <DialogHeader>
                        <DialogTitle className="text-2xl font-bold mb-4 text-teal-700">
                          Chi tiết hành trình
                        </DialogTitle>
                      </DialogHeader>
                      <div className="space-y-6">
                        <div className="flex items-center justify-between bg-gray-100 p-4 rounded-lg">
                          <div className="flex flex-col items-center">
                            <div className="text-3xl font-bold text-teal-700">{flight.departureTime}</div>
                            <div className="text-lg font-semibold">{flight.departureCode}</div>
                            <div className="text-sm text-gray-600">{flight.departureAirport}</div>
                          </div>
                          <div className="flex flex-col items-center">
                            <Plane className="text-orange-500 mb-2" />
                            <div className="text-sm font-medium">{flight.duration}</div>
                          </div>
                          <div className="flex flex-col items-center">
                            <div className="text-3xl font-bold text-teal-700">{flight.arrivalTime}</div>
                            <div className="text-lg font-semibold">{flight.arrivalCode}</div>
                            <div className="text-sm text-gray-600">{flight.arrivalAirport}</div>
                          </div>
                        </div>
                        <div className="space-y-2">
                          <div className="flex items-center gap-2">
                            <span className="font-semibold text-gray-700">Khởi hành:</span>
                            <span>{flight.departureDate}</span>
                          </div>
                          <div className="flex items-center gap-2">
                            <span className="font-semibold text-gray-700">Số hiệu chuyến bay:</span>
                            <span>{flight.flightNumber}</span>
                          </div>
                          <div className="flex items-center gap-2">
                            <span className="font-semibold text-gray-700">Hãng hàng không:</span>
                            <span>{flight.airline}</span>
                          </div>
                          <div className="flex items-center gap-2">
                            <span className="font-semibold text-gray-700">Loại máy bay:</span>
                            <span>{flight.aircraft}</span>
                          </div>
                        </div>
                      </div>
                    </DialogContent>
                  </Dialog>
                </div>

                {/* Price Buttons */}
                <div className="flex flex-col space-y-2 text-right w-44">
                  <Button
                    className={`flex-1 w-full relative ${expandedFlight === index && expandedClass === "economy" ? "bg-teal-800" : "bg-teal-700"} hover:bg-teal-800`}
                    onClick={() => handleExpand(index, "economy")}
                  >
                    <div>
                      <div className="font-semibold">Phổ thông</div>
                      <div>{flight.economyPrice.toLocaleString()} VND</div>
                    </div>
                    <span className="absolute -top-2 -right-2 bg-orange text-white text-xs px-1.5 py-0.5 rounded-full">
                      Còn {flight.seatsLeft} ghế
                    </span>
                  </Button>
                  <Button
                    className={`flex-1 w-full bg-yellow-500 hover:bg-yellow-600 text-black ${expandedFlight === index && expandedClass === "business" ? "bg-yellow-600" : ""}`}
                    onClick={() => handleExpand(index, "business")}
                  >
                    <div>
                      <div className="font-semibold">Thương gia</div>
                      <div>{flight.businessPrice.toLocaleString()} VND</div>
                    </div>
                  </Button>
                </div>
              </div>

              {/* Expanded Section */}
              {expandedFlight === index && ["economy", "business"].includes(expandedClass) && (
                <div className="mt-4 border-t pt-4">
                  <h3 className="font-semibold mb-4">Chọn giá vé</h3>
                  <div className="grid md:grid-cols-2 gap-4">
                    {(expandedClass === "economy" ? flight.economyOptions : flight.businessOptions).map((option, optionIndex) => (
                      <div key={optionIndex} className="border rounded-lg p-4 relative">
                        <div className="text-lg font-semibold mb-2">{option.name}</div>
                        <div className="text-xl font-bold text-teal-700 mb-4">
                          {option.price.toLocaleString()} VND
                        </div>
                        <div className="space-y-3 text-sm">
                          <div className="flex items-start gap-2">
                            <RefreshCw className="h-4 w-4 mt-1" />
                            <div>
                              <div className="font-medium">Thay đổi vé</div>
                              <div className="text-gray-600">
                                Phí đổi vé tối đa {option.changeFee.toLocaleString()} VND mỗi hành khách
                              </div>
                            </div>
                          </div>
                          <div className="flex items-start gap-2">
                            <RefreshCw className="h-4 w-4 mt-1" />
                            <div>
                              <div className="font-medium">Hoàn vé</div>
                              <div className="text-gray-600">
                                Phí hoàn vé tối đa {option.refundFee.toLocaleString()} VND mỗi hành khách
                              </div>
                            </div>
                          </div>
                          <div className="flex items-start gap-2">
                            <Luggage className="h-4 w-4 mt-1" />
                            <div>
                              <div className="font-medium">Hành lý ký gửi</div>
                              <div className="text-gray-600">{option.checkedBaggage}</div>
                            </div>
                          </div>
                          <div className="flex items-start gap-2">
                            <Luggage className="h-4 w-4 mt-1" />
                            <div>
                              <div className="font-medium">Hành lý xách tay</div>
                              <div className="text-gray-600">{option.carryOn}</div>
                            </div>
                          </div>
                        </div>
                        <Button
                          variant="orange"
                          className="w-full mt-4"
                          onClick={() =>
                            onSelectFlight({
                              id: flight.id,
                              selectedOptionId: option.id,
                              passengerCount,
                            })
                          }
                        >
                          Chọn
                        </Button>
                      </div>
                    ))}
                  </div>
                  <div className="text-center text-sm text-gray-500 mt-4">Vui lòng chọn giá vé để tiếp tục.</div>
                </div>
              )}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
