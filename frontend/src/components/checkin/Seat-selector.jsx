'use client';

import React, { useState } from "react";
import { MdChair } from "react-icons/md";
import { cn } from "@/lib/utils";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

export default function SeatSelector({ seats, passengers, onSeatSelect }) {
  const [selectedCustomer, setSelectedCustomer] = useState(null);
  const [tempSelectedSeat, setTempSelectedSeat] = useState(null);
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleCustomerClick = (customer) => {
    setSelectedCustomer(customer);
  };

  const handleSeatClick = (seatId) => {
    const seat = seats.find((s) => s.id === seatId);
    if (seat?.type === "available" && selectedCustomer) {
      setTempSelectedSeat(seatId);
      setIsDialogOpen(true);
    }
  };

  const confirmSeatSelection = () => {
    if (tempSelectedSeat && selectedCustomer) {
      onSeatSelect(tempSelectedSeat, selectedCustomer.id);
      setIsDialogOpen(false);
      setTempSelectedSeat(null);
    }
  };

  return (
    <div className="w-full max-w-4xl mx-auto p-4 flex flex-col md:flex-row gap-4">
      {/* Left Panel */}
      <div className="w-full md:w-72 space-y-4">
        <div className="border rounded-lg p-4">
          <h3 className="font-medium mb-2">Danh sách khách hàng</h3>
          <ul className="space-y-2">
            {passengers.map((customer) => (
              <li
                key={customer.id}
                className={cn(
                  "flex items-center justify-between text-sm p-2 rounded cursor-pointer",
                  selectedCustomer?.id === customer.id ? "bg-blue-100" : "hover:bg-gray-100"
                )}
                onClick={() => handleCustomerClick(customer)}
              >
                <span>{customer.name}</span>
                <span className={customer.seat ? "text-green-500" : "text-amber-500"}>
                  {customer.seat || "Chưa chọn chỗ"}
                </span>
              </li>
            ))}
          </ul>
        </div>

        <div className="border rounded-lg p-4">
          <div className="text-sm font-medium mb-4">Xem chú giải</div>
          <div className="space-y-3">
            <div className="flex items-center gap-2">
              <div className="w-8 h-8 flex items-center justify-center text-blue-500 bg-blue-100 border-2 border-blue-400 rounded">
                <MdChair size={24} />
              </div>
              <span className="text-sm">Chỗ ngồi đã chọn</span>
            </div>
            <div className="flex items-center gap-2">
              <div className="w-8 h-8 flex items-center justify-center text-amber-500 bg-amber-50 border-2 border-amber-200 rounded">
                <MdChair size={24} />
              </div>
              <span className="text-sm">Không còn trống</span>
            </div>
            <div className="flex items-center gap-2">
              <div className="w-8 h-8 flex items-center justify-center text-zinc-500 bg-white border-2 border-zinc-300 rounded">
                <MdChair size={24} />
              </div>
              <span className="text-sm">Chỗ ngồi còn trống</span>
            </div>
            <div className="flex items-center gap-2">
              <div className="w-8 h-8 flex items-center justify-center text-red-500 bg-red-100 border-2 border-red-400 rounded">
                <MdChair size={24} />
              </div>
              <span className="text-sm">Hàng ghế được chặn</span>
            </div>
          </div>
        </div>
      </div>

      {/* Right Panel - Seat Map */}
      <div className="flex-1 relative overflow-auto max-h-[80vh] md:max-h-none">
        {selectedCustomer ? (
          <div className="mb-4 p-2 bg-blue-100 rounded">
            Đang chọn ghế cho: {selectedCustomer.name}
          </div>
        ) : (
          <div className="mb-4 p-2 bg-yellow-100 rounded">
            Vui lòng chọn một khách hàng trước khi chọn ghế
          </div>
        )}
        <div className="grid grid-cols-[auto_repeat(6,1fr)] gap-x-2 px-12">
          {/* Column Headers */}
          <div className="h-8"></div>
          {["A", "B", "C", "D", "E", "G"].map((col) => (
            <div key={col} className="flex items-center justify-center h-8 text-sm font-medium">
              {col}
            </div>
          ))}

          {/* Seats Grid */}
          {Array.from({ length: 44 }, (_, i) => i + 1).map((row) => (
            <React.Fragment key={`row-${row}`}>
              <div className="flex items-center justify-end h-8 text-sm font-medium pr-2">
                {row}
              </div>
              {["A", "B", "C", "D", "E", "G"].map((col) => {
                const seat = seats.find((s) => s.id === `${row}${col}`);
                return (
                  <div key={`${row}${col}`} className="flex items-center justify-center">
                    <button
                      onClick={() => handleSeatClick(`${row}${col}`)}
                      disabled={!selectedCustomer || seat?.type === "blocked" || seat?.type === "unavailable"}
                      className={cn(
                        "w-8 h-8 rounded transition-colors flex items-center justify-center",
                        seat?.type === "available" && "bg-white border-zinc-300 hover:border-zinc-400 text-zinc-500",
                        seat?.type === "selected" && "bg-cyan-100 border-cyan-400 text-cyan-500",
                        seat?.type === "unavailable" && "bg-amber-50 border-amber-200 text-amber-500",
                        seat?.type === "blocked" && "bg-red-100 border-red-400 text-red-500"
                      )}
                    >
                      <MdChair size={30} />
                    </button>
                  </div>
                );
              })}
            </React.Fragment>
          ))}
        </div>
      </div>

      {/* Confirmation Dialog */}
      <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Xác nhận chọn ghế</DialogTitle>
            <DialogDescription>
              Bạn đã chọn ghế {tempSelectedSeat} cho {selectedCustomer?.name}. Bạn có chắc chắn muốn chọn ghế này không?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsDialogOpen(false)}>
              Hủy
            </Button>
            <Button variant="orange" onClick={confirmSeatSelection}>Xác nhận</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
