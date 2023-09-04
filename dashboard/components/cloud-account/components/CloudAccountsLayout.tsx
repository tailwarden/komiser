import { NextRouter } from 'next/router';
import { ReactNode, useContext } from 'react';

import GlobalAppContext from '../../layout/context/GlobalAppContext';
import Providers, { allProviders } from '../../../utils/providerHelper';

type CloudAccountsLayoutProps = {
  children: ReactNode;
  router: NextRouter;
};

function CloudAccountsLayout({ children, router }: CloudAccountsLayoutProps) {
  const { displayBanner } = useContext(GlobalAppContext);

  const cloudProviders = Object.values(allProviders);

  return (
    <>
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
                  ? 'border-l-2 border-primary bg-komiser-150 text-primary'
                  : 'text-black-400 transition-colors hover:bg-komiser-100'
              }
            `}
        >
          <div className={!router.query.view ? 'ml-[-2px]' : ''}>
            <p className="w-[192px] truncate">All Cloud Providers</p>
          </div>
        </button>

        {cloudProviders && cloudProviders.length > 0 && (
          <div className="-mx-4 -mr-6 flex flex-col gap-4 overflow-auto px-4 pr-6">
            {cloudProviders.map(view => {
              const isActive = router.query.view === view;
              return (
                <button
                  key={view}
                  onClick={() => {
                    if (isActive) return;
                    router.push(`?view=${view}`);
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
                    <p className="w-[188px] truncate">
                      {Providers.providerLabel(view)}
                    </p>
                  </div>
                </button>
              );
            })}
          </div>
        )}
      </nav>
      <main className="ml-[17rem]">{children}</main>
    </>
  );
}

export default CloudAccountsLayout;
