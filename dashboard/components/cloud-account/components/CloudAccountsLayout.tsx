import { NextRouter } from 'next/router';
import { ReactNode, useContext } from 'react';

import platform, { allProviders } from '@utils/providerHelper';
import Button from '@components/button/Button';
import GlobalAppContext from '../../layout/context/GlobalAppContext';
import { CloudAccount } from '../hooks/useCloudAccounts/useCloudAccount';

type CloudAccountsLayoutProps = {
  cloudAccounts: CloudAccount[];
  children: ReactNode;
  router: NextRouter;
};

function CloudAccountsLayout({
  cloudAccounts,
  children,
  router
}: CloudAccountsLayoutProps) {
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
                  ? 'border-l-2 border-darkcyan-500 bg-cyan-100 text-darkcyan-500'
                  : 'text-gray-700 transition-colors hover:bg-gray-50'
              }
            `}
        >
          <div className={!router.query.view ? 'ml-[-2px]' : ''}>
            <p className="w-[192px] truncate">All Cloud Providers</p>
          </div>
        </button>

        {cloudProviders && cloudProviders.length > 0 && (
          <div className="-mx-4 -mr-6 flex flex-col gap-4 overflow-auto px-4 pr-6">
            {cloudProviders
              .filter(provider =>
                cloudAccounts.some(
                  account =>
                    account.provider.toLowerCase() ===
                    provider.toLocaleLowerCase()
                )
              )
              .map(provider => {
                const isActive = router.query.view === provider;
                return (
                  <button
                    key={provider}
                    onClick={() => {
                      if (isActive) return;
                      router.push(`?view=${provider}`);
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
                      <p className="w-[188px] truncate">
                        {platform.getLabel(provider)}
                      </p>
                    </div>
                  </button>
                );
              })}
          </div>
        )}
        <div className="flex flex-col justify-end absolute bottom-10 border-t border-gray-300 p-4">
          <Button
            onClick={() => {
              router.push('/onboarding/choose-cloud/');
            }}
          >
            Connect account
          </Button>
        </div>
      </nav>
      <main className="ml-[17rem]">{children}</main>
    </>
  );
}

export default CloudAccountsLayout;
