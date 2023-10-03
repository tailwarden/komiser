import { useEffect, useRef, useState } from 'react';
import Head from 'next/head';
import Image from 'next/image';
import { useRouter } from 'next/router';

import providers from '../utils/providerHelper';

import Toast from '../components/toast/Toast';
import Modal from '../components/modal/Modal';
import EditIcon from '../components/icons/EditIcon';
import More2Icon from '../components/icons/More2Icon';
import DeleteIcon from '../components/icons/DeleteIcon';
import CloudAccountsHeader from '../components/cloud-account/components/CloudAccountsHeader';
import CloudAccountsLayout from '../components/cloud-account/components/CloudAccountsLayout';

import useCloudAccount from '../components/cloud-account/hooks/useCloudAccounts/useCloudAccount';
import CloudAccountsSidePanel from '../components/cloud-account/components/CloudAccountsSidePanel';
import CloudAccountStatus from '../components/cloud-account/components/CloudAccountStatus';
import CloudAccountDeleteContents from '../components/cloud-account/components/CloudAccountDeleteContents';
import useToast from '../components/toast/hooks/useToast';

function CloudAccounts() {
  const optionsRef = useRef<HTMLDivElement | null>(null);
  const [clickedItemId, setClickedItemId] = useState<string | null>(null);
  const [editCloudAccount, setEditCloudAccount] = useState<boolean>(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);

  const { toast, setToast, dismissToast } = useToast();
  const router = useRouter();

  const currentViewProvider = router.query.view as string;

  const {
    cloudAccounts,
    setCloudAccounts,
    openModal,
    cloudAccountItem,
    setCloudAccountItem,
    page,
    goTo,
    isNotCustomView,
    isLoading
  } = useCloudAccount();

  const [filteredCloudAccounts, setFilteredCloudAccounts] =
    useState(cloudAccounts);

  useEffect(() => {
    if (!currentViewProvider) setFilteredCloudAccounts(cloudAccounts);
    else {
      setFilteredCloudAccounts(
        cloudAccounts.filter(
          account =>
            account.provider.toLowerCase() === currentViewProvider.toLowerCase()
        )
      );
    }
  }, [currentViewProvider, cloudAccounts]);

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
    setIsDeleteModalOpen(false);
  };

  if (!cloudAccounts || isLoading) return null;

  return (
    <>
      <Head>
        <title>Cloud Accounts - Komiser</title>
        <meta name="description" content="Cloud Accounts - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {/* Wraps the cloud account page and handles the custom views sidebar */}
      <CloudAccountsLayout router={router} cloudAccounts={cloudAccounts}>
        <CloudAccountsHeader isNotCustomView={isNotCustomView} />

        {filteredCloudAccounts.map(account => {
          const { provider, name, status } = account;
          const isOpen = clickedItemId === name;

          return (
            <div
              key={name}
              onClick={() => openModal(account)}
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

              <CloudAccountStatus status={status} />

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
                      setIsDeleteModalOpen(true);
                      setCloudAccountItem(account);
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
      <Modal isOpen={isDeleteModalOpen} closeModal={() => closeRemoveModal()}>
        <div className="flex max-w-xl flex-col gap-y-6 p-8 text-black-400">
          {cloudAccountItem && (
            <CloudAccountDeleteContents
              cloudAccount={cloudAccountItem}
              onCancel={closeRemoveModal}
              setToast={setToast}
            />
          )}
        </div>
      </Modal>

      {cloudAccountItem && (
        <CloudAccountsSidePanel
          isOpen={editCloudAccount}
          closeModal={() => setEditCloudAccount(false)}
          cloudAccount={cloudAccountItem}
          cloudAccounts={cloudAccounts}
          setCloudAccounts={setCloudAccounts}
          setToast={setToast}
          page={page}
          goTo={goTo}
        />
      )}

      {/* Toast component */}
      {toast && <Toast {...toast} dismissToast={dismissToast} />}
    </>
  );
}

export default CloudAccounts;
