import { NextRouter } from 'next/router';
import { ReactNode, useContext, useState } from 'react';
import GlobalAppContext from '../../layout/context/GlobalAppContext';
import {
  InventoryItem,
  View
} from '../hooks/useInventory/types/useInventoryTypes';

type InventoryLayoutProps = {
  children: ReactNode;
  views: View[] | undefined;
  router: NextRouter;
  error: boolean;
  inventory: InventoryItem[] | undefined;
  searchedInventory: InventoryItem[] | undefined;
};

function InventoryLayout({
  children,
  views,
  router,
  error,
  inventory,
  searchedInventory
}: InventoryLayoutProps) {
  const [query, setQuery] = useState('');
  const { displayBanner } = useContext(GlobalAppContext);

  let newView = views;

  if (query && views && views.length > 0) {
    newView = views.filter(view =>
      view.name.toLowerCase().includes(query.toLowerCase())
    );
  }
  const hasInventory =
    (inventory && inventory.length > 0) ||
    (searchedInventory && searchedInventory.length > 0);

  const hasNoInventory =
    (inventory && inventory.length === 0) ||
    (searchedInventory && searchedInventory.length === 0);

  const hasNoViews = views && views.length === 0;

  const dontDisplaySidebar = (error && hasNoInventory) || hasNoViews;

  return (
    <>
      {!dontDisplaySidebar && (
        <nav
          className={`fixed ${
            displayBanner ? 'mt-[145px]' : 'mt-[73px]'
          } bottom-0 left-0 top-0 z-20 flex w-[17rem] flex-col gap-4 bg-white p-6`}
        >
          <button
            onClick={() => {
              router.push(router.pathname);
            }}
            className={`flex items-center gap-3 rounded-lg px-4 py-3 text-left text-sm font-medium
              ${
                !router.query.view
                  ? 'border-l-2 border-darkcyan-500 bg-cyan-100 text-darkcyan-500'
                  : 'text-gray-700 transition-colors hover:bg-gray-50'
              }
            `}
          >
            <div className={!router.query.view ? 'ml-[-2px]' : ''}>
              <p className="w-[192px] truncate">All resources</p>
            </div>
          </button>
          {views && views.length > 0 && (
            <>
              <div className="relative">
                <input
                  placeholder="Search views"
                  value={query}
                  maxLength={64}
                  onChange={e => setQuery(e.target.value)}
                  className="flex w-full items-center rounded-lg border border-gray-300 px-4 py-3 pl-10 text-sm font-medium text-gray-700 transition-colors hover:border-gray-500 focus-visible:outline-darkcyan-500"
                />
                <div
                  className={`absolute left-4 top-[0.95rem] ${
                    query ? 'cursor-pointer' : ''
                  }`}
                >
                  {!query ? (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke="currentColor"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="1.5"
                        d="M11.5 21a9.5 9.5 0 100-19 9.5 9.5 0 000 19zM22 22l-2-2"
                      ></path>
                    </svg>
                  ) : (
                    <svg
                      onClick={() => setQuery('')}
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      fill="none"
                      viewBox="0 0 24 24"
                      className="hover:bg-gray-50"
                    >
                      <path
                        stroke="currentColor"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="1.5"
                        d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
                      ></path>
                    </svg>
                  )}
                </div>
              </div>
              <div className="-mx-4 -mr-6 flex flex-col gap-4 overflow-auto px-4 pr-6">
                {newView &&
                  newView.length > 0 &&
                  newView.map(view => {
                    const isActive = router.query.view === view.id.toString();
                    return (
                      <button
                        key={view.id}
                        onClick={() => {
                          if (isActive) return;
                          router.push(`?view=${view.id}`);
                        }}
                        className={`flex items-center gap-3 rounded-lg px-4 py-3 text-left text-sm font-medium
              ${
                isActive
                  ? 'border-l-2 border-darkcyan-500 bg-cyan-100 text-darkcyan-500'
                  : 'text-gray-700 transition-colors hover:bg-gray-50'
              }
            `}
                      >
                        <div className={isActive ? 'ml-[-2px]' : ''}>
                          <p className="w-[188px] truncate">{view.name}</p>
                        </div>
                      </button>
                    );
                  })}
              </div>
            </>
          )}
          {query && newView && newView.length === 0 && (
            <div className="flex items-center text-xs text-gray-700">
              There are no views with this name.
            </div>
          )}
        </nav>
      )}
      <main className={!dontDisplaySidebar ? 'ml-[17rem]' : ''}>
        {children}
      </main>
    </>
  );
}

export default InventoryLayout;
