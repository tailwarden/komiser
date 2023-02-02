import { NextRouter } from 'next/router';
import { ReactNode } from 'react';
import { ViewProps } from '../hooks/useInventory';

type InventoryLayoutProps = {
  children: ReactNode;
  views: ViewProps[] | undefined;
  router: NextRouter;
};

function InventoryLayout({ children, views, router }: InventoryLayoutProps) {
  return (
    <>
      {views && views.length > 0 && (
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
            {views.map(view => {
              const isActive = router.query.view === view.id.toString();
              return (
                <button
                  key={view.id}
                  onClick={() => {
                    if (isActive) return;
                    router.push(`/?view=${view.id}`);
                  }}
                  className={`flex items-center gap-3 rounded-lg px-4 py-3 text-sm font-medium
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
          </nav>
          <main className="ml-[17rem]">{children}</main>
        </>
      )}
    </>
  );
}

export default InventoryLayout;
