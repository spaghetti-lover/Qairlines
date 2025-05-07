import React from 'react';

export function FlightSelectionNotice({ isSelectingReturn }) {
  return (
    <div className="bg-blue-100 border-l-4 border-blue-500 text-blue-700 p-4 mb-4 rounded">
      <p className="font-bold">
        {isSelectingReturn ? "Chọn chuyến bay về" : "Chọn chuyến bay đi"}
      </p>
      <p>
        {isSelectingReturn
          ? "Vui lòng chọn chuyến bay về của bạn."
          : "Vui lòng chọn chuyến bay đi của bạn."}
      </p>
    </div>
  );
}
