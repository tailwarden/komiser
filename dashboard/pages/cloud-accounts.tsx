import Head from 'next/head';
import Image from 'next/image';
import { useEffect, useRef, useState } from 'react';

import classNames from 'classnames';

import providers from '../utils/providerHelper';

import Toast from '../components/toast/Toast';
import Modal from '../components/modal/Modal';
import Button from '../components/button/Button';
import EditIcon from '../components/icons/EditIcon';
import More2Icon from '../components/icons/More2Icon';
import DeleteIcon from '../components/icons/DeleteIcon';
import AlertCircleIcon from '../components/icons/AlertCircleIcon';
import CloudAccountsHeader from '../components/cloud-account/components/CloudAccountsHeader';
import CloudAccountsLayout from '../components/cloud-account/components/CloudAccountsLayout';

import useCloudAccount from '../components/cloud-account/hooks/useCloudAccounts/useCloudAccount';

function CloudAccounts() {
  const optionsRef = useRef<HTMLDivElement | null>(null);
  const [clickedItemId, setClickedItemId] = useState<string | null>(null);
  const [editCloudAccount, setEditCloudAccount] = useState<boolean>(false);
  const [removeCloudAccount, setRemoveCloudAccount] = useState<{
    state: boolean;
    accountName: string;
  }>({
    state: false,
    accountName: ''
  });

  const { router, cloudAccounts, toast, dismissToast, isNotCustomView } =
    useCloudAccount();

  useEffect(() => {
    const handleOutsideClick = (event: MouseEvent) => {
      if (
        optionsRef.current &&
        !optionsRef.current.contains(event.target as Node)
      ) {
        setClickedItemId(null); // Close the options if clicked outside
      }
    };

    document.addEventListener('mousedown', handleOutsideClick);

    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
    };
  }, []);

  const toggleOptions = (itemId: string) => {
    setClickedItemId(prevClickedItemId => {
      if (prevClickedItemId === itemId) {
        return null; // Close on Clicking the same item's icon
      }
      return itemId;
    });
  };

  const closeRemoveModal = () => {
    setRemoveCloudAccount({
      state: false,
      accountName: ''
    });
  };

  const deleteCloudAccount = () => {
    const removalName = removeCloudAccount.accountName;
    console.log('deleting', removalName);
    // TODO: (onboarding-wizard) handle account removal API call here
  };

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
          const isOpen = clickedItemId === name;

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
                  'group relative rounded-3xl px-2 py-1 text-sm',
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
                <div className="pointer-events-none invisible absolute z-10 -ml-20 mt-2 rounded-lg bg-gray-800 p-2 text-xs text-white transition-opacity duration-300 group-hover:visible">
                  {status.message}
                </div>
              </div>

              <More2Icon
                className="h-6 w-6 cursor-pointer"
                onClick={() => toggleOptions(name)}
              />

              {isOpen && (
                <div
                  ref={optionsRef}
                  className="absolute right-0 top-0 mr-5 mt-[70px] items-center rounded-md border border-black-130 bg-white p-4 shadow-xl"
                  style={{ zIndex: 1000 }}
                >
                  <button
                    className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-black-400 hover:bg-black-150"
                    onClick={() => {
                      setEditCloudAccount(true);
                      setClickedItemId(null);
                    }}
                  >
                    <EditIcon className="mr-2 h-6 w-6" />
                    Edit cloud account
                  </button>
                  <button
                    className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-error-600 hover:bg-black-150"
                    onClick={() => {
                      setRemoveCloudAccount({
                        state: true,
                        accountName: name
                      });
                      setClickedItemId(null);
                    }}
                  >
                    <DeleteIcon className="mr-2 h-6 w-6" />
                    Remove account
                  </button>
                </div>
              )}
            </div>
          );
        })}
      </CloudAccountsLayout>

      {/* Delete Modal */}
      <Modal
        isOpen={removeCloudAccount.state}
        closeModal={() => closeRemoveModal()}
      >
        <div className="flex max-w-xl flex-col gap-y-6 p-8 text-black-400">
          <div className="flex flex-col items-center gap-y-6">
            <AlertCircleIcon className="h-16 w-16" />
            <h1 className="text-center text-xl font-semibold text-black-800">
              Are you sure you want to remove this cloud account?
            </h1>
            <h3 className="text-center">
              All related data (like custom views and tags) will be deleted and
              the {removeCloudAccount.accountName} account will be disconnected
              from Komiser.
            </h3>
          </div>
          <div className="flex flex-row place-content-end gap-x-8">
            <Button style="text" onClick={() => closeRemoveModal()}>
              Cancel
            </Button>
            <Button style="delete" onClick={() => deleteCloudAccount()}>
              Delete account
            </Button>
          </div>
        </div>
      </Modal>

      {/* Edit Drawer */}
      <Modal
        isOpen={editCloudAccount}
        closeModal={() => setEditCloudAccount(false)}
      >
        <div>Editing</div>
        <div>Replace this with the drawer</div>
      </Modal>

      {/* Toast component */}
      {toast && <Toast {...toast} dismissToast={dismissToast} />}
    </>
  );
}

export default CloudAccounts;
