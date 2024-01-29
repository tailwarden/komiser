import { useRouter } from 'next/router';
import { ErrorIcon } from '@components/icons';
import formatNumber from '../../../utils/formatNumber';
import Tooltip from '../../tooltip/Tooltip';
import {
  HiddenResource,
  InventoryStats
} from '../hooks/useInventory/types/useInventoryTypes';

type InventoryStatsCardsProps = {
  inventoryStats: InventoryStats | undefined;
  isSomeServiceUnavailable: boolean | undefined;
  error: boolean;
  statsLoading: boolean;
  hiddenResources: HiddenResource[] | undefined;
};

function InventoryStatsCards({
  inventoryStats,
  isSomeServiceUnavailable,
  error,
  statsLoading,
  hiddenResources
}: InventoryStatsCardsProps) {
  const router = useRouter();

  return (
    <>
      {!statsLoading &&
        inventoryStats &&
        inventoryStats.resources !== 0 &&
        Object.keys(inventoryStats).length !== 0 &&
        !error && (
          <div
            className={`grid-col grid md:grid-cols-2 ${
              router.query.view ? 'lg:grid-cols-4' : 'lg:grid-cols-3'
            } gap-8 ${isSomeServiceUnavailable ? 'mt-8' : ''}`}
          >
            <div className="relative flex w-full items-center gap-4 rounded-lg bg-white  px-6  py-8 text-gray-950 transition-colors">
              <div className=" rounded-lg bg-gray-50 p-4">
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
              </div>
              <div className="peer flex flex-col">
                <p className="text-xl font-medium">
                  {formatNumber(inventoryStats.resources, 'full')}
                </p>
                <p className="text-sm text-gray-500">Resources</p>
              </div>
              <Tooltip>Number of active cloud services</Tooltip>
            </div>
            <div className="relative flex w-full items-center gap-4 rounded-lg bg-white  px-6  py-8 text-gray-950 transition-colors">
              <div className=" rounded-lg bg-gray-50 p-4">
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
                    d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10z"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="1.5"
                    d="M8 3h1a28.424 28.424 0 000 18H8M15 3a28.424 28.424 0 010 18"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="1.5"
                    d="M3 16v-1a28.424 28.424 0 0018 0v1M3 9a28.424 28.424 0 0118 0"
                  ></path>
                </svg>
              </div>
              <div className="peer flex flex-col">
                <p className="text-xl font-medium">{inventoryStats.regions}</p>
                <p className="text-sm text-gray-500">Regions</p>
              </div>
              <Tooltip>
                Number of regions where you have active cloud services
              </Tooltip>
            </div>
            <div className="relative flex w-full items-center gap-4 rounded-lg bg-white  px-6  py-8 text-gray-950 transition-colors">
              <div className=" rounded-lg bg-gray-50 p-4">
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
                    d="M18.04 13.55c-.42.41-.66 1-.6 1.63.09 1.08 1.08 1.87 2.16 1.87h1.9v1.19c0 2.07-1.69 3.76-3.76 3.76H6.26c-2.07 0-3.76-1.69-3.76-3.76v-6.73c0-2.07 1.69-3.76 3.76-3.76h11.48c2.07 0 3.76 1.69 3.76 3.76v1.44h-2.02c-.56 0-1.07.22-1.44.6z"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="1.5"
                    d="M2.5 12.41V7.84c0-1.19.73-2.25 1.84-2.67l7.94-3a1.9 1.9 0 012.57 1.78v3.8M22.559 13.97v2.06c0 .55-.44 1-1 1.02h-1.96c-1.08 0-2.07-.79-2.16-1.87-.06-.63.18-1.22.6-1.63.37-.38.88-.6 1.44-.6h2.08c.56.02 1 .47 1 1.02zM7 12h7"
                  ></path>
                </svg>
              </div>
              <div className="peer flex flex-col">
                <p className="text-xl font-medium">
                  ${formatNumber(inventoryStats.costs)}
                </p>
                <p className="text-sm text-gray-500">Discoverd Cost</p>
              </div>
              {isSomeServiceUnavailable && (
                <div
                  onClick={() =>
                    window.open('https://www.tailwarden.com/', '_blank')
                  }
                  className="rounded-s absolute -top-[22px] -right-[22px] bg-white w-[44px] h-[44px] flex justify-center items-center border-2 border-gray-50"
                >
                  <ErrorIcon
                    className="inline peer cursor-pointer"
                    width={24}
                    height={24}
                  />
                  <Tooltip align="right" width="xl" bottom="sm" top="xs">
                    We couldn&apos;t determine the exact cost of your resources
                    as some cloud providers service&apos;s costs are not yet
                    supported — we suggest trying Tailwarden.
                  </Tooltip>
                </div>
              )}
              <Tooltip>Up-to-date monthly cost</Tooltip>
            </div>
            {router.query.view && hiddenResources && (
              <div className="relative flex w-full items-center gap-4 rounded-lg bg-white  px-6  py-8 text-gray-950 transition-colors">
                <div className=" rounded-lg bg-gray-50 p-4">
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
                      d="M15.03 9.47l-5.06 5.06a3.576 3.576 0 115.06-5.06z"
                    ></path>
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="1.5"
                      d="M18.32 5.77c-1.75-1.32-3.75-2.04-5.82-2.04-3.53 0-6.82 2.08-9.11 5.68-.9 1.41-.9 3.78 0 5.19.79 1.24 1.71 2.31 2.71 3.17"
                    ></path>
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="1.5"
                      d="M8.92 19.53c1.14.48 2.35.74 3.58.74 3.53 0 6.82-2.08 9.11-5.68.9-1.41.9-3.78 0-5.19-.33-.52-.69-1.01-1.06-1.47M16.01 12.7a3.565 3.565 0 01-2.82 2.82M9.97 14.53L2.5 22M22.5 2l-7.47 7.47"
                    ></path>
                  </svg>
                </div>
                <div className="peer flex flex-col">
                  <p className="text-xl font-medium">
                    {formatNumber(hiddenResources.length)}
                  </p>
                  <p className="text-sm text-gray-500">Hidden resources</p>
                </div>
                <Tooltip>
                  Resources that will be hidden from the inventory
                </Tooltip>
              </div>
            )}
          </div>
        )}
    </>
  );
}

export default InventoryStatsCards;
