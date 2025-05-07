'use client';

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";

export function PassengerListStep({ passengers, onContinue, onBack }) {
  return (
    <Card className="w-full max-w-4xl mx-auto mb-8">
      <CardContent className="p-6">
        <h2 className="text-2xl font-semibold mb-6">Danh Sách Hành Khách</h2>

        {/* Danh sách hành khách chuyến đi */}
        <h3 className="text-xl font-medium mb-4 text-orange">Chuyến đi</h3>
        {passengers.departure.length > 0 ? (
          <div className="space-y-4">
            {passengers.departure.map((passenger) => (
              <div
                key={passenger.id}
                className="flex items-center p-4 border rounded-lg hover:bg-gray-50"
              >
                <div>
                  <div className="font-medium">
                    {passenger.title} {passenger.name}
                  </div>
                  <div className="text-sm text-gray-500">
                    {passenger.type === "business" ? "Thương gia" : "Phổ thông"}
                  </div>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <p className="text-gray-500">Không có hành khách nào trong chuyến đi.</p>
        )}

        {/* Danh sách hành khách chuyến về */}
        {passengers.return.length > 0 && (
          <>
            <h3 className="text-xl font-medium mt-6 mb-4 text-green-600">Chuyến về</h3>
            <div className="space-y-4">
              {passengers.return.map((passenger) => (
                <div
                  key={passenger.id}
                  className="flex items-center p-4 border rounded-lg hover:bg-gray-50"
                >
                  <div>
                    <div className="font-medium">
                      {passenger.title} {passenger.name}
                    </div>
                    <div className="text-sm text-gray-500">
                      {passenger.type === "business" ? "Thương gia" : "Phổ thông"}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </>
        )}

        <div className="flex justify-end gap-4 mt-6">
          <Button variant="outline" onClick={onBack}>
            Quay Lại
          </Button>
          <Button variant="orange" onClick={onContinue}>
            Tiếp Tục
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
