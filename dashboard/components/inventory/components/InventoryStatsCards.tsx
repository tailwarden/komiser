import formatNumber from "../../../utils/formatNumber";
import { InventoryStats } from "../hooks/useInventory";

function InventoryStatsCards(inventoryStats: InventoryStats) {
  return (
    <div className="grid grid-col md:grid-cols-2 lg:grid-cols-3 gap-8">
      <div className="flex items-center gap-4 py-8 px-6 bg-white  text-black-900  rounded-lg w-full transition-colors">
        <div className=" bg-komiser-100 p-4 rounded-lg">
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
              strokeWidth="2"
              d="M3.17 7.44L12 12.55l8.77-5.08M12 21.61v-9.07"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M9.93 2.48L4.59 5.44c-1.21.67-2.2 2.35-2.2 3.73v5.65c0 1.38.99 3.06 2.2 3.73l5.34 2.97c1.14.63 3.01.63 4.15 0l5.34-2.97c1.21-.67 2.2-2.35 2.2-3.73V9.17c0-1.38-.99-3.06-2.2-3.73l-5.34-2.97c-1.15-.63-3.01-.63-4.15.01z"
            ></path>
          </svg>
        </div>
        <div className="flex flex-col">
          <p className="text-xl font-medium">{inventoryStats.resources}</p>
          <p className="text-sm text-black-300">Resources</p>
        </div>
      </div>
      <div className="flex items-center gap-4 py-8 px-6 bg-white  text-black-900  rounded-lg w-full transition-colors">
        <div className=" bg-komiser-100 p-4 rounded-lg">
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
              strokeWidth="2"
              d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10z"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M8 3h1a28.424 28.424 0 000 18H8M15 3a28.424 28.424 0 010 18"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M3 16v-1a28.424 28.424 0 0018 0v1M3 9a28.424 28.424 0 0118 0"
            ></path>
          </svg>
        </div>
        <div className="flex flex-col">
          <p className="text-xl font-medium">{inventoryStats.regions}</p>
          <p className="text-sm text-black-300">Regions</p>
        </div>
      </div>
      <div className="flex items-center gap-4 py-8 px-6 bg-white  text-black-900  rounded-lg w-full transition-colors">
        <div className=" bg-komiser-100 p-4 rounded-lg">
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
              strokeWidth="2"
              d="M18.04 13.55c-.42.41-.66 1-.6 1.63.09 1.08 1.08 1.87 2.16 1.87h1.9v1.19c0 2.07-1.69 3.76-3.76 3.76H6.26c-2.07 0-3.76-1.69-3.76-3.76v-6.73c0-2.07 1.69-3.76 3.76-3.76h11.48c2.07 0 3.76 1.69 3.76 3.76v1.44h-2.02c-.56 0-1.07.22-1.44.6z"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M2.5 12.41V7.84c0-1.19.73-2.25 1.84-2.67l7.94-3a1.9 1.9 0 012.57 1.78v3.8M22.559 13.97v2.06c0 .55-.44 1-1 1.02h-1.96c-1.08 0-2.07-.79-2.16-1.87-.06-.63.18-1.22.6-1.63.37-.38.88-.6 1.44-.6h2.08c.56.02 1 .47 1 1.02zM7 12h7"
            ></path>
          </svg>
        </div>
        <div className="flex flex-col">
          <p className="text-xl font-medium">
            ${formatNumber(inventoryStats.costs)}
          </p>
          <p className="text-sm text-black-300">Cost</p>
        </div>
      </div>
    </div>
  );
}

export default InventoryStatsCards;
