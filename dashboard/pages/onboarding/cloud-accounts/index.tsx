import { useState } from 'react';
import router from 'next/router';
import Head from 'next/head';
import Link from 'next/link';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '@components/onboarding-wizard/OnboardingWizardLayout';
import PlusIcon from '@components/icons/PlusIcon';
import platform from '@utils/providerHelper';
import DeleteIcon from '@components/icons/DeleteIcon';
import Modal from '@components/modal/Modal';
import CloudAccountDeleteContents from '@components/cloud-account/components/CloudAccountDeleteContents';
import Toast from '@components/toast/Toast';

import useCloudAccount from '@components/cloud-account/hooks/useCloudAccounts/useCloudAccount';
import Button from '@components/button/Button';
import { useToast } from '@components/toast/ToastProvider';
import Avatar from '@components/avatar/Avatar';

export default function CloudAccounts() {
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);

  const { toast, showToast, dismissToast } = useToast();

  const {
    cloudAccounts,
    setCloudAccounts,
    cloudAccountItem,
    setCloudAccountItem
  } = useCloudAccount();

  const closeRemoveModal = () => {
    setIsDeleteModalOpen(false);
  };

  const handleDelete = (account: any) => {
    setCloudAccountItem(account);
    setIsDeleteModalOpen(true);
  };

  const handleAfterDelete = (account: any) => {
    setCloudAccounts(
      cloudAccounts.filter((item: any) => item.id !== account.id)
    );
    setIsDeleteModalOpen(false);
  };

  return (
    <div>
      <Head>
        <title>Onboarding - Komiser</title>
        <meta name="description" content="Onboarding - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout title="Connected cloud accounts" progressBarWidth="35%">
          <div className="mb-8 leading-6 text-gray-900/60">
            <div className="font-normal">
              Here, you can add more cloud accounts or edit/delete existing
              ones, before moving to the next step.
            </div>
          </div>
          <div className="mb-4 space-y-4">
            <Link
              href={'/onboarding/choose-cloud/'}
              className="flex w-full items-center rounded border-[1.5px] border-darkcyan-500 bg-transparent p-6 text-darkcyan-500 hover:bg-darkcyan-100"
            >
              <PlusIcon className="my-4 ml-2 mr-6 h-6 w-6" />
              Add cloud account
            </Link>
            {cloudAccounts.map(account => (
              <div
                key={account.id}
                className="flex items-center justify-between rounded-lg border border-gray-300 p-6"
              >
                <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
                  <Avatar avatarName={account.provider} size={40} />
                  <div className="flex flex-col gap-1">
                    <div className="flex max-w-[14rem] items-center gap-1">
                      <p className="truncate font-medium text-gray-950">
                        {account.name}
                      </p>
                    </div>
                    <p className="flex items-center gap-2 text-xs text-gray-500">
                      {platform.getLabel(account.provider)}
                    </p>
                  </div>
                </div>
                <div className="flex gap-5">
                  <button
                    className="hidden items-center gap-2 transition-colors hover:text-darkcyan-500 md:flex"
                    onClick={() => handleDelete(account)}
                  >
                    <DeleteIcon className="h-4 w-4" />
                  </button>
                </div>
              </div>
            ))}
          </div>
          <div className="fixed bottom-0 -mx-20 flex w-[calc(100%*6/11)] justify-end border-t border-gray-300 bg-white px-20 py-4">
            <Button onClick={() => router.push('/onboarding/complete')}>
              Next
            </Button>
          </div>
        </LeftSideLayout>

        <RightSideLayout isCustom={true} customClasses="flex flex-col p-4">
          <div className="grid w-full grid-cols-7 gap-3">
            {/* Row 1 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-300"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 2 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200">
              <div className="relative bottom-3 left-3 h-full w-full scale-110 overflow-clip rounded-lg shadow-right">
                <Avatar avatarName="aws" />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 3 */}
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200">
              <div className="relative h-full w-full overflow-clip rounded-lg shadow-right">
                <Avatar avatarName="civo" />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-400"></div>

            {/* Row 4 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-400">
              <div className="relative h-full w-full overflow-clip rounded-lg bg-white shadow-right">
                <Avatar avatarName="gcp" />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 5 */}
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-400"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200">
              <div className="relative left-3 top-3  h-full w-full overflow-clip rounded-lg bg-gray-950 shadow-right">
                <Avatar avatarName="azure" />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 6 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-400"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 7 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-cyan-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>

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
    </div>
  );
}
