import { NextRouter } from 'next/router';
import { ReactNode, useState } from 'react';
import { ViewProps } from '../hooks/useInventory';

type InventoryLayoutProps = {
  children: ReactNode;
  views: ViewProps[] | undefined;
  router: NextRouter;
};

function InventoryLayout({ children, views, router }: InventoryLayoutProps) {
  const [query, setQuery] = useState('');
  let newView = views;

  if (query && views && views.length > 0) {
    newView = views.filter(view =>
      view.name.toLowerCase().includes(query.toLowerCase())
    );
  }

  return (
    <>
      <nav className="fixed top-0 left-0 bottom-0 z-20 mt-[73px] flex w-[17rem] flex-col gap-4 bg-white p-6">
        <button
          onClick={() => {
            if (!router.query.view) return;
            router.push('/');
          }}
          className={`flex items-center gap-3 rounded-lg px-4 py-3 text-left text-sm font-medium
              ${
                !router.query.view
                  ? 'border-l-2 border-primary bg-komiser-150 text-primary'
                  : 'text-black-400 transition-colors hover:bg-komiser-100'
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
                onChange={e => setQuery(e.target.value)}
                className="flex w-full items-center rounded-lg border border-black-200 px-4 py-3 pl-10 text-sm font-medium text-black-400 transition-colors hover:border-black-300 focus-visible:outline-primary"
              />
              <div
                className={`absolute top-[0.95rem] left-4 ${
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
                  >
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
                    ></path>
                  </svg>
                )}
              </div>
            </div>
            {newView &&
              newView.length > 0 &&
              newView.map(view => {
                const isActive = router.query.view === view.id.toString();
                return (
                  <button
                    key={view.id}
                    onClick={() => {
                      if (isActive) return;
                      router.push(`/?view=${view.id}`);
                    }}
                    className={`flex items-center gap-3 rounded-lg px-4 py-3 text-left text-sm font-medium
              ${
                isActive
                  ? 'border-l-2 border-primary bg-komiser-150 text-primary'
                  : 'text-black-400 transition-colors hover:bg-komiser-100'
              }
            `}
                  >
                    <div className={isActive ? 'ml-[-2px]' : ''}>
                      <p className="w-[188px] truncate">{view.name}</p>
                    </div>
                  </button>
                );
              })}
          </>
        )}
        {query && newView && newView.length === 0 && (
          <div className="flex items-center text-xs text-black-400">
            There are no custom views with this name.
          </div>
        )}
      </nav>
      <main className="ml-[17rem]">{children}</main>
    </>
  );
}

export default InventoryLayout;
