import Head from 'next/head';
import Image from 'next/image';

import classNames from 'classnames';

import providers from '../utils/providerHelper';

import Toast from '../components/toast/Toast';
import More2Icon from '../components/icons/More2Icon';
import CloudAccountsHeader from '../components/cloud-account/components/CloudAccountsHeader';
import CloudAccountsLayout from '../components/cloud-account/components/CloudAccountsLayout';

import useCloudAccount from '../components/cloud-account/hooks/useCloudAccounts/useCloudAccount';

function CloudAccounts() {
  const { router, cloudAccounts, toast, dismissToast, isNotCustomView } =
    useCloudAccount();

  return (
    <>
      <Head>
        <title>Cloud Accounts - Komiser</title>
        <meta name="description" content="Cloud Accounts - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {/* Wraps the cloud account page and handles the custom views sidebar */}
      <CloudAccountsLayout router={router}>
        <CloudAccountsHeader isNotCustomView={isNotCustomView} />

        {cloudAccounts.map(account => {
          const { provider, name, status } = account;

          return (
            <div
              key={name}
              className="relative my-5 flex w-full items-center gap-4 rounded-lg border-2 border-black-170 bg-white p-6 text-black-900 transition-colors"
            >
              <Image
                src={providers.providerImg(provider) as string}
                alt={`${name} image`}
                width={150}
                height={150}
                className="h-12 w-12 rounded-full"
              />
              <div className="mr-auto">
                <p className="font-bold">{name}</p>
                <p className="text-black-300">
                  {providers.providerLabel(provider)}
                </p>
              </div>
              <div
                className={classNames(
                  'group relative rounded-3xl py-1 px-2 text-sm',
                  {
                    'bg-green-200 text-green-600': status.state === 'Connected',
                    'bg-red-200 text-red-600':
                      status.state === 'Permission Issue',
                    'bg-komiser-200 text-komiser-600':
                      status.state === 'Syncing'
                  }
                )}
              >
                <span>{status.state}</span>
                <div className="pointer-events-none invisible absolute z-10 mt-2 -ml-20 rounded-lg bg-gray-800 p-2 text-xs text-white transition-opacity duration-300 group-hover:visible">
                  {status.message}
                </div>
              </div>
              <More2Icon className="h-6 w-6" />
            </div>
          );
        })}
      </CloudAccountsLayout>

      {/* Toast component */}
      {toast && <Toast {...toast} dismissToast={dismissToast} />}
    </>
  );
}

export default CloudAccounts;
