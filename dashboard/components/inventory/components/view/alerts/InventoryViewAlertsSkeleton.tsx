import React from 'react';

function InventoryViewAlertsSkeleton() {
  const rows = Array.from(Array(2).keys());
  return (
    <div className="flex flex-col gap-4">
      {rows.map((row, idx) => (
        <div
          key={idx}
          className="flex animate-pulse cursor-pointer select-none items-center justify-between rounded-lg border border-black-170 p-6 hover:border-black-200"
        >
          <div className="flex items-center gap-4">
            <div className="h-[42px] w-[42px] rounded-full bg-komiser-200/50"></div>
            <div className="flex flex-col gap-2">
              <div className="h-3 w-20 rounded-full bg-komiser-200/50"></div>
              <div className="h-2 w-48 rounded-full bg-komiser-200/50"></div>
            </div>
          </div>
          <div className="h-6 w-6 rounded-full bg-komiser-200/50"></div>
        </div>
      ))}
    </div>
  );
}

export default InventoryViewAlertsSkeleton;
