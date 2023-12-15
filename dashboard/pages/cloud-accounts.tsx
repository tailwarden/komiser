import { useEffect, useRef, useState } from 'react';
import Head from 'next/head';
import { useRouter } from 'next/router';

import CloudAccountItem from '@components/cloud-account/components/CloudAccountItem';
import Toast from '@components/toast/Toast';
import Modal from '@components/modal/Modal';
import CloudAccountsHeader from '@components/cloud-account/components/CloudAccountsHeader';
import CloudAccountsLayout from '@components/cloud-account/components/CloudAccountsLayout';

import useCloudAccount from '@components/cloud-account/hooks/useCloudAccounts/useCloudAccount';
import CloudAccountsSidePanel from '@components/cloud-account/components/CloudAccountsSidePanel';
import CloudAccountDeleteContents from '@components/cloud-account/components/CloudAccountDeleteContents';
import { useToast } from '@components/toast/ToastProvider';

import EmptyState from '@components/empty-state/EmptyState';
import Banner from '@components/banner/Banner';
import Button from '@components/button/Button';

function CloudAccounts() {
  const [editCloudAccount, setEditCloudAccount] = useState<boolean>(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);

  const [isTailwardenBannerDismissed, setIsTailwardenBannerDismissed] =
    useState(true);

  const { toast, showToast, dismissToast } = useToast();
  const router = useRouter();

  const currentViewProvider = router.query.view as string;

  const hideTailwardenBanner = () => {
    setIsTailwardenBannerDismissed(true);
    window.localStorage.setItem('tailwardenBannerDismissed', 'true');
  };

  useEffect(() => {
    setIsTailwardenBannerDismissed(
      window.localStorage.getItem('tailwardenBannerDismissed') === 'true'
    );
  }, []);

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

  const closeRemoveModal = () => {
    setIsDeleteModalOpen(false);
  };

  const handleAfterDelete = (account: any) => {
    setCloudAccounts(
      cloudAccounts.filter((item: any) => item.id !== account.id)
    );
    closeRemoveModal();
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
        {filteredCloudAccounts.map(account => (
          <CloudAccountItem
            key={account.id}
            account={account}
            openModal={openModal}
            setCloudAccountItem={setCloudAccountItem}
            setIsDeleteModalOpen={setIsDeleteModalOpen}
            setEditCloudAccount={setEditCloudAccount}
          />
        ))}

        {!cloudAccounts.length && (
          <div className="mt-12">
            <EmptyState
              title="We could not find a cloud account"
              message="It seems you have not connected a cloud account to Komiser, connect one right now so you can start managing it with more ease"
              action={() => {
                router.push('/onboarding/choose-cloud');
              }}
              actionLabel="Connect a cloud account"
              secondaryAction={() => {
                router.push(
                  'https://github.com/tailwarden/komiser/issues/new/choose'
                );
              }}
              secondaryActionLabel="Report an issue"
              mascotPose="thinking"
            />
          </div>
        )}
      </CloudAccountsLayout>

      {/* Delete Modal */}
      <Modal isOpen={isDeleteModalOpen} closeModal={() => closeRemoveModal()}>
        <div className="flex max-w-xl flex-col gap-y-6 p-8 text-gray-700">
          {cloudAccountItem && (
            <CloudAccountDeleteContents
              cloudAccount={cloudAccountItem}
              onCancel={closeRemoveModal}
              showToast={showToast}
              handleAfterDelete={handleAfterDelete}
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
          handleAfterDelete={handleAfterDelete}
          showToast={showToast}
          page={page}
          goTo={goTo}
        />
      )}

      {cloudAccounts.length >= 2 && !isTailwardenBannerDismissed && (
        <div className="bg-white absolute bottom-0 left-0 right-0 z-20 px-28 border-black-170 border-t text-base py-3 flex gap-4 items-center justify-center">
          For deeper insights and account-level alerts, make the switch to
          Tailwarden — our recommended cloud version for production use.{' '}
          <Button
            size="xs"
            gap="md"
            asLink
            href="https://tailwarden.com/?utm_source=komiser"
            target="_blank"
          >
            Discover Tailwarden
          </Button>
          <Button
            size="xs"
            gap="md"
            style="ghost"
            onClick={() => hideTailwardenBanner()}
          >
            X
          </Button>
        </div>
      )}
    </>
  );
}

export default CloudAccounts;
