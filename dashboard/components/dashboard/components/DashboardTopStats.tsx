import React from 'react';
import Card from '../../card/Card';
import CardSkeleton from '../../card/CardSkeleton';

function DashboardTopStats() {
  return (
    <div className="grid-col grid gap-8 md:grid-cols-2 lg:grid-cols-4">
      <CardSkeleton />
      <Card
        label="Cloud accounts"
        value={4}
        tooltip="Number of connected cloud accounts"
        icon={
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            viewBox="0 0 24 24"
            className="flex-shrink-0"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="1.5"
              d="M3.17 7.44L12 12.55l8.77-5.08M12 21.61v-9.07"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="1.5"
              d="M9.93 2.48L4.59 5.44c-1.21.67-2.2 2.35-2.2 3.73v5.65c0 1.38.99 3.06 2.2 3.73l5.34 2.97c1.14.63 3.01.63 4.15 0l5.34-2.97c1.21-.67 2.2-2.35 2.2-3.73V9.17c0-1.38-.99-3.06-2.2-3.73l-5.34-2.97c-1.15-.63-3.01-.63-4.15.01z"
            ></path>
          </svg>
        }
      />
    </div>
  );
}

export default DashboardTopStats;
