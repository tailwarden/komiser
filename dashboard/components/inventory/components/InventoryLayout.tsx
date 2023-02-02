import Link from 'next/link';
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
          <nav className="fixed top-0 left-0 bottom-0 z-20 mt-[73px] flex w-[18rem] flex-col gap-4 bg-white p-8">
            <Link href="/">
              <div
                className={`flex items-center gap-3 rounded-lg px-4 py-3 text-sm font-medium
              ${
                router.asPath === '/'
                  ? 'border-l-2 border-primary bg-komiser-150 text-primary'
                  : 'text-black-400 transition-colors hover:bg-komiser-100'
              }
            `}
              >
                <div className={router.asPath === '/' ? 'ml-[-2px]' : ''}>
                  <p className="w-[192px] truncate">All resources</p>
                </div>
              </div>
            </Link>
            {views.map(item => {
              const isActive = router.query.view === item.id.toString();
              return (
                <Link key={item.id} href={`/?view=${item.id}`} passHref>
                  <div
                    className={`flex items-center gap-3 rounded-lg px-4 py-3 text-sm font-medium
              ${
                isActive
                  ? 'border-l-2 border-primary bg-komiser-150 text-primary'
                  : 'text-black-400 transition-colors hover:bg-komiser-100'
              }
            `}
                  >
                    <div className={isActive ? 'ml-[-2px]' : ''}>
                      <p className="w-[192px] truncate">{item.name}</p>
                    </div>
                  </div>
                </Link>
              );
            })}
          </nav>
          <main className="ml-[18rem]">{children}</main>
        </>
      )}
    </>
  );
}

export default InventoryLayout;
